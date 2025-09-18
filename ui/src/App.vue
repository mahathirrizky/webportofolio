<template>
  <router-view v-slot="{ Component }">
    <PageTransition v-if="shouldTransition">
      <component :is="Component" />
    </PageTransition>
    <component v-else :is="Component" />
  </router-view>
</template>

<script setup>
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import PageTransition from './components/PageTranstion.vue'; // Ensure the path is correct

const route = useRoute();

// Define which routes should have transitions
const shouldTransition = computed(() => {
  return ['service', 'work', 'home','contact', 'resume'].includes(route.name);
});
</script>

<style scoped></style>
