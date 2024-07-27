<script lang="ts">
	import { PUBLIC_API_URL } from "$env/static/public";
</script>

<div class="card mx-auto my-4 max-w-[480px] p-4">
	<h2 class="mb-6">Login</h2>
	<form method="POST" action="{PUBLIC_API_URL}/auth/username"></form>
	<a class="btn w-full" href="{PUBLIC_API_URL}/auth/google">Google Login</a>
</div>
