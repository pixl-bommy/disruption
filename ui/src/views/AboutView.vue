<template>
  <div class="about">
    <h1>This is an about page</h1>
    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="error" class="error">{{ error }}</div>

    <div v-if="user" class="content">
      <p>{{ user.uid }}</p>
      <p>{{ user.name }}</p>
    </div>

    <button @click="openDialog">Add Item</button>

    <AddDisruptionItemDialog v-if="isShowDialog" @on-close="closeDialog" />
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

import { getUserMe } from '@/api/users'
import AddDisruptionItemDialog from '@/components/AddDisruptionItemDialog.vue'

const route = useRoute()

const loading = ref(false)
const user = ref<{ uid: string; name: string } | null>(null)
const error = ref<unknown | null>(null)

const isShowDialog = ref(false)
function openDialog() {
  isShowDialog.value = true
}
function closeDialog(action: 'saved' | 'cancelled' | 'failed') {
  console.log(action)
  isShowDialog.value = false
}

async function fetchUserData() {
  error.value = user.value = null
  loading.value = true

  try {
    user.value = await getUserMe()
  } catch (err) {
    error.value = err?.toString() ?? null
  } finally {
    loading.value = false
  }
}

watch(() => route.path, fetchUserData, { immediate: true })
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
