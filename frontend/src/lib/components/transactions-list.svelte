<script lang="ts">
	import EllipsisVertical from "lucide-svelte/icons/ellipsis-vertical";
	import { IncomingTypes, type Transaction } from "../../ambient";
	import { getToastStore, popup } from "@skeletonlabs/skeleton";
	import click from "$lib/click";
	import { getFormattedAmount, getFormattedDateShort } from "$lib";
	import { Edit, Trash } from "lucide-svelte";
	import { fly } from "svelte/transition";
	import { tick } from "svelte";

	let {
		transactions,
		incoming,
		select,
		edit,
		demo,
	}: {
		transactions: Promise<Transaction[]> | null;
		incoming: string;
		select: (t: Transaction | null) => void;
		edit: (t: Transaction | null) => void;
		demo: boolean;
	} = $props();

	let transactionsList: Transaction[] | null = $state(null);

	const toastStore = getToastStore();

	async function handleEditTransaction(event: MouseEvent) {
		event.stopPropagation();

		const id = Number.parseInt(getId(event.target!) ?? "");
		if (!id) return;

		let transaction = (await transactions)?.find((t) => t.id === id);
		if (!transaction) return;

		edit(transaction);
	}

	async function handleDeleteTransaction(event: MouseEvent) {
		event.stopPropagation();

		if (demo) {
			toastStore.trigger({
				message:
					"You are not allowed to delete transactions as a demo user",
				background: "variant-filled-warning",
			});
			return;
		}

		const id = Number.parseInt(getId(event.target!) ?? "");
		if (!id) return;

		const response = await fetch(`/api/transactions/${id}`, {
			method: "DELETE",
		});
		if (response.ok) {
			toastStore.trigger({
				background: "bg-success-400 text-black",
				message: "Transaction deleted successfully",
			});

			const updatedTransactions = transactionsList?.filter(
				(t) => t.id !== id,
			);
			if (!updatedTransactions) return;

			transactionsList = updatedTransactions;

			return;
		}

		toastStore.trigger({
			message: "Something went wrong trying to delete transaction",
			background: "bg-error-400 text-black",
		});
	}

	const getId = (target: EventTarget) => (target as HTMLElement).dataset.id;

	$effect(() => {
		if (transactions === null) return;

		const all = incoming === IncomingTypes[0];

		transactions.then((data) => {
			transactionsList = [];
			tick().then(() => {
				transactionsList = data.filter((t) => {
					const date = new Date(t.date);
					date.setUTCMilliseconds(date.getMilliseconds() + 1);
					t.date = date;
					if (all) return true;
					return (
						(incoming === IncomingTypes[1] && t.amount > 0) ||
						(incoming === IncomingTypes[2] && t.amount < 0)
					);
				});
			});
		});
	});
</script>

<div
	class="w-full overflow-x-auto"
	style="scrollbar-color: rgba(128,128,128,0.5) rgba(0,0,0,0); scrollbar-width: thin;"
>
	{#if transactions === null}
		{@render skeletonTable()}
	{/if}
	{#if transactionsList}
		<table
			class="mt-4 w-full select-none overflow-hidden rounded-md text-left [&_th]:p-4"
		>
			{@render tableHead()}
			<tbody class="transactions-table-body">
				{#each transactionsList as transaction, i}
					<tr
						class="transactions-table-row"
						use:click={() => select(transaction)}
						in:fly={{
							y: 100,
							delay: 25 * i,
							duration: 250,
						}}
					>
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
						<td data-cell="">
							<button
								class="icon"
								onclick={(event) => event.stopPropagation()}
								use:popup={{
									event: "click",
									target: `popup-${i}`,
									placement: "bottom",
								}}
							>
								<EllipsisVertical size={20} />
							</button>
							<div
								class="bg-surface-50-900-token rounded-md p-4 shadow-lg"
								data-popup="popup-{i}"
							>
								<div class="flex flex-col gap-4">
									<button
										class="flex items-center gap-3"
										onclick={handleEditTransaction}
										data-id={transaction.id}
									>
										<Edit size={20} /> Edit
									</button>
									<button
										class="flex items-center gap-3 text-error-700 disabled:opacity-50 dark:text-error-500"
										onclick={handleDeleteTransaction}
										data-id={transaction.id}
										disabled={transaction.recurring !==
											null}
									>
										<Trash size={20} /> Delete
									</button>
									<div
										class="bg-surface-50-900-token arrow"
									></div>
								</div>
							</div>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
		{#if transactionsList.length === 0}
			<p class="ml-4">
				You have no registered transactions for this month.
			</p>
		{/if}
	{/if}
</div>

{#snippet tableHead()}
	<thead>
		<tr>
			<th class="w-[10%]">Date</th>
			<th class="w-[45%] min-w-[200px]">Description</th>
			<th class="w-[20%] text-right">Amount</th>
			<th class="w-[20%]">Type</th>
			<th class="w-[5%]"></th>
		</tr>
	</thead>
{/snippet}

{#snippet tableRow(className: string = "[&>td]:p-[.125rem]")}
	<tr class={className}>
		<td>
			<div></div>
		</td>
		<td>
			<div></div>
		</td>
		<td>
			<div></div>
		</td>
		<td>
			<div></div>
		</td>
		<td>
			<div></div>
		</td>
	</tr>
{/snippet}

{#snippet skeletonTable()}
	<table class="mt-4 w-full rounded-md text-left [&_th]:p-4">
		{@render tableHead()}
		<tbody class="transactions-table-body">
			{#each { length: 5 } as i}
				{@render tableRow(`transactions-table-row skeleton ${i}`)}
				{@render tableRow()}
			{/each}
		</tbody>
	</table>
	<span class="sr-only">Loading...</span>
{/snippet}
