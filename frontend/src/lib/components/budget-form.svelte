<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import { X } from "lucide-svelte";
	import { getToastStore } from "@skeletonlabs/skeleton";
	import type { Budget } from "../../ambient";

	const dispatch = createEventDispatcher();
	const toastStore = getToastStore();

	let {
		open,
		editBudget,
		demo,
	}: { open: boolean; editBudget?: Budget | null; demo: boolean } = $props();

	let modal: HTMLDialogElement;
	let form = $state({
		name: "",
		description: "",
		amount: 0,
		categories: [{ name: "", allocatedAmount: 0 }],
	});

	$effect(() => {
		if (open) {
			modal.showModal();
		} else {
			modal.close();
		}
	});

	function addCategory() {
		form.categories = [
			...form.categories,
			{ name: "", allocatedAmount: 0 },
		];
	}

	function removeCategory(index: number) {
		form.categories = form.categories.filter((_, i) => i !== index);
	}

	async function submitBudget(event: SubmitEvent) {
		event.preventDefault();

		if (demo) {
			toastStore.trigger({
				message: "Demo users cannot create budgets",
				background: "variant-filled-warning",
			});
			return;
		}

		// Implement budget creation/update logic here
		// After successful submission:
		toastStore.trigger({
			message: "Budget saved successfully",
			background: "variant-filled-success",
		});
		dispatch("success");
		dispatch("close");
	}
</script>

<dialog class="modal" bind:this={modal}>
	<div class="modal-box w-11/12 max-w-5xl">
		<button
			class="absolute right-4 top-4 active:outline-none"
			onclick={() => dispatch("close")}
		>
			<X />
		</button>
		<h3 class="text-lg font-bold">Create Budget</h3>
		<form onsubmit={submitBudget} class="mt-4">
			<div class="form-control">
				<label class="label" for="name">
					<span class="label-text">Budget Name</span>
				</label>
				<input
					type="text"
					id="name"
					class="input-bordered input"
					bind:value={form.name}
					required
				/>
			</div>
			<div class="form-control">
				<label class="label" for="description">
					<span class="label-text">Description</span>
				</label>
				<textarea
					id="description"
					class="textarea-bordered textarea"
					bind:value={form.description}
					required
				></textarea>
			</div>
			<div class="form-control">
				<label class="label" for="amount">
					<span class="label-text">Total Budget Amount</span>
				</label>
				<input
					type="number"
					id="amount"
					class="input-bordered input"
					bind:value={form.amount}
					min="0"
					step="0.01"
					required
				/>
			</div>
			<h4 class="mt-4 font-semibold">Categories</h4>
			{#each form.categories as category, index}
				<div class="mt-2 flex gap-2">
					<input
						type="text"
						class="input-bordered input flex-grow"
						placeholder="Category name"
						bind:value={category.name}
						required
					/>
					<input
						type="number"
						class="input-bordered input w-32"
						placeholder="Amount"
						bind:value={category.allocatedAmount}
						min="0"
						step="0.01"
						required
					/>
					<button
						type="button"
						class="btn-error btn btn-sm"
						onclick={() => removeCategory(index)}
						disabled={form.categories.length === 1}
					>
						Remove
					</button>
				</div>
			{/each}
			<button
				type="button"
				class="btn-secondary btn mt-2"
				onclick={addCategory}
			>
				Add Category
			</button>
			<div class="modal-action">
				<button type="submit" class="btn-primary btn" disabled={demo}>
					Save Budget
				</button>
			</div>
		</form>
	</div>
</dialog>
