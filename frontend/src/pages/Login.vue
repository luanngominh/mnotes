<template>
   <div>
           <div>
               <input type="input" v-model="username" placeholder="Username"/>
               <input type="input" v-model="password" placeholder="password"/>
               <input type="submit" value="Login" v-on:click="login">
               <input type="submit" value="Register">
           </div>
           <div>Forgot password!</div>
   </div>
</template>

<script>
export default {
    name: "Login",
   
    data() {
        return {
            username: "",
            password: "",
        }
    },

    methods: {
        login() {
            console.log("Login with", this.username, this.password)
        },
    }
}
</script>
