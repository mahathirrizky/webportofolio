<template>
  <LandingPage>
    <div v-if="!isLoading" class="min-h-[80vh] flex flex-col justify-center py-12 xl:py-0">
      <div class="mx-[120px]">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-[60px]">
          <div v-for="(service, index) in services" :key="index" >
            <div class="flex-1 flex flex-col justify-center gap-6 group">
              <div class="w-full flex justify-between items-center">

                <div class="text-5xl font-extrabold text-outline text-transparent group-hover:text-outline-hover transition-all duration-500">{{ service.num }} </div>
                <router-link to= {{service.href}} class="h-[70px] w-[70px] rounded-full bg-white group-hover:bg-accent transition-all duration-500 flex justify-center items-center hover:-rotate-[60deg]">
                  <component 
                  :is="service.icon" 
                  class=" h-6 w-6 rotate-[330deg] text-3xl text-primary" 
                  /> 
                </router-link>
              </div>
              <h2 class="text-[42px] font-bold leading-none text-white group-hover:text-accent transition-all duration-500">{{ service.title }}</h2>
              <p class="text-white/60">{{ service.description }}</p>
              <div class="border-b border-white/20 w-full"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </LandingPage>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getPublicContent } from '../services/api';
import { ArrowDownIcon } from '@heroicons/vue/outline';
import LandingPage from '../views/LandingPage.vue';

const services = ref([]);
const isLoading = ref(true);

const processContent = (contents) => {
  const serviceContent = contents.find(c => c.page_name === 'services' && c.section === 'list' && c.key === 'items');
  if (serviceContent && serviceContent.value) {
    try {
      services.value = JSON.parse(serviceContent.value).map(service => ({ ...service, icon: ArrowDownIcon }));
    } catch (e) {
      console.error("Failed to parse services JSON:", e);
    }
  }
  isLoading.value = false;
};

onMounted(async () => {
  try {
    const response = await getPublicContent();
    processContent(response.data.contents || []);
  } catch (error) {
    console.error('Failed to fetch content:', error);
    // Fallback data
    services.value = [
      {
        num: '01',
        title: 'Web Development',
        description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos',
        href: "",
        icon: ArrowDownIcon
      },
      {
        num: '02',
        title: 'UI/UX Design',
        description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos',
        href: "",
        icon: ArrowDownIcon
      },
      {
        num: '03',
        title: 'Logo Design',
        description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos',
        href: "",
        icon: ArrowDownIcon
      },
      {
        num: '04',
        title: 'SEO',
        description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos',
        href: "",
        icon: ArrowDownIcon
      },
    ];
    isLoading.value = false;
  }
});
</script>

<style scoped>
/* Optional: Add any additional styles here */
</style>
