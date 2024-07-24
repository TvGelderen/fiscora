<script lang="ts">
    import { page } from '$app/stores';
    import TransactionForm from '$lib/components/transactionForm.svelte';
    import TransactionsList from '$lib/components/transactionsList.svelte';

    const { transactionIntervals, incomeTypes, expenseTypes } = $page.data;

    function listAllMonths() {
        const months = new Map<number, string>();
        for (let month = 0; month < 12; month++) {
            const monthName = new Date(2000, month, 1).toLocaleString(
                'default',
                { month: 'long' },
            );
            months.set(month + 1, monthName);
        }
        return months;
    }

    let month = $state(
        new Date().toLocaleString('default', { month: 'numeric' }),
    );

    $effect(() => {});
</script>

<title>Budget Buddy - Transactions</title>

<div class="grid h-full items-center gap-8 md:grid-cols-2 lg:gap-12">
    <div>
        <h1>Manage your transactions</h1>
        <p class="mt-2 text-lg">
            View, categorize, and edit your recent transactions to better
            understand your spending and savings.
        </p>
        <div class="mt-6 grid grid-cols-2 gap-4">
            <div class="card flex flex-col justify-between p-4">
                <h3>Total transactions</h3>
                <div class="mt-6">
                    <p class="text-2xl lg:text-3xl">42</p>
                    <p>This month</p>
                </div>
            </div>
            <div class="card flex flex-col justify-between p-4">
                <h3>Expenses</h3>
                <div class="mt-6">
                    <p class="text-2xl lg:text-3xl">36</p>
                    <p>This month</p>
                </div>
            </div>
        </div>
    </div>
    <TransactionForm {transactionIntervals} {incomeTypes} {expenseTypes} />
</div>

<div class="card h-screen w-full">
    {#each listAllMonths() as [key, value]}
        <p>{key}: {value}</p>
    {/each}
    <TransactionsList {month} />
</div>
