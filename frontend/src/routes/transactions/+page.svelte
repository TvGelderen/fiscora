<script lang="ts">
    import Plus from 'lucide-svelte/icons/plus';
    import X from 'lucide-svelte/icons/x';
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

    let modal: HTMLDialogElement;
    let month = $state(Number.parseInt(new Date().toLocaleString('default', { month: 'numeric' })));

    function showModal() {
        modal.showModal();
    }

    function closeModal() {
        modal.close();
    }
</script>

<title>Budget Buddy - Transactions</title>

<div class="flex justify-between items-center">
    <h2>Transactions</h2>
    <button class="btn" onclick={showModal}><Plus />&nbsp;Add transaction</button>
</div>
<div class="card w-full p-4">
    <select id="month-selector" class="select" bind:value={month}>
        {#each listAllMonths() as [idx, name]}
            <option selected="{idx === month}" value={idx}>{name}</option>
        {/each}
    </select>
    <TransactionsList {month} />
</div>

<dialog class="w-[500px] max-w-[95%]" bind:this={modal}>
    <button class="absolute top-4 right-4" onclick={closeModal}><X /></button>
    <TransactionForm {transactionIntervals} {incomeTypes} {expenseTypes} {closeModal} />
</dialog>
