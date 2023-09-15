<script lang="ts">
	import { getProductFeed } from "$lib/api/product";
	import { onMount } from "svelte";
	import Products from "$lib/components/products.svelte"
    import type { ProductDto } from "$lib/dto/product";

	let page = 0;
	let products: ProductDto[]
	onMount(async () => {
		let result = await getProductFeed({id:page})
		page++
		products=[]
		for (let product of result) {
			products.push(product)
		}
	});
</script>

<div class="flex w-full my-5">
    <div class="flex max-w-screen-xl w-full mx-auto p-6 items-center justify-center">
        <div class="w-full space-y-5">
			<Products bind:products={products} />
		</div>
	</div>
</div>