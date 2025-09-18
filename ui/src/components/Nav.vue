<template>
    <nav class="flex gap-8">
      <router-link
        v-for="(item, index) in navItems"
        :key="index"
        :to="item.path"
        class="capitalize font-medium hover:text-accent transition-all"
        active-class="text-accent border-b-2 border-accent"  
      >
        {{ item.name }}
      </router-link>
    </nav>
  </template>
  
  <script>
  export default {
    data() {
      return {
        navItems: [
          { name: 'Home', path: '/' },
          { name: 'Service', path: '/service' },
          { name: 'Resume', path: '/resume' },
          { name: 'Work', path: '/work' },
          { name: 'Contact', path: '/contact' }
        ]
      };
    }
  };
  </script>