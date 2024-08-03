<script lang="ts">
	import { page } from "$app/stores";
	import { listAllMonthNames } from "$lib";
	import { Chart } from "chart.js/auto";
	import type { PageData } from "./$types";

	const { yearInfo }: PageData = $page.data;

	const months = listAllMonthNames();
	const incomeData: number[] = [];
	const expenseData: number[] = [];
	const netIncomeData: number[] = [];

	let yearLineChartElement: HTMLCanvasElement;

	function initCharts() {
		Chart.defaults.font.family = "Martian Mono";
		setChartFontSize();

		const ctx = yearLineChartElement.getContext("2d");
		if (ctx === null) return;

		const yearLineChart = new Chart(ctx, {
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
		});

		window.addEventListener("resize", () => {
			setChartFontSize();
			yearLineChart.resize();
		});
	}

	function setChartFontSize() {
		if (window.innerWidth < 1024) {
			Chart.defaults.font.size = 10;
		} else if (window.innerWidth < 1280) {
			Chart.defaults.font.size = 12;
		} else {
			Chart.defaults.font.size = 14;
		}
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

<title>Budget Buddy - Dashboard</title>

<div class="mx-auto my-4 text-center">
	<h1 class="my-4">Dashboard</h1>

	<div class="sm:hidden">
		<p>Dashboard view is not yet supported on mobile</p>
	</div>
	<div class="hidden sm:flex">
		<canvas id="year-line-chart" bind:this={yearLineChartElement}></canvas>
	</div>
</div>
