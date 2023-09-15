<script lang="ts">
	import { checkStatus } from "$lib/api";
    import { onMount } from 'svelte';

    let isLoggedIn = false

    onMount(async () => {
		let result = await checkStatus()
        isLoggedIn = result.status
	});
</script>

{#if isLoggedIn}
    <a class="btn btn-sm" href="/">Home</a>
    <a class="btn btn-sm" href="/new">Add product</a>
    <a class="btn btn-sm" href="/about">About</a>
{:else}
    <a class="btn btn-sm" href="/login">Sign in</a>
    <a class="btn btn-sm" href="/register">Sign up</a>
{/if}