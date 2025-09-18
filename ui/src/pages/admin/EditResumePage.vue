<template>
  <Admin>
    <div class="p-4 sm:ml-64 py-8 text-white">
      <div class="max-w-7xl mx-auto">
        <h2 class="text-4xl font-bold mb-10">Edit Resume Page</h2>

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

          <!-- Experience Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg mb-12">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Experience Section</h3>
            <div class="space-y-8">
              <div v-for="(item, index) in experienceItems" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">Experience #{{ index + 1 }}</h4>
                <div class="space-y-4">
                  <div>
                    <label :for="`exp_company_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Company</label>
                    <input type="text" :id="`exp_company_${index}`" v-model="item.company" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`exp_position_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Position</label>
                    <input type="text" :id="`exp_position_${index}`" v-model="item.position" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`exp_duration_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Duration</label>
                    <input type="text" :id="`exp_duration_${index}`" v-model="item.duration" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeExperience(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove</button>
                </div>
              </div>
            </div>
            <div class="flex justify-between items-center mt-8">
              <button type="button" @click="addExperience" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New Experience</button>
              <button @click="saveExperienceContent" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All Experiences</button>
            </div>
          </div>

          <!-- Formal Education Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg mb-12">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Formal Education Section</h3>
            <div class="space-y-8">
              <div v-for="(item, index) in formalEducationItems" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">Formal Education #{{ index + 1 }}</h4>
                <div class="space-y-4">
                  <div>
                    <label :for="`formal_inst_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Institution</label>
                    <input type="text" :id="`formal_inst_${index}`" v-model="item.institution" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`formal_degree_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Degree</label>
                    <input type="text" :id="`formal_degree_${index}`" v-model="item.degree" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`formal_duration_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Duration</label>
                    <input type="text" :id="`formal_duration_${index}`" v-model="item.duration" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeFormalEducation(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove</button>
                </div>
              </div>
            </div>
            <div class="flex justify-between items-center mt-8">
              <button type="button" @click="addFormalEducation" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New Formal Education</button>
              <button @click="saveFormalEducationContent" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All Formal Education</button>
            </div>
          </div>

          <!-- Informal Education Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg mb-12">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Informal Education Section</h3>
            <div class="space-y-8">
              <div v-for="(item, index) in informalEducationItems" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">Informal Education #{{ index + 1 }}</h4>
                <div class="space-y-4">
                  <div>
                    <label :for="`informal_inst_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Institution</label>
                    <input type="text" :id="`informal_inst_${index}`" v-model="item.institution" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`informal_degree_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Degree</label>
                    <input type="text" :id="`informal_degree_${index}`" v-model="item.degree" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`informal_duration_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Duration</label>
                    <input type="text" :id="`informal_duration_${index}`" v-model="item.duration" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeInformalEducation(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove</button>
                </div>
              </div>
            </div>
            <div class="flex justify-between items-center mt-8">
              <button type="button" @click="addInformalEducation" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New Informal Education</button>
              <button @click="saveInformalEducationContent" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All Informal Education</button>
            </div>
          </div>

          <!-- Skills Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg mb-12">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">Skills Section</h3>
            <div class="space-y-8">
              <div v-for="(item, index) in skillsItems" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">Skill #{{ index + 1 }}</h4>
                <div class="space-y-4">
                  <div>
                    <label :for="`skill_icon_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Icon URL</label>
                    <div class="flex items-center gap-4">
                      <input type="text" :id="`skill_icon_${index}`" v-model="item.icon" readonly class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-gray-400 focus:ring-2 focus:ring-accent focus:border-transparent" />
                      <button type="button" @click="triggerFileUpload(index)" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300 whitespace-nowrap">Upload New</button>
                    </div>
                    <input type="file" :ref="el => fileInput[index] = el" @change="event => handleFileChange(event, index)" class="hidden" accept="image/svg+xml" />
                    <img v-if="item.icon" :src="item.icon" alt="Skill Icon Preview" class="mt-4 w-16 h-16 object-contain rounded-lg bg-gray-800 p-2" />
                  </div>
                  <div>
                    <label :for="`skill_name_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Name</label>
                    <input type="text" :id="`skill_name_${index}`" v-model="item.name" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeSkill(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove</button>
                </div>
              </div>
            </div>
            <div class="flex justify-between items-center mt-8">
              <button type="button" @click="addSkill" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New Skill</button>
              <button @click="saveSkillsContent" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All Skills</button>
            </div>
          </div>

          <!-- About Section -->
          <div class="bg-gray-800/50 p-8 rounded-2xl shadow-lg">
            <h3 class="text-2xl font-semibold mb-6 border-b border-gray-700 pb-4">About Section</h3>
            <div class="space-y-8">
              <div v-for="(item, index) in aboutItems" :key="index" class="bg-gray-900/70 p-6 rounded-xl border border-gray-700">
                <h4 class="text-xl font-semibold mb-4">About Item #{{ index + 1 }}</h4>
                <div class="space-y-4">
                  <div>
                    <label :for="`about_field_name_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Field Name</label>
                    <input type="text" :id="`about_field_name_${index}`" v-model="item.fieldName" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                  <div>
                    <label :for="`about_field_value_${index}`" class="block text-sm font-medium text-gray-400 mb-2">Field Value</label>
                    <input type="text" :id="`about_field_value_${index}`" v-model="item.fieldValue" class="w-full p-3 border border-gray-700 rounded-lg bg-gray-800 text-white focus:ring-2 focus:ring-accent focus:border-transparent" />
                  </div>
                </div>
                <div class="flex justify-end mt-6">
                  <button type="button" @click="removeAboutItem(index)" class="bg-red-600/80 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg transition-all duration-300">Remove</button>
                </div>
              </div>
            </div>
            <div class="flex justify-between items-center mt-8">
              <button type="button" @click="addAboutItem" class="bg-blue-600/80 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300">Add New About Item</button>
              <button @click="saveAboutContent" class="bg-accent hover:bg-accent-hover text-primary font-bold py-3 px-6 rounded-lg transition-all duration-300 shadow-md hover:shadow-lg">Save All About Items</button>
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
const fileInput = ref([]); // Use an array for multiple file inputs

// Data for each section (as JSON string)
const experienceItems = ref([]);
const formalEducationItems = ref([]);
const informalEducationItems = ref([]);
const skillsItems = ref([]);
const aboutItems = ref([]);

// Store original content IDs for updates
const contentMap = new Map(); // Map<key, {id, page_name, section, key}>

const triggerFileUpload = (index) => {
  if (fileInput.value[index]) {
    fileInput.value[index].click();
  }
};

const handleFileChange = async (event, index) => {
  const file = event.target.files[0];
  if (!file) return;

  message.value = 'Uploading...';
  error.value = '';

  try {
    const response = await uploadFile(file);
    if (response.data && response.data.file_path) {
      skillsItems.value[index].icon = response.data.file_path;
      message.value = 'File uploaded successfully! Saving...';
      // Automatically save the skills content to update the URL in the database
      await saveSkillsContent();
    } else {
      throw new Error('Invalid response from server after upload.');
    }
  } catch (err) {
    error.value = err.response?.data?.error || err.message || 'An error occurred during file upload.';
    message.value = '';
    console.error('File upload error:', err);
  }
};

const fetchContent = async () => {
  console.log("fetchContent started");
  isLoading.value = true;
  try {
    console.log("Calling getAdminContent");
    const response = await getAdminContent();
    console.log("getAdminContent response:", response);
    const contents = response.data.contents || [];

    const processSection = (sectionName, defaultItems, refItems) => {
      const entry = contents.find(c => c.page_name === 'resume' && c.section === sectionName && c.key === 'items');
      if (entry) {
        try {
          refItems.value = JSON.parse(entry.value);
        } catch (e) {
          console.error(`Failed to parse ${sectionName} items JSON:`, e);
          refItems.value = defaultItems;
        }
        contentMap.set(`${sectionName}_items`, { id: entry.ID, page_name: entry.PageName, section: entry.Section, key: entry.Key });
      } else {
        refItems.value = defaultItems;
      }
    };

    // Process Experience Content
    processSection('experience', [
      { company: "Tech Solutions inc", position: "Full Stack Developer", duration: "2022 - Present" },
      { company: "Web Design Studio", position: "Back-End Developer Freelance", duration: "april 2021" },
      { company: "Telkomsel", position: "Front-End Developer Freelance", duration: "2020 - 2021" },
      { company: "Tech Academy", position: "Teaching Assistent", duration: "juni 2020" }
    ], experienceItems);

    // Process Formal Education Content
    processSection('formaleducation', [
      { institution: "Fakultas Teknik Elektro Unhas", degree: "Magister Teknik Elektro", duration: "2016 - 2018" },
      { institution: "STMIK BALIKPAPAN", degree: "Sarjana Teknik Informatika", duration: "2011- 2015" }
    ], formalEducationItems);

    // Process Informal Education Content
    processSection('informaleducation', [
      { institution: "Online Course Platform", degree: "Full Stack Web Development Bootcamp", duration: "2023" },
      { institution: "Codecademy", degree: "Front-End Track", duration: "2022" },
      { institution: "Online Course", degree: "Programing Course", duration: "2020-2021" },
      { institution: "TEch Institute", degree: "Certivied Web Developer", duration: "2019" }
    ], informalEducationItems);

    // Process Skills Content
    processSection('skills', [
      { icon: '/static/html.svg', name: "html 5" },
      { icon: '/static/css.svg', name: "css" },
      { icon: '/static/js.svg', name: "javascript" },
      { icon: '/static/react.svg', name: "react.js" },
      { icon: '/static/next.svg', name: "next.js" },
      { icon: '/static/tailwind.svg', name: "tailwind css" }
    ], skillsItems);

    // Process About Content
    processSection('about', [
      { fieldName: "Name", fieldValue: "+(62) 812-3456-7890" },
      { fieldName: "Email", fieldValue: "Mahathirrizky@gmail.com" },
      { fieldName: "Nationality", fieldValue: "Indonesia" },
      { fieldName: "Experienxe", fieldValue: "12+ Years" },
      { fieldName: "Freelance", fieldValue: "Available" },
      { fieldName: "Languages", fieldValue: "English, Indonesian" }
    ], aboutItems);

  } catch (err) {
    console.error('Error fetching resume page content:', err);
    error.value = err.response?.data?.error || err.message;
  } finally {
    console.log("fetchContent finished");
    isLoading.value = false;
  }
};

const saveContent = async (pageName, section, key, value) => {
  message.value = '';
  error.value = '';
  try {
    const existing = contentMap.get(`${section}_items`); // Key for map is section_items
    let response;
    if (existing && existing.id) {
      // Update existing content
      response = await updateContent(existing.id, { PageName: pageName, Section: section, Key: key, Value: value });
    } else {
      // Create new content
      response = await createContent({ PageName: pageName, Section: section, Key: key, Value: value });
      // Update contentMap with new ID for future updates
      contentMap.set(`${section}_items`, { id: response.data.content.ID, page_name: pageName, section: section, key: key });
    }
    message.value = response.data.message || 'Content saved successfully!';
  } catch (err) {
    error.value = err.response?.data?.error || err.message;
    console.error('Error saving content:', err);
  }
};

const saveExperienceContent = async () => {
  try {
    await saveContent('resume', 'experience', 'items', JSON.stringify(experienceItems.value, null, 2));
  } catch (e) {
    error.value = 'Error saving experience content.' + e.message;
  }
};

const saveFormalEducationContent = async () => {
  try {
    await saveContent('resume', 'formaleducation', 'items', JSON.stringify(formalEducationItems.value, null, 2));
  } catch (e) {
    error.value = 'Error saving formal education content.' + e.message;
  }
};

const saveInformalEducationContent = async () => {
  try {
    await saveContent('resume', 'informaleducation', 'items', JSON.stringify(informalEducationItems.value, null, 2));
  } catch (e) {
    error.value = 'Error saving informal education content.' + e.message;
  }
};

const saveSkillsContent = async () => {
  try {
    await saveContent('resume', 'skills', 'items', JSON.stringify(skillsItems.value, null, 2));
  } catch (e) {
    error.value = 'Error saving skills content.' + e.message;
  }
};

const saveAboutContent = async () => {
  try {
    await saveContent('resume', 'about', 'items', JSON.stringify(aboutItems.value, null, 2));
  } catch (e) {
    error.value = 'Error saving about content.' + e.message;
  }
};

const addExperience = () => {
  experienceItems.value.push({ company: '', position: '', duration: '' });
};

const removeExperience = (index) => {
  experienceItems.value.splice(index, 1);
};

const addFormalEducation = () => {
  formalEducationItems.value.push({ institution: '', degree: '', duration: '' });
};

const removeFormalEducation = (index) => {
  formalEducationItems.value.splice(index, 1);
};

const addInformalEducation = () => {
  informalEducationItems.value.push({ institution: '', degree: '', duration: '' });
};

const removeInformalEducation = (index) => {
  informalEducationItems.value.splice(index, 1);
};

const addSkill = () => {
  skillsItems.value.push({ icon: '', name: '' });
};

const removeSkill = (index) => {
  skillsItems.value.splice(index, 1);
};

const addAboutItem = () => {
  aboutItems.value.push({ fieldName: '', fieldValue: '' });
};

const removeAboutItem = (index) => {
  aboutItems.value.splice(index, 1);
};

onMounted(() => {
  fetchContent();
});
</script>
