<template>
  <div class="main-app-bar">
    <PageTitle />
    <a class="button" @click="toggle">{{ isOpen ? 'O' : 'C' }}</a>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import PageTitle from './PageTitle.vue'

const emit = defineEmits<{
  (e: 'toggle', open: boolean): void
}>()

const isOpen = ref(true)

const toggle = () => {
  isOpen.value = !isOpen.value
  emit('toggle', isOpen.value)
}
</script>

<style scoped>
.main-app-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 0;
}

.main-app-bar .button {
  width: 2rem;
  height: 2rem;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 50%;
}
</style>
