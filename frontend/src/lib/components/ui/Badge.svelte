<script>
  import { cn } from '$lib/utils.js';

  export let variant = 'default'; // 'default' | 'success' | 'destructive' | 'warning' | 'outline'
  export let classNames = '';

  const baseStyles = 'inline-flex items-center gap-1.5 rounded-full px-3 py-1 text-xs font-semibold tracking-wide transition-colors';

  const variants = {
    default: 'bg-slate-800 text-slate-200 border border-slate-700',
    success: 'bg-emerald-500/10 text-emerald-400 border border-emerald-500/30',
    destructive: 'bg-rose-500/10 text-rose-400 border border-rose-500/30',
    warning: 'bg-amber-500/10 text-amber-400 border border-amber-500/30',
    outline: 'border border-slate-600 text-slate-300',
  };

  $: computedClass = cn(baseStyles, variants[variant] || variants.default, classNames);
</script>

<span class={computedClass}>
  <slot />
</span>
