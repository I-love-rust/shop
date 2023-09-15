<script lang="ts">
	import { env } from '$env/dynamic/public';
	import { createProduct, uploadImage } from '$lib/api/index';
    import Markdown from './markdown.svelte'
    import { goto } from '$app/navigation';

    import { Stepper, Step } from '@skeletonlabs/skeleton';
    import { InputChip } from '@skeletonlabs/skeleton';
    import { FileDropzone } from '@skeletonlabs/skeleton';
	import Characteristics from './characteristics.svelte';
    
    let tags: string[] = [];
    let name: string;
    let description: string;
    let image: string;
    let price: number;
    let characteristics: string;
    
    let files: FileList;
    
    function onChangeHandler(e: Event): void {
        let file = files.item(0)
    
        if (file) {
          const reader = new FileReader();
          reader.onload = (e) => {
            const imageUrl = e.target?.result;
        
            const newImage = document.getElementById('imageContainer');
            if (imageUrl && newImage) {
                image = imageUrl.toString()
            }
          };
          reader.readAsDataURL(file);
        }
    }
    
    async function onCompleteHandler(e: Event): Promise<void> {
        let link = ''
        if (image) {
            const previewLink = await uploadImage({
                image: image
            })
            link = `${env.PUBLIC_RESTAPI_URL}/${previewLink.path}`
        }
        let path = await createProduct({
            name: name,
            description: description,
            tags: tags,
            preview: link,
            price: price,
            characteristics: characteristics,
        })
        goto(`/product/${path.path}`)
    }
</script>

<div class="flex w-full my-5">
    <div class="max-w-screen-xl w-full mx-auto p-6 bg-surface-800 shadow-md rounded">
        <Stepper on:complete={onCompleteHandler}>
        	<Step>
        		<svelte:fragment slot="header">Product info</svelte:fragment>
                <div class="space-y-4">
        		        <input class="bg-surface-700 border-surface-500 shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
                            bind:value={name} id="title" type="text" placeholder="Enter product name">

                        <input class="bg-surface-700 border-surface-500 shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
                            bind:value={price} id="title" type="number" placeholder="Enter product price">

                        <InputChip bind:value={tags} name="tags" placeholder="Enter tags..." />
                        <FileDropzone name="preview" accept="image/*" bind:files on:change={onChangeHandler}>
							<svelte:fragment slot="lead"><i class="fa-solid fa-file-arrow-up text-4xl" /></svelte:fragment>
							<svelte:fragment slot="meta">PNG, JPG, and GIF allowed.</svelte:fragment>
						</FileDropzone>
                        <div class="flex justify-center">
                            <img class="max-w-xs max-h-64" id="imageContainer" alt=""/>
                        </div>
                </div>
                {#if image}
                    <img alt="preview" src={image} class="w-full object-cover aspect-[21/9]"/>
                {/if}
        	</Step>
            <Step>
        		<svelte:fragment slot="header">Product characteristics</svelte:fragment>
                <Characteristics bind:characteristics={characteristics}/>
        	</Step>
        	<Step>
        		<svelte:fragment slot="header">Product description</svelte:fragment>
                <Markdown bind:markdown_text={description} />
        	</Step>
        </Stepper>
    </div>
</div>