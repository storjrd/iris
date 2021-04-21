<template>
    <div>
        <h5 class="mb-4">Login</h5>

        <label for="emailAddress">Email Address</label>
        <input
            type="email"
            class="form-control input email"
            placeholder="example@email.com"
            v-on:keyup.enter="signUp"
            id="emailAddress"
        />

        <label for="password">Password</label>
        <input
            type="password"
            class="form-control input password"
            placeholder="••••••••••••"
            id="password"
        />

        <button
            class="btn btn-primary button signup-btn btn-block"
        >
            Login
        </button>

        <p>
            Don't have an account?
            
            <router-link to="/">Sign Up</router-link>
        </p>
    </div>
</template>