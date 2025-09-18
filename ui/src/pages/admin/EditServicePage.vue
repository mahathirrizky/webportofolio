<template>
  <Admin>
    <div class="p-4 sm:ml-64 py-8 text-white">
      <div class="max-w-7xl mx-auto">
        <h2 class="text-4xl font-bold mb-10">Edit Service Page</h2>

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

          <!-- Service Items Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Service Items</h3>
            <div class="space-y-8">
              <div v-for="(service, index) in services" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">Service #{{ index + 1 }}</h4>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label :for="`service_num_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Number</label>
                    <input type="text" :id="`service_num_${index}`" v-model="service.num" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`service_title_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Title</label>
                    <input type="text" :id="`service_title_${index}`" v-model="service.title" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div class="md:col-span-2">
                    <label :for="`service_description_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Description</label>
                    <textarea :id="`service_description_${index}`" v-model="service.description" rows="3" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent"></textarea>
                  </div>
                  <div class="md:col-span-2">
                    <label :for="`service_href_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Href</label>
                    <input type="text" :id="`service_href_${index}`" v-model="service.href" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeService(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove Service</button>
                </div>
              </div>
            </div>
            <div class="flex justify-between items-center mt-8">
              <button type="button" @click="addService" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New Service</button>
              <button @click="saveServiceContent" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All Services</button>
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

// Data for Service Items Section (as JSON string)
const serviceContentJson = ref('[]');

// Data for Service Items Section (as array of objects)
const services = ref([]);

// Store original content IDs for updates
const contentMap = new Map(); // Map<key, {id, page_name, section, key}>

const fetchContent = async () => {
  isLoading.value = true;
  try {
    const response = await getAdminContent();
    const contents = response.data.contents || [];

    // Process Service Content
    const serviceEntry = contents.find(c => c.page_name === 'services' && c.section === 'list' && c.key === 'items');
    if (serviceEntry) {
      try {
        services.value = JSON.parse(serviceEntry.value);
      } catch (e) {
        console.error("Failed to parse service items JSON:", e);
        services.value = [];
      }
      contentMap.set('list_items', { id: serviceEntry.ID, page_name: serviceEntry.PageName, section: serviceEntry.Section, key: serviceEntry.Key });
    } else {
      services.value = [
        { num: '01', title: 'Web Development', description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos', href: "" },
        { num: '02', title: 'UI/UX Design', description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos', href: "" },
        { num: '03', title: 'Logo Design', description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos', href: "" },
        { num: '04', title: 'SEO', description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos', href: "" },
      ];
    }

  } catch (err) {
    error.value = err.response?.data?.error || err.message;
    console.error('Error fetching service page content:', err);
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

const saveServiceContent = async () => {
  try {
    // No need to JSON.parse here, as services is already an array
    await saveContent('services', 'list', 'items', JSON.stringify(services.value, null, 2));
  } catch (e) {
    error.value = 'Error saving service content.' + e.message;
  }
};

const addService = () => {
  services.value.push({ num: '', title: '', description: '', href: '' });
};

const removeService = (index) => {
  services.value.splice(index, 1);
};

onMounted(fetchContent);
</script>
