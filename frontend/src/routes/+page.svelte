<script lang="ts">
	import { page } from "$app/stores";

	const { loggedIn } = $page.data;
</script>

<div class="flex w-full flex-col items-center justify-center text-center">
	<h1 class="mb-4">Take control of your finances</h1>
	<p>
		Budget Buddy helps you track your income, expenses, and savings all in
		one place. Gain valuable insights, set goals, and make smarter financial
		decisions.
	</p>

	{#if !loggedIn}
		<a class="btn mt-4 w-[200px]" href="/login">Sign up now!</a>
	{/if}
</div>
