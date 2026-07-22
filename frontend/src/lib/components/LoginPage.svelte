<script>
  import Card from './ui/Card.svelte';
  import Button from './ui/Button.svelte';
  import Badge from './ui/Badge.svelte';
  import { authStore } from '../stores/auth.svelte.js';

  let email = '';
  let password = '';
  let errorMessage = '';
  let successMessage = '';
  let isSubmitting = false;

  async function handleSubmit(event) {
    if (event) event.preventDefault();
    if (!email || !password) {
      errorMessage = 'Email dan password wajib diisi';
      successMessage = '';
      return;
    }

    errorMessage = '';
    successMessage = '';
    isSubmitting = true;

    const result = await authStore.login(email, password);
    isSubmitting = false;

    if (result.success) {
      successMessage = `Login Berhasil! Selamat datang, ${result.user.email}`;
    } else {
      errorMessage = result.error || 'Email atau password salah';
    }
  }
</script>

<div class="min-h-screen flex items-center justify-center p-4 bg-slate-950 text-slate-100 relative overflow-hidden">
  <!-- Top Notification Toasts -->
  {#if successMessage}
    <div class="fixed top-6 right-6 z-50 max-w-md animate-bounce">
      <div class="p-4 rounded-xl bg-emerald-500/20 border border-emerald-500/50 text-emerald-300 shadow-2xl backdrop-blur-md flex items-center gap-3">
        <svg class="w-6 h-6 text-emerald-400 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <div>
          <h4 class="font-bold text-sm text-emerald-200">Notifikasi Sukses</h4>
          <p class="text-xs text-emerald-300/90">{successMessage}</p>
        </div>
      </div>
    </div>
  {/if}

  {#if errorMessage}
    <div class="fixed top-6 right-6 z-50 max-w-md animate-pulse">
      <div class="p-4 rounded-xl bg-rose-500/20 border border-rose-500/50 text-rose-300 shadow-2xl backdrop-blur-md flex items-center gap-3">
        <svg class="w-6 h-6 text-rose-400 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <div>
          <h4 class="font-bold text-sm text-rose-200">Notifikasi Gagal</h4>
          <p class="text-xs text-rose-300/90">{errorMessage}</p>
        </div>
      </div>
    </div>
  {/if}

  <div class="w-full max-w-md z-10">
    <div class="text-center mb-8">
      <div class="inline-block mb-3">
        <Badge variant="outline">Enterprise Auth v1.0</Badge>
      </div>
      <h1 class="text-3xl font-extrabold tracking-tight bg-gradient-to-r from-blue-400 via-indigo-300 to-purple-400 bg-clip-text text-transparent">
        Selamat Datang
      </h1>
      <p class="text-slate-400 text-sm mt-2">
        Silakan masuk ke akun Anda untuk melanjutkan
      </p>
    </div>

    <Card classNames="border-slate-800/80 bg-slate-900/90 shadow-2xl backdrop-blur-xl">
      <form on:submit|preventDefault={handleSubmit} class="space-y-5">
        {#if successMessage}
          <div class="p-3.5 rounded-lg bg-emerald-500/10 border border-emerald-500/30 text-emerald-400 text-sm flex items-center gap-2">
            <svg class="w-5 h-5 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
            <span>{successMessage}</span>
          </div>
        {/if}

        {#if errorMessage}
          <div class="p-3.5 rounded-lg bg-rose-500/10 border border-rose-500/30 text-rose-400 text-sm flex items-center gap-2">
            <svg class="w-5 h-5 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span>{errorMessage}</span>
          </div>
        {/if}

        <div>
          <label for="email" class="block text-xs font-semibold uppercase tracking-wider text-slate-300 mb-1.5">
            Alamat Email
          </label>
          <input
            id="email"
            type="email"
            bind:value={email}
            placeholder="admin@test.com"
            disabled={isSubmitting}
            required
            class="w-full px-4 py-2.5 bg-slate-950 border border-slate-800 rounded-lg text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all disabled:opacity-50"
          />
        </div>

        <div>
          <label for="password" class="block text-xs font-semibold uppercase tracking-wider text-slate-300 mb-1.5">
            Kata Sandi
          </label>
          <input
            id="password"
            type="password"
            bind:value={password}
            placeholder="••••••••"
            disabled={isSubmitting}
            required
            class="w-full px-4 py-2.5 bg-slate-950 border border-slate-800 rounded-lg text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all disabled:opacity-50"
          />
        </div>

        <Button
          type="submit"
          variant="default"
          size="lg"
          disabled={isSubmitting}
          classNames="w-full mt-2 font-semibold shadow-blue-500/20"
        >
          {#if isSubmitting}
            <span class="inline-flex items-center gap-2">
              <svg class="animate-spin h-4 w-4 text-white" viewBox="0 0 24 24" fill="none">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Memproses...
            </span>
          {:else}
            Masuk ke Akun
          {/if}
        </Button>
      </form>
    </Card>
  </div>
</div>
