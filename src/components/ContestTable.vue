<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { storeToRefs } from "pinia";
import DataTable, {
  type DataTableOperatorFilterMetaData,
} from "primevue/datatable";
import Column from "primevue/column";
import InputText from "primevue/inputtext";
import ProgressSpinner from "primevue/progressspinner";
import Calendar from "primevue/calendar";
import TriStateCheckbox from "primevue/tristatecheckbox";
import TreeSelect from "primevue/treeselect";
import type { TreeNode } from "primevue/tree";
import { FilterMatchMode, FilterOperator } from "primevue/api";
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
  name: { operator: FilterOperator.OR, constraints: [] },
  rated: { value: null, matchMode: FilterMatchMode.EQUALS },
  startTime: {
    operator: FilterOperator.AND,
    constraints: [
      { value: null, matchMode: FilterMatchMode.DATE_AFTER },
      { value: null, matchMode: FilterMatchMode.DATE_BEFORE },
    ],
  },
  endTime: {
    operator: FilterOperator.AND,
    constraints: [
      { value: null, matchMode: FilterMatchMode.DATE_AFTER },
      { value: null, matchMode: FilterMatchMode.DATE_BEFORE },
    ],
  },
});

type ContestDecoded = Omit<ContestDetails, "startTime" | "endTime"> & {
  startTime: Date;
  endTime: Date;
  problems: ProblemSummary[];
};

const contestsDecoded = computed((): ContestDecoded[] =>
  props.contests.map(({ details, problems }) => ({
    ...details,
    startTime: new Date(details.startTime * 1000),
    endTime: new Date(details.endTime * 1000),
    problems,
  })),
);

const problemIndexes = computed(() => [
  ...Array(
    Math.max(...props.contests.map(({ problems }) => problems.length)),
  ).keys(),
]);

const nodes: TreeNode[] = [
  {
    key: "【LGR",
    label: "官方比赛",
    children: [
      { key: "Div.1】", label: "Div.1" },
      { key: "Div.2】", label: "Div.2" },
      { key: "Div.3】", label: "Div.3" },
      { key: "Div.4】", label: "Div.4" },
    ],
  },
  { key: "ICPC", label: "ICPC" },
];

type TreeSelectModelValue = Record<
  string,
  { checked: boolean; partialChecked: boolean }
>;

const selectedCategories = ref<TreeSelectModelValue>({});

function applySelectedCategories() {
  const f: DataTableOperatorFilterMetaData["constraints"] = [];
  Object.entries(selectedCategories.value).forEach((category) => {
    if (category[1].checked)
      f.push({ value: category[0], matchMode: FilterMatchMode.CONTAINS });
  });
  (filters.value.name as DataTableOperatorFilterMetaData).constraints = f;
}
</script>

<template>
  <main>
    <DataTable
      :value="contestsDecoded"
      data-key="id"
      class="p-datatable-sm"
      showGridlines
      paginator
      :rows="50"
      :rowsPerPageOptions="[20, 50, 100]"
      sortField="startTime"
      :sortOrder="-1"
      v-model:filters="filters"
      filter-display="menu"
      :global-filter-fields="[
        'name',
        ...problemIndexes.map((i) => `problems.${i}.pid`),
        ...problemIndexes.map((i) => `problems.${i}.title`),
      ]"
    >
      <template #header>
        <div style="display: flex; flex-wrap: wrap-reverse">
          <TreeSelect
            :options="nodes"
            v-model="selectedCategories"
            @hide="applySelectedCategories"
            selectionMode="checkbox"
            placeholder="Select Item"
          />
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
        v-for="[field, header] in [
          ['startTime', '开始时间'],
          ['endTime', '结束时间'],
        ] as const"
        :field="field"
        data-type="date"
        :sortable="true"
        :header="header"
        class="time"
        :key="field"
      >
        <template #body="{ data }: { data: ContestDecoded }">
          {{ data[field].toLocaleString("zh") }}
        </template>
        <template #filter="{ filterModel }">
          <Calendar
            v-model="filterModel.value"
            dateFormat="yy/mm/dd"
            placeholder="yyyy/mm/dd"
          />
        </template>
      </Column>
      <Column field="rated" header="Rated" data-type="boolean">
        <template #body="{ data }">
          <span
            class="pi"
            :class="data.rated ? 'pi-check-circle' : 'pi-times-circle'"
          />
        </template>
        <template #filter="{ filterModel }">
          <label for="rated-filter" class="font-bold">Rated </label>
          <TriStateCheckbox
            v-model="filterModel.value"
            input-id="rated-filter"
          />
        </template>
      </Column>
      <Column
        field="name"
        header="比赛名称"
        style="min-width: 16rem; position: relative"
      >
        <template #body="{ data }: { data: ContestDecoded }">
          <a
            :href="`https://www.luogu.com.cn/contest/${data.id}`"
            target="_blank"
            rel="noreferrer"
            class="link"
          >
            {{ data.name }}
            <span style="position: absolute; inset: 0" />
          </a>
        </template>
      </Column>
      <Column
        v-for="i of problemIndexes"
        :field="`problems.${i}`"
        :header="String.fromCharCode(i + 65)"
        :key="i"
        class="problem"
      >
        <template #body="{ data }: { data: ContestDecoded }">
          <ContestTableProblem
            v-if="data.problems[i]"
            :problem="data.problems[i]"
            :problem-prefix-length="problemPrefixLength[data.id]"
            :contest="data"
            :score="scores[data.id]?.[data.problems[i].pid]"
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

.problem .p-column-title {
  margin: auto;
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
