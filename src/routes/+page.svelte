<script>
  import { goto } from '$app/navigation';
  function handleRedirect() {
        goto('/AT/AThome');
    }
    function handleRedirect1() {
        goto('/UP/UPhome');
    }
    </script>

<main>

    <button on:click={handleRedirect}>4174</button>
    <button on:click={handleRedirect}>6051</button>
    <button on:click={handleRedirect}>9371</button>
    <button on:click={handleRedirect}>9372</button>
    <button on:click={handleRedirect1}>11453</button>
    <button on:click={handleRedirect}>17126</button>
    <button on:click={handleRedirect}>10229</button>
   <h1>ADMIN HOME PAGE</h1>
   <h2>hi ms.screen</h2>
   <h5>website made by jonah madover 4174</h5>
   <div class="image-container">
        <img src="/shivvy.png" alt="shiv pic" />
    </div>
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


