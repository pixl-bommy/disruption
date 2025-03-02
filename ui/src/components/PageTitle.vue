<template>
  <code class="header">$> {{ title }}</code>
</template>

<script setup lang="ts">
const props = defineProps<{
  title: string
}>()
</script>

<style scoped>
code.header {
  display: block;
  font-size: 1.5em;
  font-weight: bold;
  padding: 0.5em 0;
}
</style>
