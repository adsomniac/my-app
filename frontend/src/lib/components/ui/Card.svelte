<script>
  import { cn } from '$lib/utils.js';

  export let classNames = '';
</script>

<div class={cn("rounded-xl border border-slate-800 bg-slate-900/80 p-6 text-slate-100 shadow-xl backdrop-blur-md transition-all hover:border-slate-700", classNames)}>
  <slot />
</div>
