import { clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';

/**
 * Merges Tailwind CSS classes with clsx conditionals cleanly.
 */
export function cn(...inputs) {
  return twMerge(clsx(inputs));
}
