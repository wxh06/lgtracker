<script setup lang="ts">
import { storeToRefs } from "pinia";
import type { ProblemSummary } from "luogu-api-docs/luogu-api";
import { useUserStore } from "@/stores/user";

const props = defineProps<{
  problem: ProblemSummary;
  problemPrefixLength: number;
  ruleType: number;
  score?: number;
}>();

const user = storeToRefs(useUserStore());

const isPassed = (score: number | undefined, ruleType: number) =>
  score === undefined
    ? null
    : (ruleType != 2 && score >= 100) || (ruleType == 2 && score >= 0);

function getColor(pid: string, ruleType: number) {
  const PASSED = "yellowgreen";
  const SUBMITTED = "darkorange";
  const passedDuringContest = isPassed(props.score, ruleType);
  if (passedDuringContest || user.passed.value.has(pid)) return PASSED;
  if (passedDuringContest === false || user.submitted.value.has(pid))
    return SUBMITTED;
}
</script>

<template>
  <a
    :href="`https://www.luogu.com.cn/problem/${problem.pid}`"
    target="_blank"
    rel="noreferrer"
  >
    <span :style="{ backgroundColor: getColor(problem.pid, ruleType) }">
      {{ problem.pid }}
    </span>
  </a>
  <br />
  <small
    style="
      display: inline-block;
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
      max-width: calc(100%);
    "
    v-tooltip.bottom="problem.title"
  >
    {{ problem.title.slice(problemPrefixLength) }}
  </small>
</template>
