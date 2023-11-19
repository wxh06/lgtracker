<script setup lang="ts">
import { RouterLink } from "vue-router";
import Menubar from "primevue/menubar";
import type { MenuItem } from "primevue/menuitem";
import UserInput from "./UserInput.vue";

const items: MenuItem[] = [
  {
    label: "比赛",
    icon: "pi pi-chart-bar",
    route: "/",
  },
  {
    label: "GitHub",
    icon: "pi pi-github",
    url: "https://github.com/wxh06/lgtracker",
  },
];
</script>

<template>
  <Menubar :model="items">
    <template #item="{ item, props, hasSubmenu }">
      <RouterLink
        v-if="item.route"
        v-slot="{ href, navigate }"
        :to="item.route"
        custom
      >
        <a v-ripple :href="href" v-bind="props.action" @click="navigate">
          <span :class="item.icon" />
          <span style="margin-left: 0.5rem">{{ item.label }}</span>
        </a>
      </RouterLink>
      <a
        v-else
        v-ripple
        :href="item.url"
        :target="item.target"
        v-bind="props.action"
      >
        <span :class="item.icon" />
        <span style="margin-left: 0.5rem">{{ item.label }}</span>
        <span
          v-if="hasSubmenu"
          class="pi pi-fw pi-angle-down"
          style="margin-left: 0.5rem"
        />
      </a>
    </template>
    <template #end>
      <UserInput />
    </template>
  </Menubar>
</template>
