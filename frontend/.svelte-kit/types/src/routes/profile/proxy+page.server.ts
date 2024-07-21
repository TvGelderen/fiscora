// @ts-nocheck
import { redirect } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"

export const load = ({ locals: { user } }: Parameters<PageServerLoad>[0]) => {
    if (!user) {
        throw redirect(302, "/login");
    }

    return {
        user
    }
}
