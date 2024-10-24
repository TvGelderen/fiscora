<script lang="ts">
	import { type Transaction } from "../../ambient";
	import { getToastStore } from "@skeletonlabs/skeleton";
	import { createRandomString, getFormattedAmount, getFormattedDateShort } from "$lib";
	import { Edit, Trash, X } from "lucide-svelte";
	import { fly } from "svelte/transition";

	const toastStore = getToastStore();

	let {
		transactions,
		select,
		edit,
		add,
		remove,
		demo,
	}: {
		transactions: Transaction[];
		select: (t: Transaction | null) => void;
		edit: (t: Transaction | null) => void;
		add: (t: Transaction, idx: number) => void;
		remove: (t: Transaction, idx: number) => void;
		demo: boolean;
	} = $props();

	let modal: HTMLDialogElement;
	let transactionToDelete: Transaction | null = $state(null);

	function openDeleteModal(event: MouseEvent, transaction: Transaction) {
		event.stopPropagation();
		transactionToDelete = transaction;
		modal.showModal();
	}

	function closeDeleteModal() {
		transactionToDelete = null;
		modal.close();
	}

	async function confirmDelete() {
		if (transactionToDelete !== null) {
			const id = transactionToDelete.id;
			closeDeleteModal();
			await deleteTransaction(id);
		}
	}

	async function editTransaction(event: MouseEvent, id: number) {
		event.stopPropagation();

		let transaction = transactions.find((t) => t.id === id);
		if (!transaction) return;

		edit(transaction);
	}

	async function deleteTransaction(id: number) {
		if (demo) {
			toastStore.trigger({
				message: "You are not allowed to delete transactions as a demo user",
				background: "variant-filled-warning",
			});
			return;
		}

		const idx = transactions.findIndex((t) => t.id === id);
		const transaction = transactions.at(idx);
		if (transaction !== undefined) {
			remove(transaction, idx);
		}

		const response = await fetch(`/api/transactions/${id}`, { method: "DELETE" });
		if (!response.ok) {
			toastStore.trigger({
				message: "Something went wrong trying to delete transaction",
				background: "variant-filled-error",
			});
			if (transaction !== undefined) {
				add(transaction, idx);
			}
			return;
		}

		toastStore.trigger({
			message: "Transaction deleted successfully",
			background: "variant-filled-success",
		});
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
			{#each transactions as transaction, i (createRandomString(8))}
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
								class="icon inline rounded-md p-2 hover:bg-primary-500/25 hover:!text-black dark:hover:bg-primary-500/50 dark:hover:!text-white"
								onclick={(event) => editTransaction(event, transaction.id)}
							>
								<Edit size={20} />
							</button>
							<button
								class="icon inline rounded-md p-2 hover:bg-error-500/60 hover:!text-black dark:hover:!text-white"
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

<dialog class="max-w-md" bind:this={modal}>
	<button class="absolute right-4 top-4" onclick={closeDeleteModal}>
		<X />
	</button>
	{#if transactionToDelete !== null}
		<h3 class="mb-4">Confirm Deletion</h3>
		<p>Are you sure you want to delete this transaction? This action is permanent and cannot be undone.</p>
		{#if transactionToDelete.recurring !== null}
			<p class="mt-2 text-warning-500">
				This is a recurring transaction. Deleting it will remove all past and future occurrences.
			</p>
		{/if}
		<div class="mt-4 flex justify-end gap-2">
			<button class="!variant-filled-surface btn" onclick={closeDeleteModal}>Cancel</button>
			<button class="!variant-filled-error btn" onclick={confirmDelete}>Delete</button>
		</div>
	{/if}
</dialog>
