import { forbidden } from "$lib";
import { authorizeFetch } from "$lib/api/fetch";
import { redirect, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ locals: { session } }) => {
    if (!session) {
        return forbidden();
    }

    await authorizeFetch("auth/logout", session.accessToken)

    return redirect(302, "/");
}
