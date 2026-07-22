/**
 * API service for authentication and user endpoints
 */

export async function loginAPI(email, password) {
  const response = await fetch('/api/auth/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, password }),
  });

  const data = await response.json();
  return { status: response.status, data };
}

export async function getMeAPI(token) {
  const response = await fetch('/api/auth/me', {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
  });

  const data = await response.json();
  return { status: response.status, data };
}

export async function getProfileAPI(token) {
  const response = await fetch('/api/profile', {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
  });

  const data = await response.json();
  return { status: response.status, data };
}

export async function getAdminTestAPI(token) {
  const response = await fetch('/api/admin/test', {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
  });

  const data = await response.json();
  return { status: response.status, data };
}
