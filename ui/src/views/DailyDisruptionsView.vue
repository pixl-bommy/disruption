<template>
  <div>
    <div v-if="loading" class="loading">Loading...</div>

    <DisruptionBeanList :disruptions="dailyItems" />

    <div class="button-list" v-if="disruptions">
      <EventAddButton @click="gotoAddDisruptionItem.navigate">Add Item</EventAddButton>
      <EventItemButton
        v-for="disruption in disruptions"
        :key="disruption.id"
        :item-id="disruption.id"
        :button-text="disruption.name"
        @click="handleDisruptionClick"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useLink, useRoute } from 'vue-router'

import { fetchDisruptionItems } from '@/api/disruptions'
import type { DisruptionItemList } from '@/types/disruption'
import DisruptionBeanList from '@/components/DisruptionBeanList.vue'
import EventItemButton from '@/components/EventItemButton.vue'
import EventAddButton from '@/components/EventAddButton.vue'

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

<style scoped>
.button-list {
  display: grid;
  grid-template-columns: 1fr 1fr;
  justify-content: normal;
}
</style>
