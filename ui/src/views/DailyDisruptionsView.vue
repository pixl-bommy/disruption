<template>
  <div>
    <div v-if="loading" class="loading">Loading...</div>

    <DisruptionBeanList :disruptions="dailyItems" />

    <DisruptionEvents :disruptions="disruptions" :onDisruptionClick="handleDisruptionClick" />
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useLink, useRoute } from 'vue-router'

import { fetchDisruptionItems } from '@/api/disruptions'
import type { DisruptionItemList } from '@/types/disruption'
import DisruptionBeanList from '@/components/DisruptionBeanList.vue'
import DisruptionEvents from '@/components/DisruptionEvents.vue'

const route = useRoute()
const gotoAddDisruptionItem = useLink({ to: '/add-disruption-item' })

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
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
}

// fetch the list of disruption items when the route changes
watch(() => route.path, fetchDisruptions, { immediate: true })
</script>
