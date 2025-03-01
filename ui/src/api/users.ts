// WORKAROUND: Setting the serverApiBaseURL to 'http://localhost:3000/api/v1' for now
// This will be replaced with the actual URL when the app is deployed and is used
// to make requests to the server API using vite dev server.
const serverApiBaseURL = 'http://localhost:3000/api/v1'

export async function getUserMe() {
  const response = await fetch(`${serverApiBaseURL}/users/me`).then((res) => res.json())
  return response
}
