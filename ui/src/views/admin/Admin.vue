<template>

    <div>
      <Sidebar/>
      <slot /> 
    </div>
  </template>

  <script setup>
import Sidebar from '../../components/admin/Sidebar.vue';
</script>