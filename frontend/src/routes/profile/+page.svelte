<script lang="ts">
    import { page } from '$app/stores';

    const { user } = $page.data;
</script>

<h2>You are logged in as <span class="text-primary-500">{user.username}</span></h2>
