import type { UserInfo } from '@/types/users'
import { serverApiBaseURL } from './_root'

/**
 * Get the current logged in user.
 * @returns A promise that resolves to the user information.
 * @throws An error if the fetch request fails or the response is not of the expected format.
 */
export async function getUserMe() {
  const response = await fetch(`${serverApiBaseURL}/users/me`)

  if (!response.ok) {
    throw new Error('Failed to fetch user information')
  }

  const rawJson = await response.json()

  if (!rawJson?.uid || !rawJson?.name) {
    throw new Error('Fetched user information is not of expected format')
  }

  return rawJson as UserInfo
}
