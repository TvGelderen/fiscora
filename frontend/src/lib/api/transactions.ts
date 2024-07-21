import { authorizeFetch } from "$lib";

export async function getTransactionIntervals(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch("transactions/intervals", accessToken);
    if (!response.ok) {
        return [];
    }

    const data = await response.json() as string[];

    return data;
}

export async function getIncomeTypes(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch("transactions/income-types", accessToken);
    if (!response.ok) {
        return [];
    }

    const data = await response.json() as string[];

    return data;
}

export async function getExpenseTypes(accessToken: string): Promise<string[]> {
    const response = await authorizeFetch("transactions/expense-types", accessToken);
    if (!response.ok) {
        return [];
    }

    const data = await response.json() as string[];

    return data;
}
