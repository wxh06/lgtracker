import { reactive, ref, watch } from "vue";
import { defineStore } from "pinia";
import type {
  DataResponse,
  UserData,
  UserSummary,
} from "luogu-api-docs/luogu-api";

export const useUserStore = defineStore("user", () => {
  const uid = ref(
    parseInt(
      (import.meta.env.SSR ? undefined : localStorage.getItem("user")) ?? "0",
    ) || undefined,
  );
  const name = ref<string>();
  const passed = reactive(new Set<string>());
  const submitted = reactive(new Set<string>());

  watch(uid, (uid) => {
    if (uid) localStorage.setItem("user", uid.toString());
  });

  watch(
    uid,
    async (uid) => {
      name.value = undefined;
      passed.clear();
      submitted.clear();
      if (!uid) return;

      const r = await fetch(`/user/${uid}`, {
        headers: { "x-luogu-type": "content-only" },
      });
      if (!r.ok) throw new Error(r.statusText);
      const data: DataResponse<UserData> = await r.json();

      name.value = data.currentData.user.name;

      data.currentData.passedProblems?.forEach((problem) => {
        passed.add(problem.pid);
      });
      data.currentData.submittedProblems?.forEach((problem) => {
        submitted.add(problem.pid);
      });
    },
    { immediate: true },
  );

  async function setUserByName(username: string) {
    const r = await fetch(`/api/user/search?keyword=${username}`);
    if (!r.ok) throw new Error(r.statusText);
    const {
      users: [user],
    }: { users: [UserSummary | null] } = await r.json();
    if (user) uid.value = user.uid;
    else throw Error();
  }

  return { uid, name, passed, submitted, setUserByName };
});
