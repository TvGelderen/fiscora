import {
	getExpenseTypes,
	getIncomeTypes,
	getTransactionIntervals,
	getTransactionsYearInfo,
} from "$lib/api/transactions";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { getYear } from "$lib/api/utils";

export const load: PageServerLoad = async ({
	locals: { session, user },
	url,
}) => {
	if (!session?.accessToken || user === null) {
		throw redirect(302, "/login");
	}

	const year = getYear(url.searchParams);

	return {
		transactionIntervals: await getTransactionIntervals(
			session.accessToken,
		),
		incomeTypes: await getIncomeTypes(session.accessToken),
		expenseTypes: await getExpenseTypes(session.accessToken),
		yearInfo: await getTransactionsYearInfo(year, session.accessToken),
		demo: user.isDemo,
	};
};
