<template>
  <div class="overflow-x-auto min-w-full">
    <div class="py-2 align-middle inline-block min-w-full">
      <div class="overflow-hidden border border-gray-300 sm:rounded-lg">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr class="max-w-0">
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Email
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Date
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="min-w-full bg-white divide-y divide-gray-200">
            <tr v-for="invitation in invitations" :key="invitation.id" class="table-row min-w-full">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-md font-medium text-gray-900 truncate">
                  {{ invitation.invitee_email }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-md font-medium text-gray-900 truncate">
                  {{ date(invitation.created_at) }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <sl-button variant="neutral" circle @click="onDeleteClicked(invitation)">
                  <TrashIcon class="h-5 w-5" aria-hidden="true" />
                </sl-button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { StaffInvitation } from '@/api/model';
import type { PropType } from 'vue';
import date from 'mdninja-js/src/libs/date';
import { TrashIcon } from '@heroicons/vue/24/outline';
import SlButton from '@shoelace-style/shoelace/dist/components/button/button.js';

// props
defineProps({
  invitations: {
    type: Array as PropType<StaffInvitation[]>,
    required: true,
  },
});

// events
const $emit = defineEmits(['delete']);

// composables

// lifecycle

// variables

// computed

// watch

// functions
function onDeleteClicked(invitation: StaffInvitation) {
  $emit('delete', invitation);
}
</script>
