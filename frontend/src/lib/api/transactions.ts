import { authorizeFetch } from "$lib";
import type {
    Transaction,
    TransactionForm,
    TransactionFormErrors,
    TransactionMonthInfo,
} from "../../ambient";
import { validDate, validNumber, validString } from "./utils";

export async function getTransactionIntervals(
    accessToken: string,
): Promise<string[]> {
    const response = await authorizeFetch(
        "transactions/types/intervals",
        accessToken,
    );
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as string[];
}

export async function getIncomeTypes(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch(
        "transactions/types/income",
        accessToken,
    );
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as string[];
}

export async function getExpenseTypes(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch(
        "transactions/types/expense",
        accessToken,
    );
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as string[];
}

export async function getTransactions(
    month: number,
    year: number,
    income: string | null,
    accessToken: string,
): Promise<Transaction[]> {
    const url = `transactions?month=${month}&year=${year}${income === null ? "" : `&income=${income}`}`;
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
): Promise<Map<number, TransactionMonthInfo>> {
    const url = `transactions/summary/year?year=${year}`;
    const response = await authorizeFetch(url, accessToken);
    if (!response.ok) {
        return new Map<number, TransactionMonthInfo>();
    }

    return (await response.json()) as Map<number, TransactionMonthInfo>;
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

export function verifyForm(form: TransactionForm): TransactionFormErrors {
    const errors: TransactionFormErrors = {
        amount: null,
        description: null,
        startDate: null,
        endDate: null,
        interval: null,
        daysInterval: null,
        type: null,
    };

    if (!validNumber(form.amount)) {
        errors.amount = "Amount must be a number";
    }
    if (!validString(form.description)) {
        errors.description = "Description is required";
    }
    if (!validDate(form.startDate)) {
        errors.startDate = "Start date must be a valid date";
    }
    if (form.recurring) {
        if (!validDate(form.endDate)) {
            errors.endDate = "End date must be a valid date or null";
        }
        if (!validString(form.interval)) {
            errors.interval =
                "Recurring interval is required when a transaction recurring";
        }
        if (form.interval === "Other" && !validNumber(form.daysInterval)) {
            errors.daysInterval = "Interval in days should be set";
        }
    }
    if (!validString(form.type)) {
        errors.type = "Transaction type is required";
    }

    return errors;
}
