import { authorizeFetch } from "$lib";
import type { Transaction, TransactionMonthInfoResponse } from "../../ambient";

export async function getTransactionIntervals(
    accessToken: string,
): Promise<string[]> {
    const response = await authorizeFetch(
        "transactions/intervals",
        accessToken,
    );
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as string[];
}

export async function getIncomeTypes(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch(
        "transactions/income-types",
        accessToken,
    );
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as string[];
}

export async function getExpenseTypes(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch(
        "transactions/expense-types",
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
): Promise<TransactionMonthInfoResponse> {
    const url = `transactions/month-info?month=${month}&year=${year}`;
    const response = await authorizeFetch(url, accessToken);
    if (!response.ok) {
        return {
            income: 0,
            expense: 0,
        };
    }

    return (await response.json()) as TransactionMonthInfoResponse;
}
