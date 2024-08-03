<script>
	import {
		computePosition,
		autoUpdate,
		offset,
		shift,
		flip,
		arrow,
	} from "@floating-ui/dom";
	import Header from "$lib/components/header.svelte";
	import {
		initializeStores,
		storePopup,
		Toast,
	} from "@skeletonlabs/skeleton";
	import "../app.css";
	import { onNavigate } from "$app/navigation";
	import { page } from "$app/stores";

	let { children } = $props();

	const { demo } = $page.data;

	initializeStores();

	storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });

	onNavigate((navigation) => {
		// @ts-expect-error relatively new feature
		if (!document.startViewTransition) return;

		return new Promise((resolve) => {
			// @ts-expect-error relatively new feature
			document.startViewTransition(async () => {
				resolve();
				await navigation.complete;
			});
		});
	});
</script>

<Toast />

<Header />

<main
	class="mx-auto w-full max-w-[1200px] p-2 pb-8 md:mb-12 md:p-4 {demo &&
		'demo-container'}"
>
	{#if demo}
		<div class="demo-container-tag">Demo view</div>
	{/if}
	{@render children()}
</main>
