<template>
  <details open :style="textColor">
    <summary class="summary">
      <div class="icon-container">
        <SvgIcon />
      </div>
      <span>{{ item.name }}</span>
      <div class="actions">
        <a @click.prevent="clickHandler('edit')" data-action="delete"><EditDisruptionIcon /></a>
        <a @click.prevent="clickHandler('delete')"><DeleteDisruptionIcon /></a>
      </div>
    </summary>
    <div class="description">{{ props.item.description }}</div>
  </details>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { getDisruptionIcon } from '@/lib/getDisruptionIcon'
import type { DisruptionItem } from '@/types/disruption'
import EditDisruptionIcon from './icons/disruptions/EditDisruptionIcon.vue'
import DeleteDisruptionIcon from './icons/disruptions/DeleteDisruptionIcon.vue'

const props = defineProps<{
  item: DisruptionItem
}>()

// the icon component to display
const SvgIcon = ref(getDisruptionIcon(props.item.icon))
watch(
  () => props.item.icon,
  (newIcon) => {
    SvgIcon.value = getDisruptionIcon(newIcon)
  },
)

// icon and text color
const textColor = ref({ color: props.item.color })
watch(
  () => props.item.color,
  (newColor) => {
    textColor.value = { color: newColor }
  },
)

const clickHandler = (action: 'edit' | 'delete') => {
  console.debug('clicked', action)
}
</script>

<style scoped>
details {
  position: relative;
  margin: 0.25rem;
  padding: 0.7rem;
  text-align: left;
  font-size: 1rem;
  border-radius: 0.5rem;
  overflow: hidden;

  border: 1px solid currentColor;
  background: color-mix(in srgb, currentColor 20%, white);
}

.summary {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  outline: none;
  cursor: pointer;
}

.summary::-webkit-details-marker,
.summary::marker {
  display: none;
  content: '';
}

.description {
  color: black;
  background: white;

  padding: 0.5rem;
  border-radius: 0.5rem;
  margin-top: 0.5rem;
}

.actions {
  margin-left: auto;
}

.actions a {
  color: color-mix(in srgb, black, white 50%);
  background: none;
}

.actions a svg {
  width: 1em;
  height: 1em;
}
</style>
