<script lang="ts">
	import { page } from "$app/stores";
	import { listAllMonthNamesShort } from "$lib";
	import { Chart } from "chart.js/auto";
	import type { PageData } from "./$types";

	const { yearInfo, expenseInfo }: PageData = $page.data;

	const months = listAllMonthNamesShort();
	const incomeData: number[] = [];
	const expenseData: number[] = [];
	const netIncomeData: number[] = [];

	let yearLineChartElement: HTMLCanvasElement;
	let expenseDoughnutElement: HTMLCanvasElement;

	function initCharts() {
		Chart.defaults.font.family = "Martian Mono";
		Chart.defaults.font.size = 10;

		let ctx = yearLineChartElement.getContext("2d");
		if (ctx === null) return;

		new Chart(ctx, {
			type: "line",
			data: {
				labels: months,
				datasets: [
					{
						label: "Expense",
						backgroundColor: "rgba(213, 126, 120, .3)",
						borderColor: "rgba(213, 126, 120, .75)",
						data: expenseData,
					},
					{
						label: "Income",
						backgroundColor: "rgba(132, 203, 93, .3)",
						borderColor: "rgba(132, 203, 93, .75)",
						data: incomeData,
					},
					{
						label: "Net Income",
						backgroundColor: "rgba(234, 134, 26, .3)",
						borderColor: "rgba(234, 134, 26, .75)",
						data: netIncomeData,
						fill: true,
					},
				],
			},
			options: {
				responsive: true,
				plugins: {
					legend: {
						position: "bottom",
					},
				},
			},
		});

		ctx = expenseDoughnutElement.getContext("2d");
		if (ctx === null) return;

		console.log(Object.values(expenseInfo));

		new Chart(ctx, {
			type: "doughnut",
			data: {
				labels: Object.keys(expenseInfo),
				datasets: [
					{
						label: "Amount",
						data: Object.values(expenseInfo),
					},
				],
			},
			options: {
				responsive: true,
				plugins: {
					legend: {
						position: "left",
					},
				},
			},
		});
	}

	$effect(() => {
		for (const value of Object.values(yearInfo)) {
			incomeData.push(value.income);
			expenseData.push(value.expense);
			netIncomeData.push(value.income - value.expense);
		}

		initCharts();
	});
</script>

<svelte:head>
	<title>Budget Buddy - Dashboard</title>
</svelte:head>

<div class="mx-auto my-4 text-center">
	<h1 class="my-4">Dashboard</h1>

	<div class="sm:hidden">
		<p>Dashboard view is not yet supported on mobile</p>
	</div>
	<div class="hidden sm:block">
		<div class="grid grid-cols-2 gap-4 lg:grid-cols-3">
			<div class="card col-span-2 p-4">
				<canvas bind:this={yearLineChartElement}></canvas>
			</div>
			<div class="card col-span-1 p-4">
				<canvas bind:this={expenseDoughnutElement}></canvas>
			</div>
		</div>
	</div>
</div>

<style>
	canvas {
		width: 100% !important;
		height: auto !important;
	}
</style>
