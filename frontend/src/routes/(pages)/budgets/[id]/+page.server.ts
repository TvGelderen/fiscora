import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { getBudget } from "$lib/api/budgets";

export const load: PageServerLoad = async ({
	locals: { session, user },
	params: { id },
}) => {
	if (!session?.accessToken || user === null) {
		throw redirect(302, "/login");
	}

	return {
		budget: await getBudget(session.accessToken, id),
	};
};
