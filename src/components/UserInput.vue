<script setup lang="ts">
import { ref } from "vue";
import { storeToRefs } from "pinia";
import { useUserStore } from "@/stores/user";

const user = storeToRefs(useUserStore());
const uid = ref(user.uid.value);
</script>

<template>
  <form @submit.prevent="user.uid.value = uid">
    <input type="number" min="1" v-model="uid" />
    <input type="submit" />
  </form>
</template>

<style scoped></style>
