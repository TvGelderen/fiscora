import { getExpenseTypes, getIncomeTypes, getTransactionIntervals } from "$lib/api/transactions";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals: { session } }) => {
    if (!session?.accessToken) {
        throw redirect(302, "/login")
    }

    return {
        transactionIntervals: await getTransactionIntervals(session.accessToken),
        incomeTypes: await getIncomeTypes(session.accessToken),
        expenseTypes: await getExpenseTypes(session.accessToken),
    }
}