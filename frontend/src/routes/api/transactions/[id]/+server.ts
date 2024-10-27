import { forbidden } from "$lib";
import { authorizeFetch } from "$lib/api/fetch";
import { type RequestHandler } from "@sveltejs/kit";

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
