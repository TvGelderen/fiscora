<script lang="ts">
	import type { Transaction, TransactionForm, TransactionFormErrors } from "../../../ambient";
	import { formatDate, getISOStringUTC } from "$lib";
	import { toast } from "svelte-sonner";
	import { CalendarIcon } from "lucide-svelte";
	import * as Dialog from "$lib/components/ui/dialog";
	import * as Tabs from "$lib/components/ui/tabs";
	import * as Popover from "$lib/components/ui/popover";
	import * as Select from "$lib/components/ui/select";
	import { Switch } from "$lib/components/ui/switch";
	import { Input } from "$lib/components/ui/input";
	import { Label } from "$lib/components/ui/label";
	import { Textarea } from "$lib/components/ui/textarea";
	import { Button } from "$lib/components/ui/button";
	import { Calendar } from "$lib/components/ui/calendar";
	import { CalendarDate, type DateValue } from "@internationalized/date";

	let {
		transaction,
		transactionIntervals,
		incomeTypes,
		expenseTypes,
		demo,
		success,
		close,
	}: {
		transaction: Transaction | null;
		transactionIntervals: string[];
		incomeTypes: string[];
		expenseTypes: string[];
		demo: boolean;
		success: () => void;
		close: () => void;
	} = $props();

	const defaultForm = (): TransactionForm => {
		if (transaction !== null) {
			if (transaction.recurring === null) {
				const date = new Date(transaction.date);
				startDate = new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
			} else {
				let date = new Date(transaction.recurring.startDate!);
				startDate = new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
				date = new Date(transaction.recurring.endDate!);
				endDate = new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
			}
		}

		return {
			id: transaction?.id ?? -1,
			amount: transaction?.amount ?? 0,
			startDate: undefined,
			endDate: undefined,
			description: transaction?.description ?? "",
			recurring: !!transaction?.recurring,
			interval: transaction?.recurring?.interval ?? null,
			daysInterval: transaction?.recurring?.daysInterval ?? null,
			type: transaction?.type ?? null,
			errors: {} as TransactionFormErrors,
		};
	};

	let form: TransactionForm = $state(defaultForm());
	let startDate: DateValue | undefined = $state();
	let endDate: DateValue | undefined = $state();
	let transactionIncomeTypeOptions: { value: string; label: string }[] = $state([]);
	let transactionExpenseTypeOptions: { value: string; label: string }[] = $state([]);
	let creating = $derived(transaction === null);
	let isExpense = $derived(form.amount < 0);
	let transactionTypeOptions = $derived(isExpense ? transactionExpenseTypeOptions : transactionIncomeTypeOptions);

	async function submitTransaction(event: SubmitEvent) {
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
		if (creating) {
			response = await fetch("/api/transactions", {
				method: "POST",
				body: JSON.stringify(form),
			});
		} else {
			response = await fetch(`/api/transactions/${transaction!.id}`, {
				method: "PUT",
				body: JSON.stringify(form),
			});
		}

		if (!response.ok) {
			form = await response.json();
			return;
		}

		toast.success(`Transaction ${creating ? "created" : "updated"} successfully`);

		if (!creating) {
			transaction = null;
		}

		form = defaultForm();

		success();
	}

	$effect(() => {
		form = defaultForm();
		transactionIncomeTypeOptions = incomeTypes.map((type) => {
			return { value: type, label: type };
		});
		transactionExpenseTypeOptions = expenseTypes.map((type) => {
			return { value: type, label: type };
		});

		if (creating) {
			const date = new Date();
			startDate = new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
		}
	});
</script>

