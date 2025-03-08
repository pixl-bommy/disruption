<template>
  <span class="bean">{{ props.item.name }}</span>
  <span>&raquo;</span>
</template>

<script setup lang="ts">
import type { DisruptionItem } from '@/types/disruption'

const props = defineProps<{
  item: DisruptionItem
}>()
</script>

<style scoped>
span {
  padding: 0.35em;
}

span.bean {
  padding: 0.35em 1em;
  background-color: #f0f0f0;
  border-radius: 2em;
}
</style>
