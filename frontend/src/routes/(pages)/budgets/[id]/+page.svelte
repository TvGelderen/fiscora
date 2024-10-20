<script lang="ts">
	import { page } from "$app/stores";
	import { getFormattedAmount, getFormattedDate, getFormDate } from "$lib";
	import { getToastStore, ProgressBar } from "@skeletonlabs/skeleton";
	import type { PageData } from "./$types";
	import type { BudgetExpense, Transaction } from "../../../../ambient";
	import { ArrowLeft, ArrowRight, Plus, X } from "lucide-svelte";
	import { fly } from "svelte/transition";

	const toastStore = getToastStore();

	let { budget, demo } = $page.data as PageData;

	let availableTransactions: Transaction[] = $state([]);
	let addTransactionModal: HTMLDialogElement;
	let addTransactionPage: number = $state(0);
	let selectedTransactions: number[] = $state([]);
	let selectedTransactionsError: string = $state("");
	let selectedBudgetExpense: number = $state(-1);
	let selectedBudgetExpenseError: string = $state("");

	async function openAddTransactionModal() {
		addTransactionModal.showModal();
		await getUnassignedTransactions();
	}

	function closeAddTransactionModal() {
		addTransactionModal.close();
		addTransactionPage = 0;
		selectedTransactions = [];
		selectedBudgetExpense = -1;
	}

	function selectTransaction(id: number) {
		const idx = selectedTransactions.findIndex((i) => i === id);
		if (idx === -1) {
			selectedTransactions.push(id);
		} else {
			selectedTransactions.splice(idx, 1);
			selectedTransactions = [...selectedTransactions];
		}
	}

	function selectBudgetExpense(id: number) {
		if (selectedBudgetExpense === id) {
			selectedBudgetExpense = -1;
		} else {
			selectedBudgetExpense = id;
		}
	}

	async function handleAddTransactions() {
		selectedTransactionsError = "";
		selectedBudgetExpenseError = "";

		if (selectedTransactions.length === 0) {
			selectedTransactionsError = "Please select at least one transaction to add to the budget.";
			addTransactionPage = 0;
		}
		if (selectedBudgetExpense === -1) {
			selectedBudgetExpenseError = "Please select the expense to add the transaction(s) to.";
		}
		if (selectedTransactionsError || selectedBudgetExpenseError) {
			return;
		}

		closeAddTransactionModal();

		if (demo) {
			toastStore.trigger({
				message: "Demo users cannot create budgets",
				background: "variant-filled-warning",
			});
			return;
		}

		const response = await fetch(`/api/budgets/${budget.id}/expenses/${selectedBudgetExpense}/transactions`, {
			method: "POST",
			body: JSON.stringify(selectedTransactions),
		});
		if (!response.ok) {
			toastStore.trigger({
				background: "bg-error-400/50 text-black dark:text-white",
				message: "Something went wrong adding transactions",
			});
			return;
		}

		toastStore.trigger({
			background: "bg-success-400 text-black",
			message: "Transactions added successfully",
		});
	}

	function calculateTotalSpent(expenses: BudgetExpense[]): number {
		return expenses.reduce((total, expense) => total + expense.currentAmount, 0);
	}

	function calculateProgress(expense: BudgetExpense): number {
		return (expense.currentAmount / expense.allocatedAmount) * 100;
	}

	async function getUnassignedTransactions() {
		const response = await fetch(
			`/api/transactions/unassigned?startDate=${getFormDate(budget.startDate)}&endDate=${getFormDate(budget.endDate)}`,
		);
		const data = await response.json();
		availableTransactions = data;
	}
</script>

<svelte:head>
	<title>Fiscora - {budget.name}</title>
</svelte:head>

