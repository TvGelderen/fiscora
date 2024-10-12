export type Transaction = {
	id: number;
	amount: number;
	description: string;
	date: Date;
	startDate: Date | null;
	endDate: Date | null;
	recurring: boolean;
	interval: string | null;
	daysInterval: number | null;
	type: string | null;
	created: Date;
	updated: Date;
};

export type TransactionForm = {
	amount: number;
	description: string;
	startDate: string | null;
	endDate: string | null;
	recurring: boolean;
	interval: string | null;
	daysInterval: number | null;
	type: string | null;
	errors: TransactionFormErrors;
};

export type TransactionFormErrors = {
	amount: string | null;
	description: string | null;
	startDate: string | null;
	endDate: string | null;
	interval: string | null;
	daysInterval: string | null;
	type: string | null;
};

export type TransactionMonthInfo = {
	income: number;
	expense: number;
};

export type Budget = {
	id: number;
	name: string;
	description: string;
	amount: number;
	startDate: Date | null;
	endDate: Date | null;
	type: string;
	created: Date;
	updated: Date;
	categories: BudgetCategory[];
};

export type BudgetCategory = {
	id: number;
	name: string;
	type: string | null;
	allocatedAmount: number;
	actualAmount: number;
};

export const IncomingTypes = ["All", "Income", "Expense"];
