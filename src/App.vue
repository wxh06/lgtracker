<script setup lang="ts">
import { RouterView } from "vue-router";
import PageHeader from "./components/PageHeader.vue";
</script>

<template>
  <header>
    <PageHeader />
  </header>

  <RouterView />
</template>
