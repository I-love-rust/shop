import { toastStore } from '@skeletonlabs/skeleton';

export let successNotify = (msg: string): void => {
    toastStore.trigger({ 
        message: msg,
        background: 'variant-filled-success', 
    })
};