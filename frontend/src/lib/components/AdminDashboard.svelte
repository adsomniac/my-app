<script>
  import Card from './ui/Card.svelte';
  import Button from './ui/Button.svelte';
  import Badge from './ui/Badge.svelte';
  import { authStore } from '../stores/auth.svelte.js';
  import { getAdminTestAPI } from '../api/auth.js';

  export let user = {};

  let adminTestResult = '';
  let adminTestError = '';

  function handleLogout() {
    authStore.logout();
  }

  async function testAdminAccess() {
    adminTestResult = '';
    adminTestError = '';
    const token = $authStore.token;
    const { status, data } = await getAdminTestAPI(token);
    if (status === 200 && data.success) {
      adminTestResult = data.message;
    } else {
      adminTestError = data.message || 'Gagal mengakses area admin';
    }
  }
</script>

<div class="min-h-screen bg-slate-950 text-slate-100 p-6">
  <div class="max-w-4xl mx-auto space-y-6">
    <div class="flex items-center justify-between border-b border-slate-800 pb-4">
      <div>
        <h1 class="text-2xl font-bold text-slate-100">Admin Dashboard</h1>
        <p class="text-slate-400 text-sm">Selamat datang, Admin! ({user.email})</p>
      </div>
      <div class="flex items-center gap-3">
        <Badge variant="success">Role: {user.role}</Badge>
        <Button variant="destructive" size="sm" on:click={handleLogout}>
          Logout
        </Button>
      </div>
    </div>

    <Card classNames="border-slate-800 bg-slate-900/60">
      <div class="space-y-4">
        <h2 class="text-lg font-semibold text-slate-200">Uji Akses Otorisasi Admin (RBAC)</h2>
        <p class="text-slate-400 text-sm">
          Tekan tombol di bawah untuk menguji endpoint terproteksi <code>/api/admin/test</code>.
        </p>

        <Button variant="default" size="sm" on:click={testAdminAccess}>
          Uji Endpoint /api/admin/test
        </Button>

        {#if adminTestResult}
          <div class="p-3 bg-emerald-500/10 border border-emerald-500/30 rounded-lg text-emerald-400 text-sm">
            ✅ {adminTestResult}
          </div>
        {/if}

        {#if adminTestError}
          <div class="p-3 bg-rose-500/10 border border-rose-500/30 rounded-lg text-rose-400 text-sm">
            ❌ {adminTestError}
          </div>
        {/if}
      </div>
    </Card>
  </div>
</div>
