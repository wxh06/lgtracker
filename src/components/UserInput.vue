<script setup lang="ts">
import { ref } from "vue";
import { storeToRefs } from "pinia";
import { useUserStore } from "@/stores/user";
import UserInfo from "@/components/UserInfo.vue";

const store = useUserStore();
const user = storeToRefs(store);

const input = ref((user.uid.value ?? "").toString());

const isUid = (s: string) => /^\d+$/.test(s);

function setUser() {
  if (isUid(input.value)) user.uid.value = parseInt(input.value, 10);
  else
    store.setUserByName(input.value).catch(() => {
      // TODO: show error state
    });
}
</script>

<template>
  <form @submit.prevent="setUser">
    <label>用户名/编号 <input type="text" v-model="input" /></label>
    <input type="submit" />

    <UserInfo :user="user.details.value" />
  </form>
</template>
