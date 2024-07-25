<script lang="ts">
    import Pencil from 'lucide-svelte/icons/pencil';
    import Trash from 'lucide-svelte/icons/trash';
    import type { Transaction } from '../../ambient';

    const { month } = $props();

    async function fetchTransactions() {
        const response = await fetch(`/transactions?month=${month}&year=2024`);
        return (await response.json()) as Transaction[];
    }
</script>

{#await fetchTransactions()}
    <!-- TODO replace with skeleton -->
    <p>Loading..</p>
{:then transactions}
    <table class="w-full text-left mt-4 [&_th]:p-4 [&_td]:p-4 rounded-md">
        <thead class="hidden md:table-header-group">
            <tr>
                <th>Date</th>
                <th>Description</th>
                <th class="text-right">Amount</th>
                <th>Type</th>
                <th></th>
            </tr>
        </thead>
        <tbody class="">
        {#each transactions as transaction}
            <tr class="even:bg-surface-200/50 odd:bg-surface-200 dark:even:bg-surface-700/50 dark:odd:bg-surface-700 
                       my-4 block relative md:table-row rounded-md [&>td:not(:last-child)]:grid [&>td]:grid-cols-[16ch_auto] 
                       md:[&>td:not(:last-child)]:table-cell md:[&>td]:before:hidden [&>td]:before:content-[attr(data-cell)] [&>td]:before:font-bold [&>td]:before:capitalize">
                <td data-cell="date">{new Date(transaction.date).toLocaleDateString('default', { month: 'long', day: 'numeric',})}</td>
                <td data-cell="description">{transaction.description}</td>
                <td data-cell="amount" class="md:text-right">{transaction.incoming ? '' : '-'}{transaction.amount}</td>
                <td data-cell="type">{transaction.type}</td>
                <td data-cell="" class="block absolute top-0 right-0">
                    <button class="icon"><Pencil size={20} /></button>
                    <button class="icon"><Trash size={20} /></button>
                </td>
            </tr>
        {/each}
        </tbody>
    </table>
{/await}
