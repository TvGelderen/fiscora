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
    valid: boolean;
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
    id: string;
    name: string;
    description: string;
    amount: number;
    startDate: Date;
    endDate: Date;
    created: Date;
    updated: Date;
    expenses: BudgetExpense[];
};

export type BudgetForm = {
    name: string;
    description: string;
    amount: number;
    expenses: BudgetExpenseForm[];
    errors: BudgetFormErrors;
};

export type BudgetFormErrors = {
    valid: boolean;
    name: string | null;
    description: string | null;
    amount: string | null;
};

export type BudgetExpense = {
    id: number;
    name: string;
    allocatedAmount: number;
    currentAmount: number;
};

export type BudgetExpenseForm = {
    name: string;
    allocatedAmount: number;
    errors: BudgetExpenseFormErrors;
};

export type BudgetExpenseFormErrors = {
    valid: boolean;
    name: string | null;
    allocatedAmount: string | null;
};

export const IncomingTypes = ["All", "Income", "Expense"];
