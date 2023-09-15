<script lang="ts">
	import InputForm from "$lib/components/inputForm.svelte";
	import { ListBox, ListBoxItem } from "@skeletonlabs/skeleton";

    export let characteristics: string;

	let displayInterface = {
		Size: 0,
		Resolution: {
			Height: 0,
			Width: 0
		},
		Type: '',
		GSync: false
  	};

	let monitorInterface = {
		Voltage: 0,
		Display: displayInterface,
		GSyncPremium: false,
		Curved: false
	}

    let valueSingle: any = displayInterface;

</script>

<div class="h-full">
	<ListBox class="bg-surface-900 border-neutral-900 shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline">
		<ListBoxItem bind:group={valueSingle} name="medium" value={displayInterface}>Display</ListBoxItem>
		<ListBoxItem bind:group={valueSingle} name="medium" value={monitorInterface}>Monitor</ListBoxItem>
	</ListBox>

	<InputForm bind:characteristics={characteristics} bind:interfaceObj={valueSingle}></InputForm>
</div>