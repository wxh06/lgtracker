<script setup lang="ts">
import { reactive, ref, watch } from "vue";
import type {
  ContestDetails,
  DataResponse,
  ProblemSummary,
  UserData,
} from "luogu-api-docs/luogu-api";
import _contests from "../contests.json";

interface Contest {
  details: ContestDetails;
  problems: ProblemSummary[];
}
const contests = Object.values(_contests as Record<string, Contest>).sort(
  ({ details: { id: a } }, { details: { id: b } }) =>
    a == b ? 0 : a < b ? 1 : -1,
);

const user = ref(119491); // 该项目的诞生离不开 @5ab_juruo！
const scores = ref<Record<string, Record<string, number>>>({});
const passed = reactive(new Set<string>());
const submitted = reactive(new Set<string>());
watch(
  user,
  (uid) => {
    scores.value = {};
    passed.clear();
    submitted.clear();
    if (!uid) return;

    fetch(`/users/${uid}.json`).then((r) => {
      if (r.ok)
        r.json().then((data) => {
          scores.value = data;
        });
    });

    fetch(`/user/${uid}`, { headers: { "x-luogu-type": "content-only" } }).then(
      (r) => {
        if (r.ok)
          r.json().then((data: DataResponse<UserData>) => {
            data.currentData.passedProblems?.forEach((problem) => {
              passed.add(problem.pid);
            });
            data.currentData.submittedProblems?.forEach((problem) => {
              submitted.add(problem.pid);
            });
          });
      },
    );
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
  if (passedDuringContest || passed.has(pid)) return PASSED;
  if (passedDuringContest === false || submitted.has(pid)) return SUBMITTED;
}
</script>

<template>
  <ul>
    <li v-for="contest in contests" :key="contest.details.id">
      {{ contest.details.name }}
      <ol>
        <li v-for="problem in contest.problems" :key="problem.pid">
          <span
            :style="{ backgroundColor: getColor(contest.details, problem.pid) }"
            >{{ problem.pid }}</span
          >
          {{ problem.title }}
        </li>
      </ol>
    </li>
  </ul>
</template>
