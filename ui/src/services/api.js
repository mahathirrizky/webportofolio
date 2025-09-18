import axios from 'axios';

const API_BASE_URL = import.meta.env.MODE === 'development'
  ? 'http://localhost:8080/api' // Development backend URL
  : '/api'; // Production relative path

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add the JWT 
// token to headers
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('admin_token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor to handle token expiration or invalid tokens
api.interceptors.response.use(
  (response) => {
    console.log('API Response:', response.data); // Log all successful responses
    return response;
  },
  (error) => {
    if (error.response && error.response.status === 401) {
      // Token expired or invalid, redirect to login
      localStorage.removeItem('admin_token');
      // You might want to use router.push here, but it's outside this module's scope.
      // A simple window.location.href redirect for now.
      window.location.href = '/admin/login';
    }
    return Promise.reject(error);
  }
);

// --- Content Management API ---

// For public pages
export const getPublicContent = () => api.get('/content');

// For admin panel
export const getAdminContent = () => api.get('/admin/content');
export const createContent = (data) => api.post('/admin/content', data);
export const updateContent = (id, data) => api.put(`/admin/content/${id}`, data);
export const deleteContent = (id) => api.delete(`/admin/content/${id}`);

// --- File Upload API ---
export const uploadFile = (file) => {
  const formData = new FormData();
  formData.append('file', file);

  return api.post('/admin/upload', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};

// --- Message API ---
export const sendMessage = (data) => api.post('/messages', data, {
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded'
  }
});
export const getMessages = () => api.get('/admin/messages');
export const deleteMessage = (id) => api.delete(`/admin/messages/${id}`);


export default api;