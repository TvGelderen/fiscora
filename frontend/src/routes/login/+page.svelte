<script lang="ts">
    import { PUBLIC_API_ENDPOINT } from '$env/static/public';
</script>

<h2>Login</h2>

<div class="my-4">
    <a href="{PUBLIC_API_ENDPOINT}/api/auth?provider=google"><h1>Google Login</h1></a>
</div>

<p>Don't have an account? <a href="/register">Register here.</a></p>
