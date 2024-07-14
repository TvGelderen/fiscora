<script>
import Header from "$lib/components/header.svelte";
import { initializeStores } from "@skeletonlabs/skeleton";
import "../app.css";

let { children } = $props();

initializeStores();
</script>

<Header />

<main class="mx-auto w-full max-w-[1200px] p-4">
  {@render children()}
</main>
