export function getFormattedDate(date: Date): string {
	return new Date(date).toLocaleDateString("default", {
		weekday: "short",
		year: "numeric",
		month: "long",
		day: "numeric",
	});
}

export function getFormattedDateShort(date: Date): string {
	return new Date(date).toLocaleDateString("default", {
		month: "short",
		day: "numeric",
	});
}

export function getFormattedDateShortWithYear(date: Date): string {
	return new Date(date).toLocaleDateString("default", {
		month: "short",
		day: "numeric",
		year: "numeric",
	});
}

export function getFormDate(date: Date): string {
	return new Date(date).toISOString().split("T")[0];
}

export function toISOString(date: string): string {
	return new Date(date).toISOString();
}

export function getFormattedAmount(amount: number) {
	return `€${amount.toLocaleString("default", {
		minimumFractionDigits: 2,
		maximumFractionDigits: 2,
	})}`;
}

export const forbidden = () =>
	new Response("Forbidden", {
		status: 403,
	});

export function getCurrentMonthNumber() {
	return Number.parseInt(new Date().toLocaleString("default", { month: "numeric" }));
}

export function listAllMonths() {
	const months = new Map<number, string>();
	for (let month = 0; month < 12; month++) {
		const monthName = new Date(2000, month, 1).toLocaleString("default", {
			month: "long",
		});
		months.set(month + 1, monthName);
	}
	return months;
}

export function listAllMonthNames() {
	const months: string[] = [];
	for (let month = 0; month < 12; month++) {
		const monthName = new Date(2000, month, 1).toLocaleString("default", {
			month: "long",
		});
		months.push(monthName);
	}
	return months;
}

export function listAllMonthNamesShort() {
	const months: string[] = [];
	for (let month = 0; month < 12; month++) {
		const monthName = new Date(2000, month, 1).toLocaleString("default", {
			month: "short",
		});
		months.push(monthName);
	}
	return months;
}

export function createRandomString(length: number): string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
	let result = "";
	for (let i = 0; i < length; i++) {
		result += chars.charAt(Math.floor(Math.random() * chars.length));
	}
	return result;
}
