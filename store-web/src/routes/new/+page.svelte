<script lang="ts">
    import { goto } from '$app/navigation';
    import { checkStatus } from '$lib/api';
    import { onMount } from 'svelte';

	import Form from './form.svelte';

    let isLoggedIn = false
    onMount(async () => {
		let result = await checkStatus().catch(_ =>
            goto('/login')
        )
        if (result) {isLoggedIn = result.status}
	});
</script>


{#if isLoggedIn}
<Form />
{/if}