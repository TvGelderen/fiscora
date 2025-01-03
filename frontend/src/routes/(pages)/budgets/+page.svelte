<script lang="ts">
	import BudgetList from "./budget-list.svelte";
	import BudgetForm from "./budget-form.svelte";
	import * as Dialog from "$lib/components/ui/dialog";
	import { Plus } from "lucide-svelte";
	import type { Budget } from "../../../ambient";
	import { tick } from "svelte";

	let { data } = $props();
	let { budgets, demo } = data;

	let showFormModal: boolean = $state(false);
	let budgetState: Budget[] = $state(budgets);
	let editBudget: Budget | null = $state(null);

	function close() {
		editBudget = null;
		showFormModal = false;
	}

	function success(budget: Budget) {
		if (editBudget === null) {
			budgetState = [budget, ...budgetState];
		} else {
			const idx = budgetState.findIndex((b) => b.id === budget.id);
			if (idx !== -1) {
				budgetState[idx] = budget;
			}
		}
	}

	function add(idx: number, budget: Budget) {
		budgetState.splice(idx, 0, budget);
		budgetState = [...budgetState];
	}

	function remove(budget: Budget) {
		budgetState = budgetState.filter((t) => t.id !== budget.id);
	}

	async function edit(budget: Budget) {
		editBudget = budget;
		await tick();
		showFormModal = true;
	}
</script>

<svelte:head>
	<title>Fiscora - Budgets</title>
</svelte:head>

<div class="mx-auto mb-8 text-center lg:mb-12">
	<h1 class="mb-4">Your Budgets</h1>
	<p>Set and manage your monthly budget goals to stay on track with your financial objectives.</p>
</div>

<BudgetList budgets={budgetState} {demo} {add} {edit} {remove} />

<button
	class="fixed bottom-5 right-4 rounded-full bg-primary p-2 text-slate-50 shadow-md !shadow-slate-500 transition-all duration-300 hover:shadow-lg dark:!shadow-slate-800 sm:bottom-8 sm:right-8"
	onclick={() => (showFormModal = true)}
	disabled={demo}
>
	<Plus size={24} />
	<span class="sr-only">Add Budget</span>
</button>

<Dialog.Root
	bind:open={showFormModal}
	onOpenChange={(open) => {
		if (!open) {
			editBudget = null;
		}
	}}
>
	<BudgetForm budget={editBudget} {demo} {close} {success} />
</Dialog.Root>
