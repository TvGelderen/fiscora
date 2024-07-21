import { PUBLIC_API_URL } from "$env/static/public";

export function getApiUrl(endpoint: string): string {
    return `${PUBLIC_API_URL}${endpoint}`;
}

export async function authorizeFetch(endpoint: string, accessToken: string, method: string = "GET") {
    return await fetch(getApiUrl(endpoint), {
        method: method,
        headers: {
            Authorization: `Bearer ${accessToken}`
        }
    })
}
