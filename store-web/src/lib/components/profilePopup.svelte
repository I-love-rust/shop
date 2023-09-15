<script lang="ts">
	import { Avatar } from "@skeletonlabs/skeleton";
    import type { UserPublicDto } from "$lib/dto/index";


    export let user: UserPublicDto;
    
    export let dataPopup: string;
</script>

<div class="card p-4 w-72 shadow-xl z-50" data-popup={dataPopup}>
    <div class="space-y-4">
        <Avatar src={user.avatar} initials="{user.username}" width="w-16" />
        <div>
            <p class="font-bold">{user.name}</p>
            <p class="opacity-50">@{user.username}</p>
        </div>
        <p>{user.bio}</p>
        <a class="btn variant-soft w-full" href="/@{user.username}" target="_blank" rel="noreferrer"> View profile </a>
    </div>
    <div class="arrow bg-surface-100-800-token" />
</div>