import { type RequestHandler } from "@sveltejs/kit";
import { authorizeFetch, authorizePost } from "$lib";
import type { TransactionForm, TransactionFormErrors } from "../../../ambient";

export const GET: RequestHandler = async ({ locals: { session }, url }) => {
	if (!session) {
		return new Response("Forbidden", {
			status: 403,
		});
	}

	const month = url.searchParams.get("month");
	const year = url.searchParams.get("year");
	const income = url.searchParams.get("income");
	const fetchUrl = `transactions?month=${month}&year=${year}${income === null ? "" : `&income=${income}`}`;
	const response = await authorizeFetch(fetchUrl, session?.accessToken);
	if (response.ok) {
		return response;
	}

	return new Response("Something went wrong", {
		status: response.status,
	});
};

export const POST: RequestHandler = async ({
	locals: { session },
	request,
}) => {
	const form: TransactionForm = await request.json();
	const errors = verifyForm(form);
	const isValid = Object.values(errors).every((err) => err === null);
	if (!isValid) {
		form.errors = errors;
		return new Response(JSON.stringify(form), {
			status: 400,
			headers: { "Content-Type": "application/json" },
		});
	}

	// Fix the dates
	form.startDate = new Date(form.startDate!);
	form.endDate = form.recurring ? new Date(form.endDate!) : form.startDate;

	const response = await authorizePost(
		"transactions",
		session?.accessToken ?? "",
		JSON.stringify(form),
	);
	if (response.ok) {
		return response;
	}

	return new Response(JSON.stringify(form), {
		status: 500,
		headers: { "Content-Type": "application/json" },
	});
};

function verifyForm(form: TransactionForm): TransactionFormErrors {
	const errors: TransactionFormErrors = {
		amount: null,
		description: null,
		startDate: null,
		endDate: null,
		interval: null,
		daysInterval: null,
		type: null,
	};

	if (!validNumber(form.amount)) {
		errors.amount = "Amount must be a positive number";
	}
	if (!validString(form.description)) {
		errors.description = "Description is required";
	}
	if (!validDate(form.startDate)) {
		errors.startDate = "Start date must be a valid date";
	}
	if (form.recurring) {
		if (!validDate(form.endDate)) {
			errors.endDate = "End date must be a valid date or null";
		}
		if (!validString(form.interval)) {
			errors.interval =
				"Recurring interval is required when a transaction recurring";
		}
		if (form.interval === "Other" && !validNumber(form.daysInterval)) {
			errors.daysInterval = "Interval in days should be set";
		}
	}
	if (!validString(form.type)) {
		errors.type = "Transaction type must be a non-empty string or null";
	}

	return errors;
}

function validString(string: string | null) {
	return (
		string !== null && typeof string === "string" && string.trim() !== ""
	);
}

function validNumber(number: number | null) {
	return number !== null && typeof number === "number" && number > 0;
}

function validDate(date: Date | null) {
	return date !== null && new Date(date).toString() !== "Invalid Date";
}
