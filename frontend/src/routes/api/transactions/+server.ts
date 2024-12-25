import { forbidden } from "$lib";
import { authorizeFetch, authorizeFetchBody } from "$lib/api/fetch";
import { type RequestHandler } from "@sveltejs/kit";
import type { TransactionForm } from "../../../ambient";
import { verifyForm } from "$lib/api/transactions";

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

export const POST: RequestHandler = async ({ locals: { session }, request }) => {
    if (!session) {
        return forbidden();
    }

    let form: TransactionForm = await request.json();
    form = verifyForm(form);
    if (!form.errors.valid) {
        return new Response(JSON.stringify(form), {
            status: 400,
            headers: { "Content-Type": "application/json" },
        });
    }

    const response = await authorizeFetchBody("transactions", session.accessToken, "POST", JSON.stringify(form));
    if (!response.ok) {
        return new Response(JSON.stringify(form), {
            status: 500,
            headers: { "Content-Type": "application/json" },
        });
    }

    return response;
};
