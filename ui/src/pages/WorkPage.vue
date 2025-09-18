<template>
  <LandingPage>
    <div class="min-h-[80vh] flex flex-col justify-center py-12 xl:px-0">
      <div class="container mx-auto">
        <div v-if="!isLoading && projects.length > 0" class="flex flex-col xl:flex-row xl:gap-[30px]">
          <div class="w-full xl:w-[50%] xl:h-[460px] flex flex-col xl:justify-between order-2 xl:order-none">
            <div class="flex flex-col gap-[30px] h-[50%]">
              <div class="text-8xl leading-none font-extrabold text-transparent text-outline">
                {{ projects[currentProjectIndex].num }}
              </div>
              <h2 class="text-[42px] font-bold leading-none text-white group-hover:text-accent transition-all duration-500 capitalize">
                {{ projects[currentProjectIndex].category }}
              </h2>
              <p class="text-white/60">{{ projects[currentProjectIndex].description }}</p>
              <ul class="flex gap-4">
                <li v-for="(item, index) in projects[currentProjectIndex].stack" :key="index" class="text-xl text-accent">
                  {{ item.name }}{{ index !== projects[currentProjectIndex].stack.length - 1 ? "," : "" }}
                </li>
              </ul>
              <div class="border border-white/20"></div>
              <div class="flex items-center gap-4">
                <a :href="projects[currentProjectIndex].live" class="relative group">
                  <div class="w-[70px] h-[70px] rounded-full bg-white/5 flex justify-center items-center group mt-3">
                    <component 
                      :is="ArrowDownIcon" 
                      class="h-6 w-6 rotate-[240deg] text-3xl text-white group-hover:text-accent" 
                    /> 
                  </div>
                  <span class="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 whitespace-nowrap bg-white text-black text-sm px-2 py-1 rounded-md opacity-0 group-hover:opacity-100 transition-opacity duration-200 capitalize">Live Project</span>
                </a>
                <a :href="projects[currentProjectIndex].github" class="relative group">
                  <div class="w-[70px] h-[70px] rounded-full bg-white/5 flex justify-center items-center group mt-3">
                    <i class="mdi mdi-github text-3xl text-white group-hover:text-accent"></i>
                  </div>
                  <span class="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 whitespace-nowrap bg-white text-black text-sm px-2 py-1 rounded-md opacity-0 group-hover:opacity-100 transition-opacity duration-200 capitalize">Github repository</span>
                </a>
              </div>
            </div>
          </div>
          <div class="w-full xl:w-[50%]">
            <swiper
              :slides-per-view="1"
              :space-between="30"
              class="xl:h-[520px] mb-12"
              @swiper="onSwiper"
              @slideChange="onSlideChange"
            >
              <swiper-slide v-for="(item, index) in projects" :key="index">
                <div class="h-[460px] relative group flex justify-center items-center bg-pink-50/20">
                  <div class="absolute top-0 bottom-0 w-full h-full bg-black/10"></div>
                  <div class="w-full h-0"> <!-- 16:9 Aspect Ratio -->
                    <img :src="item.image" alt="" class="absolute top-0 left-0 w-full h-full object-cover" />
                  </div>
                </div>
              </swiper-slide>
              <div class="flex gap-2 absolute right-0 bottom-[calc(50%_-_22px)] xl:bottom-0 z-20 w-full justify-between xl:w-max xl:justify-none">
                <button @click="slidePrev" class="bg-accent hover:bg-accent-hover text-primary text-[22px] w-[44px] h-[44px] flex justify-center items-center transition-all font-bold">{{ iconslider.kiri }}</button>
                <button @click="slideNext" class="bg-accent hover:bg-accent-hover text-primary text-[22px] w-[44px] h-[44px] flex justify-center items-center transition-all font-bold">{{ iconslider.kanan }}</button>
              </div>
            </swiper>
          </div>
        </div>
      </div> 
    </div>
  </LandingPage>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getPublicContent } from '../services/api';
import LandingPage from '../views/LandingPage.vue';
import { Swiper, SwiperSlide } from 'swiper/vue';
import 'swiper/css';
import { ArrowDownIcon } from '@heroicons/vue/outline';

const iconslider = { kanan: '>', kiri: '<' };

const projects = ref([]);
const isLoading = ref(true);
const currentProjectIndex = ref(0);
let swiperInstance = ref(null);

const processContent = (contents) => {
  const projectContent = contents.find(c => c.page_name === 'work' && c.section === 'projects' && c.key === 'items');
  if (projectContent && projectContent.value) {
    try {
      projects.value = JSON.parse(projectContent.value);
    } catch (e) {
      console.error("Failed to parse projects JSON:", e);
    }
  }
  if (projects.value.length === 0) {
    currentProjectIndex.value = -1; // or handle empty state
  }
  isLoading.value = false;
};

onMounted(async () => {
  try {
    const response = await getPublicContent();
    processContent(response.data.contents || []);
  } catch (error) {
    console.error('Failed to fetch content:', error);
    isLoading.value = false;
  }
});

const onSwiper = (swiper) => {
  swiperInstance.value = swiper;
};

const onSlideChange = (swiper) => {
  currentProjectIndex.value = swiper.activeIndex;
};

const slidePrev = () => {
  if (swiperInstance.value) {
    swiperInstance.value.slidePrev();
  }
};

const slideNext = () => {
  if (swiperInstance.value) {
    swiperInstance.value.slideNext();
  }
};
</script>

<style scoped>
/* Add your styles here */
</style>
