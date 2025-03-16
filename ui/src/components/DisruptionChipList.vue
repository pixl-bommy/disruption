<template>
  <div class="disruption-list">
    <Chip
      v-for="(disruption, index) of props.dailyEvents"
      :key="disruption.id"
      :color="disruption.color"
      :icon="disruption.iconName"
      :name="disruption.name"
      :closed="index + 1 !== props.dailyEvents?.length"
      :class="{
        unsettled:
          disruption.status !== DailyEventItemStatus.Completed &&
          disruption.status !== DailyEventItemStatus.Sending,
        inprogress: disruption.status === DailyEventItemStatus.Sending,
      }"
    />
  </div>
</template>

<script setup lang="ts">
import { DailyEventItemStatus, type DailyEventItemList } from '@/types/dailyEvent'
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

.unsettled {
  opacity: 0.5;
}

.inprogress {
  animation: pulse 0.7s infinite ease-in-out;
}

@keyframes pulse {
  0% {
    opacity: 0.35;
  }
  50% {
    opacity: 0.8;
  }
  100% {
    opacity: 0.35;
  }
}
</style>
