<template>
  <nav>
    <button @click="toggleNav" class="nav-toggle">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-accent" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
      </svg>
    </button>
    
    <transition name="slide">
      <div v-if="isOpen" class="fixed inset-0 flex z-50">
        <div class="fixed inset-0 bg-primary opacity-50 " @click="closeNav"></div>
        <div class="fixed inset-y-0 right-0 w-[300px] bg-primary p-4 shadow-lg">
          <button @click="closeNav" class="absolute top-4 right-4 text-xl text-accent xl:hidden">✖</button>
          <ul class="mobile-nav">
            <div class="flex flex-col">
              
              <div class="mt-32 mb-40 text-center text 2-xl">
                
                <h1 class="text-4xl font-semibold">
                  Mahathir<span class="text-accent">.</span>
                </h1>
                <router-link 
                v-for="(item, index) in navItems"
                :key="index"
                :to="item.path" 
                @click="closeNav" 
                class="flex items-center py-2 hover:text-accent"  
              active-class="text-accent " 
              >
              <component 
              :is="item.icon" 
              class="mr-2 h-6 w-6" 
              /> 
              <span class="capitalize ">{{ item.name }}</span> <!-- Wrap item.name in a span with text-white -->
            </router-link>
          </div>
        </div>
            
        </ul>
        </div>
      </div>
    </transition>
  </nav>
</template>

<script>
import { HomeIcon, BriefcaseIcon, DocumentIcon, UserIcon, MailIcon } from '@heroicons/vue/outline';

export default {
  data() {
    return {
      navItems: [
        { name: 'Home', path: '/', icon: HomeIcon },
        { name: 'Service', path: '/service', icon: BriefcaseIcon },
        { name: 'Resume', path: '/resume', icon: DocumentIcon },
        { name: 'Work', path: '/work', icon: UserIcon },
        { name: 'Contact', path: '/contact', icon: MailIcon }
      ],
      isOpen: false
    };
  },
  methods: {
    toggleNav() {
      this.isOpen = !this.isOpen;
    },
    closeNav() {
      this.isOpen = false;
    }
  }
};
</script>

<style scoped>
/* Transition styles */
.slide-enter-active, .slide-leave-active {
  transition: transform 0.7s ease, opacity 0.3s ease;
}

.slide-enter {
  transform: translateX(100%); /* Start off-screen to the right */
  opacity: 0; /* Start transparent */
}

.slide-enter-to {
  transform: translateX(0); /* End at the original position */
  opacity: 1; /* End fully visible */
}

.slide-leave {
  transform: translateX(0); /* Keep in place when leaving */
  opacity: 1; /* Fully visible */
}

.slide-leave-to {
  transform: translateX(100%); /* Move off-screen to the right */
  opacity: 0; /* Fade out */
}
</style>
