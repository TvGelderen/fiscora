<script lang="ts">
	import { page } from "$app/stores";
	import { createDarkMode } from "$lib/theme.svelte";

	const { loggedIn } = $page.data;

	const darkMode = createDarkMode();
</script>

<div
	class="absolute inset-0 z-[-1] bg-[radial-gradient(#18182533_1px,transparent_1px)] [background-size:24px_24px] [mask-image:radial-gradient(ellipse_50%_50%_at_50%_50%,#000_70%,transparent_100%)] dark:bg-[radial-gradient(#e4e4e422_1px,transparent_1px)]"
></div>
<div class="mt-20 flex flex-col items-center justify-center gap-12 text-center">
	<div class="flex w-full max-w-[700px] flex-col items-center gap-6">
		<h1>Take control of your finances</h1>
		<p>
			Fiscora helps you track your income, expenses, and savings all in
			one place. Gain valuable insights, set goals, and make smarter
			financial decisions.
		</p>

		{#if !loggedIn}
			<a
				class="btn w-fit px-4 shadow-xl shadow-primary-500/25"
				href="/login"
			>
				Sign up now!
			</a>
		{:else}
			<a
				class="btn w-fit px-4 shadow-xl shadow-primary-500/25"
				href="/transactions"
			>
				View your transactions
			</a>
		{/if}
	</div>
	<div class="relative z-[-1] mt-[-70px] max-w-[1080px]">
		{#if darkMode.darkMode}
			<enhanced:img
				id="img"
				src="/static/img/transactions-dark.jpeg"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md border-2 border-primary-500/15"
			/>
		{:else}
			<enhanced:img
				id="img"
				src="/static/img/transactions-light.jpeg"
				loading="lazy"
				alt="Transactions page in light and darkmode"
				class="rounded-md border-2 border-primary-500/15"
			/>
		{/if}
	</div>
</div>

<style>
	#img {
		object-fit: cover;
		mask-image: linear-gradient(to top, transparent 0%, black 50%);
		-webkit-mask-image: linear-gradient(to top, transparent 0%, black 50%);
	}
</style>
