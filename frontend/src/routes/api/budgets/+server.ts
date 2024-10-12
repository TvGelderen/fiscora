import {
    authorizeFetch,
    authorizeFetchBody,
    forbidden
} from "$lib";
import { verifyForm } from "$lib/api/budgets";
import { type RequestHandler } from "@sveltejs/kit";
import type { BudgetForm } from "../../../ambient";

export const GET: RequestHandler = async ({ locals: { session } }) => {
    if (!session) {
        return forbidden();
    }

    const response = await authorizeFetch("budgets", session.accessToken);
    if (response.ok) {
        return response;
    }

    return new Response("Something went wrong", {
        status: response.status,
    });
};

export const POST: RequestHandler = async ({
    locals: { session },
    request,
}) => {
    if (!session) {
        return forbidden();
    }

    const form: BudgetForm = await request.json();
    const errors = verifyForm(form);
    if (!errors.valid) {
        console.log(errors)
        form.errors = errors;
        return new Response(JSON.stringify(form), {
            status: 400,
            headers: { "Content-Type": "application/json" },
        });
    }

    const response = await authorizeFetchBody(
        "budgets",
        session.accessToken,
        "POST",
        JSON.stringify(form),
    );
    if (response.ok) {
        return response;
    }

    return new Response(JSON.stringify(form), {
        status: 500,
        headers: { "Content-Type": "application/json" },
    });
};
