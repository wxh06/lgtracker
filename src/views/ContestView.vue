<script setup lang="ts">
import contests from "@/contests.json";
import ContestTable from "@/components/ContestTable.vue";
import Divider from "primevue/divider";
</script>

<template>
  <Divider />
  <ContestTable :contests="Object.values(contests)" />
</template>
