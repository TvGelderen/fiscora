import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = ({ locals: { user } }) => {
    return {
        loggedIn: user !== null
    }
};
