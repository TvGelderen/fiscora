import { forbidden } from "$lib";
import { authorizeFetch } from "$lib/api/fetch";
import { type RequestHandler } from "@sveltejs/kit";


export const DELETE: RequestHandler = async ({ locals: { session }, params: { id, expense_id } }) => {
    if (!session) {
        return forbidden();
    }

    if (!id || !expense_id) {
        return new Response(null, {
            status: 400,
        })
    }

    const response = await authorizeFetch(`budgets/${id}/expenses/${expense_id}`, session.accessToken, "DELETE");
    if (response.ok) {
        return response;
    }

    console.log(response)

    return new Response(null, {
        status: response.status,
    })
}