<div class="mx-auto max-w-screen-lg">
	<div class="mb-8 flex flex-col items-center justify-between gap-4 md:flex-row">
		<div>
			<h2 class="mb-4">{budget.name}</h2>
			<p class="text-secondary">{budget.description}</p>
		</div>
		<div class="grid grid-cols-[max-content_1fr] gap-x-4">
			<span class="text-secondary">Start:</span>
			<span class="text-end italic">{getFormattedDate(budget.startDate)}</span>
			<span class="text-secondary">End:</span>
			<span class="text-end italic">{getFormattedDate(budget.endDate)}</span>
		</div>
	</div>

	<section class="mb-12">
		<div class="mb-8 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
			<div class="card-primary p-4">
				<h4 class="mb-2">Total Budget</h4>
				<p class="text-2xl">
					{getFormattedAmount(budget.amount)}
				</p>
			</div>
			<div class="card-primary p-4">
				<h4 class="mb-2">Total Spent</h4>
				<p class="text-2xl">
					{getFormattedAmount(calculateTotalSpent(budget.expenses))}
				</p>
			</div>
		</div>
	</section>

	<section class="mb-12">
		<div class="mb-6 flex items-center justify-between gap-4">
			<h3>Expenses</h3>
			<button type="button" class="!variant-soft-primary btn-icon btn-icon-sm" onclick={() => {}}>
				<Plus size={20} />
			</button>
		</div>
		<div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
			{#each budget.expenses as expense}
				<div class="card-primary p-4">
					<h4 class="mb-2">{expense.name}</h4>
					<p class="mb-2">
						Spent: {getFormattedAmount(expense.currentAmount)} / Allocated:
						{getFormattedAmount(expense.allocatedAmount)}
					</p>
					<ProgressBar value={calculateProgress(expense)} max={100} height="h-2" meter="bg-primary-500" />
				</div>
			{/each}
		</div>
	</section>

	<section class="mb-12">
		<div class="mb-6 flex items-center justify-between gap-4">
			<h3>Transactions</h3>
			<button type="button" class="!variant-soft-primary btn-icon btn-icon-sm" onclick={openAddTransactionModal}>
				<Plus size={20} />
			</button>
		</div>
	</section>
</div>

<dialog class="w-full max-w-screen-md overflow-x-hidden" bind:this={addTransactionModal}>
	<button class="absolute right-4 top-4 active:outline-none" onclick={closeAddTransactionModal}>
		<X />
	</button>
	<h3 class="mb-4">Add Transaction</h3>
	<div>
		<div class={`${addTransactionPage !== 0 ? "absolute" : ""}`}>
			{#if addTransactionPage === 0}
				<div
					in:fly={{ duration: 200, x: -500, opacity: 0.5 }}
					out:fly={{ duration: 200, x: -500, opacity: 0.5 }}
				>
					<h4 class="my-4">Select the transactions to add</h4>
					{#if selectedTransactionsError}
						<p class="error-text">{selectedTransactionsError}</p>
					{/if}
					{#each availableTransactions as transaction (transaction.id)}
						<div
							class={`card-primary my-2 cursor-pointer p-4 shadow-md hover:shadow-xl ${selectedTransactions.includes(transaction.id) ? "bg-primary-400/20 dark:bg-primary-500/20" : ""}`}
							onclick={() => selectTransaction(transaction.id)}
							role="none"
						>
							<span>{getFormattedDate(transaction.date)}</span>
							<span>{transaction.description}</span>
							<span>{getFormattedAmount(transaction.amount)}</span>
						</div>
					{/each}
					{#if availableTransactions.length === 0}
						<p>No unassigned transactions found within this budget's date range.</p>
					{/if}
				</div>
			{/if}
		</div>

		<div class={`${addTransactionPage !== 1 ? "absolute" : ""}`}>
			{#if addTransactionPage === 1}
				<div in:fly={{ duration: 200, x: 500, opacity: 0.5 }} out:fly={{ duration: 200, x: 500, opacity: 0.5 }}>
					<h4 class="my-4">Select the budget expense to add them to</h4>
					{#if selectedBudgetExpenseError}
						<p class="error-text">{selectedBudgetExpenseError}</p>
					{/if}
					{#each budget.expenses as expense (expense.id)}
						<div
							class={`card-primary my-2 cursor-pointer p-4 shadow-md hover:shadow-xl ${selectedBudgetExpense === expense.id ? "bg-primary-400/20 dark:bg-primary-500/20" : ""}`}
							onclick={() => selectBudgetExpense(expense.id)}
							role="none"
						>
							<span>{expense.name}</span>
						</div>
					{/each}
				</div>
			{/if}
		</div>

		<div class="mt-8 flex items-end justify-between">
			<button
				class="flex items-center gap-2 disabled:opacity-50"
				onclick={() => addTransactionPage--}
				disabled={addTransactionPage === 0}
			>
				<ArrowLeft /> Back
			</button>
			{#if addTransactionPage !== 1}
				<button
					class="flex items-center gap-2 disabled:opacity-50"
					onclick={() => addTransactionPage++}
					disabled={addTransactionPage === 1}
				>
					Next <ArrowRight />
				</button>
			{:else}
				<div class="flex gap-2">
					<button class="!variant-filled-surface btn" onclick={closeAddTransactionModal}>Cancel</button>
					<button class="btn" onclick={handleAddTransactions}>Save</button>
				</div>
			{/if}
		</div>
	</div>
</dialog>
