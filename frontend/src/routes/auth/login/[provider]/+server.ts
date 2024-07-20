import { generateCodeVerifier, generateState } from "arctic";
import type { RequestEvent } from "./$types";
import { google } from "$lib/server/auth";
import { dev } from "$app/environment";
import { redirect } from "@sveltejs/kit";

export async function GET(event: RequestEvent): Promise<Response> {
    const state = generateState();
    const codeVerifier = generateCodeVerifier();
    const url = await google.createAuthorizationURL(state, codeVerifier);

    event.cookies.set("state", state, {
        path: "/",
        secure: !dev,
        httpOnly: true,
        maxAge: 60 * 60,
        sameSite: "lax"
    })

    event.cookies.set("code_verifier", codeVerifier, {
        path: "/",
        secure: !dev,
        httpOnly: true,
        maxAge: 60 * 60,
        sameSite: "lax"
    })

    redirect(302, url.toString());
}
