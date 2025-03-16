<template>
  <div>
    <div v-if="!dailyItems" class="loading">Loading...</div>
    <DisruptionChipList v-else :dailyEvents="dailyItems" />

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

import { fetchCreateDailyEventItem, fetchEventItems } from '@/api/dailyDisruptions'
import { fetchDisruptionItems } from '@/api/disruptions'
import {
  DailyEventItemStatus,
  type DailyEventItem,
  type DailyEventItemList,
  type DailyEventItemUnsettled,
} from '@/types/dailyEvent'
import type { DisruptionItemList } from '@/types/disruption'
import DisruptionChipList from '@/components/DisruptionChipList.vue'
import DisruptionEvents from '@/components/DisruptionEvents.vue'

const route = useRoute()

const disruptions = ref<DisruptionItemList | null>(null)
const dailyItems = ref<DailyEventItemList | null>(null)

/**
 * Handle the click event on a disruption item.
 */
async function handleDisruptionClick(payload: { itemId: string }) {
  const relatedDisruption = disruptions.value?.find((item) => item.id === payload.itemId)

  if (!relatedDisruption) {
    // NOTE: This should never happen. But maybe we want to catch this case.
    console.error('Disruption not found:', payload.itemId)
    return
  }

  const temporaryId = Date.now() + '-0'

  const newItem: DailyEventItem = {
    status: DailyEventItemStatus.Queued,
    id: temporaryId,

    disruptionId: relatedDisruption.id,
    name: relatedDisruption.name,
    color: relatedDisruption.color,
    iconName: relatedDisruption.icon,
  }

  console.log('Disruption clicked:', newItem)

  dailyItems.value = [...(dailyItems.value || []), newItem]

  try {
    dailyItems.value = (dailyItems.value || []).map((item) =>
      item.id === temporaryId ? { ...newItem, status: DailyEventItemStatus.Sending } : item,
    )

    const settledItem = await fetchCreateDailyEventItem(newItem)

    dailyItems.value = (dailyItems.value || []).map((item) =>
      item.id === temporaryId ? settledItem : item,
    )
  } catch (err) {
    console.error(err)

    dailyItems.value = (dailyItems.value || []).map((item) =>
      item.id === temporaryId ? { ...newItem, status: DailyEventItemStatus.Failed } : item,
    )
  }
}

/**
 * Fetch the list of disruption items from the API.
 */
async function fetchDisruptions() {
  disruptions.value = null

  try {
    const itemList = await fetchDisruptionItems()
    disruptions.value = itemList
  } catch (err) {
    console.error(err)
  }
}

/**
 * Fetch the list of event items from the API.
 */
async function fetchEvents() {
  disruptions.value = null

  try {
    const itemList = await fetchEventItems()
    dailyItems.value = itemList
  } catch (err) {
    console.error(err)
  }
}

// fetch the list of disruption items when the route changes
watch(
  () => route.path,
  () => {
    fetchDisruptions()
    fetchEvents()
  },
  { immediate: true },
)
</script>
