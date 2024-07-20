import { google } from "$lib/server/auth";
import { OAuth2RequestError } from "arctic";
import type { RequestEvent } from "./$types";
import { API_URL } from "$env/static/private";

export async function GET(event: RequestEvent): Promise<Response> {
    const code = event.url.searchParams.get("code");
    const state = event.url.searchParams.get("state");
    const storedState = event.cookies.get("state");
    const codeVerifier = event.cookies.get("code_verifier");

    if (!code || !state || !storedState || !codeVerifier || state !== storedState) {
        return new Response(null, { status: 400 });
    }

    try {
        const tokens = await google.validateAuthorizationCode(code, codeVerifier);
        const response = await fetch("https://openidconnect.googleapis.com/v1/userinfo", {
            headers: {
                Authorization: `Bearer ${tokens.accessToken}`
            }
        });

        const user = await response.json() as GoogleUser;

        const apiResponse = await fetch(`${API_URL}/auth/callback/google`, {
            method: "POST",
            body: JSON.stringify(user)
        });
        const data = apiResponse.json();

        console.log(data);

        // const session = await lucia.createSession(userId, {});
        // const sessionCookie = lucia.createSessionCookie(session.id);

        // event.cookies.set(sessionCookie.name, sessionCookie.value, {
        //     path: ".",
        //     ...sessionCookie.attributes
        // })

        return new Response(null, { status: 302, headers: { Location: "/" } });
    } catch (err) {
        console.log(err);

        if (err instanceof OAuth2RequestError) {
            return new Response(null, { status: 400 });
        }

        return new Response(null, { status: 500 });
    }
}

type GoogleUser = {
    sub: string
    username: string;
    email: string;
    avatar_url: string;
}
