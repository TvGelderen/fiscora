<script lang="ts">
	import { listAllMonthNamesShort } from "$lib";
	import { Chart } from "chart.js/auto";
	import { createDarkMode } from "$lib/theme.svelte";

	let { data } = $props();
	let { yearInfo, incomeInfo, expenseInfo } = data;

	const darkMode = createDarkMode();
	const months = listAllMonthNamesShort();
	const incomeData: number[] = [];
	const expenseData: number[] = [];
	const netIncomeData: number[] = [];
	const charts: Chart[] = [];

	let accumulatedNetIncomeData: number[] = [];
	let yearLineChartElement: HTMLCanvasElement;
	let expenseDoughnutElement: HTMLCanvasElement;
	let incomeDoughnutElement: HTMLCanvasElement;
	let netIncomeLineChartElement: HTMLCanvasElement;

	export const COLORS = {
		red: "rgb(255, 99, 132)",
		orange: "rgb(255, 159, 64)",
		yellow: "rgb(255, 205, 86)",
		green: "rgb(75, 192, 192)",
		blue: "rgb(54, 162, 235)",
		purple: "rgb(153, 102, 255)",
		grey: "rgb(201, 203, 207)",
	};

	export const FILL_COLORS = {
		red: "rgba(255, 99, 132, .25)",
		orange: "rgba(255, 159, 64, .25)",
		yellow: "rgba(255, 205, 86, .25)",
		green: "rgba(75, 192, 192, .25)",
		blue: "rgba(54, 162, 235, .25)",
		purple: "rgba(153, 102, 255, .25)",
		grey: "rgba(201, 203, 207, .25)",
	};

	function initCharts() {
		Chart.defaults.font.family = "Inter";
		Chart.defaults.responsive = true;
		Chart.defaults.plugins.legend.position = "bottom";
		Chart.defaults.scale.grid.color = "rgba(0,0,0,0)";
		Chart.defaults.borderColor = "rgba(0,0,0,0)";

		let ctx = yearLineChartElement.getContext("2d");
		if (ctx === null) return;

		charts.push(
			new Chart(ctx, {
				type: "line",
				data: {
					labels: months,
					datasets: [
						{
							label: "Income",
							data: incomeData,
							pointRadius: 0,
							tension: 0.25,
							borderColor: COLORS.yellow,
							backgroundColor: FILL_COLORS.yellow,
						},
						{
							label: "Expense",
							data: expenseData,
							pointRadius: 0,
							tension: 0.25,
							borderColor: COLORS.red,
							backgroundColor: FILL_COLORS.red,
						},
						{
							label: "Net Income",
							data: netIncomeData,
							pointRadius: 0,
							tension: 0.25,
							fill: true,
							borderColor: COLORS.green,
							backgroundColor: FILL_COLORS.green,
						},
					],
				},
				options: {
					scales: {
						y: {
							min: 0,
						},
					},
					interaction: {
						intersect: false,
					},
				},
			}),
		);

		ctx = expenseDoughnutElement.getContext("2d");
		if (ctx === null) return;

		charts.push(
			new Chart(ctx, {
				type: "doughnut",
				data: {
					labels: Object.keys(expenseInfo),
					datasets: [
						{
							label: "Amount",
							data: Object.values(expenseInfo),
							borderWidth: 0,
							backgroundColor: Object.values(COLORS),
						},
					],
				},
			}),
		);

		ctx = incomeDoughnutElement.getContext("2d");
		if (ctx === null) return;

		charts.push(
			new Chart(ctx, {
				type: "doughnut",
				data: {
					labels: Object.keys(incomeInfo),
					datasets: [
						{
							label: "Amount",
							data: Object.values(incomeInfo),
							borderWidth: 0,
							backgroundColor: Object.values(COLORS),
						},
					],
				},
			}),
		);

		ctx = netIncomeLineChartElement.getContext("2d");
		if (ctx === null) return;

		charts.push(
			new Chart(ctx, {
				type: "line",
				data: {
					labels: months,
					datasets: [
						{
							label: "Net Income",
							data: accumulatedNetIncomeData,
							pointRadius: 0,
							tension: 0.25,
							fill: true,
							borderColor: COLORS.green,
							backgroundColor: FILL_COLORS.green,
						},
					],
				},
				options: {
					scales: {
						y: {
							min: 0,
						},
					},
					interaction: {
						intersect: false,
					},
				},
			}),
		);
	}

	$effect(() => {
		for (const value of yearInfo) {
			incomeData.push(value[1].income);
			expenseData.push(value[1].expense);
			netIncomeData.push(value[1].income - value[1].expense);
			accumulatedNetIncomeData = [netIncomeData[0]];
			for (let i = 1; i < netIncomeData.length; i++) {
				accumulatedNetIncomeData[i] = netIncomeData[i] + accumulatedNetIncomeData[i - 1];
			}
		}
	});

	$effect(() => {
		const color = darkMode.darkMode ? "rgb(251, 231, 209)" : "rgb(2, 8, 23)";

		Chart.defaults.color = color;
		Chart.defaults.scale.ticks.color = color;

		charts.forEach((chart) => {
			chart.destroy();
		});

		initCharts();
	});
</script>

<svelte:head>
	<title>Fiscora - Dashboard</title>
</svelte:head>

<div class="mx-auto text-center">
	<div class="mb-6">
		<h1 class="mb-4">Dashboard</h1>
		<p>Your financial snapshot at a glance.</p>
		<p>Track expenses, monitor budgets, and visualize your progress towards financial goals.</p>
	</div>

	<div class="sm:hidden">
		<p>Dashboard view is not yet supported on mobile</p>
	</div>
	<div class="hidden sm:block">
		<div class="grid grid-cols-2 gap-4 lg:grid-cols-3">
			<div class="card col-span-2 p-4 shadow-lg">
				<p class="mb-2">Income, expense, and net income</p>
				<canvas bind:this={yearLineChartElement}></canvas>
			</div>
			<div class="card col-span-1 p-4 shadow-lg">
				<p class="mb-2">Average expenses</p>
				<canvas bind:this={expenseDoughnutElement}></canvas>
			</div>
			<div class="card col-span-1 p-4 shadow-lg">
				<p class="mb-2">Average income</p>
				<canvas bind:this={incomeDoughnutElement}></canvas>
			</div>
			<div class="card col-span-2 p-4 shadow-lg">
				<p class="mb-2">Accumulated net-income</p>
				<canvas bind:this={netIncomeLineChartElement}></canvas>
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
