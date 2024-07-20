import { DB_HOST, DB_NAME, DB_PASSWORD, DB_PORT, DB_USERNAME } from "$env/static/private";
import postgres from "postgres";

export const db = postgres({
    host: DB_HOST,
    port: Number.parseInt(DB_PORT),
    database: DB_NAME,
    username: DB_USERNAME,
    password: DB_PASSWORD
});

export interface DatabaseUser {
    id: string
    username: string
    password_hash: string
}
