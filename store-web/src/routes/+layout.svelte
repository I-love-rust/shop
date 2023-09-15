<script lang='ts'>
	import '@skeletonlabs/skeleton/themes/theme-vintage.css';
	import '@skeletonlabs/skeleton/styles/all.css';
	import '../app.postcss';

	import { AppShell, AppBar, Toast } from '@skeletonlabs/skeleton';
	import { Drawer, drawerStore } from '@skeletonlabs/skeleton';

	import Navigation from '$lib/appbar/navigation.svelte';
	import Trail from '$lib/appbar/trail.svelte';
	import Lead from '$lib/appbar/lead.svelte';
	import Default from '$lib/appbar/default.svelte';

	import { computePosition, autoUpdate, offset, shift, flip, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });

	function drawerOpen(): void {
		drawerStore.open({});
	}
</script>

<Toast/>

<Drawer>
	<h2 class="p-4">Navigation</h2>
	<hr />
	<Navigation />
</Drawer>

<AppShell slotSidebarLeft="bg-surface-500/5 w-0 lg:w-64">
	<svelte:fragment slot="header">
		<AppBar gridColumns="grid-cols-3" slotTrail="place-content-end">
			<svelte:fragment slot="lead">
				<Lead { drawerOpen } />
			</svelte:fragment>

			<Default />
			
			<svelte:fragment slot="trail">
				<Trail />
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<slot />
</AppShell>