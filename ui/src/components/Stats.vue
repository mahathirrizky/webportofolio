<template>
    <section>
      <div class="mx-[120px]">
        <div class="flex flex-wrap gap-6 max-w-[80vw] mx-auto xl:max-w-none">
          <div v-for="(item, index) in statsItems" :key="index" class="flex-1 flex gap-4 items-center justify-center xl:justify-start mb-4">
            <AnimatedCounter :value="item.num" :duration="5000" class="counter text-4xl xl:text-6xl font-extrabold" />
            <p :class="item.text.length < 15 ? 'max-w-[100px]' : 'max-w-[150px]'" class="leading-snug text-white/80">{{ item.text }}</p>
          </div>
        </div>
      </div>
    </section>
  </template>
  
  <script>
  import AnimatedCounter from "vue-animated-counter";
  
  export default {
    components: {
      AnimatedCounter,
    },
    props: {
      statsItems: {
        type: Array,
        required: true,
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add any additional styles here if needed */
  </style>
  