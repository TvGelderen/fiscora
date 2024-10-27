import { validDate, validNumber, validString } from "$lib/api/utils";
import { z } from "zod";

export const transactionFormSchema = z.object({
    id: z.number(),
    amount: z.number({ required_error: "Amount is required." }),
    description: z
        .string({ required_error: "Description is required." })
        .min(1, { message: "Please enter a description." })
        .max(512, { message: "Description cannot be more than 512 characters." })
        .trim(),
    type: z
        .string({ required_error: "Transaction type is required" })
        .min(1, { message: "Transaction type is required" }),
    startDate: z.string(),
    recurring: z.boolean(),
    endDate: z.string().optional(),
    interval: z.string().optional(),
    daysInterval: z.number().optional(),
}).superRefine(({ amount, recurring, startDate, endDate, interval, daysInterval }, ctx) => {
    if (amount === 0) {
        ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Amount cannot be 0.",
            path: ['amount'],
        });
    }

    if (!recurring) {
        if (!validDate(startDate)) {
            ctx.addIssue({
                code: z.ZodIssueCode.custom,
                message: "Please enter a valid date.",
                path: ['startDate'],
            });
        }
        return;
    }

    if (!validDate(startDate)) {
        ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Please enter a valid start date.",
            path: ['startDate'],
        });
    }
    if (!validDate(endDate)) {
        ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Please enter a valid end date.",
            path: ['endDate'],
        });
    }
    if (!validString(interval)) {
        ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Please enter a valid interval.",
            path: ['interval'],
        });
    }
    if (interval === "Other") {
        if (!validNumber(daysInterval)) {
            ctx.addIssue({
                code: z.ZodIssueCode.custom,
                message: "Please enter a valid number.",
                path: ['daysInterval'],
            });
        } else if (daysInterval <= 0) {
            ctx.addIssue({
                code: z.ZodIssueCode.custom,
                message: "Please enter a number greater than 0.",
                path: ['daysInterval'],
            });
        }
    }
});

export type FormSchema = typeof transactionFormSchema;
