<script lang="ts">
	import { PUBLIC_API_URL } from "$env/static/public";
</script>

<div class="card mx-auto my-4 max-w-[480px] p-4 text-center">
	<h2 class="mb-6">Login</h2>
	<a class="btn w-full" href="{PUBLIC_API_URL}/auth/google">Google Login</a>
	<a class="secondary btn mt-4 w-full" href="{PUBLIC_API_URL}/auth/demo">
		Login as demo user
	</a>
</div>
