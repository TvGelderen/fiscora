<script lang="ts">
	import type { Transaction } from "../../../ambient";
	import * as Dialog from "$lib/components/ui/dialog";
	import { getFormattedDate } from "$lib";

	const {
		transaction,
	}: {
		transaction: Transaction | null;
	} = $props();
</script>

<Dialog.Content class="w-full max-w-2xl">
	<Dialog.Header>
		<h2>Transaction details</h2>
	</Dialog.Header>
	{#if transaction !== null}
		<div class="mt-6 grid grid-cols-1 sm:grid-cols-2 lg:mt-8">
			<div class="mb-4">
				<p class="header-sm">Transaction date</p>
				<p>{getFormattedDate(transaction.date)}</p>
			</div>
			<div class="mb-4">
				<p class="header-sm">Amount</p>
				<p>€{transaction.amount}</p>
			</div>
			<div class="mb-4">
				<p class="header-sm">Type</p>
				<p>{transaction.type}</p>
			</div>
			{#if transaction.recurring != null}
				<div class="mb-4">
					<p class="header-sm">Recurring</p>
					<p>{transaction.recurring.interval}</p>
				</div>
				<div class="mb-4">
					<p class="header-sm">Start date</p>
					<p>{getFormattedDate(transaction.recurring.startDate!)}</p>
				</div>
				<div class="mb-4">
					<p class="header-sm">End date</p>
					<p>{getFormattedDate(transaction.recurring.endDate!)}</p>
				</div>
			{/if}
		</div>
		<div>
			<p class="header-sm">Description</p>
			<p>{transaction.description}</p>
		</div>
	{:else}
		<div><span class="text-error-400">No transaction selected.</span></div>
	{/if}
</Dialog.Content>
