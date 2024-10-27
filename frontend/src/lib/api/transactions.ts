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
