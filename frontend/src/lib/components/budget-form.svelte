<script lang="ts">
	import { Plus, Trash, X } from "lucide-svelte";
	import { getToastStore } from "@skeletonlabs/skeleton";
	import type {
		Budget,
		BudgetExpenseForm,
		BudgetForm,
		BudgetFormErrors,
	} from "../../ambient";

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
		success: (budget: Budget) => void;
	} = $props();

	const defaultForm = () => {
		return <BudgetForm>{
			id: budget?.id,
			name: budget?.name ?? "",
			description: budget?.description ?? "",
			amount: budget?.amount ?? 0,
			expenses:
				budget?.expenses.map(
					(expense) =>
						<BudgetExpenseForm>{
							id: expense.id,
							name: expense.name,
							allocatedAmount: expense.allocatedAmount,
							errors: {
								valid: true,
								name: null,
								allocatedAmount: null,
							},
						},
				) ?? [],
			errors: <BudgetFormErrors>{},
		};
	};

	let modal: HTMLDialogElement;
	let form: BudgetForm = $state(defaultForm());

	function addExpense() {
		form.expenses = [
			...form.expenses,
			{
				id: -1,
				name: "",
				allocatedAmount: 0,
				errors: {
					valid: true,
					name: null,
					allocatedAmount: null,
				},
			},
		];
	}

	async function removeExpense(index: number) {
		const id = form.expenses[index].id;
		form.expenses = form.expenses.filter((_, i) => i !== index);
		try {
			const response = await fetch(
				`/api/budgets/${budget?.id}/expenses/${id}`,
				{
					method: "DELETE",
				},
			);
			if (!response.ok) {
				throw Error();
			}
		} catch {
			toastStore.trigger({
				message: "Something went wrong trying to delete the expense",
				background: "variant-filled-error",
			});
		}
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

		let response: Response;
		if (budget === null) {
			response = await fetch("/api/budgets", {
				method: "POST",
				body: JSON.stringify(form),
			});
		} else {
			response = await fetch(`/api/budgets/${budget.id}`, {
				method: "PUT",
				body: JSON.stringify(form),
			});
		}

		if (!response.ok) {
			form = await response.json();
			return;
		}

		toastStore.trigger({
			background: "bg-success-400 text-black",
			message: `Budget ${budget === null ? "created" : "updated"} successfully`,
		});

		form = defaultForm();
		budget = null;

		success((await response.json()) as Budget);
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

<dialog class="w-[95%] max-w-[500px]" bind:this={modal}>
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
				class="input p-1 {form.errors.name && 'error'}"
				bind:value={form.name}
			/>
			{#if form.errors.name}
				<small class="error-text">{form.errors.name}</small>
			{/if}
		</label>
		<label class="label" for="description">
			<span>Description</span>
			<textarea
				id="description"
				name="description"
				class="input p-1 {form.errors.description && 'error'}"
				bind:value={form.description}
				maxlength="256"
				rows="3"
			></textarea>
			<span class="relative !mt-0 flex">
				<small class="absolute right-0 top-0 float-right leading-none">
					{form.description.length}/256
				</small>
				{#if form.errors.description}
					<small class="error-text leading-none">
						{form.errors.description}
					</small>
				{/if}
			</span>
		</label>
		<label class="label" for="amount">
			<span class="label-text">Total Budget Amount</span>
			<input
				id="amount"
				type="number"
				class="input p-1 {form.errors.amount && 'error'}"
				bind:value={form.amount}
				min="0"
				step="0.01"
			/>
			{#if form.errors.amount}
				<small class="error-text">{form.errors.amount}</small>
			{/if}
		</label>
		<div class="my-2 flex items-center justify-between">
			<h4>Expenses</h4>
			<button
				type="button"
				class="!variant-ghost-primary flex h-8 w-8 items-center justify-center rounded"
				onclick={addExpense}
			>
				<Plus />
			</button>
		</div>
		{#each form.expenses as expense, index}
			<div class="grid grid-cols-[1fr_130px_auto] gap-2">
				<label class="label">
					<input
						type="text"
						class="input p-1 {expense.errors.name && 'error'}"
						placeholder="Expense name"
						bind:value={expense.name}
					/>
					{#if expense.errors.name}
						<small class="error-text">
							{expense.errors.name}
						</small>
					{/if}
				</label>
				<label class="label">
					<input
						type="number"
						class="input p-1 {expense.errors.allocatedAmount &&
							'error'}"
						placeholder="Amount"
						bind:value={expense.allocatedAmount}
						min="0"
						step="0.01"
					/>
					{#if expense.errors.allocatedAmount}
						<small class="error-text">
							{expense.errors.allocatedAmount}
						</small>
					{/if}
				</label>
				<button
					type="button"
					class="!variant-filled-error btn !btn-sm h-full self-end"
					onclick={() => removeExpense(index)}
				>
					<Trash class="h-5 w-5" />
				</button>
			</div>
		{/each}
		<div class="mt-4 flex justify-end gap-4">
			<button class="!variant-filled-surface btn" onclick={close}>
				Cancel
			</button>
			<button type="submit" class="btn-primary btn" disabled={demo}>
				Save Budget
			</button>
		</div>
	</form>
</dialog>
