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
            this.loadSearchedUsers()
        },
    },
    methods: {
        async loadSearchedUsers(){
            this.errormsg = null;
            this.users = []
            if ( this.searchValue === undefined || this.searchValue === "" || this.searchValue.includes("?") || this.searchValue.includes("_")) {               
                this.users = []
                return
            }
            try { 
                let response = await this.$axios.get("/users", {
                    params: {
                        username: this.searchValue, //profileName
                    },
                    
                });
                this.users = response.data
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        getProfile(uid) {
            this.$router.replace("/users/" + uid)
        }
    },
    async mounted() {
        if (!localStorage.getItem('token')) {
            this.$router.replace("/session")
        }
        await this.loadSearchedUsers()
    },

}
</script>

<template>
    <div class="container-fluid h-100">
        <UserContainer v-for="(user,index) in users"
            :key="index"
            :uid="user.uid"
            :username="user.username"
            @selectedUser="getProfile"/>
        <p v-if="users.length == 0" class="no-result d-flex justify-content-center"> No users found.</p>
        
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    </div>
</template>

<style>
.no-result {
    color: black;
}
</style>