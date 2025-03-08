<template>
  <details open class="green">
    <summary>{{ props.item.name }}</summary>
    {{ props.item.description }}
  </details>
</template>

<script setup lang="ts">
import type { DisruptionItem } from '@/types/disruption'

const props = defineProps<{
  item: DisruptionItem
}>()
</script>

<style scoped>
details {
  position: relative;
  margin: 0.25rem;
  padding: 0.7rem;
  text-align: left;
  font-size: 1rem;
  border-radius: 0.5rem;
  border: 1px solid currentColor;
  cursor: pointer;
  overflow: hidden;
}

details summary {
  font-weight: bold;
}

details summary::-webkit-details-marker,
details summary::marker {
  display: none;
  content: '';
}
</style>
