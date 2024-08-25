<script lang="ts">
	import { page } from "$app/stores";

	let { children } = $props();

	const { demo } = $page.data;
</script>

<div class="px-2 py-4 md:mb-12 md:px-4 md:py-6 {demo && 'demo-container mt-6'}">
	{#if demo}
		<div class="demo-container-tag">Demo view</div>
	{/if}
	{@render children()}
</div>
