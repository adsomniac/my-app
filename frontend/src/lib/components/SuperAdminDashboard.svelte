<script>
  import Card from './ui/Card.svelte';
  import Button from './ui/Button.svelte';
  import Badge from './ui/Badge.svelte';
  import { authStore } from '../stores/auth.svelte.js';
  import { getAdminTestAPI } from '../api/auth.js';

  export let user = {};

  let adminTestResult = '';

  function handleLogout() {
    authStore.logout();
  }

  async function testAdminAccess() {
    const token = $authStore.token;
    const { status, data } = await getAdminTestAPI(token);
    if (status === 200 && data.success) {
      adminTestResult = data.message;
    }
  }
</script>

<div class="min-h-screen bg-slate-950 text-slate-100 p-6">
  <div class="max-w-4xl mx-auto space-y-6">
    <div class="flex items-center justify-between border-b border-purple-900/50 pb-4">
      <div>
        <h1 class="text-2xl font-bold bg-gradient-to-r from-purple-400 to-indigo-300 bg-clip-text text-transparent">
          Super Admin Dashboard
        </h1>
        <p class="text-slate-400 text-sm">Selamat datang, Super Admin! ({user.email})</p>
      </div>
      <div class="flex items-center gap-3">
        <Badge variant="warning">Role: {user.role}</Badge>
        <Button variant="destructive" size="sm" on:click={handleLogout}>
          Logout
        </Button>
      </div>
    </div>

    <Card classNames="border-purple-900/40 bg-purple-950/20">
      <div class="space-y-4">
        <h2 class="text-lg font-semibold text-purple-200">Akses Penuh Sistem</h2>
        <p class="text-slate-300 text-sm">
          Sebagai Super Admin, Anda memiliki hierarki tertinggi (Level 3) dan dapat mengakses semua endpoint sistem.
        </p>

        <Button variant="default" size="sm" on:click={testAdminAccess}>
          Uji Otorisasi Admin
        </Button>

        {#if adminTestResult}
          <div class="p-3 bg-purple-500/10 border border-purple-500/30 rounded-lg text-purple-300 text-sm">
            👑 {adminTestResult} (Super Admin Otorisasi Berlaku)
          </div>
        {/if}
      </div>
    </Card>
  </div>
</div>
