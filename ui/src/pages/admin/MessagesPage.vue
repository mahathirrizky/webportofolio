<template>
  <Admin>
    <div class="p-4 sm:ml-64 py-8 text-white">
      <div class="max-w-7xl mx-auto">
        <h2 class="text-4xl font-bold mb-10">Messages</h2>

        <div v-if="isLoading" class="text-center text-gray-400">
          <p>Loading messages...</p>
        </div>

        <div v-else-if="messages.length === 0" class="text-center text-gray-400">
          <p>No messages found.</p>
        </div>

        <div v-else class="space-y-6">
          <div v-for="message in messages" :key="message.id" class="bg-gray-800/50 p-6 rounded-2xl shadow-lg">
            <div class="flex justify-between items-start">
              <div>
                <h3 class="text-xl font-semibold text-accent">{{ message.name }}</h3>
                <p class="text-gray-400">{{ message.email }}</p>
                <p class="text-gray-400">{{ message.phone }}</p>
                <p class="text-gray-400">Service: {{ message.service }}</p>
                <p class="text-gray-500 text-sm mt-1">{{ new Date(message.created_at).toLocaleString() }}</p>
              </div>
              <button @click="deleteMessageHandler(message.id)" class="text-red-500 hover:text-red-400">
                <i class="mdi mdi-delete text-2xl"></i>
              </button>
            </div>
            <p class="text-white mt-4">{{ message.message }}</p>
          </div>
        </div>

      </div>
    </div>
  </Admin>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import Admin from '../../views/admin/Admin.vue';
import { getMessages, deleteMessage } from '../../services/api';

const isLoading = ref(true);
const messages = ref([]);

const fetchMessages = async () => {
  isLoading.value = true;
  try {
    const response = await getMessages();
    messages.value = response.data.messages || [];
  } catch (error) {
    console.error('Failed to fetch messages:', error);
  } finally {
    isLoading.value = false;
  }
};

const deleteMessageHandler = async (id) => {
  if (!confirm('Are you sure you want to delete this message?')) {
    return;
  }

  try {
    await deleteMessage(id);
    messages.value = messages.value.filter(m => m.id !== id);
  } catch (error) {
    console.error('Failed to delete message:', error);
  }
};

onMounted(fetchMessages);
</script>
