<script>
export default {
    data(){
        return{
            photoPATH: "",
            liked: false,
            totalComments: [],
            totalLikes: [],
        }
    },
    props: ["phid", "uid", "comments", "likes", "date"],

    methods: {
        loadPhoto() {

        },
        async deletePhoto() {
            try{
                await this.$axios.delete("/users/"+this.uid+"/photos/"+this.phid)
                this.$emit("deletePhoto", this.phid)
            } catch (e) {

            }
        },
    }
}
</script>

<template>
    <div class="container-fluid mt-3 mb-5 photo-slot">
        <div class="card my-card">

        </div>
    </div>
</template>

<style>
.photo-solt {
    height: 30px;
    width: 50%;
    border: dotted black 1px;
    margin: 10px;
    padding: 10px;
    background-color: rgb(147, 199, 255);
    font-weight: bold;
}
</style>