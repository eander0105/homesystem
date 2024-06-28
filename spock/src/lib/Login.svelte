<script lang="ts">
import { user, pb } from './pocketbase'

let username: string
let password: string

async function login() {
    console.log('login', username, password);
    pb.collection('users').authWithPassword(username, password);
}
async function logout() {
    console.log('logout');
    pb.authStore.clear();
}

</script>

{#if $user}
    <span>User logged in: {$user.username}</span>
    <br>
    <button on:click={logout}>Logout</button>   
{:else}
    <form on:submit|preventDefault>
        <input
            type="text"
            bind:value={username}
            placeholder="Username"
        />

        <input
            type="password"
            placeholder="Password"
            bind:value={password}
        />

        <button on:click={login}>Login</button>
    </form>
{/if}