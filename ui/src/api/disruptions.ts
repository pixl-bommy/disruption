import type { DisruptionItem, DisruptionItemList } from '@/types/disruption'
import { serverApiBaseURL } from './_root'

// The base url for the disruptions api.
const disruptionsApiBaseUrl = serverApiBaseURL + '/disruptions'

/**
 * Writes a new disruption item to the server api.
 * @param payload The data for the new disruption item.
 * @throws An error if the fetch request fails.
 */
export async function saveNewDisruptionItem(payload: { name: string; description: string }) {
  const response = await fetch(disruptionsApiBaseUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
  })

  if (!response.ok) {
    throw new Error('Failed to save new disruption item')
  }
}

/**
 * Fetches disruption items from the server api.
 * @returns A promise that resolves to an array of disruption items.
 * @throws An error if the fetch request fails or the response is not of the expected format.
 */
export async function fetchDisruptionItems(): Promise<DisruptionItemList> {
  const response = await fetch(disruptionsApiBaseUrl)

  if (!response.ok) {
    throw new Error('Failed to fetch disruption items')
  }

  // The response might not have failed, but still not be of the expected format.
  // So let's check that the response is an array.
  const rawJson = await response.json()
  if (!Array.isArray(rawJson)) {
    throw new Error('Fetched disruption items are not of expected format: ' + typeof rawJson)
  }

  // The response is an array, but we still need to check that each item is of
  // the expected format. We'll do that by filtering out any items that are not.
  return rawJson
    .map<DisruptionItem>((entry) => ({
      id: entry['uid'],
      name: entry['name'],
      description: entry['description'],
      icon: entry['icon'] ?? 'default',
      color: entry['color'] ?? fallbackColors.next().value,
    }))
    .filter((entry) => entry.id && entry.name && entry.description)
}

const fallbackColors = (function* () {
  let index = 0
  const colors = ['black', 'darkgreen', 'navy', '#9932CC', 'chocolate']

  while (true) {
    yield colors[index]
    index = (index + 1) % colors.length
  }
})()
