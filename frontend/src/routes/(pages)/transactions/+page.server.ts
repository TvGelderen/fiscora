import {
    getExpenseTypes,
    getIncomeTypes,
    getTransactionIntervals,
    getTransactions,
    getTransactionsYearInfo,
} from "$lib/api/transactions";
import { fail, redirect, type RequestEvent } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { getMonth, getYear } from "$lib/api/utils";
import { message, superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { transactionFormSchema } from "./transactionFormSchema";
import { authorizeFetchBody } from "$lib/api/fetch";

export const load: PageServerLoad = async ({
    locals: { session, user },
    url,
}) => {
    if (!session?.accessToken || user === null) {
        throw redirect(302, "/login");
    }

    const year = getYear(url.searchParams);
    const month = getMonth(url.searchParams);

    return {
        transactions: await getTransactions(month, year, session.accessToken),
        transactionForm: await superValidate(zod(transactionFormSchema)),
        transactionIntervals: await getTransactionIntervals(
            session.accessToken,
        ),
        incomeTypes: await getIncomeTypes(session.accessToken),
        expenseTypes: await getExpenseTypes(session.accessToken),
        yearInfo: await getTransactionsYearInfo(year, session.accessToken),
        demo: user.isDemo,
    };
};

export const actions = {
    default: async (event: RequestEvent) => {
        const session = event.locals.session;
        if (!session) {
            throw redirect(302, "/login");
        }

        const form = await superValidate(event, zod(transactionFormSchema));
        console.log(form)
        if (!form.valid) {
            return fail(400, { form });
        }

        try {
            let response: Response;
            if (form.data.id === -1) {
                response = await authorizeFetchBody("transactions", session.accessToken, "POST", JSON.stringify(form.data));
            } else {
                response = await authorizeFetchBody(`transactions/${form.data.id}`, session.accessToken, "PUT", JSON.stringify(form.data));
            }

            if (!response.ok) {
                console.error(await response.text());
                return fail(500);
            }
        } catch (err) {
            console.error(err);
            return fail(500);
        }

        return message(form, "Success");
    }
}

