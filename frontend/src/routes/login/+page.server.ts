import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = ({ locals: { user } }) => {
	if (user) {
		throw redirect(302, "/");
	}
};
