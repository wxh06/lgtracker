<script setup lang="ts">
import contests from "@/contests.json";
import ContestList from "@/components/ContestList.vue";
</script>

<template>
  <ContestList :contests="contests" />
</template>
