<script>
export default {
    data(){
        return{
            owner: false,
        }
    },
    props: ["cid", "userid", "username", "text", "date", "isOwner", "phid"],

    methods: {
        async removeComment() {
            try {
                await this.$axios.delete("/users/"+this.userid+"/photos/"+this.phid+"/comments/"+this.cid, {
                username : localStorage.getItem('token')}) 
                this.$emit("commentRemoved", this.cid)
            } catch (e) {
                console.log(e)
            }    
        },       
    },  
}
</script>

<template>
    <div class="container-fluid my-comment" style="border: 1px solid black;">
        <div class="d-flex justify-content-between">
            <div class="my-card-name" style="background: #d2dee9;">
                <strong>{{this.username}}</strong>:
            </div>
            <div>  <span style="opacity: 0.6;"> {{this.date}} </span> </div>
        </div> {{this.text}}
        <div class="d-flex justify-content-end my-btn">
            <button v-if="this.isOwner" @click="removeComment" type="button" class="btn btn-link" >
                remove
            </button>
        </div>
    </div>
</template>

<style>
.btn-link {
    color: black;
}
.btn-link {
    font-size: 15px;
    text-align: center;
}
.my-comment {
    padding: 8px;
    background: #dce9f5;
    border-radius: 10px;
}
.my-card-name {  
    font-style: italic;
}
</style>