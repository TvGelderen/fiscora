<script lang="ts">
import X from "lucide-svelte/icons/x";
import Menu from "lucide-svelte/icons/menu";
import User from "lucide-svelte/icons/user";
import { onMount } from "svelte";

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
    link: "/dashboard",
    title: "Dashboard",
  },
];

let navOpen = $state(false);
const toggleNav = () => (navOpen = !navOpen);
const closeNav = () => (navOpen = false);

onMount(() => {
  window.addEventListener("resize", () => {
    if (window.innerWidth >= 1024) {
      closeNav();
    }
  });

  const nav = document.getElementById("side-nav");
  if (nav) {
    const links = nav.querySelectorAll("a");
    for (const link of links) {
      link.addEventListener("click", closeNav);
    }
  }
});
</script>

<header
  class="flex h-[64px] items-center justify-between bg-surface-800 px-4 lg:justify-start"
>
  <div class="w-full flex items-center gap-6">
    <h2>Budget Buddy</h2>
    <nav class="flex h-full items-center">
      <ul class="hidden h-full items-center gap-4 text-xl lg:flex">
        {#each navLinks as link}
          <li><a class="p-2" href={link.link}>{link.title}</a></li>
        {/each}
      </ul>

      <!-- Side navbar -->
      <div
        class="absolute bottom-0 {navOpen ? 'left-0' : 'left-[-400px]'} top-0 w-full max-w-[400px] bg-surface-700 transition-all duration-300"
      >
        <button class="absolute right-2 top-2" onclick={toggleNav}><X /></button>
        <ul
          id="side-nav"
          class="flex h-full w-full flex-col items-center justify-center gap-4 text-xl"
        >
          {#each navLinks as link}
            <li><a href={link.link}>{link.title}</a></li>
          {/each}
        </ul>
      </div>
      <!-- Side navbar -->
    </nav>
  </div>

  <User class="hidden lg:block" />

  <button class="block lg:hidden" onclick={toggleNav}><Menu size={32} /></button>
</header>
