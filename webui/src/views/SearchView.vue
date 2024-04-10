<script>
export default {
    data: function() {
        return {
            users: [],
            errormsg: null,
        }
    },
    props: ["searchValue"],

    watch: {
        searchValue: function() {
            this.loadSearchUsers()
        },
    },
    methods: {
        async loadSearchUsers(){
            this.errormsg = null;
            if ( this.searchValue === undefined || this.searchValue === "" || this.searchValue.include("?") || this.searchValue.include(".")) {
                this.users = []
                return
            }
            try {
                let response = await this.$axios.get("/users", {
                    params: {
                        username: this.searchValue,
                    },
                });
                this.users = response.data
            } catch (e) {
                this.errormsg = e.toString();
            }
        }
    },
    async mounted() {
        if (!localStorage.getItem('token')) {
            this.$router.replace("/login")
        }
        await this.loadSearchUsers()
    },

}
</script>

<template>
    <div class="container-fluid h-100">
        <UserList v-for="(user,index) in users"
            :key="index"
            :uid="users.uid"
            :username="users.username"/>
        <p v-if="users.length == 0" class="no-result d-flex justify-content-center"> No users found.</p>
        
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    </div>
</template>

<style>
.no-result {
    color: black;
}
</style>