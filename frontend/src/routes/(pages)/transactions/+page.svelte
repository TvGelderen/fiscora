<script lang="ts">
	import { page } from "$app/stores";
	import type { PageData } from "./$types";
	import { tick } from "svelte";
	import TransactionList from "./transaction-list.svelte";
	import TransactionInfoModal from "./transaction-info.svelte";
	import TransactionFormModal from "./transaction-form.svelte";
	import TransactionMonthHeader from "./transaction-month-header.svelte";
	import { IncomingTypes, type Transaction, type TransactionMonthInfo } from "../../../ambient";
	import { getCurrentMonthNumber, getCurrentYear, listAllMonths } from "$lib";
	import * as Dialog from "$lib/components/ui/dialog";
	import * as Popover from "$lib/components/ui/popover/index.js";
	import { Button, buttonVariants } from "$lib/components/ui/button";
	import { Plus, CalendarIcon } from "lucide-svelte";
	import MonthPicker from "$lib/components/month-picker.svelte";
	import { zodClient } from "sveltekit-superforms/adapters";
	import { superForm } from "sveltekit-superforms";
	import { toast } from "svelte-sonner";
	import { transactionFormSchema } from "./transactionFormSchema";
	import { cn } from "$lib/utils";

	let { transactions, transactionForm, transactionIntervals, incomeTypes, expenseTypes, yearInfo, demo } =
		$page.data as PageData;

	const form = superForm(transactionForm, {
		validators: zodClient(transactionFormSchema),
		onSubmit: ({ cancel }) => {
			if (demo) {
				cancel();
				closeFormModal();
				toast.warning("Demo users are not allowed to create transactions");
			}
		},
		onError: () => {
			toast.error(`Error ${editTransaction === null ? "creating" : "updating"} transaction`);
			closeFormModal();
		},
		onUpdated: () => {
			toast.success(`Transaction ${editTransaction === null ? "created" : "updated"} successfully`);
			closeFormModal();
			updateTransactions();
		},
	});

	let showFormModal: boolean = $state(false);
	let year: number = $state(getCurrentYear());
	let month: number = $state(getCurrentMonthNumber());
	let incoming = $state(IncomingTypes[0]);
	let transactionsState: Transaction[] = $state(transactions);
	let monthInfo: TransactionMonthInfo | null = $state(null);
	let monthInfoDiff: TransactionMonthInfo | null = $state(null);
	let selectedTransaction: Transaction | null = $state(null);
	let editTransaction: Transaction | null = $state(null);

	const months = listAllMonths();

	function setSelectedTransaction(transaction: Transaction | null) {
		selectedTransaction = transaction;
	}

	async function setEditTransaction(transaction: Transaction | null) {
		editTransaction = transaction;
		await tick();
		showFormModal = true;
	}

	function closeFormModal() {
		editTransaction = null;
		showFormModal = false;
	}

	async function updateTransactions() {
		const response = await fetch(`/api/transactions?month=${month}&year=${year}`);
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

	async function handleMonthChanged() {
		await updateTransactions();
		updateMonthInfo();
	}

	async function removeTransaction(transaction: Transaction) {
		if (demo) {
			toast.warning("You are not allowed to delete transactions as a demo user");
			return;
		}

		const idx = transactions.findIndex((t) => t.id === transaction.id);
		if (idx !== -1) {
			transactionsState.splice(idx, 1);
			transactionsState = [...transactionsState];
			updateYearInfo(transaction.amount, false);
		}

		const response = await fetch(`/api/transactions/${transaction.id}`, { method: "DELETE" });
		if (!response.ok) {
			toast.error("Something went wrong trying to delete transaction");
			if (idx !== -1) {
				transactionsState.splice(idx, 0, transaction);
				transactionsState = [...transactionsState];
				updateYearInfo(transaction.amount, true);
			}
			return;
		}

		toast.success("Transaction deleted successfully");
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
		transactionsState = transactions.filter((t) => {
			return (
				incoming === IncomingTypes[0] ||
				(incoming === IncomingTypes[1] && t.amount > 0) ||
				(incoming === IncomingTypes[2] && t.amount < 0)
			);
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
					'hover:bg-primary/10'} {incoming === incomingType && 'bg-primary/10 ring-[1px] ring-primary/50'}"
				onclick={() => (incoming = incomingType)}
			>
				{incomingType}
			</button>
		{/each}
	</div>

	<Popover.Root>
		<Popover.Trigger
			class={cn(buttonVariants({ variant: "outline" }), "!h-fit w-[240px] justify-start text-left text-base")}
		>
			<CalendarIcon class="mr-2 h-5 w-5" />
			{months.get(month)}
		</Popover.Trigger>
		<Popover.Content class="w-auto p-0">
			<MonthPicker bind:year bind:month callback={handleMonthChanged} />
		</Popover.Content>
	</Popover.Root>
</div>

<TransactionList
	transactions={transactionsState}
	income={incoming}
	select={setSelectedTransaction}
	edit={setEditTransaction}
	remove={removeTransaction}
/>

<Dialog.Root
	bind:open={showFormModal}
	onOpenChange={(open) => {
		if (!open) {
			editTransaction = null;
		}
	}}
>
	<TransactionFormModal
		{form}
		{transactionIntervals}
		{incomeTypes}
		{expenseTypes}
		{demo}
		transaction={editTransaction}
	/>
</Dialog.Root>

<Dialog.Root
	open={selectedTransaction !== null}
	onOpenChange={(open) => {
		if (!open) {
			editTransaction = null;
		}
	}}
>
	<TransactionInfoModal transaction={selectedTransaction} />
</Dialog.Root>

<button
	class="fixed bottom-5 right-4 rounded-full bg-primary p-2 text-slate-50 shadow-md !shadow-slate-500 transition-all duration-300 hover:shadow-lg dark:!shadow-slate-800 sm:bottom-8 sm:right-8"
	onclick={() => (showFormModal = true)}
	disabled={demo}
>
	<Plus size={24} />
	<span class="sr-only">Add Budget</span>
</button>
