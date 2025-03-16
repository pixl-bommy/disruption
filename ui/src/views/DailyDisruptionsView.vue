<template>
  <div>
    <div v-if="!dailyItems" class="loading">Loading...</div>
    <DisruptionChipList v-else :disruptions="dailyItems" />

    <div v-if="!disruptions" class="loading">Loading...</div>
    <DisruptionEvents
      v-else
      :disruptions="disruptions"
      :onDisruptionClick="handleDisruptionClick"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

import { fetchDisruptionItems } from '@/api/disruptions'
import type { DisruptionItemList } from '@/types/disruption'
import DisruptionChipList from '@/components/DisruptionChipList.vue'
import DisruptionEvents from '@/components/DisruptionEvents.vue'

const route = useRoute()

const disruptions = ref<DisruptionItemList | null>(null)
const dailyItems = ref<DisruptionItemList | null>(null)
const loading = ref(false)

/**
 * Handle the click event on a disruption item.
 */
async function handleDisruptionClick(payload: { itemId: string }) {
  console.log('Disruption clicked:', payload.itemId)
}

/**
 * Fetch the list of disruption items from the API.
 */
async function fetchDisruptions() {
  disruptions.value = null
  loading.value = true

  try {
    const itemList = await fetchDisruptionItems()
    disruptions.value = itemList
    dailyItems.value = itemList
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
}

// fetch the list of disruption items when the route changes
watch(() => route.path, fetchDisruptions, { immediate: true })
</script>
