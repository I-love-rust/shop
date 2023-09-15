<script lang="ts">
	import { hexToRGB, textToHex } from "$lib/utils";

    export let text: string;
    let color = textToHex(text)
</script>

<style>
    .tag-btn:hover {
        background-color: var(--theme-alpha-color);
    }

    .tag-btn {
        color: white;
        text-decoration: none;
    }

    .tag-btn > span {
        color: var(--theme-color);
    }
</style>

<a style="--theme-color: {color}; --theme-alpha-color: {hexToRGB(color, 0.1)}" class="btn tag-btn py-1 px-1 hover:border-transparent" href="/search?tags[]={text}" >
    <span>#</span>{text}
</a>