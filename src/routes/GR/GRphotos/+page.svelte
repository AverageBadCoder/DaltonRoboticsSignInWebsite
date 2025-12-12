<script>
import { goto } from '$app/navigation';
import { base } from '$app/paths';

function handleRedirect() {
  goto('/DaltonRoboticsSignInWebsite/GR/GRhome');
}
</script>

<main>
  <button on:click={handleRedirect}>Back</button>
  <button on:click={null}>Upload</button>
  <h1>photos </h1>

  <div class="image-container">
    <img src="{base}/shivvy.png" alt="Shivvy" />
    <img src="{base}/ATlogo.png" alt="Atomic Theory logo" />
    <img src="{base}/granticon.png" alt="Grant" />
  </div>
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
</style>