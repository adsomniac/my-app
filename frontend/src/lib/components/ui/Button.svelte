<script>
  import { cn } from '$lib/utils.js';

  export let variant = 'default'; // 'default' | 'outline' | 'ghost' | 'destructive'
  export let size = 'default'; // 'default' | 'sm' | 'lg' | 'icon'
  export let disabled = false;
  export let classNames = '';

  const baseStyles = 'inline-flex items-center justify-center rounded-lg font-medium transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 disabled:pointer-events-none disabled:opacity-50 active:scale-95 cursor-pointer';

  const variants = {
    default: 'bg-blue-600 text-white hover:bg-blue-500 shadow-lg shadow-blue-600/20',
    outline: 'border border-slate-700 bg-transparent text-slate-200 hover:bg-slate-800 hover:text-white',
    ghost: 'bg-transparent text-slate-300 hover:bg-slate-800 hover:text-white',
    destructive: 'bg-red-600 text-white hover:bg-red-500 shadow-lg shadow-red-600/20',
  };

  const sizes = {
    default: 'h-10 px-4 py-2 text-sm',
    sm: 'h-8 px-3 text-xs rounded-md',
    lg: 'h-12 px-6 text-base rounded-xl',
    icon: 'h-10 w-10 p-2',
  };

  $: computedClass = cn(baseStyles, variants[variant] || variants.default, sizes[size] || sizes.default, classNames);
</script>

<button class={computedClass} {disabled} on:click>
  <slot />
</button>
