<template>
  <a class="button" @animationend="animationEnd">
    {{ buttonText }}
  </a>
</template>

<script setup lang="ts">
const props = defineProps<{
  itemId: string
  buttonText: string
}>()

const emit = defineEmits<{
  (e: 'click', payload: { itemId: string }): void
}>()

const animationEnd = () => {
  emit('click', props)
}
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
  border: 1px solid currentColor;
  cursor: pointer;
  overflow: hidden;
  z-index: 1;

  transition: 0.5s linear;
}

.button::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 0.5rem;
  z-index: -2;
}

.button::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 0%;
  height: 100%;
  border-radius: 0.5rem;
  background: currentColor;
  opacity: 0.5;
  z-index: -1;
}

.button:active::before {
  animation: buttonTimer 0.5s linear;
}

@keyframes buttonTimer {
  0% {
    width: 0%;
  }

  95% {
    background: currentColor;
    width: 100%;
  }

  100% {
    background: yellow;
    width: 100%;
  }
}
</style>
