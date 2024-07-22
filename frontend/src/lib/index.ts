import { PUBLIC_API_URL } from "$env/static/public";

export function getApiUrl(endpoint: string): string {
    return `${PUBLIC_API_URL}/${endpoint}`;
}

export async function authorizeFetch(endpoint: string, accessToken: string, method: string = "GET") {
    return await fetch(getApiUrl(endpoint), {
        method: method,
        headers: {
            'Authorization': `Bearer ${accessToken}`
        }
    });
}

export async function authorizePost(endpoint: string, accessToken: string, data: string) {
    return await fetch(getApiUrl(endpoint), {
        method: "POST",
        headers: {
            'Authorization': `Bearer ${accessToken}`,
            'Content-Type': 'application/json'
        },
        body: data
    });
}
