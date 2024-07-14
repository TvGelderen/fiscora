<script lang="ts">
import {
    SlideToggle,
    RadioGroup,
    RadioItem,
} from '@skeletonlabs/skeleton'

// TODO: Move to a types directory for backend use as well
enum ERecurringInterval {
    Daily = 0,
    Weekly = 1,
    Monthly = 2,
    Other = 3,
}

enum ETransactionType {
    Income = 0,
    Housing = 1,
    Food = 2,
    Transportation = 3,
    Utilities = 4,
    Entertainment = 5,
    Other = 6,
}

let recurring: boolean = $state(false)
let recurringInterval: ERecurringInterval = $state(ERecurringInterval.Monthly)
let transactionType: ETransactionType | null = $state(null)
</script>

<div class="card p-6">
    <h2>Add transaction</h2>
    <form action="/" method="post" class="mt-6 flex flex-col gap-4">
        <label class="label">
            <span>Amount</span>
            <input class="input p-1" type="number" placeholder="0" />
        </label>
        <label class="label">
            <span>Date</span>
            <input class="input p-1" type="date" placeholder="" />
        </label>
        <label class="label">
            <span>Desription</span>
            <textarea class="textarea p-1" rows="2" placeholder="Description..."
            ></textarea>
        </label>
        <label class="label flex flex-col">
            <span>Recurring</span>
            <SlideToggle
                name="slide"
                bind:checked={recurring}
                active="bg-primary-500"
                size="sm"
            />
        </label>
        {#if recurring}
            <RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
                <RadioItem
                    bind:group={recurringInterval}
                    name="justify"
                    value={ERecurringInterval.Daily}
                >Daily</RadioItem>
                <RadioItem
                    bind:group={recurringInterval}
                    name="justify"
                    value={ERecurringInterval.Weekly}
                >Weekly</RadioItem>
                <RadioItem
                    bind:group={recurringInterval}
                    name="justify"
                    value={ERecurringInterval.Monthly}
                >Monthly</RadioItem>
                <RadioItem
                    bind:group={recurringInterval}
                    name="justify"
                    value={ERecurringInterval.Other}
                >Other</RadioItem>
            </RadioGroup>
            {#if recurringInterval === ERecurringInterval.Other}
                <label class="label">
                    <span>Every (x) days</span>
                    <input
                        class="input p-1"
                        type="number"
                        placeholder="1"
                        min="1"
                    />
                </label>
            {/if}
        {/if}
        <label class="label">
            <span>Transaction category</span>
            <select class="select" size="4" bind:value={transactionType}>
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
