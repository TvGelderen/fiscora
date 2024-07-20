<script lang="ts">
import { PUBLIC_API_ENDPOINT } from '$env/static/public';

async function getUserInfo() {
    const response = await fetch(`${PUBLIC_API_ENDPOINT}/api/auth/me`);
    console.log(response);
    const data = await response.json();
    console.log(data);
}
</script>

<h1>Budget Buddy</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>

<button class="btn mt-4" onclick={getUserInfo}>Get user data</button>
