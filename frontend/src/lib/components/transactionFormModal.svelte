<script lang="ts">
    import X from 'lucide-svelte/icons/x';
    import { SlideToggle, RadioGroup, RadioItem } from '@skeletonlabs/skeleton';
    import type { TransactionForm, TransactionFormErrors } from '../../ambient';

    let {
        transactionIntervals,
        incomeTypes,
        expenseTypes,
        open,
        onclose,
    }: {
        transactionIntervals: string[];
        incomeTypes: string[];
        expenseTypes: string[];
        open: boolean;
        onclose: (() => void);
    } = $props();

    let defaultForm = {
        amount: 0,
        incoming: false,
        startDate: null,
        description: '',
        recurring: false,
        interval: null,
        daysInterval: null,
        endDate: null,
        type: null,
        errors: <TransactionFormErrors>{},
    };

    let modal: HTMLDialogElement;
    let form: TransactionForm = $state({ ...defaultForm });

    async function submitTransaction(event: SubmitEvent) {
        event.preventDefault();

        const response = await fetch('/api/transactions', {
            method: 'POST',
            body: JSON.stringify(form),
        });

        if (!response.ok) {
            form = await response.json();
            return;
        }

        onclose();
    }

    $effect(() => {
        if (open) {
            modal.showModal();
        } else {
            modal.close();
        }
    });
</script>

<dialog class="w-[500px] max-w-[95%]" bind:this={modal}>
    <button class="absolute top-4 right-4" onclick={onclose}><X /></button>
    <h2>Add transaction</h2>
    <form onsubmit={submitTransaction} class="mt-6 flex flex-col gap-4">
        <label class="label">
            <span>Amount</span>
            <input
                bind:value={form.amount}
                class="input p-1 {form.errors.amount && 'error'}"
                type="number"
                step="0.01"
            />
            {#if form.errors.amount}
                <small class="text-error-500">{form.errors.amount}</small>
            {/if}
        </label>
        <label class="label">
            <span>Date</span>
            <input
                bind:value={form.startDate}
                class="input p-1 {form.errors.startDate && 'error'}"
                type="date"
                placeholder=""
            />
            {#if form.errors.startDate}
                <small class="text-error-500">{form.errors.startDate}</small>
            {/if}
        </label>
        <label class="label">
            <span>Desription</span>
            <textarea
                bind:value={form.description}
                class="input p-1 {form.errors.description && 'error'}"
                placeholder="Description..."
                maxlength="512"
                rows="3"
            ></textarea>
            <small class="float-right">{form.description.length}/512</small>
            {#if form.errors.description}
                <small class="text-error-500">{form.errors.description}</small>
            {/if}
        </label>
        <label class="label flex flex-col">
            <span>Recurring</span>
            <SlideToggle
                name="slide"
                bind:checked={form.recurring}
                active="bg-primary-500"
                size="sm"
            />
        </label>
        {#if form.recurring}
            <RadioGroup
                active="variant-filled-primary"
                hover="hover:variant-soft-primary"
                class={form.errors.interval && 'error'}
            >
                {#each transactionIntervals as value}
                    <RadioItem bind:group={form.interval} name="justify" {value}
                        >{value}</RadioItem
                    >
                {/each}
            </RadioGroup>
            {#if form.errors.interval}
                <small class="text-error-500">{form.errors.interval}</small>
            {/if}
            {#if form.interval === 'Other'}
                <label class="label">
                    <span>Every (x) days</span>
                    <input
                        bind:value={form.daysInterval}
                        class="input p-1 {form.errors.daysInterval && 'error'}"
                        type="number"
                        placeholder="1"
                        min="1"
                    />
                    {#if form.errors.daysInterval}
                        <small class="text-error-500"
                            >{form.errors.daysInterval}</small
                        >
                    {/if}
                </label>
            {/if}
            <label class="label">
                <span>End Date</span>
                <input
                    bind:value={form.endDate}
                    class="input p-1 {form.errors.endDate && 'error'}"
                    type="date"
                    placeholder=""
                />
                {#if form.errors.endDate}
                    <small class="text-error-500">{form.errors.endDate}</small>
                {/if}
            </label>
        {/if}
        <label class="label flex flex-col">
            <span>Incoming</span>
            <SlideToggle
                name="slide"
                bind:checked={form.incoming}
                active="bg-primary-500"
                size="sm"
            />
        </label>
        <label class="label">
            <span>Transaction type</span>
            <select
                class="select {form.errors.type && 'error'}"
                bind:value={form.type}
            >
                {#if form.incoming}
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
                <small class="text-error-500">{form.errors.type}</small>
            {/if}
        </label>
        <button class="btn" type="submit">Add transaction</button>
    </form>
</dialog>
