<template>
  <div class="pt-10 sm:mx-auto sm:w-full sm:max-w-xl">
    <div class="rounded-md bg-red-50 p-2 mb-3 mt-10" v-if="error">
      <div class="flex">
        <div class="ml-3">
          <p class="text-sm text-red-700">
            {{ error }}
          </p>
        </div>
      </div>
    </div>



    <div v-if="askForEmail">
      <div class="flex flex-col">
        <input id="email" name="email" type="email" autocomplete="email" required placeholder="my@email.com"
            v-model="email" @keyup="cleanupEmail"
            class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-xs placeholder-gray-400 focus:outline-hidden focus:ring-sky-500 focus:border-sky-500 sm:text-sm"
          />

        <small class="text-gray-400 font-small">
          We will send your purchases to this email address.
        </small>

        <div>
          <PButton :loading="loading" @click="onPlaceOrderClicked()" class="mt-3">
            Continue to Payment
          </PButton>
        </div>
      </div>

      <!-- <div>
        <div class="relative flex items-start mt-1.5">
          <div class="flex h-6 items-center">
            <input v-model="subscribeToNewsletter" type="checkbox" id="subscribe_to_newsletter" aria-describedby="subscribe_to_newsletter-description"
              name="subscribe_to_newsletter" class="h-4 w-4 rounded border-gray-300 text-sky-500 focus:ring-transparent" />
          </div>
          <div class="ml-3 text-sm leading-6">
            <label for="subscribe_to_newsletter" class="text-gray-900 cursor-pointer">
              1 Free email / week to learn how to (ab)use technology for fun & profit: Programming, Hacking & Entrepreneurship. <br />
              I hate spam even more than you do.
              I'll never share your email, and you can unsubscribe at any time..
            </label>
          </div>
        </div>
      </div> -->
    </div>

    <div v-else class="flex flex-col items-center">
      <div class="flex">
        <svg class="animate-spin -ml-1 mr-3 h-12 w-12 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-50" cx="2" cy="2" r="2" stroke="currentColor" stroke-width="2"></circle>
          <path class="opacity-75" fill="#424242" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
      </div>

      <div class="flex mt-5">
        Preparing your order. Please do not reload or change page.
      </div>
    </div>

  </div>
</template>

<script lang="ts" setup>
import { useStore } from '@/app/store';
import type { PlaceOrderInput } from '@/app/model';
import { onBeforeMount, ref } from 'vue';
import { useRoute } from 'vue-router';
import PButton from '@/ui/components/p_button.vue';
import { placeOrder, trackPage } from '@/app/mdninja';

// props

// events

// composables
const $route = useRoute();
const $store = useStore();

// lifecycle
onBeforeMount(() => {
  $store.setLoading(false);
  trackPage();
  if (!$store.contact) {
    askForEmail.value = true;
  } else {
    onPlaceOrderClicked();
  }
});

// variables
let error = ref('');
let askForEmail = ref(false);
let email = ref('');
let loading = ref(false);
let subscribeToNewsletter = ref(true);

// computed

// watch

// functions
function cleanupEmail() {
  email.value = email.value.toLowerCase().trim();
}

async function onPlaceOrderClicked() {
  error.value = '';
  const emailInput = email.value.trim();
  loading.value = true;

  let products = ($route.query.products as string ?? '').split(',').filter((p) => p != '');
  const input: PlaceOrderInput = {
    products: products,
    email: emailInput === '' ? undefined: emailInput,
    subscribe_to_newsletter: subscribeToNewsletter.value,
  };

  try {
    const checkoutSessionData = await placeOrder(input);
    location.href = checkoutSessionData.stripe_checkout_url;
  } catch (err: any) {
    error.value = err.message;
    $store.setLoading(false);
    loading.value = false;
  }
}
</script>
