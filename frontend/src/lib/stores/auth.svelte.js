import { writable, derived } from 'svelte/store';
import { loginAPI, getMeAPI } from '../api/auth.js';

const TOKEN_KEY = 'auth_token';

function createAuthStore() {
  const { subscribe, set, update } = writable({
    user: null,
    token: typeof localStorage !== 'undefined' ? localStorage.getItem(TOKEN_KEY) || null : null,
    isAuthenticated: false,
    loading: true,
    error: null,
  });

  return {
    subscribe,
    login: async (email, password) => {
      update((s) => ({ ...s, loading: true, error: null }));
      try {
        const { status, data } = await loginAPI(email, password);
        if (status === 200 && data.success && data.data?.token) {
          const token = data.data.token;
          const user = data.data.user;
          localStorage.setItem(TOKEN_KEY, token);

          set({
            user,
            token,
            isAuthenticated: true,
            loading: false,
            error: null,
          });
          return { success: true, user };
        } else {
          const message = data.message || 'Login gagal';
          update((s) => ({ ...s, loading: false, error: message }));
          return { success: false, error: message };
        }
      } catch (err) {
        const message = 'Tidak dapat terhubung ke server';
        update((s) => ({ ...s, loading: false, error: message }));
        return { success: false, error: message };
      }
    },
    checkAuth: async () => {
      const token = typeof localStorage !== 'undefined' ? localStorage.getItem(TOKEN_KEY) : null;
      if (!token) {
        set({
          user: null,
          token: null,
          isAuthenticated: false,
          loading: false,
          error: null,
        });
        return false;
      }

      update((s) => ({ ...s, loading: true, error: null }));
      try {
        const { status, data } = await getMeAPI(token);
        if (status === 200 && data.success && data.data) {
          set({
            user: data.data,
            token,
            isAuthenticated: true,
            loading: false,
            error: null,
          });
          return true;
        } else {
          // Token invalid or expired
          localStorage.removeItem(TOKEN_KEY);
          set({
            user: null,
            token: null,
            isAuthenticated: false,
            loading: false,
            error: null,
          });
          return false;
        }
      } catch (err) {
        localStorage.removeItem(TOKEN_KEY);
        set({
          user: null,
          token: null,
          isAuthenticated: false,
          loading: false,
          error: null,
        });
        return false;
      }
    },
    logout: () => {
      if (typeof localStorage !== 'undefined') {
        localStorage.removeItem(TOKEN_KEY);
      }
      set({
        user: null,
        token: null,
        isAuthenticated: false,
        loading: false,
        error: null,
      });
    },
  };
}

export const authStore = createAuthStore();
