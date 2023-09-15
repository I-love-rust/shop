<script lang="ts">
    import { env } from "$env/dynamic/public";
  
    export let markdown_text = '# Hello word!';


    import { onMount } from 'svelte';
    import { uploadImage } from "$lib/api";

    onMount(() => {
        document.addEventListener('paste', handlePaste);
    });

    async function handlePaste(event: ClipboardEvent) {
        const items = event.clipboardData?.items;

        if (items){
            for (let i = 0; i < items.length; i++) {
                const item = items[i];

                if (item.type.indexOf('image') !== -1) {
                    const file = item.getAsFile();
                    if (file) {
                        const reader = new FileReader();
                        reader.onload = async (e) => {
                            const imageUrl = e.target?.result;

                            if (imageUrl){
                                let link = await uploadImage({
                                    image: imageUrl?.toString()
                                })
                                markdown_text += `![Image](${env.PUBLIC_RESTAPI_URL}/${link.path})`
                            }
                            
                        };
                        reader.readAsDataURL(file);
                    }
                }
            }
        }
    }
</script>

<div class="flex h-full">
    <textarea bind:value={markdown_text} class="textarea" style="height: inherit;" rows="20" placeholder="Lorem ipsum dolor sit amet consectetur adipisicing elit." />
</div>