<template>
    <div class="flex gap-6">
      <a
        v-for="(social, index) in socialIcons"
        :key="index"
        :href="social.url"
        target="_blank"
        rel="noopener noreferrer"
        class="text-accent text-base hover:text-primary transition"
      >
        <i :class="['mdi', social.icon, 'w-9', 'h-9', 'border', 'border-accent', 'rounded-full', 'flex', 'justify-center', 'items-center', 'hover:bg-accent', 'hover:transition-all', 'duration-500']"></i>
      </a>
    </div>
  </template>
  
  <script setup>
  const props = defineProps({
    socialIcons: {
      type: Array,
      required: true,
    },
  });
  </script>
  
  <style scoped>
  /* Optional: Add any additional styles here */
  </style>
  