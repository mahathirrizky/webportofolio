<template>
  <LandingPage>
    <div v-if="!isLoading && formContent" class="container mx-auto">
      <div class="flex flex-col xl:flex-row xl:gap-[30px]">
        <div class="xl:w-[54%] order-2 xl:order-none">
          <form @submit.prevent="handleSubmit" class="flex flex-col gap-6 p-10 bg-[#27272c] rounded-xl">
            <h3 class="text-4xl text-accent">{{ formContent.title }}</h3>
            <p class="text-white/60">
              {{ formContent.description }}
            </p>
            <div v-if="messageSent" class="p-4 bg-green-500/20 text-green-300 rounded-lg">
              Message sent successfully!
            </div>
            <div v-if="messageError" class="p-4 bg-red-500/20 text-red-300 rounded-lg">
              Failed to send message. Please try again later.
            </div>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <input type="text" v-model="form.firstname" placeholder="Firstname" required class="w-full p-3 bg-black/60 rounded-lg text-white" />
              <input type="text" v-model="form.lastname" placeholder="Lastname" required class="w-full p-3 bg-black/60 rounded-lg text-white" />
              <input type="email" v-model="form.email" placeholder="Email address" required class="w-full p-3 bg-black/60 rounded-lg text-white" />
              <div class="flex gap-2">
                <select v-model="form.countryCode" class="w-1/3 p-3 bg-black/60 rounded-lg text-white">
                  <option v-for="code in countryCodes" :key="code.code" :value="code.dial_code">{{ code.code }} {{ code.dial_code }}</option>
                </select>
                <input type="tel" v-model="form.phone" placeholder="Phone Number" required pattern="[0-9]*" class="w-2/3 p-3 bg-black/60 rounded-lg text-white" />
              </div>
              <input type="text" v-model="form.website" class="hidden" />
            </div>
            <div class="relative">
              <div @click="toggleDropdown" class="flex h-[48px] w-full items-center justify-between rounded-md border border-white/10 bg-primary px-4 text-base text-white/60 cursor-pointer">
                <span :class="{'text-white': form.service}">{{ form.service || 'Select a service' }}</span>
                <i class="mdi mdi-chevron-down"></i>
              </div>
              <ul v-if="isDropdownOpen" class="absolute z-10 w-full bg-primary rounded-md border border-white/10 mt-1">
                <li v-for="(option, index) in options" :key="index" @click="selectOption(option)" class="p-2 hover:bg-accent hover:text-black cursor-pointer">
                  {{ option }}
                </li>
              </ul>
              <input type="hidden" v-model="form.service" required />
            </div>
            <textarea v-model="form.message" required class="flex min-h-[80px] w-full border border-white/10 bg-primary px-4 py-5 text-base placeholder:text-white/60 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-accent focus-visible:ring-offset-0 disabled:cursor-not-allowed disabled:opacity-50 h-[200px] text-white rounded-md" placeholder="Type your message here."></textarea>
            <button type="submit" class="max-w-40 bg-accent hover:bg-accent-hover text-primary text-lg w-full py-3 rounded-3xl font-bold">
            {{ formContent.button_text }}
          </button>
          </form>
        </div>
        <div class="flex-1 flex items-center xl:justify-end order-1 xl:order-none mb-8 xl:mb-0">
          <ul class="flex flex-col gap-10">
            <li v-for="(item, index) in info" :key="index" class="flex items-center gap-6">
              <div class="w-[52px] h-[52px]  xl:w-[72px] xl:h-[72px] bg-[#27272c] text-accent rounded-md flex items-center justify-center">
                <FontAwesomeIcon :icon="item.icon" class="text-[28px]" />
              </div>
              <div class="flex-1">
                <p class="text-white/60">{{ item.title }}</p>
                <h3 class="text-xl">{{ item.description }}</h3>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </LandingPage>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getPublicContent, sendMessage } from '../services/api';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faPhone, faEnvelope, faMapMarkerAlt } from '@fortawesome/free-solid-svg-icons';
