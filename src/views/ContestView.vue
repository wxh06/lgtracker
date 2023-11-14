<script setup lang="ts">
import contests from "@/contests.json";
import ContestTable from "@/components/ContestTable.vue";
</script>

<template>
  <ContestTable :contests="Object.values(contests)" />
</template>
