import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { getBudgets } from "$lib/api/budgets";

export const load: PageServerLoad = async ({ locals: { session, user } }) => {
    if (!session?.accessToken || user === null) {
        throw redirect(302, "/login");
    }

    const budgets = await getBudgets(session.accessToken);
    for (const budget of budgets) {
        console.log(budget)
    }

    return {
        budgets: await getBudgets(session.accessToken),
        demo: user.isDemo,
    };
};
