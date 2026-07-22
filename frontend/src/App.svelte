<script>
  import { onMount } from 'svelte';
  import { authStore } from '$lib/stores/auth.svelte.js';
  import LoginPage from '$lib/components/LoginPage.svelte';
  import UserDashboard from '$lib/components/UserDashboard.svelte';
  import AdminDashboard from '$lib/components/AdminDashboard.svelte';
  import SuperAdminDashboard from '$lib/components/SuperAdminDashboard.svelte';

  onMount(async () => {
    await authStore.checkAuth();
  });
</script>

{#if $authStore.loading}
  <div class="min-h-screen bg-slate-950 flex flex-col items-center justify-center text-slate-200">
    <div class="inline-flex items-center gap-3 p-4 rounded-xl bg-slate-900 border border-slate-800 shadow-xl">
      <svg class="animate-spin h-6 w-6 text-blue-500" viewBox="0 0 24 24" fill="none">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <span class="text-sm font-medium">Memeriksa otentikasi...</span>
    </div>
  </div>
{:else if !$authStore.isAuthenticated || !$authStore.user}
  <LoginPage />
{:else}
  {#if $authStore.user.role === 'super_admin'}
    <SuperAdminDashboard user={$authStore.user} />
  {:else if $authStore.user.role === 'admin'}
    <AdminDashboard user={$authStore.user} />
  {:else}
    <UserDashboard user={$authStore.user} />
  {/if}
{/if}
