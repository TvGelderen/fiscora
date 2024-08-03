import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { getTransactionsYearInfo } from "$lib/api/transactions";

export const load: PageServerLoad = async ({
    locals: { session, user },
    url,
}) => {
    if (!session?.accessToken || user === null) {
        throw redirect(302, "/login");
    }

    let year = Number.parseInt(url.searchParams.get("year") ?? "0");
    if (year === 0) {
        year = 2024;
    }

    return {
        yearInfo: await getTransactionsYearInfo(year, session.accessToken),
    };
};
