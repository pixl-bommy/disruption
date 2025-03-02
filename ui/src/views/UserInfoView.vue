<template>
  <div>
    <code class="header">$> whois {{ loading ? '...' : user?.name }}</code>
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
import type { UserInfo } from '@/types/users'

const route = useRoute()

const error = ref<unknown | null>(null)
const loading = ref(false)
const user = ref<UserInfo | null>(null)

/**
 * Fetch the user data from the API.
 */
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

// fetch the user data when the route changes
watch(() => route.path, fetchUserData, { immediate: true })
</script>
