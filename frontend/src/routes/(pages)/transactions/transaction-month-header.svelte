<script lang="ts">
	import CountTo from "$lib/components/count-to.svelte";
	import type { TransactionMonthInfo } from "../../../ambient";

	let {
		monthInfo,
		monthInfoDiff,
	}: {
		monthInfo: TransactionMonthInfo | null;
		monthInfoDiff: TransactionMonthInfo | null;
	} = $props();

	let income = $state(0);
	let expense = $state(0);
	let oldIncome = $state(0);
	let oldExpense = $state(0);
	let incomeDiff = $state(0);
	let expenseDiff = $state(0);
	let oldIncomeDiff = $state(0);
	let oldExpenseDiff = $state(0);
	let netIncome = $derived(income - expense);
	let oldNetIncome = $derived(oldIncome - oldExpense);
	let netIncomeDiff = $derived(incomeDiff - expenseDiff);
	let oldNetIncomeDiff = $derived(oldIncomeDiff - oldExpenseDiff);

	$effect(() => {
		if (monthInfo === null) return;

		oldIncome = income;
		oldExpense = expense;
		income = monthInfo.income;
		expense = monthInfo.expense;

		if (monthInfoDiff === null) return;

		oldIncomeDiff = incomeDiff;
		oldExpenseDiff = expenseDiff;
		incomeDiff = monthInfoDiff.income;
		expenseDiff = monthInfoDiff.expense;
	});
</script>

<div
	class="mx-auto mb-10 grid max-w-[240px] rounded-2xl border border-primary/25 bg-primary/10 shadow-md shadow-primary/20 backdrop-blur-[1px] dark:shadow-background sm:max-w-full sm:grid-cols-3 lg:mb-16"
>
	<div class="flex flex-col items-center justify-between p-2 sm:items-start md:p-4">
		<h4 class="mb-2 md:mb-6">Total income</h4>
		<span class="mb-1 text-xl lg:text-3xl">
			<CountTo start={oldIncome} value={income} />
		</span>
		<span class="text-sm md:text-base">
			<CountTo start={oldIncomeDiff} value={incomeDiff} /> from last month
		</span>
	</div>
	<div
		class="flex flex-col items-center justify-between border-b-[1px] border-t-[1px] border-primary/25 p-2 sm:items-start sm:border-b-[0px] sm:border-l-[1px] sm:border-r-[1px] sm:border-t-[0px] md:p-4"
	>
		<h4 class="mb-2 md:mb-6">Total expense</h4>
		<span class="mb-1 text-xl lg:text-3xl">
			<CountTo start={oldExpense} value={expense} />
		</span>
		<span class="text-sm md:text-base">
			<CountTo start={oldExpenseDiff} value={expenseDiff} /> from last month
		</span>
	</div>
	<div class="flex flex-col items-center justify-between p-2 sm:items-start md:p-4">
		<h4 class="mb-2 md:mb-6">Net income</h4>
		<span class="mb-1 text-xl lg:text-3xl">
			<CountTo start={oldNetIncome} value={netIncome} />
		</span>
		<span class="text-sm md:text-base">
			<CountTo start={oldNetIncomeDiff} value={netIncomeDiff} /> from last month
		</span>
	</div>
</div>
