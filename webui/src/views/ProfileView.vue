<script>
export default {
    data: function() {
        return {
            errormsg: "",
            username: "",
            follower: [],
            following: [],
            photo: [],
            isBanned: false,
            isFollowed: false,
        }
    },
    computed:{
        currentPath() {
            return this.$route.params.uid
        },

        isOwner() {
            return this.$route.params.uid === localStorage.getItem('token')
        },
    },
    methods: {
        async follow(){
            try{
                if (this.isFollowed) {
                    await this.$axios.delete("/users/" + this.$route.params.uid + "/followers/" + localStorage.getItem('token'));

                } else {
                    await this.$axios.put("/users/" + this.$route.params.uid + "/followers/" + localStorage.getItem('token'));
                }
                this.isFollowed = !this.isFollowed
            } catch (e) {
                this.errormsg = e.toString();
            }
        },

        async ban() {
            try{
                if (this.isBanned) {
                    await this.$axios.delete("/users/" + this.$route.params.uid + "/bannedUsers/" + localStorage.getItem('token'));
                } else {
                    await this.$axios.put("/users/" + this.$route.params.uid + "/bannedUsers/" + localStorage.getItem('token'));
                }
                this.isBanned = !this.isBanned
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        
    }
}

</script>

<template>

</template>

<style>
</style>