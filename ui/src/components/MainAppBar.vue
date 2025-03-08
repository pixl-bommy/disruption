<template>
  <div class="main-app-bar">
    <code class="header">$> {{ titleFromRoute }}</code>
    <a class="button" @click="toggle">{{ isOpen ? 'O' : 'C' }}</a>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const emit = defineEmits<{
  (e: 'toggle', open: boolean): void
}>()

const titleFromRoute = ref(route.meta.title)
const isOpen = ref(true)

const toggle = () => {
  isOpen.value = !isOpen.value
  emit('toggle', isOpen.value)
}

watch(
  () => route.meta,
  () => (titleFromRoute.value = route.meta.title),
)
</script>

<style scoped>
.main-app-bar {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 0;
  max-width: 100%;
}

.main-app-bar .header {
  flex: 1 1 auto;
  display: inline-block;
  font-size: 1.5em;
  font-weight: bold;

  white-space: nowrap;
  overflow: ellipsis;
}

.main-app-bar .button {
  flex: 0 0 2rem;
  width: 2rem;
  height: 2rem;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 50%;
}
</style>
