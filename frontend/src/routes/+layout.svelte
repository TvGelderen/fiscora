<script>
import Header from '$lib/components/header.svelte'
import '../app.css'

let { children } = $props()
</script>

<Header />

<main class="mx-auto w-full h-[calc(100%-var(--header-height))] lg:h-[calc(100%-var(--header-height-lg))] max-w-[1200px] p-2 pb-8 md:p-4 md:pb-12">
    {@render children()}
</main>
