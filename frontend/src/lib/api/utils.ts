import { getCurrentMonthNumber } from "$lib";

export function getMonth(params: URLSearchParams) {
    let month = Number.parseInt(params.get("month") ?? "0");
    if (month === 0) {
        month = getCurrentMonthNumber();
    }
    return month;
}

export function getYear(params: URLSearchParams) {
    let year = Number.parseInt(params.get("year") ?? "0");
    if (year === 0) {
        year = 2024;
    }
    return year;
}

export function validString(string: string | null) {
    return (
        string !== null && typeof string === "string" && string.trim() !== ""
    );
}

export function validNumber(number: number | null) {
    return number !== null && typeof number === "number";
}

export function validDate(date: Date | string | null) {
    return date !== null && new Date(date).toString() !== "Invalid Date";
}
