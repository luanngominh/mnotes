<template>
    <Login/>
</template>

<script>
import Login from "./Login.vue"

export default {
    name: 'Home',
    data() {
        return {
            loggedIn: false
        }
    },
    components: {
        Login,
    }
}
</script>

