<script lang="ts">
	import * as Dialog from "$lib/components/ui/dialog";
	import { ArrowLeft, ArrowRight } from "lucide-svelte";
	import type { Budget, Transaction } from "../../../../ambient";
	import { getFormattedAmount, getFormattedDate } from "$lib";
	import { Button, buttonVariants } from "$lib/components/ui/button";

	const {
		budget,
		availableTransactions,
		updateTransactions,
	}: {
		budget: Budget;
		availableTransactions: Transaction[];
		updateTransactions: (idx: number[], expenseId: number) => void;
	} = $props();

	let addTransactionPage: number = $state(0);
	let selectedTransactions: number[] = $state([]);
	let selectedTransactionsError: string = $state("");
	let selectedBudgetExpense: number = $state(-1);
	let selectedBudgetExpenseError: string = $state("");

	function next() {
		if (addTransactionPage < 1) {
			addTransactionPage++;
		}
	}

	function previous() {
		if (addTransactionPage > 0) {
			addTransactionPage--;
		}
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

	function closeAddTransactionModal() {
		addTransactionPage = 0;
		selectedTransactions = [];
		selectedBudgetExpense = -1;
		selectedTransactionsError = "";
		selectedBudgetExpenseError = "";
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

		const ids = selectedTransactions;
		const expenseId = selectedBudgetExpense;

		closeAddTransactionModal();

		updateTransactions(ids, expenseId);
	}
</script>

<Dialog.Content class="max-h-[100vh] w-full max-w-screen-md overflow-y-auto overflow-x-hidden">
	<Dialog.Header>
		<h3 class="mb-4">Add Transactions</h3>
	</Dialog.Header>
	<div>
		<div class={`${addTransactionPage !== 0 ? "absolute" : ""}`}>
			{#if addTransactionPage === 0}
				<div>
					<h4 class="my-4">Select the transactions to add</h4>
					{#if selectedTransactionsError}
						<p class="error-text">{selectedTransactionsError}</p>
					{/if}
					{#each availableTransactions as transaction (transaction.id)}
						<div
							class={`card my-2 cursor-pointer p-4 shadow-md hover:shadow-lg ${selectedTransactions.includes(transaction.id) ? "bg-primary/10" : ""}`}
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
				<div>
					<h4 class="my-4">Select the budget expense to add them to</h4>
					{#if selectedBudgetExpenseError}
						<p class="error-text">{selectedBudgetExpenseError}</p>
					{/if}
					{#each budget.expenses as expense (expense.id)}
						<div
							class={`card my-2 cursor-pointer p-4 shadow-md hover:shadow-xl ${selectedBudgetExpense === expense.id ? "bg-primary/10" : ""}`}
							onclick={() => selectBudgetExpense(expense.id)}
							role="none"
						>
							<span>{expense.name}</span>
						</div>
					{/each}
				</div>
			{/if}
		</div>

		<div class="mt-8 flex items-center justify-between">
			<button
				class={`${buttonVariants({ variant: "ghost" })} flex gap-2`}
				onclick={previous}
				disabled={addTransactionPage === 0}
			>
				<ArrowLeft /> Back
			</button>
			{#if addTransactionPage !== 1}
				<button
					class={`${buttonVariants({ variant: "ghost" })} flex gap-2`}
					onclick={next}
					disabled={addTransactionPage === 1}
				>
					Next <ArrowRight />
				</button>
			{:else}
				<div class="flex gap-2">
					<Button variant="secondary" onclick={closeAddTransactionModal}>Cancel</Button>
					<Button onclick={handleAddTransactions}>Save</Button>
				</div>
			{/if}
		</div>
	</div>
</Dialog.Content>
