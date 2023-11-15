<script setup lang="ts">
import { ref } from "vue";
import { storeToRefs } from "pinia";
import InputText from "primevue/inputtext";
import InputGroup from "primevue/inputgroup";
import InputGroupAddon from "primevue/inputgroupaddon";
import Button from "primevue/button";
import ProgressSpinner from "primevue/progressspinner";
import { useUserStore } from "@/stores/user";
import UserInfo from "@/components/UserInfo.vue";

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
      <InputGroupAddon v-if="user.uid.value">
        <UserInfo v-if="user.details.value" :user="user.details.value" />
        <ProgressSpinner
          v-else
          style="max-width: 1lh; max-height: 1lh"
          stroke-width="8"
          aria-label="Loading"
        />
      </InputGroupAddon>
      <Button icon="pi pi-check" type="submit" />
    </InputGroup>
  </form>
</template>
