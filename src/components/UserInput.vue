<script setup lang="ts">
import { ref, watch } from "vue";
import { storeToRefs } from "pinia";
import { useUserStore } from "@/stores/user";

const store = useUserStore();
const user = storeToRefs(store);

const uid = ref(user.uid.value);
const name = ref(user.name.value ?? "");
const disabled = ref(Boolean(uid.value));

watch(user.uid, (id) => {
  if (id) uid.value = id;
});
watch(user.name, (username) => {
  if (username) name.value = username;
});

function setUser() {
  disabled.value = true;
  if (uid.value) user.uid.value = uid.value;
  else
    store.setUserByName(name.value).catch(() => {
      // TODO: show error state
      disabled.value = false;
    });
}

function clearUser() {
  user.uid.value = undefined;
  disabled.value = false;
}

function clearUid() {
  clearUser();
  uid.value = undefined;
}

function clearName() {
  clearUser();
  name.value = "";
}
</script>

<template>
  <form @submit.prevent="setUser">
    <label
      >用户编号 <input type="number" min="1" v-model="uid" @input="clearName"
    /></label>
    <label>用户名 <input type="text" v-model="name" @input="clearUid" /></label>

    <input type="submit" :disabled="disabled" />
  </form>
</template>
