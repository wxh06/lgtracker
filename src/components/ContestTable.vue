<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { storeToRefs } from "pinia";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import InputText from "primevue/inputtext";
import { FilterMatchMode } from "primevue/api";
import type { ContestDetails, ProblemSummary } from "luogu-api-docs/luogu-api";
import { useUserStore } from "@/stores/user";
import ContestTableProblem from "@/components/ContestTableProblem.vue";

export interface ContestWithProblems {
  details: ContestDetails;
  problems: ProblemSummary[];
}

const props = defineProps<{
  contests: ContestWithProblems[];
}>();

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

function getProblemPrefixLength(problems: ProblemSummary[]) {
  const titles = problems.map(({ title }) => title);
  const minLength = Math.min(...titles.map(({ length }) => length));
  for (let i = 0; i < minLength; i++)
    for (let j = 1; j < titles.length; j++)
      if (titles[j][i] !== titles[j - 1][i]) return i;
  return minLength;
}

const problemPrefixLength = computed(() =>
  Object.fromEntries(
    props.contests.map(({ details: { id }, problems }) => [
      id,
      getProblemPrefixLength(problems),
    ]),
  ),
);

const filters = ref({
  global: { value: null, matchMode: FilterMatchMode.CONTAINS },
});

const problemIndexes = computed(() => [
  ...Array(
    Math.max(...props.contests.map(({ problems }) => problems.length)),
  ).keys(),
]);
</script>

<template>
  <main>
    <DataTable
      :value="contests"
      data-key="details.id"
      class="p-datatable-sm"
      showGridlines
      paginator
      :rows="50"
      :rowsPerPageOptions="[20, 50, 100]"
      v-model:filters="filters"
      sortField="details.startTime"
      :sortOrder="-1"
      :global-filter-fields="[
        'details.name',
        ...problemIndexes.map((i) => `problems.${i}.pid`),
        ...problemIndexes.map((i) => `problems.${i}.title`),
      ]"
    >
      <template #header>
        <div class="flex justify-content-end">
          <span class="p-input-icon-left">
            <i class="pi pi-search" />
            <InputText
              v-model="filters['global'].value"
              placeholder="Keyword Search"
            />
          </span>
        </div>
      </template>
      <Column field="details.name" header="比赛名称" style="min-width: 16rem" />
      <Column
        field="details.startTime"
        header="开始时间"
        :sortable="true"
        style="min-width: 7rem"
      >
        <template #body="{ data }: { data: ContestWithProblems }">
          {{ new Date(data.details.startTime * 1000).toLocaleString("zh") }}
        </template>
      </Column>
      <Column
        field="details.endTime"
        header="结束时间"
        :sortable="true"
        style="min-width: 7rem"
      >
        <template #body="{ data }: { data: ContestWithProblems }">
          {{ new Date(data.details.endTime * 1000).toLocaleString("zh") }}
        </template>
      </Column>
      <Column
        v-for="i of problemIndexes"
        :field="`problems.${i}.pid`"
        :header="`${i + 1}`"
        :key="i"
        style="max-width: 8rem"
      >
        <template #body="{ data }: { data: ContestWithProblems }">
          <ContestTableProblem
            v-if="data.problems[i]"
            :problem="data.problems[i]"
            :problem-prefix-length="problemPrefixLength[data.details.id]"
            :rule-type="data.details.ruleType"
          />
        </template>
      </Column>
    </DataTable>
  </main>
</template>
