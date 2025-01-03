<script lang="ts">
	import { CalendarIcon, Plus, Trash } from "lucide-svelte";
	import type { Budget, BudgetExpenseFormErrors, BudgetForm, BudgetFormErrors } from "../../../ambient";
	import {
		formatDate,
		getCurrentMonthNumber,
		getCurrentYear,
		getFormattedAmount,
		getFormDate,
		getISOStringUTC,
		listAllMonths,
	} from "$lib";
	import { toast } from "svelte-sonner";
	import * as Dialog from "$lib/components/ui/dialog";
	import * as Tabs from "$lib/components/ui/tabs";
	import * as Popover from "$lib/components/ui/popover";
	import { Input } from "$lib/components/ui/input";
	import { Label } from "$lib/components/ui/label";
	import { Textarea } from "$lib/components/ui/textarea";
	import { Button, buttonVariants } from "$lib/components/ui/button";
	import { Calendar } from "$lib/components/ui/calendar";
	import { type DateValue } from "@internationalized/date";
	import MonthPicker from "$lib/components/month-picker.svelte";
	import { cn } from "$lib/utils";

	let {
		budget,
		demo,
		close,
		success,
	}: {
		budget: Budget | null;
		demo: boolean;
		close: () => void;
		success: (budget: Budget) => void;
	} = $props();

	const defaultForm = (): BudgetForm => {
		return {
			id: budget?.id ?? "",
			name: budget?.name ?? "",
			description: budget?.description ?? "",
			amount: budget?.amount ?? 0,
			startDate: getFormDate(budget?.startDate ?? new Date()),
			endDate: getFormDate(budget?.endDate ?? new Date()),
			expenses: budget?.expenses.map((expense) => ({
				id: expense.id,
				name: expense.name,
				allocatedAmount: expense.allocatedAmount,
				errors: {
					valid: true,
					name: null,
					allocatedAmount: null,
				},
			})) ?? [
				{
					id: -1,
					name: "",
					allocatedAmount: 0,
					errors: {} as BudgetExpenseFormErrors,
				},
			],
			errors: {} as BudgetFormErrors,
		};
	};

	let form: BudgetForm = $state(defaultForm());
	let month: number = $state(getCurrentMonthNumber());
	let year: number = $state(getCurrentYear());
	let startDate: DateValue | undefined = $state();
	let endDate: DateValue | undefined = $state();
	let budgetType = $state("monthly");

	const months = listAllMonths();

	function addExpense(event: MouseEvent) {
		event.preventDefault();

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

	async function removeExpense(event: MouseEvent, index: number) {
		event.preventDefault();

		const id = form.expenses[index].id;
		form.expenses = form.expenses.filter((_, i) => i !== index);
		try {
			const response = await fetch(`/api/budgets/${budget?.id}/expenses/${id}`, {
				method: "DELETE",
			});
			if (!response.ok) {
				throw Error();
			}
		} catch {
			toast.error("Something went wrong trying to delete the expense");
		}
	}

	async function submitBudget(event: SubmitEvent) {
		event.preventDefault();

		if (demo) {
			toast.warning("Demo users cannot create budgets");
			close();
			return;
		}

		if (startDate) {
			form.startDate = getISOStringUTC(startDate);
		}
		if (endDate) {
			form.endDate = getISOStringUTC(endDate);
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

		const created = budget === null;

		close();

		toast.success(`Budget ${created ? "created" : "updated"} successfully`);

		form = defaultForm();
		budget = null;

		success((await response.json()) as Budget);
	}

	$effect(() => {
		form = defaultForm();
	});

	$effect(() => {
		form.amount = form.expenses.reduce((acc, expense) => acc + expense.allocatedAmount, 0);
	});

	$effect(() => {
		console.log(year);
		console.log(month);
		// TODO: Change start and end date when this changes
	});
</script>

<Dialog.Content>
	<Dialog.Header>
		<h3>Create Budget</h3>
	</Dialog.Header>
	<form onsubmit={submitBudget} class="mt-6 flex flex-col gap-4">
		<Tabs.Root onValueChange={(value) => (budgetType = value!)}>
			<Tabs.List class="grid w-full grid-cols-2">
				<Tabs.Trigger value="monthly">Monthly</Tabs.Trigger>
				<Tabs.Trigger value="custom">Custom</Tabs.Trigger>
			</Tabs.List>
		</Tabs.Root>
		<div class="flex flex-col gap-2">
			<Label for="name">Budget Name</Label>
			<Input
				id="name"
				name="name"
				type="text"
				placeholder="Budget Name"
				class={form.errors.name && "error"}
				bind:value={form.name}
			/>
			{#if form.errors.name}
				<small class="text-destructive">{form.errors.name}</small>
			{/if}
		</div>
		<div class="flex flex-col gap-2">
			<Label for="description">Description</Label>
			<Textarea
				id="description"
				name="description"
				rows={3}
				maxlength={256}
				class="input p-1 {form.errors.description && 'error'}"
				bind:value={form.description}
			/>
			<span class="flex justify-between">
				<small class="text-destructive">
					{form.errors.description}
				</small>
				<small>
					{form.description.length}/256
				</small>
			</span>
		</div>
		{#if budgetType === "monthly"}
			<div class="flex items-center justify-between">
				<Label>Month</Label>
				<Popover.Root>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: "outline" }),
							"!h-fit w-[280px] justify-start px-2 text-left text-base",
						)}
					>
						<CalendarIcon class="mr-2 h-5 w-5" />
						{months.get(month)}
					</Popover.Trigger>
					<Popover.Content class="w-auto p-0">
						<MonthPicker bind:year bind:month callback={() => {}} />
					</Popover.Content>
				</Popover.Root>
			</div>
		{/if}
		<div class={budgetType === "monthly" ? "hidden" : ""}>
			<div class="flex items-center justify-between">
				<Label>Start Date</Label>
				<Popover.Root>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: "outline" }),
							`w-[280px] justify-start ${!startDate && "text-muted-foreground"} ${form.errors.startDate && "error"}`,
						)}
					>
						<CalendarIcon class="mr-2 h-4 w-4" />
						{startDate ? formatDate(startDate) : `Select a start date`}
					</Popover.Trigger>
					<Popover.Content class="w-auto p-0">
						<Calendar type="single" bind:value={startDate} initialFocus />
					</Popover.Content>
				</Popover.Root>
			</div>
			{#if form.errors.startDate}
				<small class="float-end text-destructive">{form.errors.startDate}</small>
			{/if}
		</div>
		<div class={budgetType === "monthly" ? "hidden" : ""}>
			<div class="flex items-center justify-between">
				<Label>End Date</Label>
				<Popover.Root>
					<Popover.Trigger
						class={cn(
							buttonVariants({ variant: "outline" }),
							`w-[280px] justify-start ${!endDate && "text-muted-foreground"} ${form.errors.endDate && "error"}`,
						)}
					>
						<CalendarIcon class="mr-2 h-4 w-4" />
						{endDate ? formatDate(endDate) : "Select an end date"}
					</Popover.Trigger>
					<Popover.Content class="w-auto p-0">
						<Calendar type="single" bind:value={endDate} initialFocus />
					</Popover.Content>
				</Popover.Root>
			</div>
			{#if form.errors.endDate}
				<small class="float-end text-destructive">{form.errors.endDate}</small>
			{/if}
		</div>
		<div>
			<div class="flex items-center justify-between">
				<Label>Total Budget Amount</Label>
				<span class="font-semibold">{getFormattedAmount(form.amount)}</span>
			</div>
			{#if form.errors.amount}
				<small class="float-end text-destructive">{form.errors.amount}</small>
			{/if}
		</div>
		<div class="my-2 flex items-center justify-between">
			<h4>Expenses</h4>
			<button class="btn-icon" onclick={addExpense}>
				<Plus size={20} />
			</button>
		</div>
		{#each form.expenses as expense, index}
			<div class="grid grid-cols-[1fr_130px_auto] items-start gap-2">
				<div>
					<Input
						type="text"
						placeholder="Expense name"
						class={expense.errors.name && "error"}
						bind:value={expense.name}
					/>
					{#if expense.errors.name}
						<small class="text-destructive">
							{expense.errors.name}
						</small>
					{/if}
				</div>
				<div>
					<Input
						type="number"
						min="0"
						step="0.01"
						placeholder="Amount"
						class={expense.errors.allocatedAmount && "error"}
						bind:value={expense.allocatedAmount}
					/>
					{#if expense.errors.allocatedAmount}
						<small class="text-destructive">
							{expense.errors.allocatedAmount}
						</small>
					{/if}
				</div>
				<button
					class={buttonVariants({ size: "icon", variant: "destructive" })}
					onclick={(event) => removeExpense(event, index)}
				>
					<Trash class="h-5 w-5" />
				</button>
			</div>
		{/each}
		<div class="mt-4 flex justify-end gap-4">
			<Button type="submit" disabled={demo}>
				{budget === null ? "Create" : "Update"}
			</Button>
		</div>
	</form>
</Dialog.Content>
