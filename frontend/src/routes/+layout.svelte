<script>
	import Header from "$lib/components/header.svelte";
	import { initializeStores, Toast } from "@skeletonlabs/skeleton";
	import "../app.css";

	let { children } = $props();

	initializeStores();
</script>

<Toast />

<Header />

<main class="mx-auto w-full max-w-[1200px] p-2 pb-8 md:p-4 md:pb-12">
	{@render children()}
</main>
