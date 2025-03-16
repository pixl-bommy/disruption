<template>
  <div class="disruption-list">
    <Chip
      v-for="(disruption, index) of props.dailyEvents"
      :color="disruption.color"
      :icon="disruption.iconName"
      :name="disruption.disruptionName"
      :closed="index + 1 !== props.dailyEvents?.length"
      :key="disruption.id"
    />
  </div>
</template>

<script setup lang="ts">
import type { DailyEventItemList } from '@/types/dailyEvent'
import Chip from './Chip.vue'

const props = defineProps<{
  dailyEvents: DailyEventItemList | null
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
