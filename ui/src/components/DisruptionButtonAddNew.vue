<template>
  <a class="button" @click="() => emit('click')">
    <slot />
  </a>
</template>

<script setup lang="ts">
const emit = defineEmits<{
  (e: 'click'): void
}>()
</script>

<style scoped>
.button {
  display: inline-flex;
  flex-wrap: wrap;
  align-content: center;
  justify-content: center;
  position: relative;
  margin: 0.25rem;
  padding: 0.7rem 0;
  text-align: center;
  font-size: 1rem;
  border-radius: 0.5rem;
  border: 1px solid transparent;
  cursor: pointer;
  overflow: hidden;
  z-index: 1;
  opacity: 0.5;

  transition: 0.3s linear;
}

@media (hover: hover) {
  .button:hover {
    opacity: 1;
    border-color: currentColor;
  }
}
</style>
