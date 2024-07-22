<script>
import Header from '$lib/components/header.svelte'
import '../app.css'

let { children } = $props()
</script>

<Header />

<main class="mx-auto w-full h-[calc(100%-var(--header-height))] lg:h-[calc(100%-var(--header-height-lg))] max-w-[1200px] p-2 md:p-4">
    {@render children()}
</main>
