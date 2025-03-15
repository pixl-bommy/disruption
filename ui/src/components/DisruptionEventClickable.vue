<template>
  <div class="container" :style="textColor" @click="() => emit('click', { itemId: item.id })">
    <div class="icon-container">
      <SvgIcon />
    </div>
    <span>{{ item.name }}</span>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { getDisruptionIcon } from '@/lib/getDisruptionIcon'
import type { DisruptionItem } from '@/types/disruption'

const props = defineProps<{
  item: DisruptionItem
}>()

const emit = defineEmits<{
  (e: 'click', payload: { itemId: string }): void
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
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;

  cursor: pointer;

  background: color-mix(in srgb, currentColor 20%, white);
  border: 1px solid color-mix(in srgb, currentColor 40%, white);
  border-radius: 0.8rem;
  padding: 0.7rem;

  transition: background 0.2s;
}

.container:active {
  background: color-mix(in srgb, currentColor 35%, white);
}

.icon-container {
  background: color-mix(in srgb, currentColor 40%, white);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

span {
  white-space: nowrap;
  overflow: hidden;
}

svg {
  margin: 0.5em;
  width: 2.2em;
  height: 2.2em;
  position: relative;
  overflow: visible;
}
</style>
