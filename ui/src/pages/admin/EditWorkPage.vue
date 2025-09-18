<template>
  <Admin>
    <div class="p-4 sm:ml-64 py-8 text-white">
      <div class="max-w-7xl mx-auto">
        <h2 class="text-4xl font-bold mb-10">Edit Work Page</h2>

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

          <!-- Projects Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Projects</h3>
            <div class="space-y-8">
              <div v-for="(project, index) in projectsItems" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">Project #{{ index + 1 }}</h4>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label :for="`project_num_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Number</label>
                    <input type="text" :id="`project_num_${index}`" v-model="project.num" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`project_category_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Category</label>
                    <input type="text" :id="`project_category_${index}`" v-model="project.category" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div class="md:col-span-2">
                    <label :for="`project_title_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Title</label>
                    <input type="text" :id="`project_title_${index}`" v-model="project.title" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div class="md:col-span-2">
                    <label :for="`project_description_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Description</label>
                    <textarea :id="`project_description_${index}`" v-model="project.description" rows="3" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent"></textarea>
                  </div>
                  <div class="md:col-span-2">
                    <label :for="`project_stack_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Stack (comma-separated)</label>
                    <input type="text" :id="`project_stack_${index}`" v-model="project.stack" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                    <p class="text-xs text-gray-500 mt-2">e.g., `html 5, css, javascript`</p>
                  </div>
                  <div class="md:col-span-2">
                    <label :for="`project_image_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Image</label>
                    <input type="file" @change="handleFileUpload($event, index)" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                    <img v-if="project.image" :src="project.image" alt="Project Image Preview" class="mt-4 w-32 h-auto object-cover rounded-lg border-4 border-gray-700" />
                  </div>
                  <div class="md:col-span-2">
                    <label :for="`project_live_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Live URL</label>
                    <input type="text" :id="`project_live_${index}`" v-model="project.live" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div class="md:col-span-2">
                    <label :for="`project_github_${index}`" class="block text-sm font-medium text-gray-400 mb-2">GitHub URL</label>
                    <input type="text" :id="`project_github_${index}`" v-model="project.github" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeProject(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove Project</button>
                </div>
              </div>
            </div>
            <div class="flex justify-between items-center mt-8">
              <button type="button" @click="addProject" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New Project</button>
              <button @click="saveProjectsContent" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All Projects</button>
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

// Data for Projects Section (as array of objects)
const projectsItems = ref([]);

// Store original content IDs for updates
const contentMap = new Map(); // Map<key, {id, page_name, section, key}>

const handleFileUpload = async (event, index) => {
  const file = event.target.files[0];
  if (!file) return;

  try {
    const response = await uploadFile(file);
    projectsItems.value[index].image = response.data.file_path;
    message.value = "Image uploaded successfully!";
  } catch (err) {
    error.value = err.response?.data?.error || err.message;
    console.error('Error uploading file:', err);
  }
};

const fetchContent = async () => {
  isLoading.value = true;
  try {
    const response = await getAdminContent();
    const contents = response.data.contents || [];

    // Process Projects Content
    const projectsEntry = contents.find(c => c.page_name === 'work' && c.section === 'projects' && c.key === 'items');
    if (projectsEntry) {
      try {
        const items = JSON.parse(projectsEntry.value);
        projectsItems.value = items.map(item => {
          let stack = [];
          try {
            stack = typeof item.stack === 'string' ? JSON.parse(item.stack) : item.stack;
          } catch (e) {
            console.error('Failed to parse stack JSON:', e);
          }
          item.stack = stack.map(s => s.name).join(', ');
          return item;
        });
      } catch (e) {
        console.error("Failed to parse projects items JSON:", e);
        projectsItems.value = [];
      }
      contentMap.set('projects_items', { id: projectsEntry.ID, page_name: projectsEntry.PageName, section: projectsEntry.Section, key: projectsEntry.Key });
    } else {
      projectsItems.value = [
        { num: '01', category: 'Web Development', title: 'Project 1', description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos', stack: 'html 5, css', image: '/static/thumb1.png', live: "http://github.com/mahathirrizky", github: "" },
        { num: '02', category: 'UI/UX Design', title: 'Project 2', description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos', stack: 'html 5, css', image: '/static/thumb2.png', live: "", github: "" },
        { num: '03', category: 'Logo Design', title: 'Project 3', description: 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quod ab eos', stack: 'html 5, css', image: '/static/thumb3.png', live: "", github: "" }
      ];
    }

  } catch (err) {
    error.value = err.response?.data?.error || err.message;
    console.error('Error fetching work page content:', err);
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

const saveProjectsContent = async () => {
  try {
    const projectsToSave = projectsItems.value.map(p => {
      const stackArray = p.stack.split(',').map(s => ({ name: s.trim() }));
      return {
        ...p,
        stack: JSON.stringify(stackArray, null, 2)
      };
    });
    await saveContent('work', 'projects', 'items', JSON.stringify(projectsToSave, null, 2));
  } catch (e) {
    error.value = 'Error saving projects content.' + e.message;
  }
};

const addProject = () => {
  projectsItems.value.push({
    num: '',
    category: '',
    title: '',
    description: '',
    stack: '', // Initialize as empty string
    image: '',
    live: '',
    github: ''
  });
};

const removeProject = (index) => {
  projectsItems.value.splice(index, 1);
};

onMounted(fetchContent);
</script>
