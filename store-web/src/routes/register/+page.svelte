<script lang="ts">
  import { goto } from '$app/navigation';
  import { signIn, signUp } from '$lib/api/index';
	import { errNotify } from '$lib/notify';
	import { setAccessToken, setRefreshToken } from '$lib/stores/store';
  let username = '', email='', password = ''


  $: submit = async() => {
    let user = await signUp({
      username: username,
      email: email,
      password: password
    }).catch(err => errNotify(err.description))

    if (user) {
      let result = await signIn({
        login: username,
        password: password
      })

      setAccessToken(result.access_token)
      setRefreshToken(result.refresh_token)

      goto("/")
    }
  }
</script>

<div class="flex items-center justify-center h-full">
    <div class="max-w-md w-full mx-auto p-6">
      <h2 class="text-3xl text-center mb-6">Sign up</h2>
      <form on:submit={submit} class="bg-surface-800 shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <div class="mb-4">
            <label class="block text-sm font-bold mb-2" for="Email">
              Email
            </label>
            <input bind:value={email} class="bg-surface-900 border-neutral-900 shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
              id="Email" type="text" placeholder="Enter your Email">
        </div>
        <div class="mb-4">
          <label class="block text-sm font-bold mb-2" for="username">
            Username
          </label>
          <input bind:value={username} class="bg-surface-900 border-neutral-900 shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
            id="username" type="text" placeholder="Enter your username">
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