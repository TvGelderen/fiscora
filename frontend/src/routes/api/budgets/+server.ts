import {
    forbidden,
    toISOString
} from "$lib";
import { verifyForm } from "$lib/api/budgets";
import { type RequestHandler } from "@sveltejs/kit";
import type { BudgetForm } from "../../../ambient";
import { authorizeFetch, authorizeFetchBody } from "$lib/api/fetch";

export const GET: RequestHandler = async ({ locals: { session } }) => {
    if (!session) {
        return forbidden();
    }

    const response = await authorizeFetch("budgets", session.accessToken);
    if (!response.ok) {
        return new Response("Something went wrong", {
            status: response.status,
        });
    }

    return response;
};

export const POST: RequestHandler = async ({
    locals: { session },
    request,
}) => {
    if (!session) {
        return forbidden();
    }

    let form: BudgetForm = await request.json();
    form = verifyForm(form);
    if (!form.errors.valid) {
        return new Response(JSON.stringify(form), {
            status: 400,
            headers: { "Content-Type": "application/json" },
        });
    }

    form.startDate = toISOString(form.startDate);
    form.endDate = toISOString(form.endDate);

    const response = await authorizeFetchBody(
        "budgets",
        session.accessToken,
        "POST",
        JSON.stringify(form),
    );
    if (!response.ok) {
        return new Response(JSON.stringify(form), {
            status: 500,
            headers: { "Content-Type": "application/json" },
        });
    }

    return response;
};
