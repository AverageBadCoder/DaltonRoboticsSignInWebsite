<script>
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';

  const STORAGE_KEY = 'QMbuild_sessions_v1';

  let attendeesE  = []; // [{ name }]
  let attendeesL = []; // [{ name }]
  let buildCounts = {}; // { name: number }
  let signupName = '';
  let early = true;
  let logName = ''; // manual log target
  let selectedAttendee = '';

  // allow fractional builds (0.5)
  let logAmount = 1;

  function handleRedirect() {
    goto('/DaltonRoboticsSignInWebsite/QM/QMhome');
  }

  function persist() {
    const payload = { attendeesE, attendeesL, buildCounts };
    try {
      localStorage.setItem(STORAGE_KEY, JSON.stringify(payload));
    } catch (e) {
      console.error('persist error', e);
    }
  }

  onMount(() => {
    try {
      const raw = localStorage.getItem(STORAGE_KEY);
      if (raw) {
        const parsed = JSON.parse(raw);
        attendeesE = parsed.attendeesE || [];
        attendeesL = parsed.attendeesL || [];
        buildCounts = parsed.buildCounts || {};
        // ensure numeric values (could be strings after JSON)
        Object.keys(buildCounts).forEach(k => {
          buildCounts[k] = Number(buildCounts[k]) || 0;
        });
      }
    } catch (e) {
      console.error('load error', e);
    }
  });

  function signUpE() {
    const name = (signupName || '').trim();
    if (!name) return;
    if (!attendeesE.find(a => a.name.toLowerCase() === name.toLowerCase())) {
      attendeesE = [...attendeesE, { name }];
      if (!buildCounts[name]) buildCounts[name] = 0;
      persist();
    }
    signupName = '';
  }

  function signUpL() {
    const name = (signupName || '').trim();
    if (!name) return;
    if (!attendeesL.find(a => a.name.toLowerCase() === name.toLowerCase())) {
      attendeesL = [...attendeesL, { name }];
      if (!buildCounts[name]) buildCounts[name] = 0;
      persist();
    }
    signupName = '';
  }

  function signUp() {
    if (early) {
      signUpE();
    } else {
      signUpL();
    }
  }

  function removeAttendeeE(i) {
    attendeesE = attendeesE.filter((_, idx) => idx !== i);
    persist();
  }

  function removeAttendeeL(i) {
    attendeesL = attendeesL.filter((_, idx) => idx !== i);
    persist();
  }

  // allow specifying fractional amounts (e.g. 0.5)
  function logBuildFor(name, amount = 1) {
    if (!name) return;
    const key = name.trim();
    if (!key) return;
    const amt = Number(amount) || 0;
    if (amt <= 0) return;

    // update count but do NOT auto-add if not attending
    buildCounts = { ...buildCounts, [key]: (Number(buildCounts[key]) || 0) + amt };

    // remove the person from any attending lists if they were signed up
    attendeesE = attendeesE.filter(a => a.name.toLowerCase() !== key.toLowerCase());
    attendeesL = attendeesL.filter(a => a.name.toLowerCase() !== key.toLowerCase());

    persist();
    logName = '';
    selectedAttendee = '';
    logAmount = 1;
  }

  function handleLogClick() {
    const target = selectedAttendee || logName.trim();
    if (!target) return;
    logBuildFor(target, logAmount);
  }

  // combined unique attendees for select options
  $: attendeesAll = Array.from(
    new Map(
      [...attendeesE, ...attendeesL].map(a => [a.name.toLowerCase(), a.name])
    ).values()
  );

  $: leaderboard = Object.entries(buildCounts)
    .map(([name, count]) => ({ name, count: Number(count) }))
    .sort((a, b) => b.count - a.count || a.name.localeCompare(b.name));
</script>

