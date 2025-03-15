<template>
  <div class="disruption-list">
    <Chip
      v-for="(disruption, index) of props.disruptions"
      :color="disruption.color"
      :icon="disruption.icon"
      :name="disruption.name"
      :closed="index + 1 !== props.disruptions?.length"
      :key="disruption.id"
    />
  </div>
</template>

<script setup lang="ts">
import type { DisruptionItemList } from '@/types/disruption'
import Chip from './Chip.vue'

const props = defineProps<{
  disruptions: DisruptionItemList | null
}>()
</script>

<style scoped>
.disruption-list {
  position: relative;
  font-size: 0.7rem;
  margin: 1.5em 0;

  display: flex;
  flex-wrap: wrap;
  gap: 0.5em 0;
}
</style>
