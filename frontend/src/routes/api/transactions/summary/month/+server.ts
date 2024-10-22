import { authorizeFetch } from "$lib/api/fetch";
import type { RequestHandler } from "./$types";

export const GET: RequestHandler = async ({ locals: { session }, url }) => {
    if (!session) {
        return new Response("Forbidden", {
            status: 403,
        });
    }

    const month = url.searchParams.get("month");
    const year = url.searchParams.get("year");
    const fetchUrl = `transactions/summary/month?month=${month}&year=${year}`;
    const response = await authorizeFetch(fetchUrl, session?.accessToken);
    if (response.ok) {
        return response;
    }

    return new Response("Something went wrong", {
        status: response.status,
    });
};
