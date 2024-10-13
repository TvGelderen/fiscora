import { API_URL } from "$env/static/private";

export function getApiUrl(endpoint: string): string {
    return `${API_URL}/${endpoint}`;
}

export async function authorizeFetch(
    endpoint: string,
    accessToken: string,
    method: string = "GET",
) {
    return await fetch(getApiUrl(endpoint), {
        method: method,
        headers: {
            Authorization: `Bearer ${accessToken}`,
        },
    });
}

export async function authorizeFetchBody(
    endpoint: string,
    accessToken: string,
    method: string,
    data: string,
) {
    return await fetch(getApiUrl(endpoint), {
        method: method,
        headers: {
            Authorization: `Bearer ${accessToken}`,
            "Content-Type": "application/json",
        },
        body: data,
    });
}

