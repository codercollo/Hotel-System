<script setup lang="ts">
// Create a shared state for the header title
// 'admin-page-title' is a key that lets your pages "talk" to this layout
const pageTitle = useState("admin-page-title", () => "Dashboard");
</script>

<template>
  <div class="flex h-screen overflow-hidden bg-surface-100">
    <AppSidebar />
    <div class="flex-1 flex flex-col min-w-0 overflow-y-auto">
      <!-- Admin topbar -->
      <header
        class="bg-white border-b border-surface-200 px-6 py-4 flex items-center justify-between shrink-0"
      >
        <h1 class="font-display text-xl text-brand-900">
          <!-- FIX: Replace <slot> with the shared state variable -->
          {{ pageTitle }}
        </h1>
        <div class="flex items-center gap-4">
          <button
            class="relative text-muted hover:text-brand-900 transition-colors"
          >
            <Icon name="lucide:bell" class="w-5 h-5" />
            <span
              class="absolute -top-1 -right-1 w-2 h-2 bg-accent rounded-full"
            />
          </button>
          <div
            class="w-8 h-8 rounded-full bg-forest flex items-center justify-center text-white text-xs font-medium"
          >
            A
          </div>
        </div>
      </header>
      <main class="flex-1 p-6">
        <!-- This is where your page content (admin/index.vue) will render -->
        <slot />
      </main>
    </div>
  </div>
</template>
