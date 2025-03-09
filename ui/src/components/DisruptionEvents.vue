<template>
  <div class="button-list">
    <DisruptionEventClickable
      v-for="disruption in props.disruptions"
      :key="disruption.id"
      :item="disruption"
      @click="props.onDisruptionClick"
    />
  </div>
</template>

<script setup lang="ts">
import type { DisruptionItemList } from '@/types/disruption'
import DisruptionEventClickable from './DisruptionEventClickable.vue'

const props = defineProps<{
  disruptions: DisruptionItemList | null
  onDisruptionClick: (payload: { itemId: string }) => void
}>()
</script>

<style scoped>
.button-list {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;

  gap: 0.3rem;
}

.button-list > * {
  flex: 1 1 calc(50% - 0.3rem);
}
</style>
