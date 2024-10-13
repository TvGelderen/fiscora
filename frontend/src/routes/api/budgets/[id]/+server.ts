import { forbidden } from "$lib";
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
        })
    }

    let form: BudgetForm = await request.json();
    form = verifyForm(form);
    if (!form.errors.valid) {
        return new Response(JSON.stringify(form), {
            status: 400,
            headers: { "Content-Type": "application/json" },
        });
    }

    const response = await authorizeFetchBody(
        `budgets/${id}`,
        session.accessToken,
        "PUT",
        JSON.stringify(form));
    if (response.ok) {
        return response;
    }

    return new Response(null, {
        status: response.status,
    })
}

export const DELETE: RequestHandler = async ({ locals: { session }, params }) => {
    if (!session) {
        return forbidden();
    }

    const id = params.id;
    if (!id) {
        return new Response(null, {
            status: 400,
        })
    }

    const response = await authorizeFetch(`budgets/${id}`, session.accessToken, "DELETE");
    if (response.ok) {
        return response;
    }

    return new Response(null, {
        status: response.status,
    })
}

