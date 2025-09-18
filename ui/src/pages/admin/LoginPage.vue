<template>
  <Auth>
    <div class="form-container bg-primary/20 border-accent border shadow-lg rounded-lg p-8 w-[500px] h-auto">
      <h2 class="title text-center font-bold text-2xl mb-8">Admin Login</h2>
      <form @submit.prevent="loginAdmin" class="form flex flex-col gap-4 mb-4">
        <input
          v-model="email"
          type="email"
          placeholder="Email"
          required
          class="input border border-gray-300 rounded-full p-3 bg-primary text-white focus:outline-none focus:ring-2 focus:ring-accent"
        />
        <div class="relative mb-6">
          <input
            v-model="password"
            :type="passwordFieldType"
            placeholder="Password"
            required
            class="border border-gray-300 rounded-full p-3 bg-primary text-white focus:outline-none focus:ring-2 focus:ring-accent w-full pr-10"
          />
          <i
            @click="togglePasswordVisibility"
            :class="['absolute right-3 top-1/2 transform -translate-y-1/2 cursor-pointer', passwordFieldType === 'password' ? 'fa-regular fa-eye' : 'fa-regular fa-eye-slash']"
          ></i>
        </div>
        <button
          type="submit"
          class="form-btn bg-accent text-primary py-2 rounded-full hover:bg-accent-hover transition duration-200"
        >
          Login
        </button>
      </form>
      <p v-if="message" class="text-center text-red-500 mt-4">{{ message }}</p>
    </div>
  </Auth>
</template>

<script setup>
import Auth from '../../views/auth/Auth.vue';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../../services/api'; // Import the configured axios instance

const router = useRouter();

const email = ref('');
const password = ref('');
const passwordFieldType = ref('password');
const message = ref('');

const togglePasswordVisibility = () => {
  passwordFieldType.value = passwordFieldType.value === 'password' ? 'text' : 'password';
};

const loginAdmin = async () => {
  try {
    const response = await api.post('/auth/login', {
      email: email.value,
      password: password.value
    });

    if (response.data && response.data.token) {
      localStorage.setItem('admin_token', response.data.token);
      router.push({ name: 'admin-edit-home' });
    } else {
      throw new Error(response.data.message || 'Login failed');
    }
  } catch (error) {
    message.value = error.response?.data?.error || error.message || 'An unexpected error occurred';
  }
};
</script>

<style scoped>
/* Optional: Add any additional styles here */
</style>