<Dialog.Content>
	<Dialog.Header>
		{#if creating}
			<h3>Add Transaction</h3>
		{:else}
			<h3>Update Transaction</h3>
		{/if}
	</Dialog.Header>
	<form onsubmit={submitTransaction} class="mt-6 flex flex-col gap-4">
		<div class="flex flex-col gap-2">
			<Label for="amount">Amount</Label>
			<Input
				id="amount"
				name="amount"
				type="number"
				step="0.01"
				class={form.errors.amount && "error"}
				bind:value={form.amount}
			/>
			{#if form.errors.amount}
				<small class="text-destructive">{form.errors.amount}</small>
			{/if}
		</div>
		<div class="flex flex-col gap-2">
			<Label for="description">Description</Label>
			<Textarea
				id="description"
				name="description"
				placeholder="Description..."
				rows={3}
				maxlength={512}
				class={form.errors.description && "error"}
				bind:value={form.description}
			/>
			<span class="flex justify-between">
				<small class="text-destructive">
					{form.errors.description}
				</small>
				<small>
					{form.description.length}/512
				</small>
			</span>
		</div>
		{#if creating}
			<div class="flex items-center justify-between">
				<Label for="recurring">Recurring</Label>
				<Switch id="recurring" name="slide" bind:checked={form.recurring} disabled={!creating} />
			</div>
		{/if}
		{#if form.recurring}
			<Tabs.Root value={form.interval ?? ""} onValueChange={(value) => (form.interval = value!)}>
				<Tabs.List class="grid w-full grid-cols-4">
					{#each transactionIntervals as value}
						<Tabs.Trigger {value}>
							{value}
						</Tabs.Trigger>
					{/each}
				</Tabs.List>
			</Tabs.Root>
			{#if form.errors.interval}
				<small class="text-destructive">{form.errors.interval}</small>
			{/if}
			{#if form.interval === "Other"}
				<div class="flex flex-col gap-2">
					<Label for="days-interval">
						Every {form.daysInterval ?? 1} days
					</Label>
					<Input
						id="days-interval"
						name="days-interval"
						type="number"
						min="1"
						placeholder="1"
						class={form.errors.daysInterval && "error"}
						bind:value={form.daysInterval}
					/>
					{#if form.errors.daysInterval}
						<small class="text-destructive">
							{form.errors.daysInterval}
						</small>
					{/if}
				</div>
			{/if}
		{/if}
		<div>
			<div class="flex items-center justify-between">
				{#if form.recurring}
					<Label>Start Date</Label>
				{:else}
					<Label>Date</Label>
				{/if}
				<Popover.Root>
					<Popover.Trigger>
						<Button
							variant="outline"
							class={`w-[280px] justify-start ${!startDate && "text-muted-foreground"} ${form.errors.startDate && "error"}`}
						>
							<CalendarIcon class="mr-2 h-4 w-4" />
							{startDate ? formatDate(startDate) : `Select a ${form.recurring ? "start" : ""} date`}
						</Button>
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
		{#if form.recurring}
			<div>
				<div class="flex items-center justify-between">
					<Label>End Date</Label>
					<Popover.Root>
						<Popover.Trigger>
							<Button
								variant="outline"
								class={`w-[280px] justify-start ${!endDate && "text-muted-foreground"} ${form.errors.endDate && "error"}`}
							>
								<CalendarIcon class="mr-2 h-4 w-4" />
								{endDate ? formatDate(endDate) : "Select an end date"}
							</Button>
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
		{/if}
		<div>
			<div class="flex items-center justify-between">
				<Label>Transaction type</Label>
				<Select.Root
					type="single"
					items={transactionTypeOptions}
					value={form.type || undefined}
					onValueChange={(value) => {
						form.type = value;
					}}
				>
					<Select.Trigger class={`w-[280px] px-4 ${form.errors.type && "error"}`}>
						{form.type || "Please select a transaction type"}
					</Select.Trigger>
					<Select.Content>
						{#each transactionTypeOptions as option}
							<Select.Item value={option.value} label={option.label}>{option.label}</Select.Item>
						{/each}
					</Select.Content>
				</Select.Root>
			</div>
			{#if form.errors.type}
				<small class="float-end text-destructive">{form.errors.type}</small>
			{/if}
		</div>
		<Button class="text-slate-50" type="submit" disabled={demo}>
			{#if creating}
				Add transaction
			{:else}
				Update transaction
			{/if}
		</Button>
	</form>
</Dialog.Content>
