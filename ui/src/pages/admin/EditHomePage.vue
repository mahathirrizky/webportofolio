<template>
  <Admin>
    <div class="p-4 sm:ml-64 py-8 text-white">
      <div class="max-w-7xl mx-auto">
        <h2 class="text-4xl font-bold mb-10">Edit Home Page</h2>

        <div v-if="isLoading" class="text-center text-gray-400">
          <p>Loading content...</p>
        </div>

        <div v-else>
          <!-- Success/Error Messages -->
          <div v-if="message" class="bg-green-500/20 text-green-300 p-4 rounded-lg mb-6">
            <p>{{ message }}</p>
          </div>
          <div v-if="error" class="bg-red-500/20 text-red-300 p-4 rounded-lg mb-6">
            <p>{{ error }}</p>
          </div>

          <!-- Nav Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg mb-12">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Nav Section</h3>
            <form @submit.prevent="saveNavContent" class="space-y-6">
              <div>
                <label for="brand_name" class="block text-sm font-medium text-gray-400 mb-2">Brand Name</label>
                <input type="text" id="brand_name" v-model="navContent.brand_name" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-900 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
              </div>
              <div class="flex justify-end pt-4">
                <button type="submit" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save Nav Content</button>
              </div>
            </form>
          </div>

          <!-- Hero Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg mb-12">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Hero Section</h3>
            <form @submit.prevent="saveHeroContent" class="space-y-6">
              <div>
                <label for="subtitle" class="block text-sm font-medium text-gray-400 mb-2">Subtitle</label>
                <input type="text" id="subtitle" v-model="heroContent.subtitle" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-900 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
              </div>
              <div>
                <label for="greetings" class="block text-sm font-medium text-gray-400 mb-2">Greetings</label>
                <input type="text" id="greetings" v-model="heroContent.greetings" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-900 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
              </div>
              <div>
                <label for="name" class="block text-sm font-medium text-gray-400 mb-2">Name</label>
                <input type="text" id="name" v-model="heroContent.name" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-900 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
              </div>
              <div>
                <label for="description" class="block text-sm font-medium text-gray-400 mb-2">Description</label>
                <textarea id="description" v-model="heroContent.description" rows="4" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-900 text-white focus:ring-2 focus:ring-accent focus:border-transparent"></textarea>
              </div>
              <div>
                <label for="cv_download_url" class="block text-sm font-medium text-gray-400 mb-2">CV Download Link</label>
                <div class="flex items-center gap-4">
                  <input type="text" id="cv_download_url" v-model="heroContent.cv_download_url" readonly class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-gray-400 focus:ring-2 focus:ring-accent focus:border-transparent" />
                  <button type="button" @click="triggerCvFileUpload" class="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300 whitespace-nowrap">Upload PDF</button>
                  <button type="button" v-if="heroContent.cv_download_url !== '/resume.pdf'" @click="deleteCvAndRevertToDefault" class="bg-red-600 hover:bg-red-700 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300 whitespace-nowrap">Delete PDF</button>
                </div>
                <input type="file" ref="cvFileInput" @change="handleCvFileChange" class="hidden" accept="application/pdf" />
              </div>
              <div>
                <label for="photo_url" class="block text-sm font-medium text-gray-400 mb-2">Photo</label>
                <div class="flex items-center gap-4">
                  <input type="text" id="photo_url" v-model="heroContent.photo_url" readonly class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-gray-400 focus:ring-2 focus:ring-accent focus:border-transparent" />
                  <button type="button" @click="triggerFileUpload" class="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300 whitespace-nowrap">Upload New</button>
                  <button type="button" v-if="heroContent.photo_url !== '/static/pp.png'" @click="deletePhotoAndRevertToDefault" class="bg-red-600 hover:bg-red-700 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300 whitespace-nowrap">Delete Photo</button>
                </div>
                <input type="file" ref="fileInput" @change="handleFileChange" class="hidden" accept="image/*" />
                <img v-if="heroContent.photo_url" :src="heroContent.photo_url" alt="Profile Photo Preview" class="mt-4 w-32 h-32 object-cover rounded-full border-4 border-gray-700" />
              </div>
              <div class="flex justify-end pt-4">
                <button type="submit" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save Hero Content</button>
              </div>
            </form>
          </div>

          <!-- Stats Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg mb-12">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Stats Section</h3>
            <form @submit.prevent="saveStatsContent" class="space-y-6">
              <div v-for="(item, index) in statsItems" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">Stat Item #{{ index + 1 }}</h4>
                <div class="space-y-4">
                  <div>
                    <label :for="`stat_num_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Number</label>
                    <input type="number" :id="`stat_num_${index}`" v-model="item.num" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`stat_text_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Text</label>
                    <input type="text" :id="`stat_text_${index}`" v-model="item.text" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeStat(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove</button>
                </div>
              </div>
              <div class="flex justify-between items-center mt-8">
                <button type="button" @click="addStat" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New Stat</button>
                <button type="submit" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All Stats</button>
              </div>
            </form>
          </div>

          <!-- Social Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Social Section</h3>
            <div class="space-y-8">
              <div v-for="(item, index) in socialItems" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">Social Item #{{ index + 1 }}</h4>
                <div class="space-y-4">
                  <div>
                    <label :for="`social_icon_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Icon (MDI Class)</label>
                    <input type="text" :id="`social_icon_${index}`" v-model="item.icon" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                     <p class="text-xs text-gray-500 mt-2">e.g., `mdi-github`, `mdi-linkedin`, `mdi-twitter`</p>
                    <div v-if="item.icon" class="mt-2 flex items-center gap-2">
                      <span :class="['mdi', item.icon]" class="text-accent text-2xl"></span>
                      <span class="text-gray-400 text-sm">Preview</span>
                    </div>
                  </div>
                  <div>
                    <label :for="`social_url_${index}`" class="block text-sm font-medium text-gray-400 mb-2">URL</label>
                    <input type="text" :id="`social_url_${index}`" v-model="item.url" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeSocialItem(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove</button>
                </div>
              </div>
            </div>
            <div class="flex justify-between items-center mt-8">
              <button type="button" @click="addSocialItem" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New Social Item</button>
              <button @click="saveSocialContent" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All Social Items</button>
            </div>
          </div>

        </div>
      </div>
    </div>
  </Admin>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import Admin from '../../views/admin/Admin.vue';
import { getAdminContent, updateContent, createContent, uploadFile } from '../../services/api';

const isLoading = ref(true);
const message = ref('');
const error = ref('');
const fileInput = ref(null);
const cvFileInput = ref(null);

// Data for Hero Section
const heroContent = ref({
  subtitle: '',
  greetings: '',
  name: '',
  brand_name: '',
  description: '',
  cv_download_url: '',
  photo_url: ''
});

// Data for Nav Section
const navContent = ref({
  brand_name: ''
});

// Data for Stats Section
const statsItems = ref([]);

// Data for Social Section
const socialItems = ref([]);

// Store original content IDs for updates
const contentMap = new Map(); // Map<key, {id, page_name, section, key}>

const triggerFileUpload = () => {
  fileInput.value.click();
};

const triggerCvFileUpload = () => {
  cvFileInput.value.click();
};

const handleFileChange = async (event) => {
  const file = event.target.files[0];
  if (!file) return;

  message.value = 'Uploading...';
  error.value = '';

  try {
    const response = await uploadFile(file);
    if (response.data && response.data.file_path) {
      heroContent.value.photo_url = response.data.file_path;
      message.value = 'File uploaded successfully! Remember to click "Save Hero Content" to apply changes.';
    } else {
      throw new Error('Invalid response from server after upload.');
    }
  } catch (err) {
    error.value = err.response?.data?.error || err.message || 'An error occurred during file upload.';
    message.value = '';
    console.error('File upload error:', err);
  }
};

const deletePhotoAndRevertToDefault = async () => {
  if (!confirm('Are you sure you want to delete this photo and revert to the default?')) {
    return;
  }

  message.value = 'Reverting to default photo...';
  error.value = '';

  try {
    heroContent.value.photo_url = '/static/pp.png';
    await saveHeroContent();
    message.value = 'Photo reverted to default successfully!';
  } catch (err) {
    error.value = err.response?.data?.error || err.message || 'An error occurred while reverting photo.';
    console.error('Error reverting photo:', err);
  }
};

const handleCvFileChange = async (event) => {
  const file = event.target.files[0];
  if (!file) return;

  message.value = 'Uploading CV...';
  error.value = '';

  try {
    const response = await uploadFile(file);
    if (response.data && response.data.file_path) {
      heroContent.value.cv_download_url = response.data.file_path;
      message.value = 'CV uploaded successfully! Remember to click "Save Hero Content" to apply changes.';
    } else {
      throw new Error('Invalid response from server after CV upload.');
    }
  } catch (err) {
    error.value = err.response?.data?.error || err.message || 'An error occurred during CV upload.';
    message.value = '';
    console.error('CV upload error:', err);
  }
};

const deleteCvAndRevertToDefault = async () => {
  if (!confirm('Are you sure you want to delete this CV and revert to the default?')) {
    return;
  }

  message.value = 'Reverting to default CV...';
  error.value = '';

  try {
    heroContent.value.cv_download_url = '/resume.pdf';
    await saveHeroContent();
    message.value = 'CV reverted to default successfully!';
  } catch (err) {
    error.value = err.response?.data?.error || err.message || 'An error occurred while reverting CV.';
    console.error('Error reverting CV:', err);
  }
};

const fetchContent = async () => {
  isLoading.value = true;
  try {
    const response = await getAdminContent();
    const contents = response.data.contents || [];

    // Process Hero Content
    const heroDefaults = {
      subtitle: 'FullStack Developer',
      greetings: 'Hello I am',
      name: 'Mahathir Rizky',
      brand_name: 'Person Name',
      description: 'I excel at crafting elegant digital experiences and I am proficient in various programming languages and technologies.',
      cv_download_url: '/static/resume.pdf',
      photo_url: '/static/static/pp.png'
    };
    const currentHero = {};
    contents.filter(c => c.page_name === 'home' && c.section === 'hero').forEach(c => {
      currentHero[c.key] = c.value;
      contentMap.set(`hero_${c.key}`, { id: c.ID, page_name: c.PageName, section: c.Section, key: c.Key });
    });
    heroContent.value = { ...heroDefaults, ...currentHero };

    // Process Nav Content
    const navDefaults = {
      brand_name: 'Your Brand'
    };
    const navEntry = contents.find(c => c.page_name === 'home' && c.section === 'nav' && c.key === 'brand_name');
    if (navEntry) {
      navContent.value.brand_name = navEntry.value;
      contentMap.set(`nav_brand_name`, { id: navEntry.ID, page_name: navEntry.PageName, section: navEntry.Section, key: navEntry.Key });
    } else {
      navContent.value.brand_name = navDefaults.brand_name;
    }

    // Process Stats Content
    const statsEntry = contents.find(c => c.page_name === 'home' && c.section === 'stats' && c.key === 'items');
    if (statsEntry) {
      try {
        statsItems.value = JSON.parse(statsEntry.value);
      } catch (e) {
        console.error("Failed to parse stats items JSON:", e);
        statsItems.value = [];
      }
      contentMap.set('stats_items', { id: statsEntry.ID, page_name: statsEntry.PageName, section: statsEntry.Section, key: statsEntry.Key });
    } else {
      statsItems.value = [
        { num: 12, text: 'Years of experience' },
        { num: 26, text: 'Projects completed' },
        { num: 8, text: 'Technologies mastered' },
        { num: 200, text: 'Code commits' },
      ];
    }

    // Process Social Content
    const socialEntry = contents.find(c => c.page_name === 'home' && c.section === 'social' && c.key === 'items');
    if (socialEntry) {
      try {
        socialItems.value = JSON.parse(socialEntry.value);
      } catch (e) {
        console.error("Failed to parse social items JSON:", e);
        socialItems.value = [];
      }
      contentMap.set('social_items', { id: socialEntry.ID, page_name: socialEntry.PageName, section: socialEntry.Section, key: socialEntry.Key });
    } else {
      socialItems.value = [
        { icon: 'mdi-github', url: 'https://github.com/yourusername' },
        { icon: 'mdi-linkedin', url: 'https://linkedin.com/in/yourusername' },
        { icon: 'mdi-twitter', url: 'https://twitter.com/yourusername' }
      ];
    }

  } catch (err) {
    error.value = err.response?.data?.error || err.message;
    console.error('Error fetching home page content:', err);
  } finally {
    isLoading.value = false;
  }
};

const saveContent = async (pageName, section, key, value) => {
  message.value = '';
  error.value = '';
  try {
    const existing = contentMap.get(`${section}_${key}`);
    let response;
    if (existing && existing.id) {
      // Update existing content
      response = await updateContent(existing.id, { page_name: pageName, section: section, key: key, value: value });
    } else {
      // Create new content
      response = await createContent({ page_name: pageName, section: section, key: key, value: value });
      // Update contentMap with new ID for future updates
      contentMap.set(`${section}_${key}`, { id: response.data.content.ID, page_name: pageName, section: section, key: key });
    }
    message.value = response.data.message || 'Content saved successfully!';
  } catch (err) {
    error.value = err.response?.data?.error || err.message;
    console.error('Error saving content:', err);
  }
};

const saveHeroContent = async () => {
  // Save each hero field individually
  await saveContent('home', 'hero', 'subtitle', heroContent.value.subtitle);
  await saveContent('home', 'hero', 'greetings', heroContent.value.greetings);
  await saveContent('home', 'hero', 'name', heroContent.value.name);
  await saveContent('home', 'hero', 'brand_name', heroContent.value.brand_name);
  await saveContent('home', 'hero', 'description', heroContent.value.description);
  await saveContent('home', 'hero', 'cv_download_url', heroContent.value.cv_download_url);
  await saveContent('home', 'hero', 'photo_url', heroContent.value.photo_url);
};

const saveNavContent = async () => {
  await saveContent('home', 'nav', 'brand_name', navContent.value.brand_name);
};

const saveStatsContent = async () => {
  try {
    await saveContent('home', 'stats', 'items', JSON.stringify(statsItems.value, null, 2));
  } catch (e) {
    error.value = 'Error saving stats content.' + e.message;
  }
};

const saveSocialContent = async () => {
  try {
    await saveContent('home', 'social', 'items', JSON.stringify(socialItems.value, null, 2));
  } catch (e) {
    error.value = 'Error saving social content.' + e.message;
  }
};

const addStat = () => {
  statsItems.value.push({ num: 0, text: '' });
};

const removeStat = (index) => {
  statsItems.value.splice(index, 1);
};

const addSocialItem = () => {
  socialItems.value.push({ icon: '', url: '' });
};

const removeSocialItem = (index) => {
  socialItems.value.splice(index, 1);
};

onMounted(fetchContent);
</script>
