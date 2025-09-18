<template>  
   <div>  
     <button   
       @click="toggleNav"   
       class="nav-toggle"   
       :aria-expanded="isOpen ? 'true' : 'false'"   
       aria-controls="default-sidebar"  
       :class="{'hidden': isXlOrLarger}"  
     >  
       <span class="sr-only">{{ isOpen ? 'Close' : 'Open' }} sidebar</span>  
       <svg   
         class="w-6 h-6 text-accent"   
         aria-hidden="true"   
         fill="currentColor"   
         viewBox="0 0 20 20"   
         xmlns="http://www.w3.org/2000/svg"  
       >  
         <path   
           clip-rule="evenodd"   
           fill-rule="evenodd"   
           d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"  
         ></path>  
       </svg>  
     </button>  
   
     <aside   
       id="default-sidebar"   
       class="fixed top-0 left-0 z-40 h-screen transition-transform duration-300 ease-in-out w-auto md:w-56"  
       :class="{'-translate-x-full': !isOpen && !isXlOrLarger, 'translate-x-0': isOpen || isXlOrLarger}"   
       aria-label="Sidebar"  
     >  
       <div class="h-full px-3 py-4 overflow-y-auto bg-gray-800 flex flex-col justify-between">  
         <div>  
           <button   
             @click="closeNav"   
             class="absolute top-4 right-4 text-xl text-accent xl:hidden"  
           >  
             ✖  
           </button>  
           <ul class="space-y-2 font-medium">  
             <li v-for="(item, index) in navItems" :key="index">  
               <router-link
                 :to="item.path"
                 custom
                 v-slot="{ isExactActive, navigate }"
               >  
                 <a
                   :href="item.path"
                   @click.prevent="() => { navigate(); closeNav(); }"
                   class="flex items-center py-2 hover:text-accent text-gray-300"
                   :class="{ 'text-accent': isExactActive }"
                 >
                   <component :is="item.icon" class="mr-2 h-6 w-6"></component>  
                   <span class="ms-3">{{ item.name }}</span>  
                 </a>
               </router-link>  
             </li>  
           </ul>  
         </div>  
         <div class="mt-4 hover:text-accent text-gray-300">  
           <button   
             @click="logout"   
             class="flex items-center py-2 "  
           >  
             <svg   
               class="mr-2 h-6 w-6 "   
               aria-hidden="true"   
               fill="currentColor"   
               viewBox="0 0 20 20"   
               xmlns="http://www.w3.org/2000/svg"  
             >  
               <path   
                 fill-rule="evenodd"   
                 d="M3 3a1 1 0 000 2v12a1 1 0 100 2h12a1 1 0 100-2V5a1 1 0 000-2H3zm8 6a1 1 0 11-2 0 1 1 0 012 0zm5-2a1 1 0 00-1 1v6a1 1 0 102 0V7a1 1 0 00-1-1z"   
                 clip-rule="evenodd"  
               ></path>  
             </svg>  
             <span class="ms-3">Logout</span>  
           </button>  
         </div>  
       </div>  
     </aside>  
   </div>  
 </template>  
   
 <script>  
 import { HomeIcon, BriefcaseIcon, AcademicCapIcon, MailIcon, CogIcon, DocumentTextIcon, ChatAlt2Icon } from '@heroicons/vue/outline';  
   
 export default {  
   data() {  
     return {  
       navItems: [  
         { name: 'Home', path: '/admin/edit/home', icon: DocumentTextIcon },
         { name: 'Service', path: '/admin/edit/service', icon: CogIcon },
         { name: 'Resume', path: '/admin/edit/resume', icon: AcademicCapIcon },
         { name: 'Work', path: '/admin/edit/work', icon: BriefcaseIcon },
         { name: 'Contact', path: '/admin/edit/contact', icon: MailIcon },
         { name: 'Messages', path: '/admin/messages', icon: ChatAlt2Icon },
       ],  
       isOpen: false,  
       isXlOrLarger: false  
     };  
   },  
   components: { HomeIcon, BriefcaseIcon, AcademicCapIcon, MailIcon, CogIcon, DocumentTextIcon, ChatAlt2Icon },
   mounted() {  
     this.checkScreenSize();  
     window.addEventListener('resize', this.checkScreenSize);  
   },  
   beforeDestroy() {  
     window.removeEventListener('resize', this.checkScreenSize);  
   },
   methods: {
     toggleNav() {
       if (!this.isXlOrLarger) {
         this.isOpen = !this.isOpen;
       }
     },
     closeNav() {
       if (!this.isXlOrLarger) {
         this.isOpen = false;
       }
     },
     checkScreenSize() {
       this.isXlOrLarger = window.innerWidth >= 1280; // 1280px is the typical width for xl screens
     },
     logout() {
       localStorage.removeItem('admin_token');
       this.$router.push({name:'admin-login'});
     }
   }
 };  
 </script>  
   
 <style>  
 </style>  
 