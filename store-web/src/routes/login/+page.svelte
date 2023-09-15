<script>
	import { goto } from '$app/navigation';
  import { signIn } from '$lib/api/index';
	import { errNotify } from '$lib/notify';
  import { setAccessToken, setRefreshToken } from "$lib/stores/store"

  let login = '', password = ''

  async function submit() {
    let result = await signIn({
      login: login,
      password: password
    }).catch(err => errNotify(err.description))

    if (result) {
      setAccessToken(result.access_token)
      setRefreshToken(result.refresh_token)
      goto("/")
    }
  }
</script>

<div class="flex items-center justify-center h-full">
    <div class="max-w-md w-full mx-auto p-6">
      <h2 class="text-3xl text-center mb-6">Sign in</h2>
      <form on:submit={submit} class="bg-surface-800 shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <div class="mb-4">
          <label class="block text-sm font-bold mb-2" for="login">
            Email or username
          </label>
          <input bind:value={login} class="bg-surface-900 border-neutral-900 shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
            id="login" type="text" placeholder="Enter your email or username">
        </div>
        <div class="mb-6">
          <label class="block text-sm font-bold mb-2" for="password">
            Password
          </label>
          <input bind:value={password} class="bg-surface-900 border-neutral-900 shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
            id="password" type="password" placeholder="Enter your password">
        </div>
        <div class="flex items-center justify-between">
          <button
            class="bg-primary-500 hover:bg-primary-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit">
            Sign In
          </button>
          <a class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800" href="#">
            Forgot Password?
          </a>
        </div>
      </form>
    </div>
</div>