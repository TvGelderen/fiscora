<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { Budget } from "../../ambient";
	import { Edit, Trash } from "lucide-svelte";
	import { getToastStore } from "@skeletonlabs/skeleton";

	const dispatch = createEventDispatcher();
	const toastStore = getToastStore();

	let { budgets, demo }: { budgets: Budget[]; demo: boolean } = $props();

	async function deleteBudget(id: number) {
		if (demo) {
			toastStore.trigger({
				message: "Demo users cannot delete budgets",
				background: "variant-filled-warning",
			});
			return;
		}

		// Implement delete logic here
		// After successful deletion, update the budgets list
		budgets = budgets.filter((budget) => budget.id !== id);
		toastStore.trigger({
			message: "Budget deleted successfully",
			background: "variant-filled-success",
		});
	}
</script>

<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
	{#each budgets as budget (budget.id)}
		<div class="card p-4">
			<h3 class="mb-2 text-3xl">{budget.name}</h3>
			<p class="mb-4 text-sm">{budget.description}</p>
			<div class="mb-2 flex justify-between">
				<span>Total Budget:</span>
				<span class="font-semibold">${budget.amount.toFixed(2)}</span>
			</div>
			<div class="mb-4">
				<ul class="list-inside list-disc">
					{#each budget.categories as category}
						<li>
							{category.name}: ${category.allocatedAmount.toFixed(
								2,
							)}
						</li>
					{/each}
				</ul>
			</div>
			<div class="flex justify-end gap-2">
				<button
					class="variant-soft btn btn-sm"
					onclick={() => dispatch("edit", budget)}
					disabled={demo}
				>
					<Edit size={16} />
				</button>
				<button
					class="variant-soft-error btn btn-sm"
					onclick={() => deleteBudget(budget.id)}
					disabled={demo}
				>
					<Trash size={16} />
				</button>
			</div>
		</div>
	{/each}
</div>

{#if budgets.length === 0}
	<p class="text-center">
		You haven't set any budgets yet. Create one to get started!
	</p>
{/if}
