import { authorizeFetch } from "$lib/api/fetch";
import { error } from "@sveltejs/kit";
import type { Budget, BudgetForm } from "../../ambient";
import { validNumber, validString } from "./utils";

export async function getBudgets(accessToken: string): Promise<Budget[]> {
	const response = await authorizeFetch("budgets", accessToken);
	if (!response.ok) {
		return [];
	}

	return (await response.json()) as Budget[];
}

export async function getBudget(
	accessToken: string,
	id: string,
): Promise<Budget> {
	const response = await authorizeFetch(`budgets/${id}`, accessToken);
	if (!response.ok) {
		throw error(response.status);
	}

	return (await response.json()) as Budget;
}

export function verifyForm(form: BudgetForm): BudgetForm {
	form.errors = {
		valid: true,
		name: null,
		description: null,
		amount: null,
	};

	if (!validString(form.name)) {
		form.errors.name = "Name is required";
		form.errors.valid = false;
	}
	if (!validString(form.description)) {
		form.errors.description = "Description is required";
		form.errors.valid = false;
	}
	if (!validNumber(form.amount)) {
		form.errors.amount = "Amount must be a number";
		form.errors.valid = false;
	} else if (form.amount === 0) {
		form.errors.amount = "Amount must be greater than 0";
		form.errors.valid = false;
	}

	for (const expense of form.expenses) {
		expense.errors = {
			valid: true,
			name: null,
			allocatedAmount: null,
		};

		if (!validString(expense.name)) {
			expense.errors.name = "Name is required";
			expense.errors.valid = false;
		}
		if (!validNumber(expense.allocatedAmount)) {
			expense.errors.allocatedAmount = "Amount must be a number";
			expense.errors.valid = false;
		} else if (expense.allocatedAmount === 0) {
			expense.errors.allocatedAmount = "Amount must be greater than 0";
			expense.errors.valid = false;
		}

		if (!expense.errors.valid) {
			form.errors.valid = false;
		}
	}

	return form;
}
