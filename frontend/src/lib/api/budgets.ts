import { authorizeFetch } from "$lib/index";
import type { Budget, BudgetForm, BudgetFormErrors } from "../../ambient";
import { validNumber, validString } from "./utils";

export async function getBudgets(accessToken: string): Promise<Budget[]> {
    const response = await authorizeFetch("budgets", accessToken);
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as Budget[];
}

export function verifyForm(form: BudgetForm): BudgetFormErrors {
    const errors: BudgetFormErrors = {
        valid: true,
        name: null,
        description: null,
        amount: null,
    };

    if (!validString(form.name)) {
        errors.name = "Description is required";
        errors.valid = false;
    }
    if (!validString(form.description)) {
        errors.description = "Description is required";
        errors.valid = false;
    }
    if (!validNumber(form.amount)) {
        errors.amount = "Amount must be a number";
        errors.valid = false;
    }

    return errors;
}
