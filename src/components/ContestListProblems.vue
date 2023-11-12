<script setup lang="ts">
import { storeToRefs } from "pinia";
import type { Contest, ProblemSummary } from "luogu-api-docs/luogu-api";
import { useUserStore } from "@/stores/user";

const props = defineProps<{
  contest: Contest;
  problems: ProblemSummary[];
  scores?: Record<string, number>;
}>();

const user = storeToRefs(useUserStore());

const isPassed = (score: number | undefined, ruleType: number) =>
  score === undefined
    ? undefined
    : (ruleType != 2 && score >= 100) || (ruleType == 2 && score >= 0);

function getColor(contest: Contest, pid: string) {
  const PASSED = "yellowgreen";
  const SUBMITTED = "darkorange";
  const passedDuringContest = isPassed(props.scores?.[pid], contest.ruleType);
  if (passedDuringContest || user.passed.value.has(pid)) return PASSED;
  if (passedDuringContest === false || user.submitted.value.has(pid))
    return SUBMITTED;
}
</script>

<template>
  <li>
    {{ contest.name }}
    <ol>
      <li v-for="problem in problems" :key="problem.pid">
        <span
          :style="{
            backgroundColor: getColor(contest, problem.pid),
          }"
          >{{ problem.pid }}</span
        >
        {{ problem.title }}
      </li>
    </ol>
  </li>
</template>
