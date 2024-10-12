import type { Budget, BudgetForm, BudgetFormErrors } from "../../ambient";
import { authorizeFetch, createRandomString } from "$lib/index";
import { validDate, validNumber, validString } from "./utils";

// Mock data
const mockBudgets: Budget[] = [
    {
        id: createRandomString(16),
        name: "Monthly Budget",
        description: "Overall monthly budget for household expenses",
        amount: 3000,
        startDate: new Date(),
        endDate: new Date(new Date().setMonth(new Date().getMonth() + 1)),
        created: new Date(),
        updated: new Date(),
        categories: [
            {
                id: 1,
                name: "Groceries",
                description: "Groceries",
                allocatedAmount: 500,
                currentAmount: 0,
            },
            {
                id: 2,
                name: "Utilities",
                description: "Utilities",
                allocatedAmount: 300,
                currentAmount: 0,
            },
            {
                id: 3,
                name: "Entertainment",
                description: "Entertainment",
                allocatedAmount: 200,
                currentAmount: 0,
            },
            {
                id: 4,
                name: "Savings",
                description: "Savings",
                allocatedAmount: 500,
                currentAmount: 0,
            },
        ],
    },
    {
        id: createRandomString(16),
        name: "Vacation Fund",
        description: "Saving for summer vacation",
        amount: 1500,
        startDate: new Date(),
        endDate: new Date(new Date().setMonth(new Date().getMonth() + 1)),
        created: new Date(),
        updated: new Date(),
        categories: [
            {
                id: 1,
                name: "Accommodation",
                description: "Accommodation",
                allocatedAmount: 600,
                currentAmount: 0,
            },
            {
                id: 2,
                name: "Transportation",
                description: "Transportation",
                allocatedAmount: 400,
                currentAmount: 0,
            },
            {
                id: 3,
                name: "Activities",
                description: "Activities",
                allocatedAmount: 300,
                currentAmount: 0,
            },
            {
                id: 4,
                name: "Food",
                description: "Food",
                allocatedAmount: 200,
                currentAmount: 0,
            },
        ],
    },
];

export async function getBudgets(accessToken: string): Promise<Budget[]> {
    const response = await authorizeFetch("budgets", accessToken);
    if (!response.ok) {
        return [];
    }

    return (await response.json()) as Budget[];
}

export async function createBudget(
    budget: Omit<Budget, "id">,
    accessToken: string,
): Promise<Budget> {
    // In a real implementation, you would use authorizeFetchBody here
    // const response = await authorizeFetchBody("budgets", accessToken, "POST", JSON.stringify(budget));
    // return response.json();

    // For now, we'll simulate creating a new budget
    console.log(accessToken);
    return new Promise((resolve) => {
        const newBudget: Budget = {
            ...budget,
            id: createRandomString(16),
        };
        mockBudgets.push(newBudget);
        resolve(newBudget);
    });
}

export async function updateBudget(
    budget: Budget,
    accessToken: string,
): Promise<Budget> {
    // In a real implementation, you would use authorizeFetchBody here
    // const response = await authorizeFetchBody(`budgets/${budget.id}`, accessToken, "PUT", JSON.stringify(budget));
    // return response.json();

    // For now, we'll simulate updating a budget
    console.log(accessToken);
    return new Promise((resolve, reject) => {
        const index = mockBudgets.findIndex((b) => b.id === budget.id);
        if (index !== -1) {
            mockBudgets[index] = budget;
            resolve(budget);
        } else {
            reject(new Error("Budget not found"));
        }
    });
}

// Delete a budget
export async function deleteBudget(
    id: string,
    accessToken: string,
): Promise<void> {
    // In a real implementation, you would use authorizeFetch here
    // await authorizeFetch(`budgets/${id}`, accessToken, "DELETE");

    // For now, we'll simulate deleting a budget
    console.log(accessToken);
    return new Promise((resolve, reject) => {
        const index = mockBudgets.findIndex((b) => b.id === id);
        if (index !== -1) {
            mockBudgets.splice(index, 1);
            resolve();
        } else {
            reject(new Error("Budget not found"));
        }
    });
}

export function verifyForm(form: BudgetForm): BudgetFormErrors {
    const errors: BudgetFormErrors = {
        name: null,
        description: null,
        amount: null,
        startDate: null,
        endDate: null,
    };

    if (!validString(form.name)) {
        errors.name = "Description is required";
    }
    if (!validString(form.description)) {
        errors.description = "Description is required";
    }
    if (!validNumber(form.amount)) {
        errors.amount = "Amount must be a number";
    }
    if (!validDate(form.startDate)) {
        errors.startDate = "Start date must be a valid date";
    }
    if (!validDate(form.endDate)) {
        errors.startDate = "End date must be a valid date";
    }

    return errors;
}
