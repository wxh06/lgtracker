package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

// Contest 用来描述比赛的详情与题目摘要
type Contest struct {
	Details  json.RawMessage   `json:"details"`
	Problems []json.RawMessage `json:"problems"`
}

// ContestData 将 luogu.ContestData 的部分字段替换为 json.RawMessage，用以原样输出 JSON
type ContestData struct {
	Contest         json.RawMessage `json:"contest"`
	ContestProblems []struct {
		Problem json.RawMessage `json:"problem"`
	} `json:"contestProblems"`
}

// 用户赛时各题的通过情况
var users = map[int]map[int]map[string]int{}

// fetchScoreboard 获取指定比赛的用户通过情况
func fetchScoreboard(id int) error {
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
				users[result.User.Uid] = map[int]map[string]int{}
			}
			users[result.User.Uid][id] = map[string]int{}
			for problem, score := range details {
				users[result.User.Uid][id][problem] = score.Score
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
		Contests luogu.List[luogu.Contest] `json:"contests"`
	}]]("GET", "https://www.luogu.com.cn/contest/list", nil)
	if err != nil {
		panic(err)
	}

	// 每场比赛的详情与题目摘要
	var contests = make(map[int]Contest)

	// 从文件中读取已保存的比赛数据
	if f, err := os.Open("src/contests.json"); err == nil {
		if err = json.NewDecoder(f).Decode(&contests); err != nil {
			panic(err)
		}
		if err = f.Close(); err != nil {
			panic(err)
		}
	}

	// 从文件中读取每个保存的用户通过情况
	entries, err := os.ReadDir("public/users")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		name := entry.Name()
		uid, err := strconv.Atoi(strings.Split(name, ".")[0])
		if err != nil {
			continue
		}

		f, err := os.Open("public/users/" + name)
		if err != nil {
			panic(err)
		}

		var passed map[int]map[string]int
		if err = json.NewDecoder(f).Decode(&passed); err != nil {
			panic(err)
		}
		users[uid] = passed

		if err = f.Close(); err != nil {
			panic(err)
		}
	}

	for _, contest := range data.CurrentData.Contests.Result {
		data, err := luogu.Get[luogu.DataResponse[ContestData]](fmt.Sprintf("https://www.luogu.com.cn/contest/%d", contest.Id))
		if err != nil {
			panic(err)
		}
		if data.CurrentData.ContestProblems == nil {
			continue
		}

		problems := []json.RawMessage{}
		for _, problem := range data.CurrentData.ContestProblems {
			problems = append(problems, problem.Problem)
		}
		contests[contest.Id] = Contest{data.CurrentData.Contest, problems}

		if err = fetchScoreboard(contest.Id); err != nil {
			panic(err)
		}
	}

	// 储存比赛数据
	f, err := os.Create("src/contests.json")
	if err != nil {
		panic(err)
	}
	if err = json.NewEncoder(f).Encode(contests); err != nil {
		panic(err)
	}
	if err = f.Close(); err != nil {
		panic(err)
	}

	// 存储每个用户的通过情况
	for user, passed := range users {
		f, err = os.Create(fmt.Sprintf("public/users/%d.json", user))
		if err != nil {
			panic(err)
		}
		if err = json.NewEncoder(f).Encode(passed); err != nil {
			panic(err)
		}
		if err = f.Close(); err != nil {
			panic(err)
		}
	}
}
