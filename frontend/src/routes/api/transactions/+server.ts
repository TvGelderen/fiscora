import { forbidden } from "$lib";
import { authorizeFetch } from "$lib/api/fetch";
import { type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ locals: { session }, url }) => {
    if (!session) {
        return forbidden();
    }

    const month = url.searchParams.get("month");
    const year = url.searchParams.get("year");
    const income = url.searchParams.get("income");
    const fetchUrl = `transactions?month=${month}&year=${year}&take=100${income === null ? "" : `&income=${income}`}`;
    const response = await authorizeFetch(fetchUrl, session.accessToken);
    if (!response.ok) {
        return new Response("Something went wrong", {
            status: response.status,
        });
    }

    return response;
};
