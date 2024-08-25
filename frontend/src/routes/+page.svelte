<script lang="ts">
	import { page } from "$app/stores";
	import transactionsLight from "$lib/assets/img/transactions-light.png";
	import transactionsDark from "$lib/assets/img/transactions-dark.png";
	import dashboardLight from "$lib/assets/img/dashboard-light.png";
	import dashboardDark from "$lib/assets/img/dashboard-dark.png";
	import { createDarkMode } from "$lib/theme.svelte";

	const { loggedIn } = $page.data;

	const darkMode = createDarkMode();
</script>

<div class="h-full-page grid grid-cols-5 items-center gap-12">
	<div class="col-span-2 flex w-full flex-col">
		<h1 class="mb-6">Take control of your finances</h1>
		<p>
			Budget Buddy helps you track your income, expenses, and savings all
			in one place. Gain valuable insights, set goals, and make smarter
			financial decisions.
		</p>

		{#if !loggedIn}
			<a class="btn mt-4 w-[200px]" href="/login">Sign up now!</a>
		{/if}
	</div>
	<div class="col-span-3">
		<img
			src={darkMode.darkMode ? transactionsDark : transactionsLight}
			alt="Transactions page in light and darkmode"
			class="rounded-md shadow-md shadow-primary-900/50 dark:shadow-surface-900"
		/>
	</div>
</div>

<div class="grid grid-cols-5 items-center gap-12">
	<div class="col-span-3">
		<img
			src={darkMode.darkMode ? dashboardDark : dashboardLight}
			alt="Dashboard page in light and darkmode"
			class="rounded-md shadow-md shadow-primary-900/50 dark:shadow-surface-900"
		/>
	</div>
	<div class="col-span-2 flex w-full flex-col">
		<h1 class="mb-6">Dashboard</h1>
		<p>
			Get your financial snapshot at a glance using the dashboard page.
			Track expenses, monitor budgets, and visualize your progress towards
			financial goals all in one place.
		</p>
	</div>
</div>

<div class="mb-24"></div>
