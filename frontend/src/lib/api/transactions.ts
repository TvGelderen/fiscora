import type { Transaction, TransactionForm, TransactionMonthInfo } from "../../ambient";
import { authorizeFetch } from "./fetch";
import { validDate, validNumber, validString } from "./utils";

export async function getTransactionIntervals(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch("transactions/types/intervals", accessToken);
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as string[];
}

export async function getIncomeTypes(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch("transactions/types/income", accessToken);
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as string[];
}

export async function getExpenseTypes(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch("transactions/types/expense", accessToken);
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as string[];
}

export async function getTransactions(
    month: number,
    year: number,
    accessToken: string,
    income?: string,
): Promise<Transaction[]> {
    const url = `transactions?month=${month}&year=${year}${income === undefined ? "" : `&income=${income}`}`;
    const response = await authorizeFetch(url, accessToken);
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as Transaction[];
}

export async function getUnassignedTransactions(
    startDate: string,
    endDate: string,
    accessToken: string,
): Promise<Transaction[]> {
    const url = `transactions/unassigned?startDate=${startDate}&endDate=${endDate}`;
    const response = await authorizeFetch(url, accessToken);
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as Transaction[];
}

export async function getTransactionsMonthInfo(
    month: number,
    year: number,
    accessToken: string,
): Promise<TransactionMonthInfo> {
    const url = `transactions/summary/month?month=${month}&year=${year}`;
    const response = await authorizeFetch(url, accessToken);
    if (!response.ok) {
        return {
            income: 0,
            expense: 0,
        };
    }

    return (await response.json()) as TransactionMonthInfo;
}

export async function getTransactionsYearInfo(
    year: number,
    accessToken: string,
): Promise<Map<string, TransactionMonthInfo>> {
    const url = `transactions/summary/year?year=${year}`;
    const response = await authorizeFetch(url, accessToken);
    if (!response.ok) {
        return new Map<string, TransactionMonthInfo>();
    }

    const json = await response.json();

    return new Map(Object.entries<TransactionMonthInfo>(json));
}

export async function getTransactionsYearInfoPerType(
    year: number,
    income: boolean,
    accessToken: string,
): Promise<Map<string, number>> {
    const url = `transactions/summary/year/type?year=${year}&income=${income}`;
    const response = await authorizeFetch(url, accessToken);
    if (!response.ok) {
        return new Map<string, number>();
    }

    return (await response.json()) as Map<string, number>;
}

export async function getTransactionsPerType(
    month: number,
    year: number,
    income: boolean,
    accessToken: string,
): Promise<Map<string, number>> {
    const url = `transactions/summary/month/type?month=${month}&year=${year}&income=${income}`;
    const response = await authorizeFetch(url, accessToken);
    if (!response.ok) {
        return new Map<string, number>();
    }

    return (await response.json()) as Map<string, number>;
}

export function verifyForm(form: TransactionForm): TransactionForm {
    form.errors = {
        valid: true,
        amount: null,
        description: null,
        startDate: null,
        endDate: null,
        interval: null,
        daysInterval: null,
        type: null,
    };

    if (!validNumber(form.amount)) {
        form.errors.amount = "Amount must be a number";
        form.errors.valid = false;
    } else if (form.amount === 0) {
        form.errors.amount = "Amount must not be 0";
        form.errors.valid = false;
    }
    if (!validString(form.description)) {
        form.errors.description = "Description is required";
        form.errors.valid = false;
    }
    if (!validDate(form.startDate)) {
        form.errors.startDate = "Start date must be a valid date";
        form.errors.valid = false;
    }
    if (form.recurring) {
        if (!validDate(form.endDate)) {
            form.errors.endDate = "End date must be a valid date or null";
            form.errors.valid = false;
        }
        if (!validString(form.interval)) {
            form.errors.interval = "Recurring interval is required when a transaction recurring";
            form.errors.valid = false;
        }
        if (form.interval === "Other" && !validNumber(form.daysInterval)) {
            form.errors.daysInterval = "Interval in days should be set";
            form.errors.valid = false;
        }
    }
    if (!validString(form.type)) {
        form.errors.type = "Transaction type is required";
        form.errors.valid = false;
    }

    return form;
}
