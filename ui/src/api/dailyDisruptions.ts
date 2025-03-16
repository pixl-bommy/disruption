import type { DisruptionItem, DisruptionItemList } from '@/types/disruption'

import { serverApiBaseURL } from './_root'
import {
  DailyEventItemStatus,
  type DailyEventItem,
  type DailyEventItemList,
  type DailyEventItemUnsettled,
} from '@/types/dailyEvent'

// The base url for the events api.
const disruptionsApiBaseUrl = serverApiBaseURL + '/events'

/**
 * Fetch registered events for a single day.
 * @returns A promise that resolves to an array of event items.
 * @throws An error if the fetch request fails or the response is not of the expected format.
 */
export async function fetchEventItems(date: Date = new Date()): Promise<DailyEventItemList> {
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
    .map<DailyEventItem>((entry) => ({
      id: entry['uid'],
      disruptionId: entry['disruptionId'],
      name: entry['disruptionName'],
      status: DailyEventItemStatus.Completed,
      iconName: entry['icon'] ?? 'default',
      color: entry['color'] ?? '#000000',
    }))
    .filter((entry) => entry.id && entry.disruptionId && entry.name)
}

export async function fetchCreateDailyEventItem(
  newItem: DailyEventItemUnsettled,
): Promise<DailyEventItemList> {
  const timeoutJob = new Promise((resolve) => {
    setTimeout(() => resolve(null), 1000)
  })

  const fetchingJob = fetch(`${disruptionsApiBaseUrl}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ disruptionId: newItem.disruptionId }),
  })

  const [response] = await Promise.all([fetchingJob, timeoutJob])

  if (!response.ok) {
    throw new Error('Failed to create daily event item')
  }

  return fetchEventItems()
}
