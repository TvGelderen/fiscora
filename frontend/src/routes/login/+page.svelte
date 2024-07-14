<script lang="ts">
    import { PUBLIC_API_ENDPOINT } from '$env/static/public';
</script>

<h2>Login</h2>

<form action="{PUBLIC_API_ENDPOINT}/api/auth?provider=google" method="get">
    <button type="submit">Sign in with Google</button>
</form>

<a href="{PUBLIC_API_ENDPOINT}/api/auth?provider=google" >Google</a>

<p>Don't have an account? <a href="/register">Register here.</a></p>
