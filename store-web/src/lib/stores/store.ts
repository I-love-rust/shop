import { writable } from 'svelte/store';

const accessTokenStore = writable('');

function setAccessToken(token: string) {
  accessTokenStore.set(token);
  localStorage.setItem('access_token', token);
}

function clearAccessToken() {
  localStorage.removeItem('access_token');
}



const refreshTokenStore = writable('');

function setRefreshToken(token: string) {
  accessTokenStore.set(token);
  localStorage.setItem('refresh_token', token);
}

function clearRefreshToken() {
  localStorage.removeItem('refresh_token');
}

export { accessTokenStore, setAccessToken, clearAccessToken, refreshTokenStore, setRefreshToken, clearRefreshToken};