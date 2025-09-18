<template>
    <Admin>
      <div class="p-4 sm:ml-64">
        Coantct
      </div>
    </Admin>
  </template>
  
  <script setup>
  import Admin from '../../views/admin/Admin.vue';

  import { useRouter } from 'vue-router';
  
  const router = useRouter();
  
  // Ambil nama pengguna dari token (opsional)
  
  
  const signOut = () => {
    // Hapus token dari localStorage
    localStorage.removeItem('token');
  
    // Hapus nama pengguna (opsional)

  
    // Redirect ke halaman sign-in
    router.push({ name: 'signin' });
  };
  </script>
  
  
  