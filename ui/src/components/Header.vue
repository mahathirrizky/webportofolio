<template>
    <header class="py-8 xl:py-12 text-white"> <!-- Using <header> as the root element -->
      <div class=" mx-[120px] flex justify-between items-center">
        <router-link to="/">
          <h1 class="text-4xl font-semibold">
            {{ brandName }}<span class="text-accent">.</span>
          </h1>
        </router-link>
        <div class="hidden xl:flex items-center gap-8">
          <Nav />
          <router-link to="/contact">
            <button class="btn-hire m-5 relative inline-block w-[6em] h-[2.6em] leading-[2.5em] cursor-pointer overflow-hidden border-2 border-accent transition-colors duration-700 delay-300 z-10 text-accent text-[17px] rounded-[6px] font-medium"> <!-- All utility classes moved here -->
              <span class="relative z-20">Hire Me</span>
            </button>
          </router-link>
        </div> 
        <div class="xl:hidden">
          <MobileNav />
        </div>
      </div>
    </header>
  </template>
  
  <script setup>
import Nav from './Nav.vue';
import MobileNav from './MobileNav.vue';
import { ref, onMounted } from 'vue';
import { getPublicContent } from '../services/api';

const brandName = ref('Person Name'); // Fallback value

onMounted(async () => {
  try {
    const response = await getPublicContent();
    const contents = response.data.contents || [];
    const brandNameEntry = contents.find(c => c.page_name === 'home' && c.section === 'nav' && c.key === 'brand_name');
    if (brandNameEntry) {
      brandName.value = brandNameEntry.value;
    }
  } catch (error) {
    console.error('Failed to fetch brand name:', error);
  }
});
</script>
  
  <style scoped>
  .btn-hire::before {
    content: '';
    position: absolute;
    background-color: var(--color-accent); /* Use CSS variable */
    height: 150px;
    width: 200px;
    border-radius: 9999px; /* rounded-full */
    transition-property: all;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
    transition-duration: 700ms;
    top: 100%;
    left: 100%;
  }
  .btn-hire:hover {
    color: var(--color-primary); /* Use CSS variable */
  }
  .btn-hire:hover::before {
    top: -30px;
    left: -30px;
  }
  </style>