<script lang="ts">
	import type { TransactionMonthInfo } from "../../ambient";

	let { monthInfo }: { monthInfo: Promise<TransactionMonthInfo> | null } =
		$props();
</script>

{#if monthInfo === null}
	{@render header(null)}
{:else}
	{#await monthInfo}
		{@render header(null)}
	{:then monthInfo}
		{@render header(monthInfo)}
	{/await}
{/if}

{#snippet header(info: TransactionMonthInfo | null)}
	<div
		class="mb-10 grid rounded-2xl bg-primary-500/20 shadow-md shadow-primary-900/50 dark:shadow-surface-900 sm:grid-cols-3 lg:mb-16"
	>
		<div
			class="flex flex-col items-center justify-between p-4 sm:items-start"
		>
			<h4 class="mb-6">Total income</h4>
			<span class="text-2xl lg:text-3xl">
				{#if info === null}
					€
				{:else}
					€{info.income}
				{/if}
			</span>
		</div>
		<div
			class="flex flex-col items-center justify-between border-b-[1px] border-t-[1px] border-primary-700/25 p-4 sm:items-start sm:border-b-[0px] sm:border-l-[1px] sm:border-r-[1px] sm:border-t-[0px]"
		>
			<h4 class="mb-6">Total expense</h4>
			<span class="text-2xl lg:text-3xl">
				{#if info === null}
					€
				{:else}
					€{info.expense}
				{/if}
			</span>
		</div>
		<div
			class="flex flex-col items-center justify-between p-4 sm:items-start"
		>
			<h4 class="mb-6">Net income</h4>
			<span class="text-2xl lg:text-3xl">
				{#if info === null}
					€
				{:else}
					€{info.income - info.expense}
				{/if}
			</span>
		</div>
	</div>
{/snippet}
