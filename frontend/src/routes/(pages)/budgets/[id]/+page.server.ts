import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals: { session, user } }) => {
	if (!session?.accessToken || user === null) {
		throw redirect(302, "/login");
	}
};
