<script>

import { goto } from '$app/navigation';

    function handleRedirect() {
        goto('/DaltonRoboticsSignInWebsite/UP/UPhome');
    }

</script>

<main>

    <button on:click={handleRedirect}>Back</button>
   <h1>build sessions</h1>
</main>

<style>
    button {
        margin-top: .5rem;
        margin-left: .5rem;
        margin-right: .5rem;
        padding: 0.5rem 1.5rem;
        font-size: 1rem;
        cursor: pointer;
         border-radius: 999px;
    }
    h1 {
        overflow: hidden;
        color: #ffffff;
        text-align: center;
    }
</style>