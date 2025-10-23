<script>

import { goto } from '$app/navigation';

    function handleRedirect() {
        goto('/DaltonRoboticsSignInWebsite/QM/QMhome');
    }

</script>

<main>

    <button on:click={handleRedirect}>Back</button>
    <button on:click={null}>Upload</button>
   <h1>photos (upload button doesnt work)</h1>
   <h2>imagine a photo here</h2>
   <h3>imagine a photo here</h3>
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
    h2 {
        overflow: hidden;
        color: #ffffff;
        text-align: center;
    }
    h3 {
        overflow: hidden;
        color: #ffffff;
        text-align: center;
    }
</style>