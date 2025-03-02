import { serverApiBaseURL } from './_root'

export async function getUserMe() {
  const response = await fetch(`${serverApiBaseURL}/users/me`).then((res) => res.json())
  return response
}
