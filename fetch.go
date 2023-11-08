package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wxh06/luogu-cli/pkg/luogu"
)

// 最大重试次数
const RETRIES = 3

// get 是对 luogu.Get 的封装，尝试进行请求直到成功或超过最大重试次数
func get[T any](url string) (data T, err error) {
	for i := 0; i < RETRIES; i++ {
		data, err = luogu.Get[T](url)
		if err == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return
}

// 用户赛时各题的通过情况
var users map[int]map[string]bool = map[int]map[string]bool{}

// fetchScoreboard 获取指定比赛的用户通过情况
func fetchScoreboard(id int, ruleType int) error {
	logger := log.Default()
	logger.SetPrefix(fmt.Sprintf("%d scoreboard ", id))

	for page := 1; ; page++ {
		data, err := get[luogu.GetScoreboardResponse](fmt.Sprintf("https://www.luogu.com.cn/fe/api/contest/scoreboard/%d?page=%d", id, page))
		if err != nil {
			return err
		}

		for _, result := range data.Scoreboard.Result {
			var details luogu.ScoreDetails
			// 如果 result 为空数组
			if err = json.Unmarshal(result.Details, &details); err != nil {
				continue
			}

			if users[result.User.Uid] == nil {
				users[result.User.Uid] = map[string]bool{}
			}
			for problem, score := range details {
				users[result.User.Uid][problem] = (ruleType != 2 && score.Score >= 100) || (ruleType == 2 && score.Score >= 0)
			}
		}

		logger.Printf("%d/%d", page, (data.Scoreboard.Count+data.Scoreboard.PerPage-1)/data.Scoreboard.PerPage)
		if page*data.Scoreboard.PerPage >= data.Scoreboard.Count {
			break
		}
	}

	return nil
}

func main() {
	data, err := luogu.Request[luogu.DataResponse[struct {
		Contests luogu.List[luogu.Contest[map[string]any]] `json:"contests"`
	}]]("GET", "https://www.luogu.com.cn/contest/list", nil)
	if err != nil {
		panic(err)
	}

	// 每场比赛的题目摘要
	var problems = make(map[int][]luogu.ProblemSummary)
	for _, contest := range data.CurrentData.Contests.Result {
		data, err := luogu.Get[luogu.DataResponse[luogu.ContestData[map[string]any]]](fmt.Sprintf("https://www.luogu.com.cn/contest/%d", contest.Id))
		if err != nil {
			panic(err)
		}
		if data.CurrentData.ContestProblems == nil {
			continue
		}

		problems[contest.Id] = []luogu.ProblemSummary{}
		for _, problem := range data.CurrentData.ContestProblems {
			problems[contest.Id] = append(problems[contest.Id], problem.Problem)
		}

		if err = fetchScoreboard(contest.Id, contest.RuleType); err != nil {
			panic(err)
		}
	}

	// TODO: store problem data

	// 存储每个用户的通过情况
	for user, passed := range users {
		f, err := os.Create(fmt.Sprintf("public/users/%d.json", user))
		if err != nil {
			panic(err)
		}
		if err = json.NewEncoder(f).Encode(passed); err != nil {
			panic(err)
		}
	}
}
