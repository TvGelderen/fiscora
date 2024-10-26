export type Transaction = {
    id: number;
    amount: number;
    description: string;
    date: Date;
    type: string | null;
    created: Date;
    updated: Date;
    recurring: {
        startDate: Date | null;
        endDate: Date | null;
        interval: string | null;
        daysInterval: number | null;
    } | null;
    budget: {
        id: string | null;
        name: string | null;
        expenseName: string | null;
    } | null;
};

export type TransactionForm = {
    id: number;
    amount: number;
    description: string;
    startDate: string | undefined;
    endDate: string | undefined;
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
    transactions: Transaction[] | null;
};

export type BudgetForm = {
    id: string;
    name: string;
    description: string;
    amount: number;
    startDate: string;
    endDate: string;
    expenses: BudgetExpenseForm[];
    errors: BudgetFormErrors;
};

export type BudgetFormErrors = {
    valid: boolean;
    name: string | null;
    description: string | null;
    amount: string | null;
    startDate: string | null;
    endDate: string | null;
};

export type BudgetExpense = {
    id: number;
    name: string;
    allocatedAmount: number;
    currentAmount: number;
};

export type BudgetExpenseForm = {
    id: number;
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
