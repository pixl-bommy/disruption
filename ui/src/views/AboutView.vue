<template>
  <div class="about">
    <h1>This is an about page</h1>
    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="error" class="error">{{ error }}</div>

    <div v-if="user" class="content">
      <p>{{ user.uid }}</p>
      <p>{{ user.name }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

import { getUserMe } from '@/api/users'

const route = useRoute()

const loading = ref(false)
const user = ref(null)
const error = ref(null)

async function fetchUserData() {
  error.value = user.value = null
  loading.value = true

  try {
    user.value = await getUserMe()
  } catch (err) {
    error.value = err.toString()
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
