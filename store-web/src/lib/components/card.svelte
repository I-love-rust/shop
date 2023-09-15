<script lang="ts">
	import { Avatar } from "@skeletonlabs/skeleton";
    import type { ProductDto } from "$lib/dto/product";
    import Tag from '$lib/components/tag.svelte'

    export let product: ProductDto
</script>

<div class="card bg-initial card-hover overflow-hidden inline-flex h-44">
    <header class="flex items-center p-2">
        {#if product.preview != ''}
        <a class="card card-hover overflow-hidden" href="/product/{product.path}">
            <img src="{product.preview}" class="bg-black/50 h-40" alt="" />
        </a>
        {/if}
    </header>
    <div class="p-4 space-y-1">
        <a class="card card-hover overflow-hidden" href="/product/{product.path}">
            <h3 class="h3" data-toc-ignore>{product.name}</h3>
        </a>
        <!-- {#each product.tags as tag}
            <Tag bind:text={tag} />
        {/each} -->
        <p>{product.price}$</p>
    </div>
</div>