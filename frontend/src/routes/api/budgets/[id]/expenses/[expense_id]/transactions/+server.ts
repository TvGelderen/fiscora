import { forbidden } from "$lib";
import { authorizeFetchBody } from "$lib/api/fetch";
import { type RequestHandler } from "@sveltejs/kit";


export const POST: RequestHandler = async ({ locals: { session }, params: { id, expense_id }, request }) => {
    if (!session) {
        return forbidden();
    }

    if (!id || !expense_id) {
        return new Response(null, {
            status: 400,
        })
    }

    return await authorizeFetchBody(
        `budgets/${id}/expenses/${expense_id}/transactions`,
        session.accessToken,
        "POST",
        JSON.stringify(await request.json()));
}

