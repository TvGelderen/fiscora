<script lang="ts">
	import X from "lucide-svelte/icons/x";
	import { SlideToggle, RadioGroup, RadioItem, getToastStore } from "@skeletonlabs/skeleton";
	import type { Transaction, TransactionForm, TransactionFormErrors } from "../../ambient";
	import { popup } from "@skeletonlabs/skeleton";
	import { getFormDate } from "$lib";

	const toastStore = getToastStore();

	let {
		transaction,
		transactionIntervals,
		incomeTypes,
		expenseTypes,
		open,
		demo,
		close,
		success,
	}: {
		transaction: Transaction | null;
		transactionIntervals: string[];
		incomeTypes: string[];
		expenseTypes: string[];
		open: boolean;
		demo: boolean;
		close: () => void;
		success: () => void;
	} = $props();

	const defaultForm = (): TransactionForm => {
		let startDate: string | null = null;
		let endDate: string | null = null;

		if (transaction !== null) {
			if (transaction.recurring === null) {
				startDate = getFormDate(transaction.date);
			} else {
				startDate = getFormDate(transaction.recurring.startDate!);
				endDate = getFormDate(transaction.recurring.endDate!);
			}
		}

		return {
			id: transaction?.id ?? -1,
			amount: transaction?.amount ?? 0,
			startDate: startDate,
			description: transaction?.description ?? "",
			recurring: !!transaction?.recurring,
			interval: transaction?.recurring?.interval ?? null,
			daysInterval: transaction?.recurring?.daysInterval ?? null,
			endDate: endDate,
			type: transaction?.type ?? null,
			errors: <TransactionFormErrors>{},
		};
	};

	let modal: HTMLDialogElement;
	let form: TransactionForm = $state(defaultForm());

	async function submitTransaction(event: SubmitEvent) {
		event.preventDefault();

		if (demo) {
			toastStore.trigger({
				message: "Demo users cannot create budgets",
				background: "variant-filled-warning",
			});
			return;
		}

		let response: Response;
		if (transaction === null) {
			response = await fetch("/api/transactions", {
				method: "POST",
				body: JSON.stringify(form),
			});
		} else {
			response = await fetch(`/api/transactions/${transaction.id}`, {
				method: "PUT",
				body: JSON.stringify(form),
			});
		}

		if (!response.ok) {
			form = await response.json();
			return;
		}

		const created = transaction === null;

		toastStore.trigger({
			message: `Transaction ${created ? "created" : "updated"} successfully`,
			background: "variant-filled-success",
		});

		if (!created) {
			transaction = null;
		}

		form = defaultForm();

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

<dialog class="w-full max-w-lg" bind:this={modal}>
	<button class="absolute right-4 top-4" onclick={close}>
		<X />
	</button>
	{#if transaction === null}
		<h3>Add Transaction</h3>
	{:else}
		<h3>Update Transaction</h3>
	{/if}
	<form onsubmit={submitTransaction} class="mt-6 flex flex-col gap-4">
		<label class="label" for="amount">
			<span>Amount</span>
			<input
				id="amount"
				name="amount"
				bind:value={form.amount}
				class="input p-1 {form.errors.amount && 'error'}"
				type="number"
				step="0.01"
			/>
			{#if form.errors.amount}
				<small class="error-text">{form.errors.amount}</small>
			{/if}
		</label>
		<label class="label mb-4" for="description">
			<span>Desription</span>
			<textarea
				id="description"
				name="description"
				bind:value={form.description}
				class="input p-1 {form.errors.description && 'error'}"
				placeholder="Description..."
				maxlength="512"
				rows="3"
			></textarea>
			<span class="relative !mt-0 flex">
				<small class="absolute right-0 top-0 float-right leading-none">
					{form.description.length}/512
				</small>
				{#if form.errors.description}
					<small class="error-text leading-none">
						{form.errors.description}
					</small>
				{/if}
			</span>
		</label>
		{#if transaction === null}
			<label class="label flex items-center justify-between">
				<span>Recurring</span>
				<SlideToggle
					name="slide"
					bind:checked={form.recurring}
					active="bg-primary-500"
					size="sm"
					disabled={transaction !== null}
				/>
			</label>
		{/if}
		{#if form.recurring}
			<RadioGroup
				active="variant-filled-primary"
				hover="hover:variant-soft-primary"
				class={form.errors.interval && "error"}
			>
				{#each transactionIntervals as value}
					<RadioItem bind:group={form.interval} name="justify" {value}>
						{value}
					</RadioItem>
				{/each}
			</RadioGroup>
			{#if form.errors.interval}
				<small class="error-text">{form.errors.interval}</small>
			{/if}
			{#if form.interval === "Other"}
				<label class="label">
					<span>Every {form.daysInterval ?? 1} days</span>
					<input
						bind:value={form.daysInterval}
						class="input p-1 {form.errors.daysInterval && 'error'}"
						type="number"
						placeholder="1"
						min="1"
					/>
					{#if form.errors.daysInterval}
						<small class="error-text">
							{form.errors.daysInterval}
						</small>
					{/if}
				</label>
			{/if}
		{/if}
		<div class={`grid ${form.recurring ? "grid-cols-2" : "grid-cols-1"} gap-2`}>
			<label class="label" for="startDate">
				{#if form.recurring}
					<span>Start Date</span>
				{:else}
					<span>Date</span>
				{/if}
				<input
					id="startDate"
					name="startDate"
					bind:value={form.startDate}
					class="input p-1 {form.errors.startDate && 'error'}"
					type="date"
					placeholder=""
				/>
				{#if form.errors.startDate}
					<small class="error-text">{form.errors.startDate}</small>
				{/if}
			</label>
			{#if form.recurring}
				<label class="label">
					<span>End Date</span>
					<input
						bind:value={form.endDate}
						class="input p-1 {form.errors.endDate && 'error'}"
						type="date"
						placeholder=""
					/>
					{#if form.errors.endDate}
						<small class="error-text">{form.errors.endDate}</small>
					{/if}
				</label>
			{/if}
		</div>
		<label class="label">
			<span>Transaction type</span>
			<select class="select {form.errors.type && 'error'}" bind:value={form.type}>
				{#if form.amount > 0}
					{#each incomeTypes as value}
						<option {value}>{value}</option>
					{/each}
				{:else}
					{#each expenseTypes as value}
						<option {value}>{value}</option>
					{/each}
				{/if}
			</select>
			{#if form.errors.type}
				<small class="error-text">{form.errors.type}</small>
			{/if}
		</label>
		<button
			class="btn"
			type="submit"
			disabled={demo}
			use:popup={{
				event: "hover",
				target: "form-button-tooltip",
				placement: "bottom",
				middleware: {
					offset: 10,
				},
			}}
		>
			{#if transaction === null}
				Add transaction
			{:else}
				Update transaction
			{/if}
		</button>
	</form>
	{#if demo}
		<div class="bg-surface-200-700-token rounded-md p-4 shadow-lg" data-popup="form-button-tooltip">
			You are not allowed to {transaction === null ? "create" : "update"} transactions as a demo user
			<div class="bg-surface-200-700-token arrow"></div>
		</div>
	{/if}
</dialog>
