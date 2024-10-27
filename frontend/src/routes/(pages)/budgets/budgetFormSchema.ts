import { validDate } from "$lib/api/utils";
import { z } from "zod";

export const budgetFormSchema = z.object({
    id: z.number(),
    name: z
        .string({ required_error: "Name is required." })
        .min(1, { message: "Please enter a name." }),
    description: z
        .string({ required_error: "Description is required." })
        .min(1, { message: "Please enter a description." })
        .max(256, { message: "Description cannot be more than 256 characters." })
        .trim(),
    amount: z
        .number({ required_error: "Amount is required." })
        .positive({ message: "Amount must be greate than 0." }),
    startDate: z.string(),
    endDate: z.string(),
}).superRefine((val, ctx) => {
    if (val.amount === 0) {
        ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Amount cannot be 0.",
            path: ['amount'],
        });
    }

    if (!validDate(val.startDate)) {
        ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Please enter a valid start date.",
            path: ['startDate'],
        });
    }
    if (!validDate(val.endDate)) {
        ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Please enter a valid end date.",
            path: ['endDate'],
        });
    }
});

export type FormSchema = typeof budgetFormSchema;
