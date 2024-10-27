<script lang="ts">
	import type { Transaction } from "../../../ambient";
	import { formatDate, getISOStringUTC } from "$lib";
	import { CalendarIcon } from "lucide-svelte";
	import * as Dialog from "$lib/components/ui/dialog";
	import * as Form from "$lib/components/ui/form";
	import * as Tabs from "$lib/components/ui/tabs";
	import * as Popover from "$lib/components/ui/popover";
	import * as Select from "$lib/components/ui/select";
	import { Switch } from "$lib/components/ui/switch";
	import { Input } from "$lib/components/ui/input";
	import { Textarea } from "$lib/components/ui/textarea";
	import { Button } from "$lib/components/ui/button";
	import { Calendar } from "$lib/components/ui/calendar";
	import { CalendarDate, type DateValue } from "@internationalized/date";
	import type { SuperForm } from "sveltekit-superforms";

	let {
		form,
		transaction,
		transactionIntervals,
		incomeTypes,
		expenseTypes,
		demo,
	}: {
		form: SuperForm<
			{
				amount: number;
				id: number;
				type: string;
				description: string;
				recurring: boolean;
				startDate: string;
				interval?: string | undefined;
				daysInterval?: number | undefined;
				endDate?: string | undefined;
			},
			any
		>;
		transaction: Transaction | null;
		transactionIntervals: string[];
		incomeTypes: string[];
		expenseTypes: string[];
		demo: boolean;
	} = $props();

	const { form: formData, errors, enhance } = form;

	let startDate: DateValue | undefined = $state();
	let endDate: DateValue | undefined = $state();
	let transactionIncomeTypeOptions: { value: string; label: string }[] = $state([]);
	let transactionExpenseTypeOptions: { value: string; label: string }[] = $state([]);
	const isExpense = $derived($formData.amount < 0);
	const transactionTypeOptions = $derived(isExpense ? transactionExpenseTypeOptions : transactionIncomeTypeOptions);
	const creating = $derived($formData.id === -1);

	$effect(() => {
		let date = new Date();
		if (transaction === null) {
			startDate = new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
		} else if (transaction.recurring === null) {
			date = new Date(transaction.date);
			startDate = new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
		} else {
			date = new Date(transaction.recurring.startDate!);
			startDate = new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
			date = new Date(transaction.recurring.endDate!);
			endDate = new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
		}

		const startDateString = transaction
			? transaction.recurring
				? new Date(transaction.recurring.startDate!).toISOString()
				: new Date(transaction.date).toISOString()
			: new Date().toISOString();
		const endDateString =
			transaction && transaction.recurring ? new Date(transaction.recurring.endDate!).toISOString() : "";

		$formData = {
			id: transaction?.id ?? -1,
			amount: transaction?.amount ?? 0,
			startDate: startDateString,
			endDate: endDateString,
			description: transaction?.description ?? "",
			recurring: !!transaction?.recurring,
			interval: transaction?.recurring?.interval ?? "",
			daysInterval: transaction?.recurring?.daysInterval ?? 1,
			type: transaction?.type ?? "",
		};
	});

	$effect(() => {
		transactionIncomeTypeOptions = incomeTypes.map((type) => {
			return { value: type, label: type };
		});
		transactionExpenseTypeOptions = expenseTypes.map((type) => {
			return { value: type, label: type };
		});
	});

	$effect(() => {
		$inspect($formData);
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
	<form method="POST" class="mt-6 flex flex-col gap-4" use:enhance>
		<input hidden name="id" value={$formData.id} />
		<Form.Field {form} name="amount">
			<Form.Control let:attrs>
				<Form.Label>Amount</Form.Label>
				<Input {...attrs} type="number" step=".01" bind:value={$formData.amount} />
				<Form.FieldErrors />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="description">
			<Form.Control let:attrs>
				<Form.Label>Description</Form.Label>
				<Textarea {...attrs} rows={3} maxlength={512} bind:value={$formData.description} />
				<span class="flex justify-between">
					<Form.FieldErrors />
					<small>
						{$formData.description.length}/512
					</small>
				</span>
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="recurring" class="flex items-center justify-between">
			<Form.Control let:attrs>
				<Form.Label>Recurring</Form.Label>
				<Switch
					{...attrs}
					includeInput
					aria-readonly={!creating}
					disabled={!creating}
					bind:checked={$formData.recurring}
				/>
			</Form.Control>
		</Form.Field>
		{#if $formData.recurring}
			<Form.Field {form} name="interval">
				<Form.Control let:attrs>
					<Form.Label>Interval</Form.Label>
					<Tabs.Root
						{...attrs}
						value={$formData.interval ?? ""}
						onValueChange={(value) => ($formData.interval = value!)}
					>
						<Tabs.List class="grid w-full grid-cols-4">
							{#each transactionIntervals as value}
								<Tabs.Trigger {value}>
									{value}
								</Tabs.Trigger>
							{/each}
						</Tabs.List>
					</Tabs.Root>
					<Form.FieldErrors />
					<input hidden name={attrs.name} value={$formData.interval} />
				</Form.Control>
			</Form.Field>
			{#if $errors.interval}
				<small class="text-destructive">{$errors.interval}</small>
			{/if}
			{#if $formData.interval === "Other"}
				<Form.Field {form} name="daysInterval">
					<Form.Control let:attrs>
						<Form.Label for="daysInterval">Every {$formData.daysInterval ?? 1} days</Form.Label>
						<Input
							{...attrs}
							type="number"
							class={$errors.daysInterval && "error"}
							bind:value={$formData.daysInterval}
						/>
						<Form.FieldErrors />
					</Form.Control>
				</Form.Field>
			{/if}
		{/if}
		<Form.Field {form} name="startDate" class="flex flex-col">
			<Form.Control let:attrs>
				<Form.Label>
					{#if $formData.recurring}
						Start Date
					{:else}
						Date
					{/if}
				</Form.Label>
				<Popover.Root openFocus>
					<Popover.Trigger {...attrs} asChild let:builder>
						<Button
							variant="outline"
							class={`w-[280px] justify-start ${!startDate && "text-muted-foreground"}`}
							builders={[builder]}
						>
							<CalendarIcon class="mr-2 h-4 w-4" />
							{startDate ? formatDate(startDate) : `Select a ${$formData.recurring ? "start" : ""} date`}
						</Button>
					</Popover.Trigger>
					<Popover.Content class="w-auto p-0">
						<Calendar
							initialFocus
							bind:value={startDate}
							onValueChange={(value) => {
								$formData.startDate = value ? getISOStringUTC(value) : "";
							}}
						/>
					</Popover.Content>
				</Popover.Root>
				<Form.FieldErrors />
				<input hidden name={attrs.name} value={$formData.startDate} />
			</Form.Control>
		</Form.Field>
		{#if $formData.recurring}
			<Form.Field {form} name="endDate" class="flex flex-col">
				<Form.Control let:attrs>
					<Form.Label>End Date</Form.Label>
					<Popover.Root openFocus>
						<Popover.Trigger {...attrs} asChild let:builder>
							<Button
								variant="outline"
								class={`w-[280px] justify-start ${!endDate && "text-muted-foreground"}`}
								builders={[builder]}
							>
								<CalendarIcon class="mr-2 h-4 w-4" />
								{endDate ? formatDate(endDate) : "Select an end date"}
							</Button>
						</Popover.Trigger>
						<Popover.Content class="w-auto p-0">
							<Calendar
								initialFocus
								bind:value={endDate}
								onValueChange={(value) => {
									$formData.endDate = value ? getISOStringUTC(value) : "";
								}}
							/>
						</Popover.Content>
					</Popover.Root>
					<Form.FieldErrors />
					<input hidden name={attrs.name} value={$formData.endDate} />
				</Form.Control>
			</Form.Field>
		{/if}
		<Form.Field {form} name="type" class="flex flex-col">
			<Form.Control let:attrs>
				<Form.Label>Transaction type</Form.Label>
				<Select.Root
					items={transactionTypeOptions}
					selected={transactionTypeOptions.find((option) => option.value === $formData.type)}
					onSelectedChange={(option) => {
						$formData.type = option?.value ?? "";
					}}
				>
					<Select.Trigger {...attrs} class="w-[280px] px-4">
						<Select.Value placeholder="Select a transaction type" />
					</Select.Trigger>
					<Select.Content>
						{#each transactionTypeOptions as option}
							<Select.Item value={option.value} label={option.label}>{option.label}</Select.Item>
						{/each}
					</Select.Content>
				</Select.Root>
				<Form.FieldErrors />
				<input hidden name={attrs.name} value={$formData.type} />
			</Form.Control>
		</Form.Field>
		<Button class="text-slate-50" type="submit" disabled={demo}>
			{#if creating}
				Create transaction
			{:else}
				Update transaction
			{/if}
		</Button>
	</form>
</Dialog.Content>
