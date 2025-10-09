<script>
    import { goto } from '$app/navigation';

    function handleRedirect() {
        goto('/DaltonRoboticsSignInWebsite/AT/ATphotos');
    }
    function handleRedirect1() {
        goto('/DaltonRoboticsSignInWebsite/AT/ATbuildSessions');
    }
    function handleRedirect2() {
        goto('/DaltonRoboticsSignInWebsite/AT/ATtaskList');
    }
    import { base } from '$app/paths';
</script>

<main>
    <div class="button-row">
        <button on:click={handleRedirect2}>Task List</button>
        <button on:click={handleRedirect1}>Build Sessions</button>
        <button on:click={handleRedirect}>Photos</button>
    </div>
    <h1>Welcome to the Atomic Theory 4174 home page</h1>
    <div class="image-container">

        <img src="{base}/ATlogo.png" alt="Atomic Theory logo">
    </div>
</main>

<style>
    :global(body) {
        background-color: #010a35;
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
