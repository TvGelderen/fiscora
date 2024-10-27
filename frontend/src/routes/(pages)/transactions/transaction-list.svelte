<script lang="ts">
	import { getFormattedAmount, getFormattedDateShort } from "$lib";
	import * as AlertDialog from "$lib/components/ui/alert-dialog";
	import { Edit, Trash } from "lucide-svelte";
	import { fly } from "svelte/transition";
	import { type Transaction } from "../../../ambient";
	import { buttonVariants } from "$lib/components/ui/button";

	let {
		transactions,
		income,
		select,
		edit,
		remove,
	}: {
		transactions: Transaction[];
		income: string;
		select: (t: Transaction | null) => void;
		edit: (t: Transaction | null) => void;
		remove: (t: Transaction) => void;
	} = $props();

	let transactionToDelete: Transaction | null = $state(null);

	function openDeleteModal(event: MouseEvent, transaction: Transaction) {
		event.stopPropagation();
		transactionToDelete = transaction;
	}

	function closeDeleteModal() {
		transactionToDelete = null;
	}

	function confirmDelete() {
		if (transactionToDelete !== null) {
			const transaction = transactionToDelete;
			closeDeleteModal();
			remove(transaction);
		}
	}

	async function editTransaction(event: MouseEvent, transaction: Transaction) {
		event.stopPropagation();
		edit(transaction);
	}
</script>

<div
	class="w-full overflow-x-auto"
	style="scrollbar-color: rgba(128,128,128,0.5) rgba(0,0,0,0); scrollbar-width: thin;"
>
	<table class="mt-4 w-full select-none overflow-hidden rounded-md text-left [&_th]:p-4">
		<thead>
			<tr>
				<th class="w-[10%]">Date</th>
				<th class="w-[45%] min-w-[200px]">Description</th>
				<th class="w-[20%] text-right">Amount</th>
				<th class="w-[20%]">Type</th>
				<th class="w-[5%]"></th>
			</tr>
		</thead>
		<tbody class="transactions-table-body">
			{#each transactions as transaction, i (`${transaction.id}-${income}`)}
				<tr
					class="transactions-table-row"
					onclick={() => select(transaction)}
					in:fly={{
						y: 100,
						delay: 25 * i,
						duration: 200,
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
						<div class="flex justify-end gap-1">
							<button
								class={buttonVariants({ size: "icon", variant: "ghost" })}
								onclick={(event) => editTransaction(event, transaction)}
							>
								<Edit size={20} />
							</button>
							<button
								class={`${buttonVariants({ size: "icon", variant: "destructive" })} bg-transparent`}
								onclick={(event) => openDeleteModal(event, transaction)}
							>
								<Trash size={20} />
							</button>
						</div>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
	{#if transactions.length === 0}
		<p class="ml-4">You have no registered transactions for this month.</p>
	{/if}
</div>

<AlertDialog.Root open={transactionToDelete !== null}>
	<AlertDialog.Content>
		<AlertDialog.Header>
			<h3>Confirm Deletion</h3>
		</AlertDialog.Header>
		{#if transactionToDelete !== null}
			<p>Are you sure you want to delete this transaction? This action is permanent and cannot be undone.</p>
			{#if transactionToDelete.recurring !== null}
				<p class="mt-2 text-orange-200">
					This is a recurring transaction. Deleting it will remove all past and future occurrences.
				</p>
			{/if}
		{/if}
		<AlertDialog.Footer>
			<AlertDialog.Cancel onclick={closeDeleteModal}>Cancel</AlertDialog.Cancel>
			<AlertDialog.Action class={buttonVariants({ variant: "destructive" })} onclick={confirmDelete}>
				Delete
			</AlertDialog.Action>
		</AlertDialog.Footer>
	</AlertDialog.Content>
</AlertDialog.Root>
