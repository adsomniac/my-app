<script>
  import { onMount } from 'svelte';

  let apiStatus = 'checking...';
  let dbStatus = 'checking...';
  let apiMessage = '';
  let isOnline = navigator.onLine;

  async function checkHealth() {
    try {
      const res = await fetch('/api/health');
      if (res.ok) {
        const data = await res.json();
        apiStatus = data.status || 'ok';
        dbStatus = data.database || 'disconnected';
        apiMessage = data.message || '';
      } else {
        apiStatus = 'error';
        dbStatus = 'unknown';
      }
    } catch (err) {
      apiStatus = 'offline / unreachable';
      dbStatus = 'unknown';
    }
  }

  onMount(() => {
    checkHealth();

    window.addEventListener('online', () => (isOnline = true));
    window.addEventListener('offline', () => (isOnline = false));
  });
</script>

<main class="container">
  <header class="header">
    <div class="badge-row">
      <span class="badge {isOnline ? 'badge-online' : 'badge-offline'}">
        {isOnline ? '● Online' : '○ Offline Mode'}
      </span>
      <span class="badge badge-pwa">⚡ PWA Ready</span>
    </div>
    <h1>Monolith App</h1>
    <p class="subtitle">Golang + Svelte.js + PostgreSQL + PWA</p>
  </header>

  <section class="grid">
    <div class="card">
      <div class="card-icon">🚀</div>
      <h3>Backend Status</h3>
      <p class="status-val {apiStatus === 'ok' ? 'text-success' : 'text-warning'}">
        {apiStatus.toUpperCase()}
      </p>
      <span class="card-desc">{apiMessage || 'Golang HTTP Monolith Server'}</span>
    </div>

    <div class="card">
      <div class="card-icon">🐘</div>
      <h3>PostgreSQL</h3>
      <p class="status-val {dbStatus === 'connected' ? 'text-success' : 'text-muted'}">
        {dbStatus.toUpperCase()}
      </p>
      <span class="card-desc">Database connection pool</span>
    </div>

    <div class="card">
      <div class="card-icon">📲</div>
      <h3>PWA Capabilities</h3>
      <p class="status-val text-info">ACTIVE</p>
      <span class="card-desc">Service Worker & Manifest enabled</span>
    </div>
  </section>

  <footer class="footer">
    <button class="btn" on:click={checkHealth}>Refresh Status</button>
  </footer>
</main>

<style>
  :global(body) {
    margin: 0;
    font-family: 'Inter', system-ui, -apple-system, sans-serif;
    background-color: #0f172a;
    color: #f8fafc;
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .container {
    max-width: 800px;
    width: 90%;
    margin: 2rem auto;
    padding: 2.5rem;
    background: rgba(30, 41, 59, 0.7);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 24px;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.4);
  }

  .header {
    text-align: center;
    margin-bottom: 2.5rem;
  }

  .badge-row {
    display: flex;
    justify-content: center;
    gap: 0.75rem;
    margin-bottom: 1rem;
  }

  .badge {
    padding: 0.35rem 0.85rem;
    border-radius: 9999px;
    font-size: 0.8rem;
    font-weight: 600;
  }

  .badge-online {
    background: rgba(34, 197, 94, 0.2);
    color: #4ade80;
    border: 1px solid rgba(34, 197, 94, 0.4);
  }

  .badge-offline {
    background: rgba(239, 68, 68, 0.2);
    color: #f87171;
    border: 1px solid rgba(239, 68, 68, 0.4);
  }

  .badge-pwa {
    background: rgba(168, 85, 247, 0.2);
    color: #c084fc;
    border: 1px solid rgba(168, 85, 247, 0.4);
  }

  h1 {
    font-size: 2.5rem;
    margin: 0.5rem 0;
    background: linear-gradient(135deg, #38bdf8, #818cf8);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .subtitle {
    color: #94a3b8;
    margin: 0;
    font-size: 1.05rem;
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
    gap: 1.5rem;
    margin-bottom: 2rem;
  }

  .card {
    background: rgba(15, 23, 42, 0.6);
    border: 1px solid rgba(255, 255, 255, 0.07);
    padding: 1.5rem;
    border-radius: 16px;
    text-align: center;
    transition: transform 0.2s ease, border-color 0.2s ease;
  }

  .card:hover {
    transform: translateY(-4px);
    border-color: rgba(56, 189, 248, 0.4);
  }

  .card-icon {
    font-size: 2rem;
    margin-bottom: 0.5rem;
  }

  h3 {
    margin: 0.25rem 0;
    font-size: 1.1rem;
    color: #e2e8f0;
  }

  .status-val {
    font-size: 1.25rem;
    font-weight: 700;
    margin: 0.5rem 0;
  }

  .text-success { color: #4ade80; }
  .text-warning { color: #fbbf24; }
  .text-info { color: #38bdf8; }
  .text-muted { color: #64748b; }

  .card-desc {
    font-size: 0.825rem;
    color: #94a3b8;
  }

  .footer {
    text-align: center;
  }

  .btn {
    background: linear-gradient(135deg, #0284c7, #4f46e5);
    color: white;
    border: none;
    padding: 0.85rem 2rem;
    font-size: 1rem;
    font-weight: 600;
    border-radius: 12px;
    cursor: pointer;
    transition: opacity 0.2s ease;
  }

  .btn:hover {
    opacity: 0.9;
  }
</style>
