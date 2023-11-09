<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import UserInput from "@/components/UserInput.vue";
</script>

<template>
  <header>
    <nav>
      <RouterLink to="/">比赛</RouterLink>
    </nav>
    <UserInput />
  </header>

  <RouterView />
</template>
