<script lang="ts">
	import { ArrowLeft, ArrowRight, X } from "lucide-svelte";
	import { fly } from "svelte/transition";
	import type { Budget, Transaction } from "../../ambient";
	import { getFormattedAmount, getFormattedDate } from "$lib";

	const {
		open,
		budget,
		availableTransactions,
		close,
		updateTransactions,
	}: {
		open: boolean;
		budget: Budget;
		availableTransactions: Transaction[];
		close: () => void;
		updateTransactions: (idx: number[], expenseId: number) => void;
	} = $props();

	let modal: HTMLDialogElement;
	let addTransactionPage: number = $state(0);
	let selectedTransactions: number[] = $state([]);
	let selectedTransactionsError: string = $state("");
	let selectedBudgetExpense: number = $state(-1);
	let selectedBudgetExpenseError: string = $state("");

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
		close();
		modal.close();
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

	$effect(() => {
		if (open) {
			modal.showModal();
		}
	});
</script>

<dialog class="w-full max-w-screen-md overflow-x-hidden" bind:this={modal}>
	<button class="absolute right-4 top-4 active:outline-none" onclick={closeAddTransactionModal}>
		<X />
	</button>
	<h3 class="mb-4">Add Transactions</h3>
	<div>
		<div class={`${addTransactionPage !== 0 ? "absolute" : ""}`}>
			{#if addTransactionPage === 0}
				<div
					in:fly={{ duration: 100, x: -500, opacity: 0.5 }}
					out:fly={{ duration: 100, x: -500, opacity: 0.5 }}
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