import LandingPage from '../views/LandingPage.vue';

const formContent = ref(null);
const info = ref([]);
const isLoading = ref(true);
const messageSent = ref(false);
const messageError = ref(false);

const countryCodes = ref([
  { code: 'US', dial_code: '+1' },
  { code: 'GB', dial_code: '+44' },
  { code: 'ID', dial_code: '+62' },
  { code: 'AF', dial_code: '+93' },
  { code: 'AL', dial_code: '+355' },
  { code: 'DZ', dial_code: '+213' },
  { code: 'AS', dial_code: '+1-684' },
  { code: 'AD', dial_code: '+376' },
  { code: 'AO', dial_code: '+244' },
  { code: 'AI', dial_code: '+1-264' },
  { code: 'AQ', dial_code: '+672' },
  { code: 'AG', dial_code: '+1-268' },
  { code: 'AR', dial_code: '+54' },
  { code: 'AM', dial_code: '+374' },
  { code: 'AW', dial_code: '+297' },
  { code: 'AU', dial_code: '+61' },
  { code: 'AT', dial_code: '+43' },
  { code: 'AZ', dial_code: '+994' },
  { code: 'BS', dial_code: '+1-242' },
  { code: 'BH', dial_code: '+973' },
  { code: 'BD', dial_code: '+880' },
  { code: 'BB', dial_code: '+1-246' },
  { code: 'BY', dial_code: '+375' },
  { code: 'BE', dial_code: '+32' },
  { code: 'BZ', dial_code: '+501' },
  { code: 'BJ', dial_code: '+229' },
  { code: 'BM', dial_code: '+1-441' },
  { code: 'BT', dial_code: '+975' },
  { code: 'BO', dial_code: '+591' },
  { code: 'BA', dial_code: '+387' },
  { code: 'BW', dial_code: '+267' },
  { code: 'BR', dial_code: '+55' },
  { code: 'IO', dial_code: '+246' },
  { code: 'VG', dial_code: '+1-284' },
  { code: 'BN', dial_code: '+673' },
  { code: 'BG', dial_code: '+359' },
  { code: 'BF', dial_code: '+226' },
  { code: 'BI', dial_code: '+257' },
  { code: 'KH', dial_code: '+855' },
  { code: 'CM', dial_code: '+237' },
  { code: 'CA', dial_code: '+1' },
  { code: 'CV', dial_code: '+238' },
  { code: 'KY', dial_code: '+1-345' },
  { code: 'CF', dial_code: '+236' },
  { code: 'TD', dial_code: '+235' },
  { code: 'CL', dial_code: '+56' },
  { code: 'CN', dial_code: '+86' },
  { code: 'CX', dial_code: '+61' },
  { code: 'CC', dial_code: '+61' },
  { code: 'CO', dial_code: '+57' },
  { code: 'KM', dial_code: '+269' },
  { code: 'CK', dial_code: '+682' },
  { code: 'CR', dial_code: '+506' },
  { code: 'HR', dial_code: '+385' },
  { code: 'CU', dial_code: '+53' },
  { code: 'CW', dial_code: '+599' },
  { code: 'CY', dial_code: '+357' },
  { code: 'CZ', dial_code: '+420' },
  { code: 'CD', dial_code: '+243' },
  { code: 'DK', dial_code: '+45' },
  { code: 'DJ', dial_code: '+253' },
  { code: 'DM', dial_code: '+1-767' },
  { code: 'DO', dial_code: '+1-809, +1-829, +1-849' },
  { code: 'TL', dial_code: '+670' },
  { code: 'EC', dial_code: '+593' },
  { code: 'EG', dial_code: '+20' },
  { code: 'SV', dial_code: '+503' },
  { code: 'GQ', dial_code: '+240' },
  { code: 'ER', dial_code: '+291' },
  { code: 'EE', dial_code: '+372' },
  { code: 'ET', dial_code: '+251' },
  { code: 'FK', dial_code: '+500' },
  { code: 'FO', dial_code: '+298' },
  { code: 'FJ', dial_code: '+679' },
  { code: 'FI', dial_code: '+358' },
  { code: 'FR', dial_code: '+33' },
  { code: 'GF', dial_code: '+594' },
  { code: 'PF', dial_code: '+689' },
  { code: 'GA', dial_code: '+241' },
  { code: 'GM', dial_code: '+220' },
  { code: 'GE', dial_code: '+995' },
  { code: 'DE', dial_code: '+49' },
  { code: 'GH', dial_code: '+233' },
  { code: 'GI', dial_code: '+350' },
  { code: 'GR', dial_code: '+30' },
  { code: 'GL', dial_code: '+299' },
  { code: 'GD', dial_code: '+1-473' },
  { code: 'GP', dial_code: '+590' },
  { code: 'GU', dial_code: '+1-671' },
  { code: 'GT', dial_code: '+502' },
  { code: 'GG', dial_code: '+44-1481' },
  { code: 'GN', dial_code: '+224' },
  { code: 'GW', dial_code: '+245' },
  { code: 'GY', dial_code: '+592' },
  { code: 'HT', dial_code: '+509' },
  { code: 'HN', dial_code: '+504' },
  { code: 'HK', dial_code: '+852' },
  { code: 'HU', dial_code: '+36' },
  { code: 'IS', dial_code: '+354' },
  { code: 'IN', dial_code: '+91' },
  { code: 'IR', dial_code: '+98' },
  { code: 'IQ', dial_code: '+964' },
  { code: 'IE', dial_code: '+353' },
  { code: 'IM', dial_code: '+44-1624' },
  { code: 'IL', dial_code: '+972' },
  { code: 'IT', dial_code: '+39' },
  { code: 'CI', dial_code: '+225' },
  { code: 'JM', dial_code: '+1-876' },
  { code: 'JP', dial_code: '+81' },
  { code: 'JE', dial_code: '+44-1534' },
  { code: 'JO', dial_code: '+962' },
  { code: 'KZ', dial_code: '+7' },
  { code: 'KE', dial_code: '+254' },
  { code: 'KI', dial_code: '+686' },
  { code: 'XK', dial_code: '+383' },
  { code: 'KW', dial_code: '+965' },
  { code: 'KG', dial_code: '+996' },
  { code: 'LA', dial_code: '+856' },
  { code: 'LV', dial_code: '+371' },
  { code: 'LB', dial_code: '+961' },
  { code: 'LS', dial_code: '+266' },
  { code: 'LR', dial_code: '+231' },
  { code: 'LY', dial_code: '+218' },
  { code: 'LI', dial_code: '+423' },
  { code: 'LT', dial_code: '+370' },
  { code: 'LU', dial_code: '+352' },
  { code: 'MO', dial_code: '+853' },
  { code: 'MK', dial_code: '+389' },
  { code: 'MG', dial_code: '+261' },
  { code: 'MW', dial_code: '+265' },
  { code: 'MY', dial_code: '+60' },
  { code: 'MV', dial_code: '+960' },
  { code: 'ML', dial_code: '+223' },
  { code: 'MT', dial_code: '+356' },
  { code: 'MH', dial_code: '+692' },
  { code: 'MQ', dial_code: '+596' },
  { code: 'MR', dial_code: '+222' },
  { code: 'MU', dial_code: '+230' },
  { code: 'YT', dial_code: '+262' },
  { code: 'MX', dial_code: '+52' },
  { code: 'FM', dial_code: '+691' },
  { code: 'MD', dial_code: '+373' },
  { code: 'MC', dial_code: '+377' },
  { code: 'MN', dial_code: '+976' },
  { code: 'ME', dial_code: '+382' },
  { code: 'MS', dial_code: '+1-664' },
  { code: 'MA', dial_code: '+212' },
  { code: 'MZ', dial_code: '+258' },
  { code: 'MM', dial_code: '+95' },
  { code: 'NA', dial_code: '+264' },
  { code: 'NR', dial_code: '+674' },
  { code: 'NP', dial_code: '+977' },
  { code: 'NL', dial_code: '+31' },
  { code: 'NC', dial_code: '+687' },
  { code: 'NZ', dial_code: '+64' },
  { code: 'NI', dial_code: '+505' },
  { code: 'NE', dial_code: '+227' },
  { code: 'NG', dial_code: '+234' },
  { code: 'NU', dial_code: '+683' },
  { code: 'NF', dial_code: '+672' },
  { code: 'KP', dial_code: '+850' },
  { code: 'MP', dial_code: '+1-670' },
  { code: 'NO', dial_code: '+47' },
  { code: 'OM', dial_code: '+968' },
  { code: 'PK', dial_code: '+92' },
  { code: 'PW', dial_code: '+680' },
  { code: 'PS', dial_code: '+970' },
  { code: 'PA', dial_code: '+507' },
  { code: 'PG', dial_code: '+675' },
  { code: 'PY', dial_code: '+595' },
  { code: 'PE', dial_code: '+51' },
  { code: 'PH', dial_code: '+63' },
  { code: 'PN', dial_code: '+870' },
  { code: 'PL', dial_code: '+48' },
  { code: 'PT', dial_code: '+351' },
  { code: 'PR', dial_code: '+1-787, +1-939' },
  { code: 'QA', dial_code: '+974' },
  { code: 'CG', dial_code: '+242' },
  { code: 'RE', dial_code: '+262' },
  { code: 'RO', dial_code: '+40' },
  { code: 'RU', dial_code: '+7' },
  { code: 'RW', dial_code: '+250' },
  { code: 'BL', dial_code: '+590' },
  { code: 'SH', dial_code: '+290' },
  { code: 'KN', dial_code: '+1-869' },
  { code: 'LC', dial_code: '+1-758' },
  { code: 'MF', dial_code: '+590' },
  { code: 'PM', dial_code: '+508' },
  { code: 'VC', dial_code: '+1-784' },
  { code: 'WS', dial_code: '+685' },
  { code: 'SM', dial_code: '+378' },
  { code: 'ST', dial_code: '+239' },
  { code: 'SA', dial_code: '+966' },
  { code: 'SN', dial_code: '+221' },
  { code: 'RS', dial_code: '+381' },
  { code: 'SC', dial_code: '+248' },
  { code: 'SL', dial_code: '+232' },
  { code: 'SG', dial_code: '+65' },
  { code: 'SX', dial_code: '+1-721' },
  { code: 'SK', dial_code: '+421' },
  { code: 'SI', dial_code: '+386' },
  { code: 'SB', dial_code: '+677' },
  { code: 'SO', dial_code: '+252' },
  { code: 'ZA', dial_code: '+27' },
  { code: 'KR', dial_code: '+82' },
  { code: 'SS', dial_code: '+211' },
  { code: 'ES', dial_code: '+34' },
  { code: 'LK', dial_code: '+94' },
  { code: 'SD', dial_code: '+249' },
  { code: 'SR', dial_code: '+597' },
  { code: 'SJ', dial_code: '+47' },
  { code: 'SZ', dial_code: '+268' },
  { code: 'SE', dial_code: '+46' },
  { code: 'CH', dial_code: '+41' },
  { code: 'SY', dial_code: '+963' },
  { code: 'TW', dial_code: '+886' },
  { code: 'TJ', dial_code: '+992' },
  { code: 'TZ', dial_code: '+255' },
  { code: 'TH', dial_code: '+66' },
  { code: 'TG', dial_code: '+228' },
  { code: 'TK', dial_code: '+690' },
  { code: 'TO', dial_code: '+676' },
  { code: 'TT', dial_code: '+1-868' },
  { code: 'TN', dial_code: '+216' },
  { code: 'TR', dial_code: '+90' },
  { code: 'TM', dial_code: '+993' },
  { code: 'TC', dial_code: '+1-649' },
  { code: 'TV', dial_code: '+688' },
  { code: 'UG', dial_code: '+256' },
  { code: 'UA', dial_code: '+380' },
  { code: 'AE', dial_code: '+971' },
  { code: 'GB', dial_code: '+44' },
  { code: 'US', dial_code: '+1' },
  { code: 'UY', dial_code: '+598' },
  { code: 'UZ', dial_code: '+998' },
  { code: 'VU', dial_code: '+678' },
  { code: 'VA', dial_code: '+39-06, +379' },
  { code: 'VE', dial_code: '+58' },
  { code: 'VN', dial_code: '+84' },
  { code: 'WF', dial_code: '+681' },
  { code: 'EH', dial_code: '+212' },
  { code: 'YE', dial_code: '+967' },
  { code: 'ZM', dial_code: '+260' },
  { code: 'ZW', dial_code: '+263' }
]);

