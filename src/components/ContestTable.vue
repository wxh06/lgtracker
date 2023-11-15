<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { storeToRefs } from "pinia";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import InputText from "primevue/inputtext";
import ProgressSpinner from "primevue/progressspinner";
import { FilterMatchMode } from "primevue/api";
import type { ContestDetails, ProblemSummary } from "luogu-api-docs/luogu-api";
import { useUserStore } from "@/stores/user";
import ContestTableProblem from "./ContestTableProblem.vue";
import UserInfo from "./UserInfo.vue";

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
  if (problems.length <= 1) return 0;
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
        <div style="display: flex; flex-wrap: wrap-reverse">
          <span class="p-input-icon-left">
            <i class="pi pi-search" />
            <InputText
              v-model="filters['global'].value"
              placeholder="Keyword Search"
            />
          </span>
          <span style="display: flex; margin-left: auto; align-items: center">
            <UserInfo
              v-if="user.details.value"
              style="margin: 0.25rem 0"
              :user="user.details.value"
            />
            <ProgressSpinner
              v-else-if="user.uid.value"
              style="height: 2rem; width: 2rem"
              stroke-width="8"
            />
          </span>
        </div>
      </template>
      <Column
        field="details.name"
        header="比赛名称"
        style="min-width: 16rem; position: relative"
      >
        <template #body="{ data: { details } }: { data: ContestWithProblems }">
          <a
            :href="`https://www.luogu.com.cn/contest/${details.id}`"
            target="_blank"
            rel="noreferrer"
            class="link"
          >
            {{ details.name }}
            <span style="position: absolute; inset: 0" />
          </a>
        </template>
      </Column>
      <Column
        field="details.startTime"
        header="开始时间"
        :sortable="true"
        class="time"
      >
        <template #body="{ data }: { data: ContestWithProblems }">
          {{ new Date(data.details.startTime * 1000).toLocaleString("zh") }}
        </template>
      </Column>
      <Column
        field="details.endTime"
        header="结束时间"
        :sortable="true"
        class="time"
      >
        <template #body="{ data }: { data: ContestWithProblems }">
          {{ new Date(data.details.endTime * 1000).toLocaleString("zh") }}
        </template>
      </Column>
      <Column
        v-for="i of problemIndexes"
        :field="`problems.${i}`"
        :header="`${i + 1}`"
        :key="i"
        class="problem"
      >
        <template #body="{ data }: { data: ContestWithProblems }">
          <ContestTableProblem
            v-if="data.problems[i]"
            :problem="data.problems[i]"
            :problem-prefix-length="problemPrefixLength[data.details.id]"
            :contest="data.details"
            :score="scores[data.details.id]?.[data.problems[i].pid]"
          />
        </template>
      </Column>
    </DataTable>
  </main>
</template>

<style>
th.time {
  white-space: nowrap;
}

td.problem {
  max-width: 8rem;
  padding: 0;
  height: 0;
}

td.problem > * {
  height: 100%;
  padding: 0.1rem;
}
</style>
