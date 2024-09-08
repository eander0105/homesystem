<script lang="ts">

import api from '../lib/utils/api';

let username: string
let password: string

async function login() {
    console.log('login', username, password);
    api.get('/hello').then((res: any) => {
        console.log(res);
    });
}
async function logout() {
    console.log('logout');
    // pb.authStore.clear();
}

</script>

{#if false}
    <span>User logged in: </span>
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