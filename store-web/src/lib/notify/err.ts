import { toastStore } from '@skeletonlabs/skeleton';

export let errNotify = (msg: string): void => {
    toastStore.trigger({ 
        message: msg,
        background: 'variant-filled-error', 
    })
};