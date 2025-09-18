<template>
  <Admin>
    <div class="p-4 sm:ml-64 py-8 text-white">
      <div class="max-w-7xl mx-auto">
        <h2 class="text-4xl font-bold mb-10">Edit Contact Page</h2>

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

          <!-- Form Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg mb-12">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Form Section</h3>
            <form @submit.prevent="saveFormContent" class="space-y-6">
              <div>
                <label for="form_title" class="block text-sm font-medium text-gray-400 mb-2">Title</label>
                <input type="text" id="form_title" v-model="formContent.title" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-900 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
              </div>
              <div>
                <label for="form_description" class="block text-sm font-medium text-gray-400 mb-2">Description</label>
                <textarea id="form_description" v-model="formContent.description" rows="4" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-900 text-white focus:ring-2 focus:ring-accent focus:border-transparent"></textarea>
              </div>
              <div>
                <label for="form_button_text" class="block text-sm font-medium text-gray-400 mb-2">Button Text</label>
                <input type="text" id="form_button_text" v-model="formContent.button_text" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-900 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
              </div>
              <div class="flex justify-end pt-4">
                <button type="submit" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save Form Content</button>
              </div>
            </form>
          </div>

          <!-- Info Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Info Section</h3>
            <div class="space-y-8">
              <div v-for="(item, index) in infoItems" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">Info Item #{{ index + 1 }}</h4>
                <div class="space-y-4">
                  <div>
                    <label :for="`info_icon_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Icon (FontAwesome Name)</label>
                    <input type="text" :id="`info_icon_${index}`" v-model="item.icon" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                    <p class="text-xs text-gray-500 mt-2">e.g., `faPhone`, `faEnvelope`, `faMapMarkerAlt`</p>
                  </div>
                  <div>
                    <label :for="`info_title_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Title</label>
                    <input type="text" :id="`info_title_${index}`" v-model="item.title" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`info_description_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Description</label>
                    <textarea :id="`info_description_${index}`" v-model="item.description" rows="3" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent"></textarea>
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeInfoItem(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove</button>
                </div>
              </div>
            </div>
            <div class="flex justify-between items-center mt-8">
              <button type="button" @click="addInfoItem" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New Info Item</button>
              <button @click="saveInfoContent" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All Info Items</button>
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
import { getAdminContent, updateContent, createContent } from '../../services/api';

const isLoading = ref(true);
const message = ref('');
const error = ref('');

// Data for Form Section
const formContent = ref({
  title: '',
  description: '',
  button_text: ''
});

// Data for Info Section
const infoItems = ref([]);

// Store original content IDs for updates
const contentMap = new Map(); // Map<key, {id, page_name, section, key}>

const fetchContent = async () => {
  isLoading.value = true;
  try {
    const response = await getAdminContent();
    const contents = response.data.contents || [];

    // Process Form Content
    const formDefaults = {
      title: 'Let’s work together',
      description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Eligendi vero quaerat dignissimos autem ab.',
      button_text: 'Send message'
    };
    const currentForm = {};
    contents.filter(c => c.page_name === 'contact' && c.section === 'form').forEach(c => {
      currentForm[c.key] = c.value;
      contentMap.set(`form_${c.key}`, { id: c.ID, page_name: c.PageName, section: c.Section, key: c.Key });
    });
    formContent.value = { ...formDefaults, ...currentForm };

    // Process Info Content
    const infoEntry = contents.find(c => c.page_name === 'contact' && c.section === 'info' && c.key === 'items');
    if (infoEntry) {
      try {
        infoItems.value = JSON.parse(infoEntry.value);
      } catch (e) {
        console.error('Failed to parse info items JSON:', e);
        infoItems.value = [];
      }
      contentMap.set('info_items', { id: infoEntry.ID, page_name: infoEntry.PageName, section: infoEntry.Section, key: infoEntry.Key });
    } else {
      infoItems.value = [
        { icon: 'faPhone', title: 'Phone', description: '+1 234 567 890' },
        { icon: 'faEnvelope', title: 'Email', description: 'example@example.com' },
        { icon: 'faMapMarkerAlt', title: 'Address', description: '1234 Street Name, City Name, Country Name' }
      ];
    }

  } catch (err) {
    error.value = err.response?.data?.error || err.message;
    console.error('Error fetching contact page content:', err);
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
      response = await updateContent(existing.id, { PageName: pageName, Section: section, Key: key, Value: value });
    } else {
      // Create new content
      response = await createContent({ PageName: pageName, Section: section, Key: key, Value: value });
      // Update contentMap with new ID for future updates
      contentMap.set(`${section}_${key}`, { id: response.data.content.ID, page_name: pageName, section: section, key: key });
    }
    message.value = response.data.message || 'Content saved successfully!';
  } catch (err) {
    error.value = err.response?.data?.error || err.message;
    console.error('Error saving content:', err);
  }
};

const saveFormContent = async () => {
  // Save each form field individually
  await saveContent('contact', 'form', 'title', formContent.value.title);
  await saveContent('contact', 'form', 'description', formContent.value.description);
  await saveContent('contact', 'form', 'button_text', formContent.value.button_text);
};

const saveInfoContent = async () => {
  try {
    await saveContent('contact', 'info', 'items', JSON.stringify(infoItems.value, null, 2));
  } catch (e) {
    error.value = 'Error saving info content.' + e.message;
  }
};

const addInfoItem = () => {
  infoItems.value.push({ icon: '', title: '', description: '' });
};

const removeInfoItem = (index) => {
  infoItems.value.splice(index, 1);
};

onMounted(fetchContent);
</script>