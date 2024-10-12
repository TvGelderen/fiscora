import type { Budget } from "../../ambient";

// Mock data
const mockBudgets: Budget[] = [
	{
		id: 1,
		name: "Monthly Budget",
		description: "Overall monthly budget for household expenses",
		amount: 3000,
		startDate: new Date(),
		endDate: new Date(new Date().setMonth(new Date().getMonth() + 1)),
		type: "Monthly",
		created: new Date(),
		updated: new Date(),
		categories: [
			{
				id: 1,
				name: "Groceries",
				type: null,
				allocatedAmount: 500,
				actualAmount: 0,
			},
			{
				id: 2,
				name: "Utilities",
				type: null,
				allocatedAmount: 300,
				actualAmount: 0,
			},
			{
				id: 3,
				name: "Entertainment",
				type: null,
				allocatedAmount: 200,
				actualAmount: 0,
			},
			{
				id: 4,
				name: "Savings",
				type: null,
				allocatedAmount: 500,
				actualAmount: 0,
			},
		],
	},
	{
		id: 2,
		name: "Vacation Fund",
		description: "Saving for summer vacation",
		amount: 1500,
		startDate: new Date(),
		endDate: new Date(new Date().setMonth(new Date().getMonth() + 1)),
		type: "Monthly",
		created: new Date(),
		updated: new Date(),
		categories: [
			{
				id: 1,
				name: "Accommodation",
				type: null,
				allocatedAmount: 600,
				actualAmount: 0,
			},
			{
				id: 2,
				name: "Transportation",
				type: null,
				allocatedAmount: 400,
				actualAmount: 0,
			},
			{
				id: 3,
				name: "Activities",
				type: null,
				allocatedAmount: 300,
				actualAmount: 0,
			},
			{
				id: 4,
				name: "Food",
				type: null,
				allocatedAmount: 200,
				actualAmount: 0,
			},
		],
	},
];

// Get all budgets
export async function getBudgets(accessToken: string): Promise<Budget[]> {
	// In a real implementation, you would use authorizeFetch here
	// const response = await authorizeFetch("budgets", accessToken);
	// return response.json();

	// For now, we'll return the mock data
	console.log(accessToken);
	return new Promise((resolve) => {
		setTimeout(() => resolve(mockBudgets), 500); // Simulate network delay
	});
}

// Create a new budget
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
		setTimeout(() => {
			const newBudget: Budget = {
				...budget,
				id: Math.max(...mockBudgets.map((b) => b.id)) + 1,
			};
			mockBudgets.push(newBudget);
			resolve(newBudget);
		}, 500);
	});
}

// Update an existing budget
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
		setTimeout(() => {
			const index = mockBudgets.findIndex((b) => b.id === budget.id);
			if (index !== -1) {
				mockBudgets[index] = budget;
				resolve(budget);
			} else {
				reject(new Error("Budget not found"));
			}
		}, 500);
	});
}

// Delete a budget
export async function deleteBudget(
	id: number,
	accessToken: string,
): Promise<void> {
	// In a real implementation, you would use authorizeFetch here
	// await authorizeFetch(`budgets/${id}`, accessToken, "DELETE");

	// For now, we'll simulate deleting a budget
	console.log(accessToken);
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			const index = mockBudgets.findIndex((b) => b.id === id);
			if (index !== -1) {
				mockBudgets.splice(index, 1);
				resolve();
			} else {
				reject(new Error("Budget not found"));
			}
		}, 500);
	});
}
