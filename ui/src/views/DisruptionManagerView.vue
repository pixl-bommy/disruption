<template>
  <div>
    <code class="header">$> manage disruptions</code>
    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="disruptions">
      <ul v-for="disruption in disruptions" :key="disruption.id">
        <li>{{ disruption.name }}: {{ disruption.description }}</li>
      </ul>
    </div>

    <div v-if="disruptions">
      <button @click="gotoAddDisruptionItem.navigate">Add Item</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useLink, useRoute } from 'vue-router'

import { fetchDisruptionItems } from '@/api/disruptions'
import type { DisruptionItemList } from '@/types/disruption'

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
