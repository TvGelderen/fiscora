import { forbidden, toISOString } from "$lib";
import { verifyForm } from "$lib/api/budgets";
import { authorizeFetch, authorizeFetchBody } from "$lib/api/fetch";
import { type RequestHandler } from "@sveltejs/kit";
import type { BudgetForm } from "../../../../ambient";

export const PUT: RequestHandler = async ({ locals: { session }, request, params: { id } }) => {
    if (!session) {
        return forbidden();
    }

    if (!id) {
        return new Response(null, {
            status: 400,
        });
    }

    let form: BudgetForm = await request.json();
    form = verifyForm(form);
    if (!form.errors.valid) {
        return new Response(JSON.stringify(form), {
            status: 400,
            headers: { "Content-Type": "application/json" },
        });
    }

    form.startDate = toISOString(form.startDate!);
    form.endDate = toISOString(form.endDate!);

    const response = await authorizeFetchBody(`budgets/${id}`, session.accessToken, "PUT", JSON.stringify(form));
    if (!response.ok) {
        return new Response(null, {
            status: response.status,
        });
    }

    return response;
};

export const DELETE: RequestHandler = async ({ locals: { session }, params }) => {
    if (!session) {
        return forbidden();
    }

    const id = params.id;
    if (!id) {
        return new Response(null, {
            status: 400,
        });
    }

    const response = await authorizeFetch(`budgets/${id}`, session.accessToken, "DELETE");
    if (!response.ok) {
        return new Response(null, {
            status: response.status,
        });
    }

    return response;
};
