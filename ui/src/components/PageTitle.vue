<template>
  <code class="header">$> {{ titleFromRoute }}</code>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const titleFromRoute = ref(route.meta.title)

watch(
  () => route.meta,
  () => (titleFromRoute.value = route.meta.title),
)
</script>

<style scoped>
code.header {
  display: block;
  font-size: 1.5em;
  font-weight: bold;
}
</style>
