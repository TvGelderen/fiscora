<script lang="ts">
	import Plus from "lucide-svelte/icons/plus";
	import { page } from "$app/stores";
	import TransactionsList from "$lib/components/transactionsList.svelte";
	import TransactionInfoModal from "$lib/components/transactionInfoModal.svelte";
	import TransactionFormModal from "$lib/components/transactionFormModal.svelte";
	import type {
		Transaction,
		TransactionMonthInfoResponse,
	} from "../../ambient";

	const { transactionIntervals, incomeTypes, expenseTypes } = $page.data;

	const transactionTypes = ["All", "Income", "Expense"];

	function listAllMonths() {
		const months = new Map<number, string>();
		for (let month = 0; month < 12; month++) {
			const monthName = new Date(2000, month, 1).toLocaleString(
				"default",
				{ month: "long" },
			);
			months.set(month + 1, monthName);
		}
		return months;
	}

	let showFormModal = $state(false);
	let month = $state(
		Number.parseInt(
			new Date().toLocaleString("default", { month: "numeric" }),
		),
	);
	let type = $state(transactionTypes[0]);
	let transactions: Promise<Transaction[]> | null = $state(null);
	let selectedTransaction: Transaction | null = $state(null);
	let income = $state(0);
	let expense = $state(0);
	let netIncome = $derived(income - expense);

	function selectTransaction(transaction: Transaction | null) {
		selectedTransaction = transaction;
	}

	async function fetchTransactions() {
		const url = `/api/transactions?month=${month}&year=2024${type !== transactionTypes[0] ? `&income=${type === transactionTypes[1]}` : ""}`;
		const response = await fetch(url);
		return (await response.json()) as Transaction[];
	}

	async function fetchTransactionsMonthInfo() {
		const url = `/api/transactions/month-info?month=${month}&year=2024`;
		const response = await fetch(url);
		return (await response.json()) as TransactionMonthInfoResponse;
	}

	$effect(() => {
		fetchTransactionsMonthInfo().then((data) => {
			income = data.income;
			expense = data.expense;
		});
	});

	$effect(() => {
		transactions = fetchTransactions();
	});
</script>

<title>Budget Buddy - Transactions</title>

<div class="mx-auto mb-10 mt-4 text-center lg:mb-16">
	<h1 class="mb-4">Your transactions</h1>
	<p>
		Add, view, and edit your transactions to stay on top of your financial
		journey.
	</p>
	<p>Track your finances with ease and gain valuable insights.</p>
</div>

<div
	class="mb-10 grid rounded-2xl bg-primary-500/20 shadow-md shadow-primary-900/50 dark:shadow-surface-900 sm:grid-cols-3 lg:mb-16"
>
	<div class="flex flex-col items-center justify-between p-4 sm:items-start">
		<h4 class="mb-6">Total income</h4>
		<span class="text-2xl lg:text-3xl">€{income}</span>
	</div>
	<div
		class="flex flex-col items-center justify-between border-b-[1px] border-t-[1px] border-primary-700/25 p-4 sm:items-start sm:border-b-[0px] sm:border-l-[1px] sm:border-r-[1px] sm:border-t-[0px]"
	>
		<h4 class="mb-6">Total expense</h4>
		<span class="text-2xl lg:text-3xl">€{expense}</span>
	</div>
	<div class="flex flex-col items-center justify-between p-4 sm:items-start">
		<h4 class="mb-6">Net income</h4>
		<span class="text-2xl lg:text-3xl">€{netIncome}</span>
	</div>
</div>

<div class="my-4 flex flex-col items-center justify-between sm:flex-row">
	<div class="flex gap-2">
		{#each transactionTypes as transactionType}
			<button
				class="rounded-full px-4 py-2 transition-colors {type !==
					transactionType && 'hover:bg-primary-500/20'} {type ===
					transactionType && 'variant-ghost-primary'}"
				onclick={() => (type = transactionType)}
			>
				{transactionType}
			</button>
		{/each}
	</div>
	<button
		class="secondary btn mt-4 sm:mt-0"
		onclick={() => (showFormModal = true)}
	>
		<Plus />&nbsp;Add transaction
	</button>
</div>
<div class="">
	<select id="month-selector" class="select" bind:value={month}>
		{#each listAllMonths() as [idx, name]}
			<option selected={idx === month} value={idx}>{name}</option>
		{/each}
	</select>
	<TransactionsList {transactions} {selectTransaction} />
</div>

<TransactionFormModal
	{transactionIntervals}
	{incomeTypes}
	{expenseTypes}
	open={showFormModal}
	onclose={() => (showFormModal = false)}
/>

<TransactionInfoModal
	transaction={selectedTransaction}
	onclose={() => selectTransaction(null)}
/>
