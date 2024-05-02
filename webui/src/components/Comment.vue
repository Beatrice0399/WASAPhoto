<script>
export default {
    props: ["cid", "username", "text", "date"],

    methods: {
        async addComment() {
            await this.$axios.post("/users/"+this.uid+"/photos/"+this.phid+"/comments")
        }
    },  

}
</script>

<template>

</template>

<style>

</style>