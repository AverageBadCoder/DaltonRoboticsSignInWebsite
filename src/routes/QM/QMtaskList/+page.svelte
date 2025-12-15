<script>
    import { goto } from '$app/navigation';
    import { derived } from 'svelte/store';
    let todos = [];
    let newTodo = '';
    let newDesc = '';

    // filter: 'all' | 'active' | 'completed'
    let filter = 'all';

    function handleRedirect() {
        goto('/DaltonRoboticsSignInWebsite/QM/QMhome');
    }

    function addTodo() {
        if (newTodo.trim()) {
            todos = [
                ...todos,
                { text: newTodo.trim(), description: newDesc.trim(), checked: false, important: false }
            ];
            newTodo = '';
            newDesc = '';
        }
    }

    function toggleChecked(index) {
        todos = todos.map((todo, i) =>
            i === index ? { ...todo, checked: !todo.checked } : todo
        );
    }

    function toggleImportant(index) {
        todos = todos.map((todo, i) =>
            i === index ? { ...todo, important: !todo.important } : todo
        );
    }

    function removeTodo(index) {
        todos = todos.filter((_, i) => i !== index);
    }

    // derived visible list based on filter
    $: visibleTodos = todos
        .filter(t => {
            if (filter === 'active') return !t.checked;
            if (filter === 'completed') return t.checked;
            return true; // all
        })
        // move important tasks to top
        .slice()
        .sort((a, b) => {
            const ai = a.important ? 1 : 0;
            const bi = b.important ? 1 : 0;
            return bi - ai; // important first
        });

    function setFilter(f) {
        filter = f;
    }
</script>

<main>
    <button on:click={handleRedirect}>Back</button>
    <h1>task list</h1>

    <div class="todo-container">
        <input
            type="text"
            bind:value={newTodo}
            placeholder="Task title..."
            on:keydown={(e) => e.key === 'Enter' && addTodo()}
        />
        <!-- svelte-ignore element_invalid_self_closing_tag -->
        <textarea
            bind:value={newDesc}
            placeholder="Description (optional)"
            rows="2"
        />
        <div class="controls">
            <button on:click={addTodo}>Add</button>

            <!-- filter controls -->
            <div class="filter">
                <button class:active={filter === 'all'} on:click={() => setFilter('all')}>All</button>
                <button class:active={filter === 'active'} on:click={() => setFilter('active')}>Active</button>
                <button class:active={filter === 'completed'} on:click={() => setFilter('completed')}>Completed</button>
            </div>
        </div>

        <ul>
            {#each visibleTodos as todo, i}
                <li class:important={todo.important}>
                    <input
                        type="checkbox"
                        checked={todo.checked}
                        on:change={() => toggleChecked(todos.indexOf(todo))}
                    />
                    <div class="task-content">
                      <span class:checked={todo.checked} class="task-title">{todo.text}</span>
                      {#if todo.description}
                        <div class="task-desc">{todo.description}</div>
                      {/if}
                    </div>

                    <!-- important toggle button -->
                    <button
                        type="button"
                        class="important-btn"
                        aria-label={todo.important ? 'Unmark important' : 'Mark important'}
                        title={todo.important ? 'Unmark important' : 'Mark important'}
                        on:click={() => toggleImportant(todos.indexOf(todo))}
                    >⭐</button>

                    <button class="remove-btn" on:click={() => removeTodo(todos.indexOf(todo))}>🗑</button>
                </li>
            {/each}
        </ul>
    </div>
</main>

<style>
@font-face {
  font-family: "Anta";
  src: url("/fonts/Anta-Regular.ttf") format("truetype");
  font-weight: 100 900;
  font-style: normal;
  font-display: swap;
}

/* force everywhere including form controls and pseudo elements */
html, body, #svelte, *, *::before, *::after,
input, button, textarea, select, option,
label, a, h1, h2, h3, h4, p, span, div {
  font-family: "Anta", system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif !important;
}

/* placeholder rules */
::placeholder { font-family: "Anta", sans-serif !important; }
:-ms-input-placeholder { font-family: "Anta", sans-serif !important; }
::-ms-input-placeholder { font-family: "Anta", sans-serif !important; }

    main {
        background: #0048ff;
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
    h1 {
        overflow: hidden;
        color: #ffffff;
        text-align: center;
    }
    .todo-container {
        max-width: 600px;
        margin: 2rem auto;
        background: #fff;
        padding: 1rem;
        border-radius: 12px;
        box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    }
    .controls {
        display: flex;
        align-items: center;
        gap: 1rem;
        margin-bottom: 0.5rem;
    }
    .filter button {
        padding: 0.25rem 0.6rem;
        border-radius: 6px;
        background: transparent;
        border: 1px solid #ddd;
    }
    .filter button.active {
        background: #003989;
        color: #fff;
        border-color: #003989;
    }
    input[type="text"], textarea {
        padding: 0.5rem;
        font-size: 1rem;
        border-radius: 8px;
        border: 1px solid #ccc;
        width: 100%;
        margin-bottom: 0.5rem;
        box-sizing: border-box;
    }
    textarea {
        resize: vertical;
    }
    ul {
        list-style: none;
        padding: 0;
        margin-top: 1rem;
    }
    li {
        display: flex;
        align-items: center;
        justify-content: flex-start;
        gap: 0.5rem;
        padding: 0.5rem 0;
        border-bottom: 1px solid #eee;
    }

    /* highlight important tasks with red left border and light red background */
    li.important {
        background: #fff5f5;
        border-left: 4px solid #e11d48;
    }

    /* allow flex child to shrink so text can wrap instead of overflowing */
    .task-content {
        display: flex;
        flex-direction: column;
        flex: 1;
        min-width: 0;
    }
    /* enable wrapping and word breaking for long titles/descriptions */
    .task-title {
        font-weight: 600;
        white-space: normal;
        word-break: break-word;
        overflow-wrap: anywhere;
    }

    /* make title red when important */
    li.important .task-title {
        color: #b91c1c;
    }

    .task-desc {
        font-size: 0.9rem;
        color: #666;
        margin-top: 0.25rem;
        white-space: normal;
        word-break: break-word;
        overflow-wrap: anywhere;
    }
    .checked {
        text-decoration: line-through;
        color: #aaa;
    }
    .remove-btn {
        display: inline-flex;
        align-items: center;
        justify-content: center;
        padding: 0;
        background: #ffffff;
        color: #000000;
        border: none;
        border-radius: 50%;
        width: 2rem;
        height: 2rem;
        font-size: 1rem;
        line-height: 1;
        cursor: pointer;
        margin-left: 1rem;
    }

    .important-btn {
      /* override global button styles so it is always visible */
      padding: 0;
      width: 2rem;
      height: 2rem;
      display: inline-flex;
      align-items: center;
      justify-content: center;
      background: transparent;
      border: none;
      font-size: 1.1rem;
      line-height: 1;
      color: #e11d48;
      cursor: pointer;
      margin-left: 0.25rem;
    }
    .important-btn[title="Unmark important"] { transform: scale(1.05); }

    /* ...existing styles... */
</style>

<slot />