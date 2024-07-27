// @ts-nocheck
import {
	getExpenseTypes,
	getIncomeTypes,
	getTransactionIntervals,
} from "$lib/api/transactions";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load = async ({ locals: { session } }: Parameters<PageServerLoad>[0]) => {
	if (!session?.accessToken) {
		throw redirect(302, "/login");
	}

	return {
		transactionIntervals: await getTransactionIntervals(
			session.accessToken,
		),
		incomeTypes: await getIncomeTypes(session.accessToken),
		expenseTypes: await getExpenseTypes(session.accessToken),
	};
};
