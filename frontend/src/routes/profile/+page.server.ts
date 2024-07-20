import type { PageServerLoadEvent } from "./$types";
import { redirect } from "@sveltejs/kit";
import { PUBLIC_API_ENDPOINT } from '$env/static/public';

export const load = async (event: PageServerLoadEvent) => {
    const session = event.cookies.get("Session");
    if (!session) {
        throw redirect(307, "/login");
    }

    const apiURL = `${PUBLIC_API_ENDPOINT}/api/test`;

    const response = await fetch(apiURL, {
        method: 'GET',
        credentials: 'include',
    });

    console.log(response);

    const data = await response.json();

    console.log(data);
}
