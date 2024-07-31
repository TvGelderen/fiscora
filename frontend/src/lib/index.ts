import { PUBLIC_API_URL } from "$env/static/public";

export function getApiUrl(endpoint: string): string {
    return `${PUBLIC_API_URL}/${endpoint}`;
}

export function getFormattedDate(date: Date): string {
    return new Date(date).toLocaleDateString("default", {
        weekday: "short",
        year: "numeric",
        month: "long",
        day: "numeric",
    });
}

export function getFormattedDateShort(date: Date): string {
    return new Date(date).toLocaleDateString("default", {
        month: "short",
        day: "numeric",
    });
}

export function getFormDate(date: Date): string {
    return new Date(date).toISOString().split("T")[0];
}

export function toISOString(date: string): string {
    return new Date(date).toISOString();
}

export function getFormattedAmount(amount: number) {
    return amount.toLocaleString(undefined, {
        minimumFractionDigits: 2,
        maximumFractionDigits: 2,
    });
}

export const forbidden = () => new Response("Forbidden", {
    status: 403,
});

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
