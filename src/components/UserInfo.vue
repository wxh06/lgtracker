<script setup lang="ts">
import type { UserSummary } from "luogu-api-docs/luogu-api";

defineProps<{ user?: UserSummary }>();

const getClassName = (color: string) =>
  "lg-fg-" +
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
    :style="{ textDecoration: 'none' }"
    v-if="user"
  >
    <span :class="getClassName(user.color)">{{ user.name }}</span>
  </a>
</template>

<style>
@import "https://cdn.luogu.com.cn/css/luogu3.css";
</style>
