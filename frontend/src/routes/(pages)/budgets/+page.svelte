<script lang="ts">
	import { page } from "$app/stores";
	import BudgetList from "$lib/components/budget-list.svelte";
	import BudgetForm from "$lib/components/budget-form.svelte";
	import { Plus } from "lucide-svelte";
	import type { Budget } from "../../../ambient";
	import type { PageData } from "./$types";

	let { budgets, demo } = $page.data as PageData;

	let budgetState: Budget[] = $state(budgets);
	let editBudget: Budget | null = $state(null);
	let showFormModal: boolean = $state(false);

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

		close();
	}

	function add(idx: number, budget: Budget) {
		budgetState.splice(idx, 0, budget);
		budgetState = [...budgetState];
	}

	function remove(budget: Budget) {
		budgetState = budgetState.filter((t) => t.id !== budget.id);
	}

	function edit(budget: Budget) {
		editBudget = budget;
		showFormModal = true;
	}

	function openFormModal() {
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
	class="variant-filled-primary btn-icon btn-lg fixed bottom-4 right-4 rounded-full shadow-lg transition-colors duration-300 hover:shadow-xl sm:bottom-8 sm:right-8"
	onclick={openFormModal}
	disabled={demo}
>
	<Plus size={24} />
	<span class="sr-only">Add Budget</span>
</button>

<BudgetForm open={showFormModal} budget={editBudget} {demo} {close} {success} />
