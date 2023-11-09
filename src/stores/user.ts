import { reactive, ref, watch } from "vue";
import { defineStore } from "pinia";
import type { DataResponse, UserData } from "luogu-api-docs/luogu-api";

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
      passed.clear();
      submitted.clear();
      if (!uid) return;

      const r = await fetch(`/user/${uid}`, {
        headers: { "x-luogu-type": "content-only" },
      });
      if (!r.ok) return;
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

  return { uid, name, passed, submitted };
});
