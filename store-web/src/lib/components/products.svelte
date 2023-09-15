<script lang="ts">
	import { ProgressRadial } from '@skeletonlabs/skeleton';
	import Card from "$lib/components/card.svelte"
	import type { ProductDto } from '$lib/dto/product';

    export let products: ProductDto[]
</script>

<div class="w-full text-token grid grid-cols-1 md:grid-cols-3 gap-4">
    {#if products}
        {#each products as product}
            <Card bind:product={product} />
        {/each}
    {:else}
        <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
        <ProgressRadial stroke={100} meter="stroke-primary-500" track="stroke-primary-500/30" />
        </div>
    {/if}
</div>