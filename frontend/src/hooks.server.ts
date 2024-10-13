import { authorizeFetch } from "$lib/api/fetch";
import type { User } from "$lib/types";
import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
    const accessToken = event.cookies.get("AccessToken");
    if (!accessToken) {
        return resolve(event);
    }

    event.locals.session = { accessToken };

    const response = await authorizeFetch("users/me", accessToken);
    if (response.ok) {
        const user = (await response.json()) as User;
        user.isDemo = user.username === "demo";
        event.locals.user = user;

        if (user.username === "demo" && event.request.method !== "GET") {
            return new Response(null, {
                status: 401,
            });
        }
    }

    return resolve(event);
};
