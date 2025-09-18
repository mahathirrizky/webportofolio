   <template>
     <div class="w-full h-full relative">
       <div class="w-[298px] h-[298px] xl:w-[498px] xl:h-[498px] mix-blend-lighten transition-transform duration-700 delay-5000 transform hover:scale-105">
         <img 
           :src="imageSrc" 
           class="object-contain rounded-full transition-opacity duration-700 opacity-0" 
           ref="profileImage"
         />
       </div>
     </div>
   </template>

   <script setup>
   import { ref, onMounted, onUnmounted, watch } from 'vue';
   import { useRoute } from 'vue-router';

   const props = defineProps({
     imageSrc: {
       type: String,
       required: true,
     },
   });

   const profileImage = ref(null);
   let timeoutId = null;

   const route = useRoute();

   // Watch for changes in imageSrc to re-apply fade-in effect if image changes
   watch(() => props.imageSrc, (newVal, oldVal) => {
     if (newVal !== oldVal) {
       addDelayEffect();
     }
   });

   const addDelayEffect = () => {
     if (profileImage.value) {
       // Remove opacity-100 to trigger fade-out if already visible
       profileImage.value.classList.remove('opacity-100');
       if (timeoutId) {
         clearTimeout(timeoutId);
       }
       timeoutId = setTimeout(() => {
         if (profileImage.value) {
           profileImage.value.classList.add('opacity-100');
         }
       }, 500); // Delay 500 ms
     }
   };

   onMounted(() => {
     addDelayEffect();
   });

   onUnmounted(() => {
     if (timeoutId) {
       clearTimeout(timeoutId);
     }
   });
   </script>

   <style scoped>
   /* Optional: Add any additional styles here */
   img.opacity-100 {
     opacity: 1;
   }
   </style>
