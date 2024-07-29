<script lang="ts">
	import Plus from "lucide-svelte/icons/plus";
	import { page } from "$app/stores";
	import TransactionsList from "$lib/components/transactionsList.svelte";
	import TransactionInfoModal from "$lib/components/transactionInfoModal.svelte";
	import TransactionFormModal from "$lib/components/transactionFormModal.svelte";
	import {
		IncomingTypes,
		type Transaction,
		type TransactionMonthInfo,
	} from "../../ambient";
	import TransactionMonthHeader from "$lib/components/transactionMonthHeader.svelte";

	const { transactionIntervals, incomeTypes, expenseTypes } = $page.data;

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
	let incoming = $state(IncomingTypes[0]);
	let transactions: Promise<Transaction[]> | null = $state(null);
	let monthInfo: Promise<TransactionMonthInfo> | null = $state(null);
	let selectedTransaction: Transaction | null = $state(null);

	function selectTransaction(transaction: Transaction | null) {
		selectedTransaction = transaction;
	}

	async function fetchTransactions() {
		const url = `/api/transactions?month=${month}&year=2024`;
		const response = await fetch(url);
		return (await response.json()) as Transaction[];
	}

	async function fetchTransactionsMonthInfo() {
		const url = `/api/transactions/month-info?month=${month}&year=2024`;
		const response = await fetch(url);
		return (await response.json()) as TransactionMonthInfo;
	}

	$effect(() => {
		monthInfo = fetchTransactionsMonthInfo();
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

<TransactionMonthHeader {monthInfo} />

<div class="my-4 flex flex-col items-center justify-between sm:flex-row">
	<div class="flex gap-2">
		{#each IncomingTypes as incomingType}
			<button
				class="rounded-full px-4 py-2 transition-colors {incoming !==
					incomingType && 'hover:bg-primary-500/20'} {incoming ===
					incomingType && 'variant-ghost-primary'}"
				onclick={() => (incoming = incomingType)}
			>
				{incomingType}
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
	<TransactionsList {transactions} {incoming} {selectTransaction} />
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
