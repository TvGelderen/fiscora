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
    {#each transactions as transaction}
        <div class="m-4 rounded bg-surface-600 p-4">
            <span
                >{new Date(transaction.date).toLocaleDateString('default', {
                    weekday: 'short',
                    month: 'long',
                    day: 'numeric',
                })}</span
            >
            <span
                class={transaction.incoming ? 'text-green-500' : 'text-red-500'}
                >{transaction.amount}</span
            >
        </div>
    {/each}
{/await}