<main>
  <div class="top-row">
    <div class="left-controls">
      <div class="signup">
        <button on:click={handleRedirect}>Back</button>
        <input
          placeholder="Your name to attend"
          bind:value={signupName}
          on:keydown={(e) => e.key === 'Enter' && signUp()}
        />

        <label class="slot-label">
          <input
            type="checkbox"
            bind:checked={early}
          />
          <span>Early</span>
        </label>

        <button on:click={signUp}>Sign up for next build</button>
      </div>

      <div class="log-controls">
        <select bind:value={selectedAttendee}>
          <option value=''>— select attendee —</option>
          {#each attendeesAll as name}
            <option value={name}>{name}</option>
          {/each}
        </select>

        <input
          placeholder="Or enter name to log build"
          bind:value={logName}
          on:keydown={(e) => e.key === 'Enter' && handleLogClick()}
        />

        <input
          type="number"
          step="0.5"
          min="0.5"
          bind:value={logAmount}
          style="width:6rem"
          title="Number of builds (0.5 allowed)"
        />

        <button on:click={handleLogClick}>Log Build</button>
      </div>
    </div>
  </div>

  <!-- change: wrap early+late in left column and make grid two columns so late appears below early -->
  <section class="lists">
    <div class="left-column">
      <div class="attendance">
        <h2>Attending early build</h2>
        {#if attendeesE.length === 0}
          <p class="muted">No one signed up yet</p>
        {:else}
          <ul>
            {#each attendeesE as a, i}
              <li>
                <span>{a.name}</span>
                <button class="small remove" on:click={() => removeAttendeeE(i)} aria-label="remove">✖</button>
              </li>
            {/each}
          </ul>
        {/if}
      </div>
  <!-- used ai (model: gpt 5 mini, date: 12/9) to stack early and late and keep font sizes same cause it was REALLY annoying me and i couldn't find solution -->
      <div class="attendance attendance-late">
        <h2>Attending late build</h2>
        {#if attendeesL.length === 0}
          <p class="muted">No one signed up yet</p>
        {:else}
          <ul>
            {#each attendeesL as a, i}
              <li>
                <span>{a.name}</span>
                <button class="small remove" on:click={() => removeAttendeeL(i)} aria-label="remove">✖</button>
              </li>
            {/each}
          </ul>
        {/if}
      </div>
    </div>

    <div class="leaderboard">
      <h2>Leaderboard — most builds</h2>
      {#if leaderboard.length === 0}
        <p class="muted">No builds logged yet</p>
      {:else}
        <ol>
          {#each leaderboard as item}
            <li><strong>{item.name}</strong> <span class="count">{item.count % 1 === 0 ? item.count : item.count.toFixed(1)}</span></li>
          {/each}
        </ol>
      {/if}
    </div>
  </section>
</main>

<style>
  main {
    padding: 1rem;
    color: #fff;
    min-height: 100vh;
    background: linear-gradient(180deg,#021034,#00122a);
    box-sizing: border-box;
  }

  .top-row { margin-bottom: 1rem; }

  .left-controls {
    display: flex;
    gap: 1.25rem;
    flex-wrap: wrap;
    align-items: center;
  }

  .signup, .log-controls {
    display: flex;
    gap: .5rem;
    align-items: center;
  }

  .slot-label {
    display: inline-flex;
    align-items: center;
    gap: .35rem;
    font-size: 1rem;
    color: #fff;
  }

  input[type="text"], input[type="number"], select, input {
    padding: .5rem;
    border-radius: 8px;
    border: 1px solid #ccc;
    font-size: 1rem;
  }

  button {
    padding: .5rem .9rem;
    border-radius: 8px;
    border: none;
    background: #0b63ff;
    color: #fff;
    cursor: pointer;
    font-size: 1rem;
  }

  .lists {
    display: grid;
    grid-template-columns: 1fr 340px; /* left column (early+late stacked) | leaderboard */
    gap: 1.5rem;
    align-items: start; /* keep leaderboard top aligned with early */
  }

  /* left column stacks early then late */
  .left-column {
    display: flex;
    flex-direction: column;
    gap: 1.5rem; /* space between early and late boxes */
  }

  /* keep heading/font sizes unchanged (no modifications to font-size) */
  .attendance, .leaderboard {
    background: #fff;
    color: #111;
    padding: 1rem;
    border-radius: 10px;
  }

  /* same heading size for early/late */
  .attendance h2, .attendance h3, .leaderboard h2 {
    font-size: 1.1rem;
    margin: 0 0 .5rem 0;
  }

  .attendance ul, .leaderboard ol {
    margin: .5rem 0 0 0;
    padding-left: 1.2rem;
  }
  .attendance li, .leaderboard li {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: .3rem 0;
  }
  .muted { color: #666; }
  .count {
    background: #efefef;
    padding: .15rem .5rem;
    border-radius: 6px;
    font-weight: 600;
  }

  /* ensure leaderboard top lines up with early build top */
  .leaderboard { align-self: start; }

  @media (max-width: 900px) {
    .lists { grid-template-columns: 1fr; }
  }
</style>