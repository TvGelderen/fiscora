<script lang="ts">
import {
    SlideToggle,
    RadioGroup,
    RadioItem,
} from '@skeletonlabs/skeleton';
    import { json } from '@sveltejs/kit';

// TODO: Move to a types directory for backend use as well
enum ERecurringInterval {
    Daily = 0,
    Weekly = 1,
    Monthly = 2,
    Other = 3,
};

enum ETransactionType {
    Income = 0,
    Housing = 1,
    Food = 2,
    Transportation = 3,
    Utilities = 4,
    Entertainment = 5,
    Other = 6,
};

type Form = {
    amount: number;
    incoming: boolean;
    description: string;
    startDate: Date;
    endDate: Date | null;
    recurring: boolean;
    recurringInterval: ERecurringInterval;
    daysInterval: number | null;
    transactionType: ETransactionType | null;
}

let formData: Form = $state({
    amount: 0,
    incoming: false,
    startDate: new Date(),
    description: '',
    recurring: false,
    recurringInterval: ERecurringInterval.Monthly,
    daysInterval: null,
    endDate: null,
    transactionType: null
});

function valid() {
    
}

async function submitTransaction(event: SubmitEvent) {
    event.preventDefault();

    const formElement = event.target as HTMLFormElement;

    const response = await fetch(formElement.action, {
        method: "POST",
        body: JSON.stringify(formData)
    });

    const responseData = await response.json();

    console.log(response);
    console.log(responseData);
}
</script>

<div class="card p-6">
    <h2>Add transaction</h2>
    <form onsubmit={submitTransaction} action="/transactions" method="POST" class="mt-6 flex flex-col gap-4">
        <label class="label">
            <span>Amount</span>
            <input bind:value={formData.amount} class="input p-1" type="number" placeholder="0" />
        </label>
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
            <span>Date</span>
            <input bind:value={formData.startDate} class="input p-1" type="date" placeholder="" step="30" />
        </label>
        <label class="label">
            <span>Desription</span>
            <textarea bind:value={formData.description} class="textarea p-1" rows="3" placeholder="Description..."
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
            <RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
                <RadioItem
                    bind:group={formData.recurringInterval}
                    name="justify"
                    value={ERecurringInterval.Daily}
                >Daily</RadioItem>
                <RadioItem
                    bind:group={formData.recurringInterval}
                    name="justify"
                    value={ERecurringInterval.Weekly}
                >Weekly</RadioItem>
                <RadioItem
                    bind:group={formData.recurringInterval}
                    name="justify"
                    value={ERecurringInterval.Monthly}
                >Monthly</RadioItem>
                <RadioItem
                    bind:group={formData.recurringInterval}
                    name="justify"
                    value={ERecurringInterval.Other}
                >Other</RadioItem>
            </RadioGroup>
            {#if formData.recurringInterval === ERecurringInterval.Other}
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
                <input bind:value={formData.endDate} class="input p-1" type="date" placeholder="" />
            </label>
        {/if}
        <label class="label">
            <span>Transaction category</span>
            <select class="select" size="4" bind:value={formData.transactionType}>
                <option value={ETransactionType.Income}>Income</option>
                <option value={ETransactionType.Housing}>Housing</option>
                <option value={ETransactionType.Food}>Food</option>
                <option value={ETransactionType.Transportation}>Transportation</option>
                <option value={ETransactionType.Utilities}>Utilities</option>
                <option value={ETransactionType.Entertainment}>Entertainment</option>
                <option value={ETransactionType.Other}>Other</option>
            </select>
        </label>
        <button class="btn" type="submit">Add transaction</button>
    </form>
</div>
