<style scoped>
button {
  display: none;
}

.item-list {
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  justify-content: center;
  align-items: stretch;
}
</style>

<template>
  <div>
    <div v-if="loading" class="loading">Loading...</div>

    <button @click="gotoAddDisruptionItem.navigate">Add Item</button>

    <div class="item-list" v-if="disruptions">
      <DisruptionDetailsItem
        v-for="disruption in disruptions"
        :key="disruption.id"
        :item="disruption"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useLink, useRoute } from 'vue-router'

import { fetchDisruptionItems } from '@/api/disruptions'
import type { DisruptionItemList } from '@/types/disruption'

import DisruptionDetailsItem from '@/components/DisruptionDetailsItem.vue'

const route = useRoute()
const gotoAddDisruptionItem = useLink({ to: '/add-disruption-item' })

const disruptions = ref<DisruptionItemList | null>(null)
const loading = ref(false)

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
