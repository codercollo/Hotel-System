<script setup lang="ts">
definePageMeta({ layout: "admin", requiresAuth: true, roles: ["admin"] });

const pageTitle = useState("admin-page-title");
pageTitle.value = "My Account";

// Adjust to your actual auth composable
const { user, logout } = useAuth();

// ── Form state ────────────────────────────────────────────────────────────────
const saving = ref(false);
const saved = ref(false);
const savingPw = ref(false);
const savedPw = ref(false);
const pwError = ref<string | null>(null);

const form = reactive({
  name: user.value?.name ?? "",
  email: user.value?.email ?? "",
  phone: (user.value as any)?.phone ?? "",
});

const pw = reactive({
  current: "",
  next: "",
  confirm: "",
});

// ── Derived ───────────────────────────────────────────────────────────────────
const initials = computed(() => {
  const name = form.name || form.email || "";

  return name
    .split(/[\s@.]+/)
    .filter((s): s is string => s.length > 0)
    .slice(0, 2)
    .map((s) => s.charAt(0).toUpperCase())
    .join("");
});

// ── Handlers ──────────────────────────────────────────────────────────────────
async function saveProfile() {
  saving.value = true;
  saved.value = false;
  try {
    // TODO: PATCH /api/v1/profile  { name, email, phone }
    await new Promise((r) => setTimeout(r, 700)); // simulate
    saved.value = true;
    setTimeout(() => (saved.value = false), 2500);
  } finally {
    saving.value = false;
  }
}

async function changePassword() {
  pwError.value = null;
  savingPw.value = true;
  savedPw.value = false;
  try {
    if (pw.next !== pw.confirm) {
      pwError.value = "New passwords don't match.";
      return;
    }
    if (pw.next.length < 8) {
      pwError.value = "Password must be at least 8 characters.";
      return;
    }
    // TODO: POST /api/v1/profile/password  { current, password }
    await new Promise((r) => setTimeout(r, 700)); // simulate
    savedPw.value = true;
    pw.current = pw.next = pw.confirm = "";
    setTimeout(() => (savedPw.value = false), 2500);
  } finally {
    savingPw.value = false;
  }
}

async function handleLogout() {
  await logout();
  navigateTo("/login");
}
</script>

<template>
  <div class="max-w-2xl mx-auto space-y-6">
    <!-- Avatar + identity card -->
    <div class="card p-6 flex items-center gap-5">
      <div
        class="w-16 h-16 rounded-2xl bg-forest flex items-center justify-center text-white text-xl font-display shrink-0"
      >
        {{ initials }}
      </div>
      <div class="flex-1 min-w-0">
        <p class="font-display text-xl text-brand-900 truncate">
          {{ form.name || "Admin" }}
        </p>
        <p class="text-sm font-sans text-muted truncate">{{ form.email }}</p>
        <span
          class="mt-1 inline-flex items-center gap-1 text-[10px] font-sans uppercase tracking-widest text-forest bg-forest/10 px-2 py-0.5 rounded-full"
        >
          <Icon name="lucide:shield-check" class="w-3 h-3" /> Administrator
        </span>
      </div>
      <button
        class="shrink-0 flex items-center gap-2 text-xs font-sans text-red-500 hover:text-red-600 border border-red-200 hover:border-red-300 hover:bg-red-50 px-3 py-2 rounded-xl transition-all"
        @click="handleLogout"
      >
        <Icon name="lucide:log-out" class="w-3.5 h-3.5" />
        Sign out
      </button>
    </div>

    <!-- Profile form -->
    <div class="card p-6">
      <h3 class="font-display text-lg text-brand-900 mb-5">
        Profile Information
      </h3>

      <div class="space-y-4">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label
              class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >
              Full Name
            </label>
            <input
              v-model="form.name"
              type="text"
              class="input w-full"
              placeholder="Your name"
            />
          </div>
          <div>
            <label
              class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >
              Phone
            </label>
            <input
              v-model="form.phone"
              type="tel"
              class="input w-full"
              placeholder="+1 000 000 0000"
            />
          </div>
        </div>

        <div>
          <label
            class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
          >
            Email Address
          </label>
          <input
            v-model="form.email"
            type="email"
            class="input w-full"
            placeholder="admin@hotel.com"
          />
        </div>
      </div>

      <div class="mt-5 flex items-center gap-3">
        <UiButton :loading="saving" @click="saveProfile">
          <Icon name="lucide:save" class="w-3.5 h-3.5" />
          Save Changes
        </UiButton>
        <Transition name="fade">
          <span
            v-if="saved"
            class="text-xs font-sans text-forest flex items-center gap-1"
          >
            <Icon name="lucide:check-circle" class="w-3.5 h-3.5" /> Saved
          </span>
        </Transition>
      </div>
    </div>

    <!-- Password form -->
    <div class="card p-6">
      <h3 class="font-display text-lg text-brand-900 mb-5">Change Password</h3>

      <div class="space-y-4">
        <div>
          <label
            class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
          >
            Current Password
          </label>
          <input
            v-model="pw.current"
            type="password"
            class="input w-full"
            placeholder="••••••••"
          />
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label
              class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >
              New Password
            </label>
            <input
              v-model="pw.next"
              type="password"
              class="input w-full"
              placeholder="••••••••"
            />
          </div>
          <div>
            <label
              class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >
              Confirm Password
            </label>
            <input
              v-model="pw.confirm"
              type="password"
              class="input w-full"
              placeholder="••••••••"
            />
          </div>
        </div>
      </div>

      <UiAlert
        v-if="pwError"
        variant="error"
        :message="pwError"
        class="mt-4 text-xs"
      />

      <div class="mt-5 flex items-center gap-3">
        <UiButton :loading="savingPw" @click="changePassword">
          <Icon name="lucide:lock" class="w-3.5 h-3.5" />
          Update Password
        </UiButton>
        <Transition name="fade">
          <span
            v-if="savedPw"
            class="text-xs font-sans text-forest flex items-center gap-1"
          >
            <Icon name="lucide:check-circle" class="w-3.5 h-3.5" /> Password
            updated
          </span>
        </Transition>
      </div>
    </div>

    <!-- Danger zone -->
    <div class="card p-6 border-red-100">
      <h3 class="font-display text-lg text-red-600 mb-2">Danger Zone</h3>
      <p class="text-xs font-sans text-muted mb-4">
        These actions are irreversible. Proceed with caution.
      </p>
      <UiButton
        variant="outline"
        size="sm"
        class="text-red-600 border-red-200 hover:bg-red-50 hover:border-red-300"
      >
        <Icon name="lucide:trash-2" class="w-3.5 h-3.5" />
        Delete Account
      </UiButton>
    </div>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
