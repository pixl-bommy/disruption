<template>
  <div class="container" :class="{ closed: props.closed }" :style="textColor">
    <div class="icon-container">
      <SvgIcon />
    </div>
    <span class="label">{{ props.name }}</span>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { getDisruptionIcon } from '@/lib/getDisruptionIcon'

const props = defineProps<{
  name: string
  icon: string
  color: string
  closed?: boolean
}>()

// the icon component to display
const SvgIcon = ref(getDisruptionIcon(props.icon))
watch(
  () => props.icon,
  (newIcon) => {
    SvgIcon.value = getDisruptionIcon(newIcon)
  },
)

// icon and text color
const textColor = ref({ color: props.color })
watch(
  () => props.color,
  (newColor) => {
    textColor.value = { color: newColor }
  },
)
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 0.75em;

  background: color-mix(in srgb, currentColor 20%, white);
  border: 1px solid color-mix(in srgb, currentColor 40%, white);
  border-radius: 1.5em;

  transition: all 0.3s ease;
}

.label {
  white-space: nowrap;
  overflow: hidden;
  padding-right: 0.75em;
  transition: all 0.3s ease;
}

.container.closed:not(:hover) {
  gap: 0;
}

.container.closed:not(:hover) .label {
  width: 0;
  padding-right: 0;
}

.icon-container {
  background: color-mix(in srgb, currentColor 40%, white);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

svg {
  margin: 0.5em;
  width: 1.75em;
  height: 1.75em;
  position: relative;
  overflow: visible;
}
</style>
