<template>
  <LandingPage>
    <div v-if="!isLoading" class="min-h-[80vh] pt-10 pb-0 xl:py-0 flex">
      <div class="mx-[120px] w-full">
        <Tabs :tabs="tabs">
          <template #default="{ activeTab }">
            <div v-if="activeTab === 0">
              <div class="flex flex-col gap-[30px] text-center xl:text-left">
                <h3 class="text-4xl font-bold">{{ experience.title }}</h3>
                <p class="max-w-[600px] text-white/60 mx-auto xl:mx-0">{{ experience.description }}</p>
                <div class="overflow-y-auto max-h-[400px] scroller ">
                  <ul class="grid grid-cols-1 lg:grid-cols-2 gap-[30px]">
                    <li v-for="(item, index) in experience.items" :key="index" class="bg-[#232329] h-[184px] py-6 px-10 rounded-xl flex flex-col justify-center items-center lg:items-start gap-1">
                      <span class="text-accent">{{ item.duration }}</span>
                      <h3 class="text-xl max-w-[260px] min-h-[60px] text-center lg:text-left">{{ item.position }}</h3>
                      <div class="flex items-center gap-3">
                        <span class="w-[6px] h-[6px] rounded-full bg-accent"></span>
                        <p class="text-white/60">{{ item.company }}</p>
                      </div>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <div v-if="activeTab === 1">
              <div class="flex flex-col gap-[30px] text-center xl:text-left">
                <h3 class="text-4xl font-bold">{{ formaleducation.title }}</h3>
                <p class="max-w-[600px] text-white/60 mx-auto xl:mx-0">{{ formaleducation.description }}</p>
                <div class="overflow-y-auto max-h-[400px] scroller ">
                  <ul class="grid grid-cols-1 lg:grid-cols-2 gap-[30px]">
                    <li v-for="(item, index) in formaleducation.items" :key="index" class="bg-[#232329] h-[184px] py-6 px-10 rounded-xl flex flex-col justify-center items-center lg:items-start gap-1">
                      <span class="text-accent">{{ item.duration }}</span>
                      <h3 class="text-xl max-w-[260px] min-h-[60px] text-center lg:text-left">{{ item.degree}}</h3>
                      <div class="flex items-center gap-3">
                        <span class="w-[6px] h-[6px] rounded-full bg-accent"></span>
                        <p class="text-white/60">{{ item.institution }}</p>
                      </div>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <div v-if="activeTab === 2">
              <div class="flex flex-col gap-[30px] text-center xl:text-left">
                <h3 class="text-4xl font-bold">{{ informaleducation.title }}</h3>
                <p class="max-w-[600px] text-white/60 mx-auto xl:mx-0">{{ informaleducation.description }}</p>
                <div class="overflow-y-auto max-h-[400px] scroller ">
                  <ul class="grid grid-cols-1 lg:grid-cols-2 gap-[30px]">
                    <li v-for="(item, index) in informaleducation.items" :key="index" class="bg-[#232329] h-[184px] py-6 px-10 rounded-xl flex flex-col justify-center items-center lg:items-start gap-1">
                      <span class="text-accent">{{ item.duration }}</span>
                      <h3 class="text-xl max-w-[260px] min-h-[60px] text-center lg:text-left">{{ item.degree}}</h3>
                      <div class="flex items-center gap-3">
                        <span class="w-[6px] h-[6px] rounded-full bg-accent"></span>
                        <p class="text-white/60">{{ item.institution }}</p>
                      </div>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <div v-if="activeTab === 3 " class="w-full h-full">
              <div class="flex flex-col gap-[30px]">
                <div class="flex flex-col gap-[40px] text-center xl:text-left">

                  <h3 class="text-4xl font-bold">{{ skills.title }}</h3>
                  <p class="max-w-[600px] text-white/60 mx-auto xl:mx-0">{{ skills.description }}</p>
                </div>
                <ul class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4 xl:gap-[60px]">
                  <li v-for="(item, index) in skills.skillList" :key="index" class=" w-full h-[150px] bg-[#232329] rounded-xl flex justify-center items-center">
                    <div class="relative group">
                      <img
                          :src="item.icon"
                          alt=""
                          class="h-[45px] w-[45px]"
                        />
                        <span class="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 whitespace-nowrap bg-white text-black text-sm px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity duration-200 capitalize">
                          {{ item.name }}
                        </span>
                    </div>
                  </li>
                </ul>
              </div>
            </div>
            <div v-if="activeTab === 4">
              <div class="w-full text-center xl:text-left">

                <div class=" flex flex-col gap-[30px]">
                  <h3 class="text-4xl font-bold">{{ about.title }}</h3>
                  <p class="max-w-[600px] text-white/60 mx-auto xl:mx-0">{{ about.description }}</p>
                  <ul class="grid grid-cols-1 xl:grid-cols-2 gap-y-6 max-w-[620px] mx-auto xl:mx-0">
                    <li v-for="(item, index) in about.info" :key="index" class="flex items-center justify-center xl:justify-start gap-4">
                      <span class="text-white/60">{{ item.fieldName }}</span>
                      <span class="text-xl">{{ item.fieldValue }}</span>
                      
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </template>
        </Tabs>
      </div>
    </div>
  </LandingPage>
