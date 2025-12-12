<script>
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  /**
   * @type {{ name: any; email: any; } | null}
   */
  let user = null;

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
  function handleRedirect() {
        goto('/DaltonRoboticsSignInWebsite/AT/AThome');
    }
    function handleRedirect1() {
        goto('/DaltonRoboticsSignInWebsite/UP/UPhome');
    }
    function handleRedirect2() {
        goto('/DaltonRoboticsSignInWebsite/QM/QMhome');
    }
    function handleRedirect3() {
        goto('/DaltonRoboticsSignInWebsite/SM/SMhome');
    }
    function handleRedirect4() {
        goto('/DaltonRoboticsSignInWebsite/NS/NShome');
    }
    function handleRedirect5() {
        goto('/DaltonRoboticsSignInWebsite/GR/GRhome');
    }
    function handleRedirect6() {
        goto('/DaltonRoboticsSignInWebsite/EH/EHhome');
    }
    import { base } from '$app/paths';
</script>

<main>

    <button on:click={handleRedirect}>4174</button>
    <button on:click={handleRedirect2}>6051</button>
    <button on:click={handleRedirect5}>9371</button>
    <button on:click={handleRedirect3}>9372</button>
    <button on:click={handleRedirect1}>11453</button>
    <button on:click={handleRedirect4}>17126</button>
    <button on:click={handleRedirect6}>10229</button>
   <h1>ADMIN HOME PAGE</h1>
   <h2>hi ms.screen</h2>
   <h5>website made by jonah madover 4174</h5>
   <div class="image-container">
        <img src="{base}/shivvy.png" alt="shiv pic" />
    </div>
    <button on:click={login}>Sign in with Google</button>
    {#if user}
  <p>{user.name} ({user.email})</p>
{/if}
</main>


<style>
  :global(body) {
        background-color: #ffffff;
        margin: 0;
        min-height: 100vh;
    }
    main {
        background-color: transparent;
        min-height: 100vh;
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
    .image-container {
        display: flex;
        justify-content: center;
        align-items: center;
        margin-top: 2rem;
    }
    .image-container img {
        max-width: 400px;
        width: 100%;
        height: auto;
        border-radius: 50%; /* Make the border a circle */
        box-shadow: 0 4px 16px rgba(0,0,0,0.2);
        background: #fff;
        padding: 1rem;
        object-fit: cover; /* Ensures the image covers the circle */
        aspect-ratio: 1 / 1; /* Keeps the image square for the circle */
    }
</style>


