// src/router/authMiddleware.js
export const authMiddleware = (to, from, next) => {
  const adminToken = localStorage.getItem('admin_token');

  // If trying to access admin login page and already logged in, redirect to dashboard
  if (to.name === 'admin-login' && adminToken) {
    next({ name: 'admin-dashboard' });
    return;
  }

  // If accessing any admin route (including /admin/dashboard)
  // and requires authentication (which admin-dashboard does via meta field)
  if (to.matched.some(record => record.meta.requiresAdminAuth)) {
    if (!adminToken) {
      next({ name: 'admin-login' });
    } else {
      next(); // Proceed to the intended admin route (e.g., /admin/dashboard)
    }
  } else {
    // For non-admin routes, just proceed
    next();
  }
};