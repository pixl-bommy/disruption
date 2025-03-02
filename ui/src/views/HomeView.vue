<template>
  <div class="home">
    <h1>This is an home page</h1>
    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="disruptions">
      <div v-for="disruption in disruptions" :key="disruption.id">
        <p>{{ disruption.name }}: {{ disruption.description }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { fetchDisruptionItems } from '@/api/disruptions'
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const loading = ref(false)
const disruptions = ref<{ id: string; name: string; description: string }[] | null>(null)

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

watch(() => route.path, fetchDisruptions, { immediate: true })
</script>

<style>
@media (min-width: 1024px) {
  .home {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
