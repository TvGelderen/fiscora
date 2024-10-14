<script lang="ts">
	import { page } from "$app/stores";
	import { getFormattedAmount, getFormattedDate } from "$lib";
	import { ProgressBar } from "@skeletonlabs/skeleton";
	import type { PageData } from "./$types";
	import type { BudgetExpense } from "../../../../ambient";

	let { budget } = $page.data as PageData;

	function calculateTotalSpent(expenses: BudgetExpense[]): number {
		return expenses.reduce(
			(total, expense) => total + expense.currentAmount,
			0,
		);
	}

	function calculateProgress(expense: BudgetExpense): number {
		return (expense.currentAmount / expense.allocatedAmount) * 100;
	}
</script>

<svelte:head>
	<title>Fiscora - {budget.name || "Budget Details"}</title>
</svelte:head>

<div class="container mx-auto px-4 py-8">
	<h2 class="mb-4">{budget.name}</h2>
	<p class="mb-6">{budget.description}</p>

	<div class="mb-8 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
		<div class="card p-4">
			<h3 class="mb-2">Total Budget</h3>
			<p class="text-2xl">
				${getFormattedAmount(budget.amount)}
			</p>
		</div>
		<div class="card p-4">
			<h3 class="mb-2">Total Spent</h3>
			<p class="text-2xl">
				${getFormattedAmount(calculateTotalSpent(budget.expenses))}
			</p>
		</div>
		<div class="card p-4">
			<h3 class="mb-2 text-lg">Start Date</h3>
			<p class="text-xl">{getFormattedDate(budget.startDate)}</p>
		</div>
		<div class="card p-4">
			<h3 class="mb-2 text-lg">End Date</h3>
			<p class="text-xl">{getFormattedDate(budget.endDate)}</p>
		</div>
	</div>

	<h2 class="mb-6">Expenses</h2>
	<div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
		{#each budget.expenses as expense}
			<div class="card p-4">
				<h4 class="mb-2">{expense.name}</h4>
				<p class="mb-2">
					Spent: {getFormattedAmount(expense.currentAmount)} / Allocated:
					{getFormattedAmount(expense.allocatedAmount)}
				</p>
				<ProgressBar
					value={calculateProgress(expense)}
					max={100}
					height="h-2"
					meter="bg-primary-500"
				/>
			</div>
		{/each}
	</div>

	<h2 class="my-4">Transactions</h2>
</div>
