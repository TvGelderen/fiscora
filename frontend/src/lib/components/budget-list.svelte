<script lang="ts">
	import type { Budget } from "../../ambient";
	import { Edit, Trash, X } from "lucide-svelte";
	import { getToastStore } from "@skeletonlabs/skeleton";

	const toastStore = getToastStore();

	let {
		budgets,
		demo,
		edit: edit,
		remove: remove,
	}: {
		budgets: Budget[];
		demo: boolean;
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

		if (demo) {
			toastStore.trigger({
				message: "Demo users cannot delete budgets",
				background: "variant-filled-warning",
			});

			closeDeleteConfirmation();
			return;
		}

		const response = await fetch(`/api/budgets/${budgetToDelete.id}`, {
			method: "DELETE",
		});
		if (response.ok) {
			toastStore.trigger({
				message: "Budget deleted successfully",
				background: "bg-success-400 text-black",
			});

			remove(budgetToDelete);
		} else {
			toastStore.trigger({
				message: "Error deleting budget",
				background: "variant-filled-error",
			});
		}

		closeDeleteConfirmation();
	}
</script>

<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
	{#each budgets as budget (budget.id)}
		<div class="card flex flex-col justify-between p-4 shadow-lg">
			<div>
				<a href="/budgets/{budget.id}">
					<h3 class="mb-2 text-3xl">{budget.name}</h3>
				</a>
				<p class="mb-4 text-sm">{budget.description}</p>
				<div class="mb-2 flex justify-between">
					<span>Total Budget:</span>
					<span class="font-semibold">
						${budget.amount.toFixed(2)}
					</span>
				</div>
				<div class="mb-4">
					<ul class="text-secondary list-inside list-disc">
						{#each budget.expenses as expense}
							<li>
								{expense.name}: ${expense.allocatedAmount.toFixed(
									2,
								)}
							</li>
						{/each}
					</ul>
				</div>
			</div>
			<div class="flex justify-end gap-2">
				<button
					class="btn-icon btn-icon-sm hover:variant-filled-primary"
					onclick={() => edit(budget)}
					disabled={demo}
				>
					<Edit size={20} />
				</button>
				<button
					class="btn-icon btn-icon-sm hover:variant-filled-error"
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
	<p class="text-center">
		You haven't set any budgets yet. Create one to get started!
	</p>
{/if}

<dialog class="max-w-md" bind:this={modal}>
	<button class="absolute right-4 top-4" onclick={closeDeleteConfirmation}>
		<X />
	</button>
	{#if budgetToDelete}
		<h3 class="mb-4">Confirm Deletion</h3>
		<p class="mb-4">
			Are you sure you want to delete the budget "{budgetToDelete.name}"?
			This action is permanent and cannot be undone.
		</p>
		<div class="flex justify-end gap-2">
			<button
				class="!variant-filled-surface btn"
				onclick={closeDeleteConfirmation}
			>
				Cancel
			</button>
			<button class="!variant-filled-error btn" onclick={deleteBudget}>
				Delete
			</button>
		</div>
	{/if}
</dialog>
