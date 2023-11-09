<script setup lang="ts">
import { ref, watch } from "vue";
import type { ContestDetails, ProblemSummary } from "luogu-api-docs/luogu-api";
import _contests from "../contests.json";

interface Contest {
  details: ContestDetails;
  problems: ProblemSummary[];
}

const contests = Object.values(_contests as Record<string, Contest>).sort(
  ({ details: { id: a } }, { details: { id: b } }) =>
    a == b ? 0 : a < b ? 1 : -1,
);

const isPassed = (score: number | undefined, ruleType: number) =>
  score && ((ruleType != 2 && score >= 100) || (ruleType == 2 && score >= 0));

const user = ref(119491); // 该项目的诞生离不开 @5ab_juruo！
const scores = ref<Record<string, Record<string, number>>>({});
watch(
  user,
  async (uid) => {
    scores.value = {};
    const r = await fetch(`/users/${uid}.json`);
    if (!r.ok) return;
    scores.value = await r.json();
  },
  { immediate: true },
);
</script>

<template>
  <ul>
    <li v-for="contest in contests" :key="contest.details.id">
      {{ contest.details.name }}
      <ol>
        <li v-for="problem in contest.problems" :key="problem.pid">
          <span
            :style="{
              backgroundColor:
                ((isPassed(
                  scores[contest.details.id]?.[problem.pid],
                  contest.details.ruleType,
                ) &&
                  'yellowgreen') ??
                  'inherit') ||
                'darkorange',
            }"
          >
            {{ problem.pid }}
          </span>
          {{ problem.title }}
        </li>
      </ol>
    </li>
  </ul>
</template>
