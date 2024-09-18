<script lang="ts">
import { Card, Input, Heading, Button, A, Helper } from 'flowbite-svelte';
import { EyeOutline, EyeSlashOutline } from 'flowbite-svelte-icons';

import api from '../lib/utils/api';

let signIn: boolean = true

let firstname: string
let lastname: string

let username: string
let email: string

let password: string
let passwordConfirm: string

let showPassword: boolean = false

let singupError: string | null = null

$: passwordMatches = () => {
    if (password !== passwordConfirm) {
        return 'Passwords do not match';
    }

    return null;
}

async function login() {
    console.log('login', username, password);
    api.get('/hello').then((res: any) => {
        console.log(res);
    });
}

async function signup() {
    console.log('signup', firstname, lastname, username, email, password, passwordConfirm);
    singupError = 'username already exists';
}

async function logout() {
    console.log('logout');
    // pb.authStore.clear();
}

</script>


<Card class="m-auto">
    {#if signIn}
        <!-- Sign in form -->
        <form on:submit|preventDefault>
            <Heading class="mb-6" tag="h3">Sign in</Heading>
            <Input id="username" class="mb-4" type="text" placeholder="Username" required bind:value={username} />
            <div class="mb-6">
                <Input id="show-password" type={showPassword ? 'text' : 'password'} placeholder="Password" >
                    <button slot="right" on:click={() => (showPassword = !showPassword)} class="pointer-events-auto">
                        {#if showPassword}
                            <EyeOutline class="w-6 h-6" />
                        {:else}
                            <EyeSlashOutline class="w-6 h-6" />
                        {/if}
                    </button>
                </Input>
            </div>
            <Button class="w-full" type="submit" on:click={login}>Sign in</Button>
        </form>
        <Helper class="mt-6 text-left">Don't have an account? <A on:click={() => signIn = false}>Sign up</A></Helper>
    {:else}
        <!-- Sign up form -->
        <form on:submit|preventDefault>
            <Heading class="mb-6" tag="h3">Sign up</Heading>
            <div class="grid gap-2 mb-4 md:grid-cols-2">
                <Input id="firstname" type="text" placeholder="Firstname" require bind:value={firstname} />
                <Input id="lastname" type="text" placeholder="Lastname" require bind:value={lastname} />
            </div>
            <Input id="signup-username" class="mb-4" type="text" placeholder="Username" required bind:value={username} />
            <Input id="signup-email" class="mb-4" type="email" placeholder="Email" required bind:value={email} />
            <Input id="signup-password" class="mb-4" type='password' placeholder="Your password here" require bind:value={password} />
            <div class="mb-4">
                <Input id="signup-confirm" type='password' placeholder="Confirm your password" require bind:value={passwordConfirm} />
                {#if passwordMatches()}
                    <Helper color="red">Passoword does not match</Helper>
                {/if}
            </div>
            <Button class="w-full mb-2" type="submit" on:click={signup}>Sign up</Button>
            {#if singupError}
                <Helper class="h-0" color="red">{singupError}</Helper>
            {/if}
        </form>
        <Helper class="mt-6 text-left">Already have an account? <A on:click={() => signIn = true}>Sign up</A></Helper>
    {/if}
</Card>
