<script lang="ts">
	import { page } from "$app/stores";
	import { getFormattedAmount, getFormattedDate, getFormattedDateShort, getFormDate } from "$lib";
	import { ProgressBar } from "@skeletonlabs/skeleton";
	import type { PageData } from "./$types";
	import type { Budget, BudgetExpense, Transaction } from "../../../../ambient";
	import { Plus, Trash, X } from "lucide-svelte";
	import BudgetAddTransactions from "$lib/components/budget-add-transactions.svelte";
	import { toast } from "svelte-sonner";

	let { budget, demo } = $page.data as PageData;

	let removeModal: HTMLDialogElement;
	let transactionToRemoveId: number | null = $state(null);
	let budgetState: Budget = $state(budget);
	let availableTransactions: Transaction[] = $state([]);
	let showAddTransactions: boolean = $state(false);

	async function updateTransactions(ids: number[], expenseId: number) {
		if (budgetState === null) {
			return;
		}

		if (demo) {
			toast.warning("Demo users cannot create budgets");
			return;
		}

		if (budgetState.transactions !== null) {
			const transactions = [
				...budgetState.transactions,
				...availableTransactions.filter((t) => ids.includes(t.id)),
			];

			transactions.sort((a, b) => (a.date > b.date ? -1 : a.date < b.date ? 1 : 0));

			budgetState.transactions = transactions;
		} else {
			budgetState.transactions = availableTransactions.filter((t) => ids.includes(t.id));
		}

		const response = await fetch(`/api/budgets/${budget.id}/expenses/${expenseId}/transactions`, {
			method: "POST",
			body: JSON.stringify(ids),
		});
		if (!response.ok) {
			toast.error("Something went wrong adding transactions");
			return;
		}

		availableTransactions = availableTransactions.filter((t) => !ids.includes(t.id));

		const transactions = (await response.json()) as Transaction[];

		toast.success("Transactions added successfully");

		budgetState.transactions = transactions;
	}

	function openRemoveModal(event: MouseEvent, transactionId: number) {
		event.preventDefault();
		removeModal.showModal();
		transactionToRemoveId = transactionId;
	}

	function closeRemoveModal() {
		removeModal.close();
		transactionToRemoveId = null;
	}

	async function confirmRemove() {
		if (transactionToRemoveId === null || budgetState === null || budgetState.transactions === null) {
			return;
		}

		const id = transactionToRemoveId;

		closeRemoveModal();

		if (demo) {
			toast.warning("Demo users cannot remove transactions");
			return;
		}

		const transactionIdx = budgetState.transactions.findIndex((t) => t.id === id);
		const transaction = budgetState.transactions.at(transactionIdx);

		budgetState.transactions = budgetState.transactions.filter((t) => t.id !== id);

		const response = await fetch(`/api/transactions/${id}/budget`, { method: "DELETE" });
		if (!response.ok) {
			toast.error("Something went wrong removing transaction");

			if (transaction !== undefined) {
				budgetState.transactions.splice(transactionIdx, 0, transaction);
				budgetState.transactions = [...budgetState.transactions];
			}

			return;
		}

		toast.success("Transaction removed successfully");
	}

	function calculateTotalSpent(expenses: BudgetExpense[]): number {
		return expenses.reduce((total, expense) => total + expense.currentAmount, 0);
	}

	function calculateProgress(expense: BudgetExpense): number {
		return (expense.currentAmount / expense.allocatedAmount) * 100;
	}

	async function updateAvailableTransactions() {
		const response = await fetch(
			`/api/transactions/unassigned?startDate=${getFormDate(budget.startDate)}&endDate=${getFormDate(budget.endDate)}`,
		);
		availableTransactions = (await response.json()) as Transaction[];
	}

	$effect(() => {
		updateAvailableTransactions();
	});
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
			<button
				type="button"
				class="!variant-soft-primary btn-icon btn-icon-sm"
				onclick={() => (showAddTransactions = true)}
			>
				<Plus size={20} />
			</button>
		</div>
		<div>
			{#if (budget.transactions !== null || budgetState.transactions !== null) && (budget.transactions?.length !== 0 || budgetState.transactions?.length !== 0)}
				<table class="mt-4 w-full select-none overflow-hidden rounded-md text-left [&_th]:p-4">
					<thead>
						<tr>
							<th class="w-[10%]">Date</th>
							<th class="w-[35%] min-w-[200px]">Description</th>
							<th class="w-[15%] text-right">Amount</th>
							<th class="w-[20%]">Type</th>
							<th class="w-[20%]">Expense</th>
						</tr>
					</thead>
					<tbody class="transactions-table-body">
						{#each budgetState?.transactions ?? budget.transactions! as transaction}
							<tr class="transactions-table-row">
								<td data-cell="date">
									{getFormattedDateShort(transaction.date)}
								</td>
								<td data-cell="description">
									{transaction.description}
								</td>
								<td data-cell="amount">
									{getFormattedAmount(transaction.amount)}
								</td>
								<td data-cell="type">{transaction.type}</td>
								<td data-cell="expense">{transaction.budget?.expenseName ?? "-"}</td>
								<td data-cell="">
									<div class="flex justify-end gap-1">
										<button
											class="icon hover:bg-error-500/60 inline rounded-md p-2 hover:!text-black dark:hover:!text-white"
											onclick={(event) => openRemoveModal(event, transaction.id)}
										>
											<Trash size={20} />
										</button>
									</div>
								</td>
							</tr>
						{/each}
					</tbody>
				</table>
			{:else}
				<div>No transactions found for this budget.</div>
			{/if}
		</div>
	</section>
</div>

<BudgetAddTransactions
	open={showAddTransactions}
	budget={budgetState ?? budget}
	{availableTransactions}
	close={() => (showAddTransactions = false)}
	{updateTransactions}
/>

<dialog bind:this={removeModal}>
	<button class="absolute right-4 top-4" onclick={closeRemoveModal}>
		<X />
	</button>
	{#if transactionToRemoveId !== null}
		<h3 class="mb-4">Confirm Removal</h3>
		<p>Are you sure you want to remove this transaction from this budget?</p>
		<div class="mt-4 flex justify-end gap-2">
			<button class="!variant-filled-surface btn" onclick={closeRemoveModal}>Cancel</button>
			<button class="!variant-filled-error btn" onclick={confirmRemove}>Delete</button>
		</div>
	{/if}
</dialog>
