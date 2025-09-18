// src/main.js
import { createApp } from 'vue';
import './style.css';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router';
import '@fortawesome/fontawesome-free/css/all.css';
import '@mdi/font/css/materialdesignicons.min.css';

// Public Pages
import ServicePage from './pages/ServicePage.vue';
import HomePage from './pages/HomaPage.vue';
import ContactPage from './pages/ContactPage.vue';
import ResumePage from './pages/ResumePage.vue';
import WorkPage from './pages/WorkPage.vue';

// Admin Pages
import AdminLoginPage from './pages/admin/LoginPage.vue';

import ContactPageadmin from './pages/admin/ContactPageadmin.vue';
import ResumePageadmin from './pages/admin/ResumePageadmin.vue';
import ServicePageadmin from './pages/admin/ServicePageadmin.vue';
import WorkPageadmin from './pages/admin/WorkPageadmin.vue';
import EditHomePage from './pages/admin/EditHomePage.vue'; // New import
import EditServicePage from './pages/admin/EditServicePage.vue'; // New import
import EditResumePage from './pages/admin/EditResumePage.vue'; // New import
import EditWorkPage from './pages/admin/EditWorkPage.vue'; // New import
import EditContactPage from './pages/admin/EditContactPage.vue'; // New import
import MessagesPage from './pages/admin/MessagesPage.vue'; // New import

// Import middleware
import { authMiddleware } from './authMiddleware';

// Routes registration
const routes = [
  { path: '/', component: HomePage, name: 'home' },
  { path: '/service', component: ServicePage, name: 'service' },
  { path: '/contact', component: ContactPage, name: 'contact' },
  { path: '/resume', component: ResumePage, name: 'resume' },
  { path: '/work', component: WorkPage, name: 'work' },

  // Admin Routes
  { path: '/admin/login', component: AdminLoginPage, name: 'admin-login' },
  {
    path: '/admin',
    redirect: to => {
      const adminToken = localStorage.getItem('admin_token');
      if (adminToken) {
        return { name: 'admin-edit-home' };
      } else {
        return { name: 'admin-login' };
      }
    }
  },
  
  {
    path: '/admin/edit/home',
    component: EditHomePage,
    name: 'admin-edit-home',
    meta: { requiresAdminAuth: true }
  },
  {
    path: '/admin/edit/service',
    component: EditServicePage,
    name: 'admin-edit-service',
    meta: { requiresAdminAuth: true }
  },
  {
    path: '/admin/edit/resume',
    component: EditResumePage,
    name: 'admin-edit-resume',
    meta: { requiresAdminAuth: true }
  },
  {
    path: '/admin/edit/work',
    component: EditWorkPage,
    name: 'admin-edit-work',
    meta: { requiresAdminAuth: true }
  },
  {
    path: '/admin/edit/contact',
    component: EditContactPage,
    name: 'admin-edit-contact',
    meta: { requiresAdminAuth: true }
  },
  {
    path: '/admin/messages',
    component: MessagesPage,
    name: 'admin-messages',
    meta: { requiresAdminAuth: true }
  },
  {
    path: '/admin/contact',
    component: ContactPageadmin,
    name: 'admin-contact',
    meta: { requiresAdminAuth: true }
  },
  {
    path: '/admin/resume',
    component: ResumePageadmin,
    name: 'admin-resume',
    meta: { requiresAdminAuth: true }
  },
  {
    path: '/admin/service',
    component: ServicePageadmin,
    name: 'admin-service',
    meta: { requiresAdminAuth: true }
  },
  {
    path: '/admin/work',
    component: WorkPageadmin,
    name: 'admin-work',
    meta: { requiresAdminAuth: true }
  },
];

// Create Router
const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Use middleware
router.beforeEach(authMiddleware);

// Create app
createApp(App) // Root Component
  .use(router) // Use Router
  .mount('#app');