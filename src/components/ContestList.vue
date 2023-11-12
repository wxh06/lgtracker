<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { storeToRefs } from "pinia";
import type { ContestDetails, ProblemSummary } from "luogu-api-docs/luogu-api";
import { useUserStore } from "@/stores/user";
import ContestListProblems from "@/components/ContestListProblems.vue";

export interface ContestWithProblems {
  details: ContestDetails;
  problems: ProblemSummary[];
}

const { contests } = defineProps<{
  contests: Record<string, ContestWithProblems>;
}>();

const keyword = ref("");
const filteredContests = computed(() =>
  Object.values(contests)
    .filter(({ details: { name } }) =>
      name.toLowerCase().includes(keyword.value.toLowerCase()),
    )
    .sort(({ details: { id: a } }, { details: { id: b } }) =>
      a == b ? 0 : a < b ? 1 : -1,
    ),
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
</script>

<template>
  <label>
    关键词
    <input type="search" v-model="keyword" />
  </label>
  <main>
    <ul>
      <ContestListProblems
        v-for="contest in filteredContests"
        :contest="contest.details"
        :problems="contest.problems"
        :scores="scores[contest.details.id]"
        :key="contest.details.id"
      />
    </ul>
  </main>
</template>
