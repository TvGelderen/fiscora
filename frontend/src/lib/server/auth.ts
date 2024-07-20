import { Lucia } from "lucia";
import { dev } from "$app/environment";
import { PostgresJsAdapter } from "@lucia-auth/adapter-postgresql";
import { db, type DatabaseUser } from "./db";
import { Google } from "arctic";
import { GOOGLE_ID, GOOGLE_SECRET, GOOGLE_CALLBACK } from "$env/static/private";

const adapter = new PostgresJsAdapter(db, {
    user: "postgres",
    session: "auth_session"
});

export const lucia = new Lucia(adapter, {
    sessionCookie: {
        attributes: {
            secure: !dev
        }
    },
    getUserAttributes: (attributes) => {
        return {
            username: attributes.username
        };
    }
});

declare module "lucia" {
    interface Register {
        Lucia: typeof lucia;
        DatabaseUserAttributes: DatabaseUser
    }
}

export const google = new Google(GOOGLE_ID, GOOGLE_SECRET, GOOGLE_CALLBACK);
