import { type RequestHandler } from "@sveltejs/kit";
import { forbidden } from "$lib";
import { authorizeFetch } from "$lib/api/fetch";

export const GET: RequestHandler = async ({ locals: { session }, url }) => {
    if (!session) {
        return forbidden();
    }

    const startDate = url.searchParams.get("startDate");
    const endDate = url.searchParams.get("endDate");
    if (startDate === null || endDate === null) {
        return new Response("Invalid date format", {
            status: 400,
        });
    }

    const fetchUrl = `transactions/unassigned?startDate=${startDate}&endDate=${endDate}`;
    const response = await authorizeFetch(fetchUrl, session.accessToken);
    if (!response.ok) {
        return new Response("Something went wrong", {
            status: response.status,
        });
    }

    return response;
};
