<script>
    import { goto } from '$app/navigation';

    let todos = [];
    let newTodo = '';

    function handleRedirect() {
        goto('/AT/AThome');
    }

    function addTodo() {
        if (newTodo.trim()) {
            todos = [...todos, { text: newTodo.trim(), checked: false }];
            newTodo = '';
        }
    }

    function toggleChecked(index) {
        todos = todos.map((todo, i) =>
            i === index ? { ...todo, checked: !todo.checked } : todo
        );
    }

    function removeTodo(index) {
        todos = todos.filter((_, i) => i !== index);
    }
</script>

<main>
    <button on:click={handleRedirect}>Back</button>
    <h1>task list</h1>

    <div class="todo-container">
        <input
            type="text"
            bind:value={newTodo}
            placeholder="Add a new task..."
            on:keydown={(e) => e.key === 'Enter' && addTodo()}
        />
        <button on:click={addTodo}>Add</button>
        <ul>
            {#each todos as todo, i}
                <li>
                    <input
                        type="checkbox"
                        checked={todo.checked}
                        on:change={() => toggleChecked(i)}
                    />
                    <span class:checked={todo.checked}>{todo.text}</span>
                    <button class="remove-btn" on:click={() => removeTodo(i)}>✕</button>
                </li>
            {/each}
        </ul>
    </div>
</main>

<style>
    main {
        background: #010a35;
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
        max-width: 400px;
        margin: 2rem auto;
        background: #fff;
        padding: 1rem;
        border-radius: 12px;
        box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    }
    input[type="text"] {
        padding: 0.5rem;
        font-size: 1rem;
        border-radius: 8px;
        border: 1px solid #ccc;
        width: 70%;
        margin-right: 0.5rem;
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
    .checked {
        text-decoration: line-through;
        color: #aaa;
    }
    .remove-btn {
        background: #ff0000;
        color: #fff;
        border: none;
        border-radius: 50%;
        width: 2rem;
        height: 2rem;
        font-size: 1rem;
        cursor: pointer;
        margin-left: 1rem;
    }
</style>