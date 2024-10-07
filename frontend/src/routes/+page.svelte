<script lang="ts">
	import { page } from "$app/stores";
	import { createDarkMode } from "$lib/theme.svelte";

	const { loggedIn } = $page.data;

	const darkMode = createDarkMode();
</script>

<div
	class="relative my-10 flex min-h-[50dvh] flex-col items-center justify-center gap-12 text-center lg:flex-row lg:justify-evenly lg:text-left"
>
	<div class="flex w-full max-w-[700px] flex-col items-center lg:items-start">
		<h1 class="mb-6">Take control of your finances</h1>
		<p>
			Fiscora helps you track your income, expenses, and savings all in
			one place. Gain valuable insights, set goals, and make smarter
			financial decisions.
		</p>

		{#if !loggedIn}
			<a class="btn mt-4 w-fit px-4" href="/login">Sign up now!</a>
		{:else}
			<a class="btn mt-4 w-fit px-4" href="/transactions">
				View your transactions
			</a>
		{/if}
	</div>
	<div class="max-w-[700px]">
		{#if darkMode.darkMode}
			<enhanced:img
				src="/static/img/transactions-dark.png"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md border-2 border-primary-500/25 shadow-md shadow-primary-900/50 dark:shadow-surface-900"
			/>
		{:else}
			<enhanced:img
				src="/static/img/transactions-light.png"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md border-2 border-primary-500/25 shadow-md shadow-primary-900/50 dark:shadow-surface-900"
			/>
		{/if}
	</div>
</div>

<div
	class="relative my-10 flex flex-col items-center justify-center gap-12 text-center lg:flex-row lg:justify-evenly lg:text-left"
>
	<div class="order-last max-w-[700px] lg:order-1">
		{#if darkMode.darkMode}
			<enhanced:img
				src="/static/img/dashboard-dark.png"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md border-2 border-primary-500/25 shadow-md shadow-primary-900/50 dark:shadow-surface-900"
			/>
		{:else}
			<enhanced:img
				src="/static/img/dashboard-light.png"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md border-2 border-primary-500/25 shadow-md shadow-primary-900/50 dark:shadow-surface-900"
			/>
		{/if}
	</div>
	<div class="order-2 flex w-full max-w-[700px] flex-col">
		<h1 class="mb-6">Dashboard</h1>
		<p>
			Get your financial snapshot at a glance using the dashboard page.
			Track expenses, monitor budgets, and visualize your progress towards
			financial goals all in one place.
		</p>
	</div>
</div>
