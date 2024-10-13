<script lang="ts">
	import { page } from "$app/stores";
</script>

<div
	class="flex h-[calc(100vh_-_var(--header-height))] items-center justify-center"
>
	<h1
		class="mr-4 border-r-2 border-surface-400 pr-4 text-6xl dark:border-surface-200 lg:mr-8 lg:pr-8 lg:text-8xl"
	>
		{$page.status}
	</h1>
	<span class="text-2xl lg:text-4xl">
		{$page.error?.message}
	</span>
</div>
