<script lang="ts">
	import { page } from "$app/stores";
	import BudgetList from "$lib/components/budget-list.svelte";
	import BudgetForm from "$lib/components/budget-form.svelte";
	import { Plus } from "lucide-svelte";
	import type { Budget } from "../../../ambient";
	import type { PageData } from "./$types";

	let { budgets, demo } = $page.data as PageData;

	let budgetArray: Budget[] = $state(budgets);
	let editBudget: Budget | null = $state(null);
	let showFormModal: boolean = $state(false);

	function closeFormModal() {
		editBudget = null;
		showFormModal = false;
	}

	function handleSuccess(budget: Budget) {
		if (editBudget === null) {
			budgetArray = [budget, ...budgetArray];
		} else {
			const idx = budgetArray.findIndex((b) => b.id === budget.id);
			if (idx !== -1) {
				budgetArray.splice(idx, 1, budget);
				budgetArray = [...budgetArray];
			}
		}

		closeFormModal();
	}

	function removeBudget(budget: Budget) {
		budgetArray = budgetArray.filter((t) => t.id !== budget.id);
	}

	function setEditBudget(budget: Budget) {
		editBudget = budget;
		showFormModal = true;
	}
</script>

<svelte:head>
	<title>Fiscora - Budgets</title>
</svelte:head>

<div class="mx-auto mb-8 text-center lg:mb-12">
	<h1 class="mb-4">Your Budgets</h1>
	<p>
		Set and manage your monthly budget goals to stay on track with your
		financial objectives.
	</p>
</div>

<div class="my-4 flex justify-end">
	<button
		class="secondary btn"
		onclick={() => (showFormModal = true)}
		disabled={demo}
	>
		<Plus />&nbsp;Add Budget
	</button>
</div>

<BudgetList
	budgets={budgetArray}
	{demo}
	edit={setEditBudget}
	remove={removeBudget}
/>

<BudgetForm
	open={showFormModal}
	budget={editBudget}
	{demo}
	close={closeFormModal}
	success={handleSuccess}
/>
