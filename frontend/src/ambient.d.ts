export type TransactionForm = {
    amount: number;
    description: string;
    startDate: Date;
    endDate: Date | null;
    recurring: boolean;
    transactionInterval: string | null;
    daysInterval: number | null;
    incoming: boolean;
    transactionType: string | null;
    errors: TransactionFormErrors;
};

export type TransactionFormErrors = {
    amount: string | null;
    description: string | null;
    startDate: string | null;
    endDate: string | null;
    transactionInterval: string | null;
    daysInterval: string | null;
    transactionType: string | null;
};
