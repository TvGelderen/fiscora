<script lang="ts">
    import { SlideToggle, RadioGroup, RadioItem } from '@skeletonlabs/skeleton';

    const {
        transactionIntervals,
        incomeTypes,
        expenseTypes,
    }: {
        transactionIntervals: string[];
        incomeTypes: string[];
        expenseTypes: string[];
    } = $props();

    console.log(transactionIntervals);

    type Form = {
        amount: number;
        incoming: boolean;
        description: string;
        startDate: Date;
        endDate: Date | null;
        recurring: boolean;
        recurringInterval: string | null;
        daysInterval: number | null;
        transactionType: string | null;
    };

    let formData: Form = $state({
        amount: 0,
        incoming: false,
        startDate: new Date(),
        description: '',
        recurring: false,
        recurringInterval: null,
        daysInterval: null,
        endDate: null,
        transactionType: null,
    });

    async function submitTransaction(event: SubmitEvent) {
        event.preventDefault();

        const formElement = event.target as HTMLFormElement;

        const response = await fetch(formElement.action, {
            method: 'POST',
            body: JSON.stringify(formData),
        });

        const responseData = await response.json();

        console.log(responseData);
    }
</script>

<div class="card p-6">
    <h2>Add transaction</h2>
    <form
        onsubmit={submitTransaction}
        action="/transactions"
        method="POST"
        class="mt-6 flex flex-col gap-4"
    >
        <label class="label">
            <span>Amount</span>
            <input
                bind:value={formData.amount}
                class="input p-1"
                type="number"
                placeholder="0"
            />
        </label>
        <label class="label">
            <span>Date</span>
            <input
                bind:value={formData.startDate}
                class="input p-1"
                type="date"
                placeholder=""
                step="30"
            />
        </label>
        <label class="label">
            <span>Desription</span>
            <textarea
                bind:value={formData.description}
                class="textarea p-1"
                rows="3"
                placeholder="Description..."
            ></textarea>
        </label>
        <label class="label flex flex-col">
            <span>Recurring</span>
            <SlideToggle
                name="slide"
                bind:checked={formData.recurring}
                active="bg-primary-500"
                size="sm"
            />
        </label>
        {#if formData.recurring}
            <RadioGroup
                active="variant-filled-primary"
                hover="hover:variant-soft-primary"
            >
                {#each transactionIntervals as value}
                    <RadioItem
                        bind:group={formData.recurringInterval}
                        name="justify"
                        {value}>{value}</RadioItem
                    >
                {/each}
            </RadioGroup>
            {#if formData.recurringInterval === 'other'}
                <label class="label">
                    <span>Every (x) days</span>
                    <input
                        bind:value={formData.daysInterval}
                        class="input p-1"
                        type="number"
                        placeholder="1"
                        min="1"
                    />
                </label>
            {/if}
            <label class="label">
                <span>End Date</span>
                <input
                    bind:value={formData.endDate}
                    class="input p-1"
                    type="date"
                    placeholder=""
                />
            </label>
        {/if}
        <label class="label flex flex-col">
            <span>Incoming</span>
            <SlideToggle
                name="slide"
                bind:checked={formData.incoming}
                active="bg-primary-500"
                size="sm"
            />
        </label>
        <label class="label">
            <span>Transaction type</span>
            <select
                class="select"
                size="4"
                bind:value={formData.transactionType}
            >
                {#if formData.incoming}
                    {#each incomeTypes as value}
                        <option {value}>{value}</option>
                    {/each}
                {:else}
                    {#each expenseTypes as value}
                        <option {value}>{value}</option>
                    {/each}
                {/if}
            </select>
        </label>
        <button class="btn" type="submit">Add transaction</button>
    </form>
</div>
