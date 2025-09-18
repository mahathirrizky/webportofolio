<!-- src/views/LandingPage.vue -->
<template>
    <div class="max-w-[2400px] relative">
    <Header />
        <slot /> 
    </div>
  </template>
  
  <script setup>

  import Header from   '../components/Header.vue';
  </script>
  
  <style scoped>
 
  </style>
  