<script setup lang="ts">
import Badge from "primevue/badge";
import type { UserSummary } from "luogu-api-docs/luogu-api";

defineProps<{ user: UserSummary }>();

const getClassName = (type: "fg" | "bg", color: string) =>
  `lg-${type}-` +
  {
    Cheater: "brown",
    Gray: "gray",
    Blue: "bluelight",
    Green: "green",
    Orange: "orange",
    Red: "red",
    Purple: "purple",
  }[color];
</script>

<template>
  <a
    :href="`https://www.luogu.com.cn/user/${user.uid}`"
    :class="getClassName('fg', user.color)"
    style="text-decoration: none"
  >
    {{ user.name }}
  </a>
  <Badge
    v-if="user.badge"
    :value="user.badge"
    :class="getClassName('bg', user.color)"
    style="margin-left: 0.25rem"
  />
  <!-- TODO: color -->
  <span
    v-if="user.ccfLevel"
    class="pi pi-verified"
    style="margin-left: 0.25rem"
  />
</template>

<style>
@import "https://cdn.luogu.com.cn/css/luogu3.css";
</style>
