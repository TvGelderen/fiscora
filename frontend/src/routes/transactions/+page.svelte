<script lang="ts">
    import Plus from 'lucide-svelte/icons/plus';
    import X from 'lucide-svelte/icons/x';
    import { page } from '$app/stores';
    import TransactionForm from '$lib/components/transactionForm.svelte';
    import TransactionsList from '$lib/components/transactionsList.svelte';
    import TransactionInfoModal from '$lib/components/transactionInfoModal.svelte';
    import type { Transaction } from '../../ambient';

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

    let modal: HTMLDialogElement;
    let month = $state(Number.parseInt(new Date().toLocaleString('default', { month: 'numeric' })));
    let transactions: Transaction[] = $state([]);
    let selectedTransaction: Transaction | null = $state(null);

    let income = $derived(transactions.filter(transaction => transaction.incoming).reduce((acc, cur) => acc += cur.amount, 0))
    let expense = $derived(transactions.filter(transaction => !transaction.incoming).reduce((acc, cur) => acc += cur.amount, 0))
    let netIncome = $derived(income - expense);

    function selectTransaction(transaction: Transaction | null) {
        selectedTransaction = transaction;
    }

    function showModal() {
        modal.showModal();
    }

    function closeModal() {
        modal.close();
    }

    async function fetchTransactions() {
        const response = await fetch(`/transactions?month=${month}&year=2024`);
        return (await response.json()) as Transaction[];
    }

    $effect(() => {
        console.log("$effect");
        fetchTransactions().then(data => {
            transactions = data;
        });
    })
</script>

<title>Budget Buddy - Transactions</title>

<div class="mx-auto text-center mb-10 lg:mb-16">
    <h1 class="mb-4">Your transactions</h1>
    <p>Add, view, and edit your transactions to stay on top of your financial journey.</p>
    <p>Track your finances with ease and gain valuable insights.</p>
</div>

<div class="grid sm:grid-cols-3 mb-10 lg:mb-16 rounded-3xl bg-primary-500/20 [&>div]:p-4 [&>div>span]:text-3xl">
    <div class="flex flex-col justify-between">
        <h4 class="mb-6">Total income</h4>
        <span>€{income}</span>
    </div>
    <div class="flex flex-col justify-between border-t-[1px] border-b-[1px] sm:border-t-[0px] sm:border-b-[0px] sm:border-l-[1px] sm:border-r-[1px] border-primary-700/25">
        <h4 class="mb-6">Total expense</h4>
        <span>€{expense}</span>
    </div>
    <div class="flex flex-col justify-between">
        <h4 class="mb-6">Net income</h4>
        <span>€{netIncome}</span>
    </div>
</div>

<div class="flex justify-between items-center flex-col sm:flex-row my-4">
    <h2>Transactions</h2>
    <button class="btn secondary mt-4 sm:mt-0 " onclick={showModal}><Plus />&nbsp;Add transaction</button>
</div>
<div class="card w-full p-4">
    <select id="month-selector" class="select" bind:value={month}>
        {#each listAllMonths() as [idx, name]}
            <option selected="{idx === month}" value={idx}>{name}</option>
        {/each}
    </select>
    <TransactionsList {transactions} {selectTransaction} />
</div>

<dialog class="w-[500px] max-w-[95%]" bind:this={modal}>
    <button class="absolute top-4 right-4" onclick={closeModal}><X /></button>
    <TransactionForm {transactionIntervals} {incomeTypes} {expenseTypes} {closeModal} />
</dialog>

<TransactionInfoModal transaction={selectedTransaction} onclose={() => selectTransaction(null)} />
