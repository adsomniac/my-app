<script>
  import Card from './ui/Card.svelte';
  import Button from './ui/Button.svelte';
  import Badge from './ui/Badge.svelte';
  import { authStore } from '../stores/auth.svelte.js';

  export let user = {};

  function handleLogout() {
    authStore.logout();
  }
</script>

<div class="min-h-screen bg-slate-950 text-slate-100 p-6">
  <div class="max-w-4xl mx-auto space-y-6">
    <div class="flex items-center justify-between border-b border-slate-800 pb-4">
      <div>
        <h1 class="text-2xl font-bold text-slate-100">User Dashboard</h1>
        <p class="text-slate-400 text-sm">Selamat datang kembali, {user.email}!</p>
      </div>
      <div class="flex items-center gap-3">
        <Badge variant="outline">Role: {user.role}</Badge>
        <Button variant="destructive" size="sm" on:click={handleLogout}>
          Logout
        </Button>
      </div>
    </div>

    <Card classNames="border-slate-800 bg-slate-900/60">
      <div class="space-y-4">
        <h2 class="text-lg font-semibold text-slate-200">Informasi Pengguna</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
          <div class="p-3 bg-slate-950 rounded-lg border border-slate-800/80">
            <span class="text-slate-400 block text-xs">ID Pengguna</span>
            <span class="font-mono text-slate-200">{user.id}</span>
          </div>
          <div class="p-3 bg-slate-950 rounded-lg border border-slate-800/80">
            <span class="text-slate-400 block text-xs">Alamat Email</span>
            <span class="text-slate-200">{user.email}</span>
          </div>
        </div>
      </div>
    </Card>
  </div>
</div>