const form = ref({
  firstname: '',
  lastname: '',
  email: '',
  countryCode: '+62',
  phone: '',
  service: '',
  message: '',
  website: '' // Honeypot field
});

const handleSubmit = async () => {
  try {
    const formData = new URLSearchParams();
    formData.append('name', `${form.value.firstname} ${form.value.lastname}`);
    formData.append('email', form.value.email);
    formData.append('phone', `${form.value.countryCode} ${form.value.phone}`);
    formData.append('service', form.value.service);
    formData.append('message', form.value.message);
    formData.append('website', form.value.website);

    await sendMessage(formData);
    messageSent.value = true;
    messageError.value = false;
    form.value = {
      firstname: '',
      lastname: '',
      email: '',
      countryCode: '+62',
      phone: '',
      service: '',
      message: '',
      website: ''
    };
  } catch (error) {
    console.error('Failed to send message:', error);
    messageError.value = true;
    messageSent.value = false;
  }
};

const iconMap = {
  faPhone,
  faEnvelope,
  faMapMarkerAlt
};

const processContent = (contents) => {
  const contactContents = contents.filter(c => c.page_name === 'contact');
  
  const form = {};
  contactContents.filter(c => c.section === 'form').forEach(c => {
    form[c.key] = c.value;
  });
  formContent.value = form;

  const infoContent = contactContents.find(c => c.section === 'info' && c.key === 'items');
  if (infoContent && infoContent.value) {
    try {
      info.value = JSON.parse(infoContent.value).map(item => ({ ...item, icon: iconMap[item.icon] }));
    } catch (e) {
      console.error("Failed to parse info JSON:", e);
    }
  }
  isLoading.value = false;
};

onMounted(async () => {
  try {
    const response = await getPublicContent();
    processContent(response.data.contents || []);
  } catch (error) {
    console.error('Failed to fetch content:', error);
    // Fallback data
    formContent.value = {
        title: 'Let’s work together',
        description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Eligendi vero quaerat dignissimos autem ab.',
        button_text: 'Send message'
    };
    info.value = [
      {
        icon: faPhone,
        title: 'Phone',
        description: '+1 234 567 890'
      },
      {
        icon: faEnvelope,
        title: 'Email',
        description: 'example@example.com'
      },
      {
        icon: faMapMarkerAlt,
        title: 'Address',
        description: '1234 Street Name, City Name, Country Name'
      }
    ];
    isLoading.value = false;
  }
});

const selected = ref('');
const isDropdownOpen = ref(false);
const options = ['Web Development', 'UI/UX Design', 'Logo Design'];

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value;
};

const selectOption = (option) => {
  form.value.service = option;
  isDropdownOpen.value = false;
};
</script>


