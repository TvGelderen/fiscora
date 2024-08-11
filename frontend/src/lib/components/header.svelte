<script lang="ts">
	import X from "lucide-svelte/icons/x";
	import Sun from "lucide-svelte/icons/sun";
	import Moon from "lucide-svelte/icons/moon";
	import Menu from "lucide-svelte/icons/menu";
	import User from "lucide-svelte/icons/user";
	import { onMount } from "svelte";
	import click from "$lib/click";
	import Logo from "./logo.svelte";
	import { createDarkMode } from "$lib/theme.svelte";

	type NavLink = {
		link: string;
		title: string;
	};

	let navLinks: NavLink[] = [
		{
			link: "/",
			title: "Home",
		},
		{
			link: "/transactions",
			title: "Transactions",
		},
		{
			link: "/dashboard",
			title: "Dashboard",
		},
	];

	let navOpen = $state(false);

	const darkMode = createDarkMode();

	const toggleNav = () => (navOpen = !navOpen);
	const closeNav = () => (navOpen = false);

	const toggleTheme = () => {
		darkMode.toggle();
		const theme = darkMode.darkMode ? "dark" : "light";
		const html = document.querySelector("html");
		localStorage.setItem("theme", theme);
		if (html) {
			html.classList.value = theme;
		}
	};

	onMount(() => {
		window.addEventListener("resize", () => {
			if (window.innerWidth >= 1024) {
				closeNav();
			}
		});

		const theme = localStorage.getItem("theme");
		if (theme) {
			darkMode.set(theme === "dark");
		} else {
			const prefersDark = window.matchMedia(
				"(prefers-color-scheme: dark)",
			);
			if (prefersDark) {
				darkMode.set(true);
			} else {
				darkMode.set(false);
			}
		}
	});
</script>

<header
	class="z-10 flex h-[var(--header-height)] items-center justify-between px-4 lg:h-[var(--header-height-lg)] lg:justify-start"
>
	<div class="flex w-full items-center justify-between">
		<Logo />
		<nav class="flex h-full items-center">
			<ul class="text-md mr-4 hidden h-full items-center gap-2 lg:flex">
				{#each navLinks as link}
					<li>
						<a class="p-2" href={link.link} aria-label={link.title}>
							{link.title}
						</a>
					</li>
				{/each}
			</ul>

			<button onclick={toggleTheme} class="icon mr-4" id="theme-toggle">
				{#if darkMode.darkMode}
					<Sun />
				{:else}
					<Moon />
				{/if}
			</button>

			<a href="/login"><User class="hidden lg:block" /></a>
		</nav>
	</div>

	<button class="block lg:hidden" onclick={toggleNav} aria-label="menu">
		<Menu size={32} />
	</button>
</header>

<!-- Side navbar -->
{#if navOpen}
	<div
		class="absolute inset-0 z-[100] bg-surface-400/50 backdrop-blur-sm"
		use:click={closeNav}
	></div>
{/if}
<div
	class="absolute bottom-0 {navOpen
		? 'left-0'
		: 'left-[-400px]'} bg-surface-200-700-token top-0 z-[100] w-full max-w-[400px] transition-all duration-300"
>
	<button
		class="absolute right-2 top-2"
		onclick={closeNav}
		aria-label="close-nav"
	>
		<X />
	</button>
	<ul
		id="side-nav"
		class="flex h-full w-full flex-col items-center justify-center gap-4 text-xl"
	>
		{#each navLinks as link}
			<li>
				<a href={link.link} use:click={closeNav}>
					{link.title}
				</a>
			</li>
		{/each}
	</ul>
</div>
<!-- Side navbar -->
