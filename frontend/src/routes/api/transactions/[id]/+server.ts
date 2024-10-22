import { type RequestHandler } from "@sveltejs/kit";
import { forbidden, toISOString } from "$lib";
import type { TransactionForm } from "../../../../ambient";
import { verifyForm } from "$lib/api/transactions";
import { authorizeFetch, authorizeFetchBody } from "$lib/api/fetch";

export const PUT: RequestHandler = async ({ locals: { session }, request, params: { id } }) => {
    if (!session) {
        return forbidden();
    }

    if (!id) {
        return new Response(null, {
            status: 400,
        });
    }

    let form: TransactionForm = await request.json();
    form = verifyForm(form);
    if (!form.errors.valid) {
        return new Response(JSON.stringify(form), {
            status: 400,
            headers: { "Content-Type": "application/json" },
        });
    }

    form.startDate = toISOString(form.startDate!);
    form.endDate = form.recurring ? toISOString(form.endDate!) : null;

    const response = await authorizeFetchBody(`transactions/${id}`, session.accessToken, "PUT", JSON.stringify(form));
    if (!response.ok) {
        return new Response(null, {
            status: 500,
        });
    }

    return response;
};

export const DELETE: RequestHandler = async ({ locals: { session }, params: { id } }) => {
    if (!session) {
        return forbidden();
    }

    if (!id) {
        return new Response(null, {
            status: 400,
        });
    }

    const response = await authorizeFetch(`transactions/${id}`, session.accessToken, "DELETE");
    if (!response.ok) {
        return new Response(null, {
            status: 500,
        });
    }

    return response;
};
