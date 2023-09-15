<script lang="ts">
    import { marked } from 'marked'
	import SvelteMarkdown from 'svelte-markdown';
    import Tag from '$lib/components/tag.svelte'
	import type { ProductDto } from '$lib/dto';
	import { generateBill } from '$lib/api/';
	import { errNotify, successNotify } from '$lib/notify';
	import Characteristics from '../../new/characteristics.svelte';

    export let data;
    const product: ProductDto = data.data;

    async function buy() {
        let result = await generateBill({
            product: product.id,
        }).catch(err => errNotify(err.description))
        if (result) {
            successNotify("Successful purchase")
        }
    }

    function getObjectProperties(obj: any) {
        return Object.entries(obj);
    }
</script>

<div class="flex w-full my-5">
    <div class="flex max-w-screen-xl mx-auto p-6 bg-surface-800">
        <div class="h-54 space-y-5 flex">
            {#if product.preview != ''}
                <img class="rounded-lg object-cover h-96" src={product.preview} alt="" />
            {/if}
            <div class="p-5 space-y-3">
                <h1>{product.name}</h1>
                {#each product.tags as tag}
                    <Tag bind:text={tag} />
                {/each}
                <hr/>
                <table>
                    <tbody>
                        {#each getObjectProperties(JSON.parse(product.characteristics)) as [key, value]}
                            <tr>
                                <td>{key}</td>
                                <td>{value}</td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
                <hr/>
                <div>
                    <SvelteMarkdown source={marked.lexer(product.description)} />
                </div>
                <hr/>
                <div class="flex">
                    <h2>{product.price}$</h2>
                    <button on:click={buy}
                        class="bg-primary-500 hover:bg-primary-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                        type="submit">
                        Buy
                    </button>
                </div>
            </div>
        </div>
    </div>
</div>