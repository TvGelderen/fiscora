// See https://kit.svelte.dev/docs/types#app

import type { User, Session } from '$lib/types';

// for information about these interfaces
declare global {
    namespace App {
        // interface Error {}
        interface Locals {
            user: User | null;
            session: Session | null;
        }
        // interface PageData {}
        // interface PageState {}
        // interface Platform {}
    }
}
