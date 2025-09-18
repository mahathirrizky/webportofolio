<template>
    <div class="xl:flex-row xl:flex xl:gap-1">
      <div class="flex flex-col w-full xl:mx-0 gap-5 items-center">
        <div
          v-for="(tab, index) in tabs"
          :key="index"
          class="p-1 mr-5 cursor-pointer"
          :class="{
            'inline-flex  h-auto w-full max-w-[350px] xl:max-w-[450px] items-center justify-center rounded-md py-2 text-primary bg-accent': isActiveTab(index),
            'inline-flex h-auto w-full max-w-[350px] xl:max-w-[450px] items-center justify-center rounded-md py-2 bg-[#27272c] text-white': !isActiveTab(index)
          }"
          @click="selectTab(index)"
        >
          {{ tab.name }}
        </div>
      </div>
      <div class="mt-4 min-h-[70vh] w-full xl:mt-0 xl:ml-3">
        <slot :activeTab="activeTab"></slot>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const props = defineProps({
    tabs: {
      type: Array,
      required: true,
    },
  });
  
  const activeTab = ref(0);
  
  const selectTab = (index) => {
    activeTab.value = index;
  };
  
  const isActiveTab = (index) => {
    return activeTab.value === index;
  };
  </script>
  
  <style scoped>

  </style>
  