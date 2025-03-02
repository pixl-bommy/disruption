import { serverApiBaseURL } from './_root'

export async function saveNewDisruptionItem(payload: { name: string; description: string }) {
  const apiUrl = serverApiBaseURL + '/disruptions'
  const response = await fetch(apiUrl, {
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

export async function fetchDisruptionItems() {
  const apiUrl = serverApiBaseURL + '/disruptions'
  const response = await fetch(apiUrl)

  if (!response.ok) {
    throw new Error('Failed to fetch disruption items')
  }

  const rawJson = await response.json()

  if (!Array.isArray(rawJson)) {
    throw new Error('Fetched disruption items are not of expected format: ' + typeof rawJson)
  }

  return rawJson.map((entry) => ({
    id: entry['uid'],
    name: entry['name'],
    description: entry['description'],
  }))
}
