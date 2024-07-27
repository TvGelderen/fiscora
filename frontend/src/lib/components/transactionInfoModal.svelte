<script lang="ts">
    import X from 'lucide-svelte/icons/x';
    import type { Transaction } from '../../ambient';
    import { getFormattedDate } from '$lib';

    const {
        transaction,
        onclose,
    }: {
        transaction: Transaction | null;
        onclose: () => void;
    } = $props();

    let modal: HTMLDialogElement;

    $effect(() => {
        if (transaction === null && modal.open) {
            modal.close();
        } else if (transaction !== null && !modal.open) {
            modal.showModal();
        }
    });
</script>

<dialog class="bg-surface-200-700-token w-[640px] max-w-[95%] p-4 lg:p-6" bind:this={modal}>
    {#if transaction}
        <button class="icon absolute right-4 top-4" onclick={onclose}><X /></button>
        <h2 class="mb-4">Transaction details</h2>
        <div class="mt-6 grid grid-cols-1 sm:grid-cols-2 lg:mt-8">
            <div class="mb-4">
                <p class="header-sm">Transaction date</p>
                <p>{getFormattedDate(transaction.date)}</p>
            </div>
            <div class="mb-4">
                <p class="header-sm">Amount</p>
                <p>€{transaction.incoming ? '' : '-'}{transaction.amount}</p>
            </div>
            <div class="mb-4">
                <p class="header-sm">Type</p>
                <p>{transaction.type}</p>
            </div>
            {#if transaction.recurring}
                <div class="mb-4">
                    <p class="header-sm">Recurring</p>
                    <p>{transaction.interval}</p>
                </div>
                <div class="mb-4">
                    <p class="header-sm">Start date</p>
                    <p>{getFormattedDate(transaction.startDate!)}</p>
                </div>
                <div class="mb-4">
                    <p class="header-sm">End date</p>
                    <p>{getFormattedDate(transaction.endDate!)}</p>
                </div>
            {/if}
        </div>
        <div>
            <p class="header-sm">Description</p>
            <p>{transaction.description}</p>
        </div>
    {:else}
        <div><span class="text-error-400">No transaction selected.</span></div>
    {/if}
</dialog>
