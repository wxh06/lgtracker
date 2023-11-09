<script setup lang="ts">
import { ref, watch } from "vue";
import { storeToRefs } from "pinia";
import type { ContestDetails, ProblemSummary } from "luogu-api-docs/luogu-api";
import _contests from "@/contests.json";
import { useUserStore } from "@/stores/user";

interface Contest {
  details: ContestDetails;
  problems: ProblemSummary[];
}

const contests = Object.values(_contests as Record<string, Contest>).sort(
  ({ details: { id: a } }, { details: { id: b } }) =>
    a == b ? 0 : a < b ? 1 : -1,
);

const user = storeToRefs(useUserStore());
const scores = ref<Record<string, Record<string, number>>>({});
watch(
  user.uid,
  async (uid) => {
    scores.value = {};
    if (!uid) return;

    const r = await fetch(`/users/${uid}.json`);
    if (r.ok) scores.value = await r.json();
  },
  { immediate: true },
);

const isPassed = (score: number | undefined, ruleType: number) =>
  score === undefined
    ? undefined
    : (ruleType != 2 && score >= 100) || (ruleType == 2 && score >= 0);

function getColor(contest: ContestDetails, pid: string) {
  const PASSED = "yellowgreen";
  const SUBMITTED = "darkorange";
  const passedDuringContest = isPassed(
    scores.value[contest.id]?.[pid],
    contest.ruleType,
  );
  if (passedDuringContest || user.passed.value.has(pid)) return PASSED;
  if (passedDuringContest === false || user.submitted.value.has(pid))
    return SUBMITTED;
}
</script>

<template>
  <main>
    <ul>
      <li v-for="contest in contests" :key="contest.details.id">
        {{ contest.details.name }}
        <ol>
          <li v-for="problem in contest.problems" :key="problem.pid">
            <span
              :style="{
                backgroundColor: getColor(contest.details, problem.pid),
              }"
              >{{ problem.pid }}</span
            >
            {{ problem.title }}
          </li>
        </ol>
      </li>
    </ul>
  </main>
</template>
