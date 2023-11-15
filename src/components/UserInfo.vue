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
  <span>
    <a
      :href="`https://www.luogu.com.cn/user/${user.uid}`"
      :class="getClassName('fg', user.color)"
      class="vertical-center"
      style="text-decoration: none"
    >
      {{ user.name }}
    </a>
    <Badge
      v-if="user.badge"
      :value="user.badge"
      :class="getClassName('bg', user.color)"
      class="vertical-center margin-left"
    />
    <!-- TODO: color -->
    <span
      v-if="user.ccfLevel"
      class="vertical-center margin-left pi pi-verified"
    />
  </span>
</template>

<style scoped>
@import "https://cdn.luogu.com.cn/css/luogu3.css";

.vertical-center {
  vertical-align: middle;
}

.margin-left {
  margin-left: 0.25rem;
}
</style>
