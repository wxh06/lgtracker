<script setup lang="ts">
import { ref } from "vue";
import { storeToRefs } from "pinia";
import InputText from "primevue/inputtext";
import InputGroup from "primevue/inputgroup";
import InputGroupAddon from "primevue/inputgroupaddon";
import Button from "primevue/button";
import { useUserStore } from "@/stores/user";

const store = useUserStore();
const user = storeToRefs(store);

const input = ref((user.uid.value ?? "").toString());

const isUid = (s: string) => /^\d+$/.test(s);

function setUser() {
  user.uid.value = undefined;
  if (isUid(input.value)) user.uid.value = parseInt(input.value, 10);
  else
    store.setUserByName(input.value).catch(() => {
      // TODO: show error state
    });
}
</script>

<template>
  <form @submit.prevent="setUser">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-user" />
      </InputGroupAddon>
      <InputText
        id="user"
        placeholder="UID / 用户名"
        aria-label="用户编号或用户名"
        v-model="input"
      />
      <Button icon="pi pi-check" type="submit" />
    </InputGroup>
  </form>
</template>
