import type { DisruptionItem, DisruptionItemList } from '@/types/disruption'

import { serverApiBaseURL } from './_root'

// The base url for the events api.
const disruptionsApiBaseUrl = serverApiBaseURL + '/events'

/**
 * Fetch registered events for a single day.
 * @returns A promise that resolves to an array of event items.
 * @throws An error if the fetch request fails or the response is not of the expected format.
 */
export async function fetchEventItems(date: Date = new Date()): Promise<DisruptionItemList> {
  const timestamp = date.getTime()
  const offset = date.getTimezoneOffset()
  const apiRequestUrl = `${disruptionsApiBaseUrl}/${timestamp}?offset=${offset}`

  const response = await fetch(apiRequestUrl)

  if (!response.ok) {
    throw new Error('Failed to fetch event items')
  }

  // The response might not have failed, but still not be of the expected format.
  // So let's check that the response is an array.
  const rawJson = await response.json()
  if (!Array.isArray(rawJson)) {
    throw new Error('Fetched event items are not of expected format: ' + typeof rawJson)
  }

  return rawJson
}