</template>
  
  <script setup>
import { ref, onMounted } from 'vue';
import { getPublicContent } from '../services/api';
import Tabs from '../components/Tabs.vue';
import LandingPage from '../views/LandingPage.vue';

const experience = ref({ title: '', description: '', items: [] });
const formaleducation = ref({ title: '', description: '', items: [] });
const informaleducation = ref({ title: '', description: '', items: [] });
const skills = ref({ title: '', description: '', skillList: [] });
const about = ref({ title: '', description: '', info: [] });
const isLoading = ref(true);

const processContent = (contents) => {
  const resumeContents = contents.filter(c => c.page_name === 'resume');
  
  const processSection = (sectionName) => {
    const title = resumeContents.find(c => c.section === sectionName && c.key === 'title')?.value || '';
    const description = resumeContents.find(c => c.section === sectionName && c.key === 'description')?.value || '';
    const itemsContent = resumeContents.find(c => c.section === sectionName && c.key === 'items')?.value;
    let items = [];
    if (itemsContent) {
      try {
        items = JSON.parse(itemsContent);
      } catch (e) {
        console.error(`Failed to parse ${sectionName} JSON:`, e);
      }
    }
    return { title, description, items };
  };

  const skillsData = processSection('skills');
  skills.value = { title: skillsData.title, description: skillsData.description, skillList: skillsData.items };

  const aboutData = processSection('about');
  about.value = { title: aboutData.title, description: aboutData.description, info: aboutData.items };

  experience.value = processSection('experience');
  formaleducation.value = processSection('formaleducation');

  // Sort formal education by year
  if (formaleducation.value.items) {
    formaleducation.value.items.sort((a, b) => {
      const yearA = parseInt(a.duration.split('-')[0], 10);
      const yearB = parseInt(b.duration.split('-')[0], 10);
      return yearB - yearA;
    });
  }

  informaleducation.value = processSection('informaleducation');

  // Sort informal education by year
  if (informaleducation.value.items) {
    informaleducation.value.items.sort((a, b) => {
      const yearA = parseInt(a.duration.split('-')[0], 10);
      const yearB = parseInt(b.duration.split('-')[0], 10);
      return yearB - yearA;
    });
  }

  isLoading.value = false;
};

onMounted(async () => {
  try {
    const response = await getPublicContent();
    processContent(response.data.contents || []);
  } catch (error) {
    console.error('Failed to fetch content:', error);
    // Fallback data can be set here if needed
    isLoading.value = false;
  }
});

const tabs = [
  { name: 'Experience' },
  { name: 'Formal Education' },
  { name: 'Informal Education' },
  { name: 'Skills' },
  { name: 'About me' },
];
</script>

<style scoped>
.scroller {
  scrollbar-color: #00ff99 #1E293B;
}
</style>