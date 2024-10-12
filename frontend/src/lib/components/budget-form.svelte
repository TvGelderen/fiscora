<script lang="ts">
	import { X } from "lucide-svelte";
	import { getToastStore } from "@skeletonlabs/skeleton";
	import type { Budget } from "../../ambient";

	const toastStore = getToastStore();

	let {
		open,
		budget,
		demo,
		close,
		success,
	}: {
		open: boolean;
		budget: Budget | null;
		demo: boolean;
		close: () => void;
		success: () => void;
	} = $props();

	const defaultForm = () => {
		return {
			name: budget?.name ?? "",
			description: budget?.description ?? "",
			amount: budget?.amount ?? 0,
			categories: budget?.categories.map((category) => ({
				name: category.name,
				allocatedAmount: category.allocatedAmount,
			})) ?? [{ name: "", allocatedAmount: 0 }],
		};
	};

	let modal: HTMLDialogElement;
	let form = $state(defaultForm());

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

		toastStore.trigger({
			background: "bg-success-400 text-black",
			message: `Budget ${budget === null ? "created" : "updated"} successfully`,
			timeout: 1500,
		});

		success();
	}

	$effect(() => {
		form = defaultForm();
	});

	$effect(() => {
		if (open) {
			modal.showModal();
		} else {
			modal.close();
		}
	});
</script>

<dialog class="max-w-screen-sm" bind:this={modal}>
	<button class="absolute right-4 top-4 active:outline-none" onclick={close}>
		<X />
	</button>
	<h3>Create Budget</h3>
	<form onsubmit={submitBudget} class="mt-6 flex flex-col gap-4">
		<label class="label" for="name">
			<span>Budget Name</span>
			<input
				id="name"
				name="name"
				type="text"
				class="input-bordered input"
				bind:value={form.name}
				required
			/>
		</label>
		<label class="label" for="description">
			<span>Description</span>
			<textarea
				id="description"
				name="description"
				class="textarea-bordered textarea"
				bind:value={form.description}
				required
			></textarea>
		</label>
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
</dialog>
