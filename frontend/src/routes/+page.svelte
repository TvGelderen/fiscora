<script lang="ts">
	import { page } from "$app/stores";
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
		{#if darkMode.darkMode}
			<enhanced:img
				src="/static/img/transactions-dark.png"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md shadow-md shadow-primary-900/50 dark:shadow-surface-900"
			/>
		{:else}
			<enhanced:img
				src="/static/img/transactions-light.png"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md shadow-md shadow-primary-900/50 dark:shadow-surface-900"
			/>
		{/if}
	</div>
</div>

<div class="grid grid-cols-5 items-center gap-12">
	<div class="col-span-3">
		{#if darkMode.darkMode}
			<enhanced:img
				src="/static/img/dashboard-dark.png"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md shadow-md shadow-primary-900/50 dark:shadow-surface-900"
			/>
		{:else}
			<enhanced:img
				src="/static/img/dashboard-light.png"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md shadow-md shadow-primary-900/50 dark:shadow-surface-900"
			/>
		{/if}
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
