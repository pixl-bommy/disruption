<template>
  <div class="disruption-list">
    <DisruptionBean v-for="disruption of props.disruptions" :item="disruption" />
  </div>
</template>

<script setup lang="ts">
import type { DisruptionItemList } from '@/types/disruption'
import DisruptionBean from './DisruptionBean.vue'

const props = defineProps<{
  disruptions: DisruptionItemList | null
}>()
</script>

<style scoped>
.disruption-list {
  display: flex;
  flex-wrap: wrap;

  font-size: 0.7rem;
  gap: 0.3em 0.1em;

  margin: 1.5em 0;
}
</style>
