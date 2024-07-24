<script lang="ts">
import X from 'lucide-svelte/icons/x';
import Menu from 'lucide-svelte/icons/menu';
import User from 'lucide-svelte/icons/user';
import { onMount } from 'svelte';
import click from '$lib/click';
import Logo from './logo.svelte';

type NavLink = {
    link: string
    title: string
};

let navLinks: NavLink[] = [
    {
        link: '/',
        title: 'Home',
    },
    {
        link: '/transactions',
        title: 'Transactions',
    },
    {
        link: '/dashboard',
        title: 'Dashboard',
    },
];

let navOpen = $state(false);
const toggleNav = () => (navOpen = !navOpen);
const closeNav = () => (navOpen = false);

onMount(() => {
    window.addEventListener('resize', () => {
        if (window.innerWidth >= 1024) {
            closeNav();
        }
    });
});
</script>

<header class="flex h-[var(--header-height)] items-center justify-between px-4 lg:h-[var(--header-height-lg)] lg:justify-start">
    <div class="flex w-full items-center justify-between">
        <Logo />
        <nav class="flex h-full items-center">
            <ul class="hidden h-full items-center gap-2 text-xl lg:flex mr-4">
                {#each navLinks as link}
                    <li><a class="p-2" href={link.link}>{link.title}</a></li>
                {/each}
            </ul>

            <a href="/login"><User class="hidden lg:block" /></a>

            <!-- Side navbar -->
            {#if navOpen}
                <div class="bg-surface-400/50 absolute inset-0 backdrop-blur-sm" use:click={closeNav}></div>
            {/if}
            <div class="absolute bottom-0 {navOpen ? 'left-0' : 'left-[-400px]'} bg-surface-700 top-0 w-full max-w-[400px] transition-all duration-300 z-10">
                <button class="absolute right-2 top-2" onclick={closeNav}><X /></button>
                <ul id="side-nav" class="flex h-full w-full flex-col items-center justify-center gap-4 text-xl">
                    {#each navLinks as link}
                        <li><a href={link.link} use:click={closeNav}>{link.title}</a></li>
                    {/each}
                </ul>
            </div>
            <!-- Side navbar -->
        </nav>
    </div>

    <button class="block lg:hidden" onclick={toggleNav}><Menu size={32} /></button>
</header>
