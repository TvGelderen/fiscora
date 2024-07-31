<script lang="ts">
	import type { TransactionMonthInfo } from "../../ambient";
	import CountTo from "./countTo.svelte";

	let { monthInfo }: { monthInfo: Promise<TransactionMonthInfo> | null } =
		$props();

	let income = $state(0);
	let expense = $state(0);
	let netIncome = $derived(income - expense);
	let oldIncome = $state(0);
	let oldExpense = $state(0);
	let oldNetIncome = $derived(oldIncome - oldExpense);

	$effect(() => {
		if (monthInfo === null) return;

		monthInfo.then((info) => {
			oldIncome = income;
			oldExpense = expense;
			income = info.income;
			expense = info.expense;
		});
	});
</script>

<div
	class="mb-10 grid rounded-2xl bg-primary-500/20 shadow-md shadow-primary-900/50 dark:shadow-surface-900 sm:grid-cols-3 lg:mb-16"
>
	<div class="flex flex-col items-center justify-between p-4 sm:items-start">
		<h4 class="mb-6">Total income</h4>
		<span class="text-2xl lg:text-3xl">
			{#key income}
				€<CountTo start={oldIncome} value={income} />
			{/key}
		</span>
	</div>
	<div
		class="flex flex-col items-center justify-between border-b-[1px] border-t-[1px] border-primary-700/25 p-4 sm:items-start sm:border-b-[0px] sm:border-l-[1px] sm:border-r-[1px] sm:border-t-[0px]"
	>
		<h4 class="mb-6">Total expense</h4>
		<span class="text-2xl lg:text-3xl">
			{#key expense}
				€<CountTo start={oldExpense} value={expense} />
			{/key}
		</span>
	</div>
	<div class="flex flex-col items-center justify-between p-4 sm:items-start">
		<h4 class="mb-6">Net income</h4>
		<span class="text-2xl lg:text-3xl">
			{#key netIncome}
				€<CountTo start={oldNetIncome} value={netIncome} />
			{/key}
		</span>
	</div>
</div>
