<script setup lang="ts">
import { storeToRefs } from "pinia";
import type { Contest, ProblemSummary } from "luogu-api-docs/luogu-api";
import { useUserStore } from "@/stores/user";

const props = defineProps<{
  problem: ProblemSummary;
  problemPrefixLength: number;
  contest: Contest;
  score?: number;
}>();

const user = storeToRefs(useUserStore());

const isPassed = (score: number | undefined, ruleType: number) =>
  score === undefined
    ? null
    : (ruleType != 2 && score >= 100) || (ruleType == 2 && score >= 0);

function getColor(pid: string, ruleType: number) {
  const PASSED = "LawnGreen";
  const SUBMITTED = "Gold";
  const passedDuringContest = isPassed(props.score, ruleType);
  if (passedDuringContest || user.passed.value.has(pid)) return PASSED;
  if (passedDuringContest === false || user.submitted.value.has(pid))
    return SUBMITTED;
}
</script>

<template>
  <a
    :href="`https://www.luogu.com.cn/problem/${problem.pid}?contestId=${contest.id}`"
    target="_blank"
    rel="noreferrer"
    class="link"
    :style="{ backgroundColor: getColor(problem.pid, contest.ruleType) }"
    style="
      text-align: center;
      display: flex;
      flex-direction: column;
      justify-content: space-evenly;
    "
  >
    {{ problem.pid }}
    <small
      style="overflow: hidden; white-space: nowrap; text-overflow: ellipsis"
      v-tooltip.bottom="problem.title"
    >
      {{ problem.title.slice(problemPrefixLength) }}
    </small>
  </a>
</template>
