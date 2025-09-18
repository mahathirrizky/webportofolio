<template>
    <transition
      name="fade"
      @before-enter="beforeEnter"
      @enter="enter"
      @leave="leave"
    >
      <slot />
    </transition>
  </template>
  
  <script setup>
  const beforeEnter = (el) => {
    el.style.opacity = 0; // Start with opacity 0
  };
  
  const enter = (el, done) => {
    el.offsetHeight; // Trigger reflow
    el.style.transition = 'opacity 0.2s ease-in-out';
    
    // Add a delay before starting the fade-in
    setTimeout(() => {
      el.style.opacity = 1; // Fade in
      done(); // Call done when the transition is complete
    }, 1000); // 1 second delay
  };
  
  const leave = (el, done) => {
    el.style.transition = 'opacity 0.2s ease-in-out';
    
    // Add a delay before starting the fade-out
    setTimeout(() => {
      el.style.opacity = 0; // Fade out
      done(); // Call done when the transition is complete
    });
  };
  </script>
  
  <style scoped>
  .fade-enter-active, .fade-leave-active {
    transition: opacity 0.2s ease-in-out;
  }
  .fade-enter, .fade-leave-to {
    opacity: 0;
  }
  </style>
  