<script>
export default {
    data(){
        return{
            photoPATH: "",
            liked: false,
            owner: false,
            totalComments: [],
            totalLikes: [],
            countLikes: 0
        }
    },
    props: ["phid", "uid", "username", "path", "comments", "likes", "date"],

    methods: {
        isOwner() {
            return this.uid === localStorage.getItem('token') ? this.owner=true : this.owner=false    
        },
        loadPhoto() { 
            this.photoPATH = __API_URL__+ "/users/"+this.uid+"/photos/"+this.phid
            this.loadLikes()
            this.isOwner()  
        },
        async deletePhoto() {
            try{
                await this.$axios.delete("/users/"+this.uid+"/photos/"+this.phid, {
                username : localStorage.getItem('token')}) 
                this.$emit("deletePhoto", this.phid)
            } catch (e) {
                console.log(e)
            }
        },
        async addLike() {
            try {
                if (!this.liked) {
                    await this.$axios.put("/users/"+this.uid+"/photos/"+this.phid+"/likes/"+localStorage.getItem('token'))
                    this.countLikes += 1
                } else {
                    await this.$axios.delete("/users/"+this.uid+"/photos/"+this.phid+"/likes/"+localStorage.getItem('token'))
                    this.countLikes -= 1
                }
                this.liked = !this.liked
            } catch (e) {

            }
        },
        
        loadLikes() {
            if (this.likes != null) {
                for (let i = 0; i < this.likes.length; i++) {
                    this.countLikes += 1
                    this.totalLikes.push(this.likes[i].uid)
                    if (this.likes[i].uid == localStorage.getItem('token')) {
                        this.liked = true;
                    }
                }
            }
        },
    },

    async mounted() {
        await this.loadPhoto() 
    },
}
</script>

<template>
    <div class="container-fluid mt-3 mb-5 photo-slot col-10 d-flex justify-content-center">
        <div class="card w-50 my-card ">
            <div class="d-flex justify-content-between">
                <div>
                    {{uid}} @{{username}} 
                </div>
                <div>
                    {{date}}
                    <button v-if="this.owner" @click="deletePhoto" class="my-trnsp-btn me-2" type="button">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
                    </button>
                </div>
            </div>

            <div class="d-flex justify-content-center">
                <img :src="photoPATH" class="card-img-top img-fluid">
            </div>

            <div class="d-flex">
            <button @click="addLike" class="my-trnsp-btn me-2" type="button">
                        <svg v-if="this.liked" class="feather"><use href="/feather-sprite-v4.29.0.svg#heart" style="fill: red"/></svg>
                        <svg v-else class="feather"><use href="/feather-sprite-v4.29.0.svg#heart" /> </svg> {{countLikes}}
                    </button>
            <button v-if="isOwner" @click="addComment" class="my-trnsp-btn me-2" type="button">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
                    </button>
            </div>
        </div>
    </div>
</template>

<style>
.feather {
    fill: red; /* Specifica il colore del riempimento (fill) per l'SVG */
}
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