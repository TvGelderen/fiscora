<script lang="ts">
	import { page } from "$app/stores";
	import { listAllMonthNames } from "$lib";
	import { Chart } from "chart.js/auto";
	import type { PageData } from "./$types";

	const { yearInfo }: PageData = $page.data;

	const months = listAllMonthNames();
	const incomeData: number[] = [];
	const expenseData: number[] = [];

	let yearLineChartElement: HTMLCanvasElement;

	function initCharts() {
		const ctx = yearLineChartElement.getContext("2d");
		if (ctx === null) return;

		new Chart(ctx, {
			type: "line",
			data: {
				labels: months,
				datasets: [
					{
						label: "Expense",
						backgroundColor: "rgba(225, 204,230, .3)",
						borderColor: "rgb(205, 130, 158)",
						data: expenseData,
					},
					{
						label: "Income",
						backgroundColor: "rgba(184, 185, 210, .3)",
						borderColor: "rgb(35, 26, 136)",
						data: incomeData,
					},
				],
			},
		});
	}

	$effect(() => {
		for (const value of Object.values(yearInfo)) {
			incomeData.push(value.income);
			expenseData.push(value.expense);
		}

		initCharts();
	});
</script>

<title>Budget Buddy - Dashboard</title>

<div class="mx-auto my-4 text-center">
	<h1 class="my-4">Dashboard</h1>

	<canvas id="year-line-chart" bind:this={yearLineChartElement}></canvas>
</div>
