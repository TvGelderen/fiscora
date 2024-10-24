<script lang="ts">
	import Plus from "lucide-svelte/icons/plus";
	import { page } from "$app/stores";
	import TransactionList from "$lib/components/transaction-list.svelte";
	import TransactionInfoModal from "$lib/components/transaction-info.svelte";
	import TransactionFormModal from "$lib/components/transaction-form.svelte";
	import { IncomingTypes, type Transaction, type TransactionMonthInfo } from "../../../ambient";
	import TransactionMonthHeader from "$lib/components/transaction-month-header.svelte";
	import { getCurrentMonthNumber, listAllMonths } from "$lib";
	import type { PageData } from "./$types";

	let { transactions, transactionIntervals, incomeTypes, expenseTypes, yearInfo, demo } = $page.data as PageData;

	let showFormModal = $state(false);
	let month = $state(getCurrentMonthNumber());
	let incoming = $state(IncomingTypes[0]);
	let transactionsState: Transaction[] = $state(transactions);
	let monthInfo: TransactionMonthInfo | null = $state(null);
	let monthInfoDiff: TransactionMonthInfo | null = $state(null);
	let selectedTransaction: Transaction | null = $state(null);
	let editTransaction: Transaction | null = $state(null);

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

	async function updateTransactions() {
		const response = await fetch(`/api/transactions?month=${month}&year=2024`);
		const value = (await response.json()) as Transaction[];
		transactionsState = value;
	}

	function updateMonthInfo() {
		const info = yearInfo.get(month.toString());
		if (!info) return;

		monthInfo = info;

		const prevMonth = yearInfo.get((month - 1).toString());
		if (!prevMonth) return;

		monthInfoDiff = {
			income: monthInfo.income - prevMonth.income,
			expense: monthInfo.expense - prevMonth.expense,
		};
	}

	async function handleSuccess() {
		closeFormModal();
		await updateTransactions();
	}

	function addTransaction(transaction: Transaction, idx: number) {
		transactionsState.splice(idx, 0, transaction);
		transactionsState = [...transactionsState];
		updateYearInfo(transaction.amount, true);
	}

	function removeTransaction(transaction: Transaction, idx: number) {
		transactionsState.splice(idx, 1);
		transactionsState = [...transactionsState];
		updateYearInfo(transaction.amount, false);
	}

	function updateYearInfo(amount: number, add: boolean) {
		if (monthInfo === null) return;

		if (add) {
			if (amount > 0) {
				monthInfo.income += amount;
			} else {
				monthInfo.expense -= amount;
			}
		} else {
			if (amount > 0) {
				monthInfo.income -= amount;
			} else {
				monthInfo.expense += amount;
			}
		}

		yearInfo.set(month.toString(), monthInfo);
		updateMonthInfo();
	}

	$effect(() => {
		updateMonthInfo();
		updateTransactions();
	});

	$effect(() => {
		const all = incoming === IncomingTypes[0];
		transactionsState = transactions.filter((t) => {
			const date = new Date(t.date);
			date.setUTCMilliseconds(date.getMilliseconds() + 1);
			t.date = date;
			if (all) return true;
			return (incoming === IncomingTypes[1] && t.amount > 0) || (incoming === IncomingTypes[2] && t.amount < 0);
		});
	});
</script>

<svelte:head>
	<title>Fiscora - Transactions</title>
</svelte:head>

<div class="mx-auto mb-8 text-center lg:mb-12">
	<h1 class="mb-4">Your transactions</h1>
	<p>Add, view, and edit your transactions to stay on top of your financial journey.</p>
	<p>Track your finances with ease and gain valuable insights.</p>
</div>

<TransactionMonthHeader {monthInfo} {monthInfoDiff} />

<div class="flex flex-wrap items-center justify-center gap-4 sm:flex-nowrap sm:justify-between sm:gap-6">
	<div class="order-last flex gap-2 sm:order-first">
		{#each IncomingTypes as incomingType}
			<button
				class="rounded-full px-4 py-2 backdrop-blur-[1px] transition-colors {incoming !== incomingType &&
					'hover:bg-primary-500/20'} {incoming === incomingType && 'variant-ghost-primary'}"
				onclick={() => (incoming = incomingType)}
			>
				{incomingType}
			</button>
		{/each}
	</div>

	<select id="month-selector" class="select w-[240px]" bind:value={month}>
		{#each listAllMonths() as [idx, name]}
			<option selected={idx === month} value={idx}>{name}</option>
		{/each}
	</select>
</div>

<TransactionList
	transactions={transactionsState}
	select={setSelectedTransaction}
	edit={setEditTransaction}
	add={addTransaction}
	remove={removeTransaction}
	{demo}
/>

<TransactionFormModal
	{transactionIntervals}
	{incomeTypes}
	{expenseTypes}
	{demo}
	transaction={editTransaction}
	open={showFormModal}
	close={closeFormModal}
	success={handleSuccess}
/>

<TransactionInfoModal transaction={selectedTransaction} close={() => setSelectedTransaction(null)} />

<button
	class="variant-filled-primary btn-icon btn-lg fixed bottom-4 right-4 rounded-full shadow-lg transition-colors duration-300 hover:shadow-xl sm:bottom-8 sm:right-8"
	onclick={() => (showFormModal = true)}
	disabled={demo}
>
	<Plus size={24} />
	<span class="sr-only">Add Budget</span>
</button>
