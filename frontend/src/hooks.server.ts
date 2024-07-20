import { lucia } from "$lib/server/auth";
import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
    const sessionId = event.cookies.get(lucia.sessionCookieName);
    if (!sessionId) {
        event.locals.user = null
        event.locals.session = null;
        return resolve(event);
    }

    let sessionCookie = lucia.createBlankSessionCookie();

    const { session, user } = await lucia.validateSession(sessionId);
    if (session && session.fresh) {
        sessionCookie = lucia.createSessionCookie(session.id);
    }

    event.cookies.set(sessionCookie.name, sessionCookie.value, {
        path: ".",
        ...sessionCookie.attributes
    });

    event.locals.user = user;
    event.locals.session = session;

    return resolve(event);
}
