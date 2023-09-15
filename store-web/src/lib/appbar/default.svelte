<script lang="ts">
    let searchtext = ""

    function search(event: { key: string; }) {
        if (event.key === 'Enter') {
            window.location.href = `/search?text=${searchtext}`;
        }
    }
</script>

<div>
    <input class="bg-surface-700 border-surface-500 shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
                id="title" bind:value={searchtext} on:keydown={search} type="search" placeholder="Search">
</div>