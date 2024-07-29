<script lang="ts">
	import EllipsisVertical from "lucide-svelte/icons/ellipsis-vertical";
	import { IncomingTypes, type Transaction } from "../../ambient";
	import click from "$lib/click";
	import { getFormattedDateShort } from "$lib";

	const {
		transactions,
		incoming,
		selectTransaction,
	}: {
		transactions: Promise<Transaction[]> | null;
		incoming: string;
		selectTransaction: (t: Transaction | null) => void;
	} = $props();
</script>

<div class="w-full overflow-auto">
	{#if transactions === null}
		{@render skeletonTable()}
	{:else}
		{#await transactions}
			{@render skeletonTable()}
		{:then transactions}
			<table class="mt-4 w-full rounded-md text-left [&_th]:p-4">
				{@render tableHead()}
				<tbody class="transactions-table-body">
					{#each transactions.filter((t) => {
						if (incoming === IncomingTypes[0]) return true;
						return (incoming === IncomingTypes[1] && t.incoming) || (incoming === IncomingTypes[2] && !t.incoming);
					}) as transaction}
						<tr
							class="transactions-table-row"
							use:click={() => selectTransaction(transaction)}
						>
							<td data-cell="date">
								{getFormattedDateShort(transaction.date)}
							</td>
							<td data-cell="description">
								{transaction.description}
							</td>
							<td data-cell="amount">
								{transaction.incoming
									? ""
									: "-"}{transaction.amount}
							</td>
							<td data-cell="type">{transaction.type}</td>
							<td data-cell="">
								<button class="icon">
									<EllipsisVertical size={20} />
								</button>
							</td>
						</tr>
						{@render rowSpacer()}
					{/each}
				</tbody>
			</table>
			{#if transactions.length === 0}
				<p class="ml-4">
					You have no registered transactions for this month.
				</p>
			{/if}
		{/await}
	{/if}
</div>

{#snippet tableHead()}
	<thead>
		<tr>
			<th class="w-[10%]">Date</th>
			<th class="w-[50%] min-w-[200px]">Description</th>
			<th class="w-[20%] text-right">Amount</th>
			<th class="w-[15%]">Type</th>
			<th class="w-[5%]"></th>
		</tr>
	</thead>
{/snippet}

{#snippet rowSpacer(className: string = "[&>td]:p-[.125rem]")}
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
				{@render rowSpacer(`transactions-table-row skeleton ${i}`)}
				{@render rowSpacer()}
			{/each}
		</tbody>
	</table>
{/snippet}
