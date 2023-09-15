<script lang="ts">
	import { drawerStore } from '@skeletonlabs/skeleton';
	function drawerClose(): void {
		drawerStore.close();
	}
</script>

<nav class="list-nav p-4 h-full">
	<ul>
		<li><a href="/" on:click={drawerClose}>Homepage</a></li>
	</ul>
</nav>