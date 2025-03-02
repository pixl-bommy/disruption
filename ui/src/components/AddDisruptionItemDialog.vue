<template>
  <div class="add-disruption-item">
    <h1>Add a disruption item</h1>
    <form action="" @submit="saveAndClose">
      <p>Bezeichnung</p>
      <input v-model="disruptionName" type="text" required />
      <p>Beschreibung</p>
      <textarea v-model="disruptionDescription" rows="8" required></textarea>
      <button type="submit">Save</button>
      <button @click="rejectAndClose">Cancel</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { saveNewDisruptionItem } from '@/api/disruptions'
import { ref } from 'vue'

const emit = defineEmits<{
  (event: 'onClose', action: 'saved' | 'cancelled' | 'failed'): void
}>()

const disruptionDescription = ref('')
const disruptionName = ref('')

async function saveAndClose() {
  // read data from the form
  const data = {
    name: disruptionName.value,
    description: disruptionDescription.value,
  }

  // send data to the server
  try {
    await saveNewDisruptionItem(data)
  } catch (err) {
    emit('onClose', 'failed')
    return
  }

  emit('onClose', 'saved')
}

function rejectAndClose() {
  emit('onClose', 'cancelled')
}
</script>
