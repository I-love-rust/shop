<script lang="ts">
	import Products from "$lib/components/products.svelte"
	import { onMount } from "svelte";
    import type { ProductDto } from "$lib/dto/product";
    import { page } from '$app/stores';
	import { searchProduct } from "$lib/api/";
    
    const tags = $page.url.searchParams.get('tags[]')
    const author = $page.url.searchParams.get('author')
    const text = $page.url.searchParams.get('text')

    let lastID = 0;
	let products: ProductDto[]
	onMount(async () => {
		let result = await searchProduct({
            page: lastID,
            tags: tags,
            text: text,
            authorID: Number(author),
        })
		lastID++
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