import { logout } from '$lib/stores/auth';
import type { PageLoad } from './$types';
import { goto } from '$app/navigation';

export const load: PageLoad = async() => {
    await logout();
    goto('/').then(() => {
        window.location.reload();
    });
};