<script lang="ts">
    import EllipsisVertical from 'lucide-svelte/icons/ellipsis-vertical';
    import type { Transaction } from '../../ambient';
    import click from '$lib/click';
    import { getFormattedDateShort } from '$lib';

    const {
        transactions,
        selectTransaction,
    }: {
        transactions: Transaction[];
        selectTransaction: (t: Transaction | null) => void;
    } = $props();
</script>

<div class="w-full overflow-auto">
    <table class="mt-4 w-full rounded-md text-left [&_th]:p-4">
        <thead>
            <tr>
                <th>Date</th>
                <th>Description</th>
                <th class="text-right">Amount</th>
                <th>Type</th>
                <th></th>
            </tr>
        </thead>
        <tbody class="transactions-table-body">
            {#each transactions as transaction}
                <tr class="transactions-table-row" use:click={() => selectTransaction(transaction)}>
                    <td data-cell="date">{getFormattedDateShort(transaction.date)}</td>
                    <td data-cell="description">{transaction.description}</td>
                    <td data-cell="amount">{transaction.incoming ? '' : '-'}{transaction.amount}</td
                    >
                    <td data-cell="type">{transaction.type}</td>
                    <td data-cell="">
                        <button class="icon"><EllipsisVertical size={20} /></button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
    {#if transactions.length === 0}
        <p class="ml-4">You have no registered transactions for this month.</p>
    {/if}
</div>
