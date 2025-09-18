<template>
  <LandingPage>
    <div v-if="!isLoading && heroContent" class="h-full">
      <div class="mx-[120px] ">
        <div class="flex flex-col xl:flex-row items-center justify-between xl:pt-8 xl:pb-24">
          <div class="text-center xl:text-left order-2 xl:order-none">
            <span class="text-xl">{{ heroContent.subtitle }}</span>
            <h1 class="h1 mb-6">{{ heroContent.greetings }}<br /><span class="text-accent">{{ heroContent.name }}</span></h1>
            <p class="max-w-[500px] mb-9 text-white/80">
              {{ heroContent.description }}
            </p>
            <div class="flex flex-col xl:flex-row items-center gap-8">
                <a :href="heroContent.cv_download_url" download class="cursor-pointer group relative gap-1.5 px-6 py-3 border border-accent bg-primary text-accent rounded-full hover:bg-accent-hover hover:text-primary transition font-semibold shadow-md uppercase no-underline">
                  <div class="flex items-center gap-2">
                    Download CV
                    <DownloadIcon class="w-5 h-5" />
                  </div>
                </a>
                <div class="mb-8 xl:mb-0"> 
                  <Social :socialIcons="socialIcons" />
                </div>
            </div>
          </div>
          <div class="order-1 xl:order-none mb-8 xl:mb-0">
            <div class="relative w-[298px] h-[298px] xl:w-[498px] xl:h-[498px]">
              <div class="animated-border absolute inset-0 rounded-full"></div>
              <div class="absolute inset-[3px] rounded-full bg-primary flex items-center justify-center overflow-hidden">
                <Photo :imageSrc="photoUrl" />
              </div>
            </div>
          </div>
        </div>
        <Stats :statsItems="statsItems" />
      </div>
    </div>
  </LandingPage>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getPublicContent } from '../services/api';
import LandingPage from '../views/LandingPage.vue';
import Social from '../components/Social.vue';
import Photo from '../components/Photo.vue';
import Stats from '../components/Stats.vue';
import { DownloadIcon } from '@heroicons/vue/outline';

const heroContent = ref(null);
const statsItems = ref([]);
const photoUrl = ref('/static/pp.png'); // Default photo URL
const socialIcons = ref([]);
const isLoading = ref(true);

const processContent = (contents) => {
  const defaults = {
    subtitle: 'FullStack Developer',
    greetings: 'Hello I\'m',
    name: 'Mahathir Rizky',
    description: 'I excel at crafting elegant digital experiences and I am proficient in various programming languages and technologies.',
    cv_download_url: '#', // Default CV URL
    photo_url: '/static/pp.png' // Default photo URL
  };


  const heroFromDB = {};
  contents.filter(c => c.page_name === 'home' && c.section === 'hero').forEach(c => {
    heroFromDB[c.key] = c.value;
  });

  heroContent.value = { ...defaults, ...heroFromDB };
  photoUrl.value = heroContent.value.photo_url || '/static/pp.png';

  const statsContent = contents.find(c => c.page_name === 'home' && c.section === 'stats' && c.key === 'items');
  if (statsContent && statsContent.value) {
    try {
      statsItems.value = JSON.parse(statsContent.value);
    } catch (e) {
      console.error("Failed to parse stats JSON:", e);
    }
  } else {
    statsItems.value = [
      { num: 12, text: 'Years of experience' },
      { num: 26, text: 'Projects completed' },
      { num: 8, text: 'Technologies mastered' },
      { num: 200, text: 'Code commits' },
    ];
  }

  const socialContent = contents.find(c => c.page_name === 'home' && c.section === 'social' && c.key === 'items');
  if (socialContent && socialContent.value) {
    try {
      socialIcons.value = JSON.parse(socialContent.value);
    } catch (e) {
      console.error("Failed to parse social icons JSON:", e);
    }
  } else {
    socialIcons.value = [
      { icon: 'mdi-github', url: 'https://github.com/yourusername' },
      { icon: 'mdi-linkedin', url: 'https://linkedin.com/in/yourusername' },
      { icon: 'mdi-twitter', url: 'https://twitter.com/yourusername' }
    ];
  }

  isLoading.value = false;
};

onMounted(async () => {
  try {
    const response = await getPublicContent();
    processContent(response.data.contents || []);
  } catch (error) {
    console.error('Failed to fetch content:', error);
    // Optionally, load fallback data
    heroContent.value = {
        subtitle: 'FullStack Developer',
        greetings: 'Hello I\'m',
        name: 'Mahathir Rizky',
        description: 'I excel at crafting elegant digital experiences and I am proficient in various programming languages and technologies.',
        cv_download_url: '#'
    };
    isLoading.value = false;
  }
});
</script>

<style scoped>
.animated-border {
  background: conic-gradient(
    #00ff99 0deg 60deg, transparent 60deg 72deg,
    #00ff99 72deg 132deg, transparent 132deg 144deg,
    #00ff99 144deg 204deg, transparent 204deg 216deg,
    #00ff99 216deg 276deg, transparent 276deg 288deg,
    #00ff99 288deg 348deg, transparent 348deg 360deg
  );
  animation: rotate 40s linear infinite;
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
