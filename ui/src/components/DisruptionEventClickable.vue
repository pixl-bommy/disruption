<template>
  <div class="container" :style="textColor" @click="() => emit('click', { itemId: item.id })">
    <div class="icon-container">
      <SvgIcon />
    </div>
    <span>{{ item.name }}</span>
  </div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, type Component } from 'vue'
import type { DisruptionItem } from '@/types/disruption'
import DefaultIcon from './icons/disruptions/DefaultDisruptionIcon.vue'

const props = defineProps<{
  item: DisruptionItem
}>()

const emit = defineEmits<{
  (e: 'click', payload: { itemId: string }): void
}>()

const textColor = {
  color: props.item.color,
}

const SvgIcon = ((iconName: string) => {
  if (!iconName || typeof iconName !== 'string') return DefaultIcon

  const iconFileName = iconName[0].toUpperCase() + iconName.slice(1)

  return defineAsyncComponent<Component>({
    loader: () => import(`@/components/icons/disruptions/${iconFileName}DisruptionIcon.vue`),
    errorComponent: DefaultIcon,
  })
})(props.item.icon)
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;

  background: color-mix(in srgb, currentColor 20%, white);
  border: 1px solid color-mix(in srgb, currentColor 40%, white);
  border-radius: 0.8rem;
  padding: 0.7rem;

  cursor: pointer;
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
