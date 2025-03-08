<template>
  <div class="button-list">
    <DisruptionButtonAddNew @click="props.onAddItemClick">Add Item</DisruptionButtonAddNew>
    <DisruptionButton
      v-for="disruption in props.disruptions"
      :key="disruption.id"
      :item-id="disruption.id"
      :button-text="disruption.name"
      @click="props.onDisruptionClick"
    />
  </div>
</template>

<script setup lang="ts">
import type { DisruptionItemList } from '@/types/disruption'
import DisruptionButton from './DisruptionButton.vue'
import DisruptionButtonAddNew from './DisruptionButtonAddNew.vue'

const props = defineProps<{
  disruptions: DisruptionItemList | null
  onAddItemClick: () => void
  onDisruptionClick: (payload: { itemId: string }) => void
}>()
</script>

<style scoped>
.button-list {
  display: grid;
  grid-template-columns: 1fr 1fr;
  justify-content: normal;
}
</style>
