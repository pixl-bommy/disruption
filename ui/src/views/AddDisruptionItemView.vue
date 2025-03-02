<template>
  <div>
    <PageTitle title="add disruption item" />
    <form action="" @submit="saveAndClose">
      <p>Bezeichnung</p>
      <input v-model="disruptionName" type="text" required />
      <p>Beschreibung</p>
      <textarea v-model="disruptionDescription" rows="8" required></textarea>
      <div>
        <button @click="rejectAndClose">Cancel</button>
        <button type="submit">Save</button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { saveNewDisruptionItem } from '@/api/disruptions'
import PageTitle from '@/components/PageTitle.vue'

// routing
const router = useRouter()

// component data
const disruptionDescription = ref('')
const disruptionName = ref('')

/**
 * Save the new disruption item and close the dialog.
 * TODO: Maybe we do not want to close automatically, if the save operation failed.
 */
async function saveAndClose() {
  // read data from the form
  const data = {
    name: disruptionName.value,
    description: disruptionDescription.value,
  }

  // store new disruption item using the API
  try {
    await saveNewDisruptionItem(data)
  } catch (err) {
    // TODO: handle error
  }

  // close the form by navigating back
  router.back()
}

/**
 * Close the dialog without saving the new disruption item.
 * TODO: Maybe we want to ask the user for confirmation before closing.
 */
function rejectAndClose() {
  // close the form by navigating back
  router.back()
}
</script>
