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
	} from "../../../ambient";
	import TransactionMonthHeader from "$lib/components/transactionMonthHeader.svelte";
	import { getToastStore } from "@skeletonlabs/skeleton";
	import { getCurrentMonthNumber, listAllMonths } from "$lib";

	const { transactionIntervals, incomeTypes, expenseTypes, yearInfo, demo } =
		$page.data;

	let showFormModal = $state(false);
	let month = $state(getCurrentMonthNumber());
	let incoming = $state(IncomingTypes[0]);
	let transactions: Promise<Transaction[]> | null = $state(null);
	let monthInfo: TransactionMonthInfo | null = $state(null);
	let monthInfoDiff: TransactionMonthInfo | null = $state(null);
	let selectedTransaction: Transaction | null = $state(null);
	let editTransaction: Transaction | null = $state(null);

	const toastStore = getToastStore();

	function setSelectedTransaction(transaction: Transaction | null) {
		selectedTransaction = transaction;
	}

	function setEditTransaction(transaction: Transaction | null) {
		editTransaction = transaction;
		showFormModal = true;
	}

	function closeFormModal() {
		editTransaction = null;
		showFormModal = false;
	}

	async function fetchTransactions() {
		const url = `/api/transactions?month=${month}&year=2024`;
		const response = await fetch(url);
		return (await response.json()) as Transaction[];
	}

	$effect(() => {
		transactions = fetchTransactions();
		monthInfo = yearInfo[month];
		if (month === 1 || monthInfo === null) return;

		const prevMonth = yearInfo[month - 1];
		if (prevMonth === undefined) return;

		monthInfoDiff = {
			income: monthInfo.income - prevMonth.income,
			expense: monthInfo.expense - prevMonth.expense,
		};
	});

	async function handleSuccess(action: string) {
		toastStore.trigger({
			background: "bg-success-400 text-black",
			message: `Transaction ${action} successfully`,
			timeout: 1500,
		});

		const response = await fetchTransactions();
		transactions = new Promise((r) => r(response));
	}
</script>

<svelte:head>
	<title>Fiscora - Transactions</title>
</svelte:head>

<div class="mx-auto mb-8 text-center lg:mb-12">
	<h1 class="mb-4">Your transactions</h1>
	<p>
		Add, view, and edit your transactions to stay on top of your financial
		journey.
	</p>
	<p>Track your finances with ease and gain valuable insights.</p>
</div>

<TransactionMonthHeader {monthInfo} {monthInfoDiff} />

<div class="my-4 flex flex-col items-center justify-between sm:flex-row">
	<div class="flex gap-2">
		{#each IncomingTypes as incomingType}
			<button
				class="rounded-full px-4 py-2 backdrop-blur-[1px] transition-colors {incoming !==
					incomingType && 'hover:bg-primary-500/20'} {incoming ===
					incomingType && 'variant-ghost-primary'}"
				onclick={() => (incoming = incomingType)}
			>
				{incomingType}
			</button>
		{/each}

		<select id="month-selector" class="select ml-4" bind:value={month}>
			{#each listAllMonths() as [idx, name]}
				<option selected={idx === month} value={idx}>{name}</option>
			{/each}
		</select>
	</div>
	<button
		class="secondary btn mt-4 sm:mt-0"
		onclick={() => (showFormModal = true)}
	>
		<Plus />&nbsp;Add transaction
	</button>
</div>
<div>
	<TransactionsList
		{transactions}
		{incoming}
		selectTransaction={setSelectedTransaction}
		editTransaction={setEditTransaction}
		{demo}
	/>
</div>

<TransactionFormModal
	{transactionIntervals}
	{incomeTypes}
	{expenseTypes}
	transaction={editTransaction}
	open={showFormModal}
	handleClose={closeFormModal}
	{handleSuccess}
	{demo}
/>

<TransactionInfoModal
	transaction={selectedTransaction}
	onclose={() => setSelectedTransaction(null)}
/>
