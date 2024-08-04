import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import {
	getExpensesPerType,
	getTransactionsYearInfo,
} from "$lib/api/transactions";
import { getCurrentMonthNumber } from "$lib";

export const load: PageServerLoad = async ({
	locals: { session, user },
	url,
}) => {
	if (!session?.accessToken || user === null) {
		throw redirect(302, "/login");
	}

	let month = Number.parseInt(url.searchParams.get("month") ?? "0");
	if (month === 0) {
		month = getCurrentMonthNumber();
	}
	let year = Number.parseInt(url.searchParams.get("year") ?? "0");
	if (year === 0) {
		year = 2024;
	}

	return {
		yearInfo: await getTransactionsYearInfo(year, session.accessToken),
		expenseInfo: await getExpensesPerType(month, year, session.accessToken),
	};
};
