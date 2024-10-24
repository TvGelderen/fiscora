import { getTransactionsYearInfo } from "$lib/api/transactions";
import { getYear } from "$lib/api/utils";
import type { RequestHandler } from "./$types";

export const GET: RequestHandler = async ({ locals: { session }, url }) => {
    if (!session) {
        return new Response("Forbidden", {
            status: 403,
        });
    }

    const year = getYear(url.searchParams);
    const map = await getTransactionsYearInfo(year, session.accessToken);
    if (map.size === 0) {
        return new Response("Something went wrong", {
            status: 500,
        });
    }

    return new Response(JSON.stringify(map), {
        status: 200
    });
};
