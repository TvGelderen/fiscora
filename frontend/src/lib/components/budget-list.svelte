<script lang="ts">
	import type { Budget } from "../../ambient";
	import { Edit, Trash, X } from "lucide-svelte";
	import { getToastStore } from "@skeletonlabs/skeleton";
	import { getFormattedAmount, getFormattedDateShortWithYear } from "$lib";

	const toastStore = getToastStore();

	let {
		budgets,
		demo,
		add,
		edit,
		remove,
	}: {
		budgets: Budget[];
		demo: boolean;
		add: (idx: number, budget: Budget) => void;
		edit: (budget: Budget) => void;
		remove: (budget: Budget) => void;
	} = $props();

	let modal: HTMLDialogElement;
	let budgetToDelete: Budget | null = $state(null);

	function openDeleteConfirmation(budget: Budget) {
		budgetToDelete = budget;
		modal.showModal();
	}

	function closeDeleteConfirmation() {
		modal.close();
		budgetToDelete = null;
	}

	async function deleteBudget() {
		if (budgetToDelete === null) return;

		const id = budgetToDelete.id;
		const idx = budgets.findIndex((b) => b.id === id);
		const budget = budgets.at(idx);

		remove(budgetToDelete);

		closeDeleteConfirmation();

		if (demo) {
			toastStore.trigger({
				message: "Demo users cannot delete budgets",
				background: "variant-filled-warning",
			});
			return;
		}

		const response = await fetch(`/api/budgets/${id}`, { method: "DELETE" });
		if (!response.ok) {
			toastStore.trigger({
				message: "Error deleting budget",
				background: "variant-filled-error",
			});

			if (budget !== undefined) {
				add(idx, budget);
			}

			return;
		}

		toastStore.trigger({
			message: "Budget deleted successfully",
			background: "variant-filled-success",
		});
	}
</script>

<div class="flex flex-wrap justify-center gap-6">
	{#each budgets as budget (budget.id)}
		<div class="card-primary flex w-full max-w-sm flex-col justify-between p-6">
			<div>
				<a href="/budgets/{budget.id}" class="hover:underline">
					<h3 class="mb-2">
						{budget.name}
					</h3>
				</a>
				<p class="text-secondary mb-4 text-sm">
					{budget.description}
				</p>
				<div class="mb-4 flex flex-col gap-2 rounded-lg bg-surface-100 p-3 dark:bg-surface-600">
					<div class="flex items-center justify-between">
						<span class="text-sm font-medium">Total Budget:</span>
						<span class="text-lg font-bold">
							{getFormattedAmount(budget.amount)}
						</span>
					</div>
					<div class="text-secondary flex justify-between text-sm">
						<span>
							Start: {getFormattedDateShortWithYear(budget.startDate)}
						</span>
						<span>
							End: {getFormattedDateShortWithYear(budget.endDate)}
						</span>
					</div>
				</div>
				<div class="mb-2">
					<ul class="list-inside list-disc space-y-2">
						{#each budget.expenses as expense}
							<li class="flex items-center justify-between">
								<span class="text-secondary text-base">
									{expense.name}:
								</span>
								<span class="font-semibold">
									{getFormattedAmount(expense.allocatedAmount)}
								</span>
							</li>
						{/each}
					</ul>
				</div>
			</div>
			<div class="mt-4 flex justify-end gap-1">
				<button class="btn-icon hover:!variant-soft-primary" onclick={() => edit(budget)} disabled={demo}>
					<Edit size={20} />
				</button>
				<button
					class="btn-icon hover:!variant-filled-error"
					onclick={() => openDeleteConfirmation(budget)}
					disabled={demo}
				>
					<Trash size={20} />
				</button>
			</div>
		</div>
	{/each}
</div>

{#if budgets.length === 0}
	<p class="text-center">You haven't set any budgets yet. Create one to get started!</p>
{/if}

<dialog class="max-w-md" bind:this={modal}>
	<button class="absolute right-4 top-4" onclick={closeDeleteConfirmation}>
		<X />
	</button>
	{#if budgetToDelete}
		<h3 class="mb-4">Confirm Deletion</h3>
		<p class="mb-4">Are you sure you want to delete this budget? This action is permanent and cannot be undone.</p>
		<div class="flex justify-end gap-2">
			<button class="!variant-filled-surface btn" onclick={closeDeleteConfirmation}>Cancel</button>
			<button class="!variant-filled-error btn" onclick={deleteBudget}>Delete</button>
		</div>
	{/if}
</dialog>
