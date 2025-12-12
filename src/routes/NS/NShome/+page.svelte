<script>
    import { goto } from '$app/navigation';
      import { onMount } from 'svelte';
    import { base } from '$app/paths';
  /**
   * @type {{ name: any; email: any; } | null}
   */
  let user = null;

    function handleRedirect() {
        goto('/DaltonRoboticsSignInWebsite/NS/NSphotos');
    }
    function handleRedirect1() {
        goto('/DaltonRoboticsSignInWebsite/NS/NSbuildSessions');
    }
    function handleRedirect2() {
        goto('/DaltonRoboticsSignInWebsite/NS/NStaskList');
    }

  // use the ngrok URL so the static site hits your public backend
  const BACKEND_URL = 'https://unmordantly-stirruplike-naida.ngrok-free.dev';
  function login() {
    location.href = BACKEND_URL + '/auth/google/login';
  }

  onMount(async () => {
    // capture token returned in fragment: https://your-gh-pages/#token=...
    const hash = new URLSearchParams(location.hash.replace(/^#/, ''));
    if (hash.has('token')) {
      const token = hash.get('token');
      if (token !== null) {
        localStorage.setItem('token', token);
      }
      history.replaceState(null, '', location.pathname + location.search);
    }
    const token = localStorage.getItem('token');
    if (token) {
      // changed: use BACKEND_URL instead of hardcoded localhost
      const res = await fetch(BACKEND_URL + '/api/me', {
        headers: { Authorization: 'Bearer ' + token }
      });
      if (res.ok) user = await res.json();
    }
  });
</script>

<main>
    <div class="button-row">
        <button on:click={handleRedirect2}>Task List</button>
        <button on:click={handleRedirect1}>Build Sessions</button>
        <button on:click={handleRedirect}>Photos</button>
    </div>
    <h1>Welcome to the Natural Selection 17126 home page</h1>
    <div class="image-container">

        <img src="{base}/ATlogo.png" alt="Atomic Theory logo">
    </div>
    <button on:click={login}>Sign in with Google</button>
    {#if user}
  <p>{user.name} ({user.email})</p>
{/if}
</main>

<style>
    :global(body) {
        background-color: #124502;
        margin: 0;
        min-height: 100vh;
    }
    main {
        background-color: transparent;
        min-height: 100vh;
    }
    .button-row {
        display: flex;
        justify-content: flex-start;
        gap: 1rem;
        margin-bottom: 2rem;
        margin-top: 1rem;
    }
    h1 {
        overflow: hidden;
        color: #ffffff;
        text-align: center;
    }
    .image-container {
        display: flex;
        justify-content: center;
        align-items: center;
        margin-top: 2rem;
    }
    .image-container img {
        max-width: 600px;
        width: 100%;
        height: auto;
        border-radius: 50%; /* Make the border a circle */
        box-shadow: 0 4px 16px rgba(0,0,0,0.2);
        background: #fff;
        padding: 1rem;
        object-fit: cover; /* Ensures the image covers the circle */
        aspect-ratio: 1 / 1; /* Keeps the image square for the circle */
    }
    button {
        margin-top: .5rem;
        margin-left: .5rem;
        margin-right: .5rem;
        padding: 0.5rem 1.5rem;
        font-size: 1rem;
        cursor: pointer;
         border-radius: 999px;
    }
</style>
