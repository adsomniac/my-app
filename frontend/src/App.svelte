<script>
  import { onMount } from 'svelte';
  import Card from '$lib/components/ui/Card.svelte';
  import Button from '$lib/components/ui/Button.svelte';
  import Badge from '$lib/components/ui/Badge.svelte';
  import { 
    Activity, 
    Database, 
    Wifi, 
    WifiOff, 
    RefreshCw, 
    Server, 
    CheckCircle2, 
    XCircle, 
    AlertCircle, 
    Layers 
  } from 'lucide-svelte';

  let apiStatus = 'checking...';
  let dbStatus = 'checking...';
  let apiMessage = '';
  let isOnline = typeof navigator !== 'undefined' ? navigator.onLine : true;
  let loading = false;

  async function checkHealth() {
    loading = true;
    try {
      const res = await fetch('/api/health');
      if (res.ok) {
        const data = await res.json();
        apiStatus = data.status || 'ok';
        dbStatus = data.database || 'disconnected';
        apiMessage = data.message || 'Server active';
      } else {
        apiStatus = 'error';
        dbStatus = 'unknown';
        apiMessage = 'Received error status from server';
      }
    } catch (err) {
      apiStatus = 'offline';
      dbStatus = 'unknown';
      apiMessage = 'Failed to connect to backend server';
    } finally {
      loading = false;
    }
  }

  function handleOnlineStatus() {
    isOnline = navigator.onLine;
  }

  onMount(() => {
    checkHealth();
    window.addEventListener('online', handleOnlineStatus);
    window.addEventListener('offline', handleOnlineStatus);

    return () => {
      window.removeEventListener('online', handleOnlineStatus);
      window.removeEventListener('offline', handleOnlineStatus);
    };
  });
</script>

<main class="min-h-screen bg-slate-950 text-slate-100 flex flex-col items-center justify-between p-6 sm:p-12 relative overflow-hidden">
  <!-- Glow background highlights -->
  <div class="absolute top-1/4 left-1/2 -translate-x-1/2 -translate-y-1/2 w-96 h-96 bg-blue-600/20 rounded-full blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-1/4 left-1/3 w-80 h-80 bg-purple-600/15 rounded-full blur-3xl pointer-events-none"></div>

  <!-- Header -->
  <header class="w-full max-w-4xl flex items-center justify-between z-10 border-b border-slate-800 pb-6">
    <div class="flex items-center gap-3">
      <div class="p-2.5 bg-blue-600/10 border border-blue-500/30 rounded-xl text-blue-400">
        <Layers class="w-6 h-6" />
      </div>
      <div>
        <h1 class="text-xl font-bold text-white tracking-tight">Golang Monolith SPA</h1>
        <p class="text-xs text-slate-400">Modular MVC Backend & Svelte Frontend</p>
      </div>
    </div>

    <div class="flex items-center gap-3">
      {#if isOnline}
        <Badge variant="success">
          <Wifi class="w-3.5 h-3.5" />
          Online
        </Badge>
      {:else}
        <Badge variant="destructive">
          <WifiOff class="w-3.5 h-3.5" />
          Offline
        </Badge>
      {/if}

      <Button variant="outline" size="sm" on:click={checkHealth} disabled={loading}>
        <RefreshCw class={`w-4 h-4 mr-1.5 ${loading ? 'animate-spin' : ''}`} />
        Refresh
      </Button>
    </div>
  </header>

  <!-- Hero Content & Status Grid -->
  <div class="w-full max-w-4xl my-auto py-12 z-10 flex flex-col gap-8">
    <div class="text-center space-y-3">
      <Badge variant="outline" classNames="px-4 py-1.5 text-xs text-blue-400 border-blue-500/30 bg-blue-500/5">
        <Activity class="w-3.5 h-3.5 text-blue-400" />
        System Health Monitor
      </Badge>
      <h2 class="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
        Aplikasi Full-Stack Siap Production
      </h2>
      <p class="text-sm sm:text-base text-slate-400 max-w-xl mx-auto">
        Integrasi sempurna antara Golang Backend (Modular MVC) dan Svelte Frontend (shadcn-svelte + Tailwind CSS).
      </p>
    </div>

    <!-- Status Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <!-- API Server Status -->
      <Card classNames="flex flex-col justify-between">
        <div class="flex items-start justify-between">
          <div class="flex items-center gap-3">
            <div class="p-3 bg-slate-800 rounded-lg text-blue-400 border border-slate-700">
              <Server class="w-6 h-6" />
            </div>
            <div>
              <h3 class="text-base font-semibold text-slate-100">API Server</h3>
              <p class="text-xs text-slate-400">REST API Service Status</p>
            </div>
          </div>
          {#if apiStatus === 'ok'}
            <Badge variant="success">
              <CheckCircle2 class="w-3.5 h-3.5" />
              Active
            </Badge>
          {:else if apiStatus === 'checking...'}
            <Badge variant="warning">
              <AlertCircle class="w-3.5 h-3.5 animate-pulse" />
              Checking...
            </Badge>
          {:else}
            <Badge variant="destructive">
              <XCircle class="w-3.5 h-3.5" />
              Error
            </Badge>
          {/if}
        </div>
        <div class="mt-6 pt-4 border-t border-slate-800/80 flex items-center justify-between text-xs">
          <span class="text-slate-400">Response Message:</span>
          <span class="font-mono text-slate-200">{apiMessage || 'No status'}</span>
        </div>
      </Card>

      <!-- Database Status -->
      <Card classNames="flex flex-col justify-between">
        <div class="flex items-start justify-between">
          <div class="flex items-center gap-3">
            <div class="p-3 bg-slate-800 rounded-lg text-purple-400 border border-slate-700">
              <Database class="w-6 h-6" />
            </div>
            <div>
              <h3 class="text-base font-semibold text-slate-100">Database</h3>
              <p class="text-xs text-slate-400">PostgreSQL Connection</p>
            </div>
          </div>
          {#if dbStatus === 'connected'}
            <Badge variant="success">
              <CheckCircle2 class="w-3.5 h-3.5" />
              Connected
            </Badge>
          {:else if dbStatus === 'checking...'}
            <Badge variant="warning">
              <AlertCircle class="w-3.5 h-3.5 animate-pulse" />
              Checking...
            </Badge>
          {:else}
            <Badge variant="default">
              <AlertCircle class="w-3.5 h-3.5 text-slate-400" />
              {dbStatus}
            </Badge>
          {/if}
        </div>
        <div class="mt-6 pt-4 border-t border-slate-800/80 flex items-center justify-between text-xs">
          <span class="text-slate-400">Database Status:</span>
          <span class="font-mono text-slate-200 capitalize">{dbStatus}</span>
        </div>
      </Card>
    </div>
  </div>

  <!-- Footer -->
  <footer class="w-full max-w-4xl border-t border-slate-800/80 pt-6 text-center text-xs text-slate-500 z-10">
    <p>Golang Monolith + Svelte (shadcn-svelte UI) &bull; Refactored Clean Architecture</p>
  </footer>
</main>
