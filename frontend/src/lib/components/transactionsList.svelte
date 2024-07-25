<script lang="ts">
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
    <table class="w-full text-left [&_th]:p-4 [&_td]:p-4">
        <thead>
            <tr>
                <th>Date</th>
                <th>Description</th>
                <th>Amount</th>
                <th>Type</th>
                <th></th>
            </tr>
        </thead>
        <tbody class="[&>tr:nth-child(even)]:bg-surface-200/50 [&>tr:nth-child(odd)]:bg-surface-200 dark:[&>tr:nth-child(even)]:bg-surface-700/50 dark:[&>tr:nth-child(odd)]:bg-surface-700">
        {#each transactions as transaction}
            <tr class="m-4 rounded-md">
                <td>
                    {new Date(transaction.date).toLocaleDateString('default', {
                        weekday: 'short',
                        month: 'long',
                        day: 'numeric',})}
                </td>
                <td>{transaction.description}</td>
                <td class="{transaction.incoming ? 'text-green-500' : 'text-red-500'} text-right">{transaction.amount}</td>
                <td>{transaction.type}</td>
                <td>Edit</td>
            </tr>
        {/each}
        </tbody>
    </table>
{/await}
