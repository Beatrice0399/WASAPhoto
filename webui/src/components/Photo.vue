<script>
export default {
    data(){
        return{
            photoPATH: "",
            photoid: 0,
            liked: false,
            owner: false,
            totalComments: [],
            totalLikes: [],
            countLikes: 0,
            countComments: 0,
            newComment: ""
        }
    },
    props: ["phid", "uid", "username", "path", "comments", "likes", "date"],

    methods: {
        isOwner() {
            return this.uid === localStorage.getItem('token') ? this.owner=true : this.owner=false    
        },
        loadPhoto() { 
            this.photoPATH = __API_URL__+ "/users/"+this.uid+"/photos/"+this.phid
            this.photoid = this.phid
            this.loadComments()
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
        async addComment() {
            await this.$axios.post("/users/"+this.uid+"/photos/"+this.phid+"/comments", {

            })
        },
        loadComments() {
            if (this.comments != null) {
                for (let i = 0; i < this.comments.length; i++) {
                    this.countComments += 1
                    this.totalComments.push(this.comments[i])
                    if (this.comments[i].uid === localStorage.getItem('token')) {
                        this.comments[i].isOwner = false
                    }
                }
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
        isOwnerComment(myvar) {
            return  myvar == localStorage.getItem('token') ? true : false    
        },
        removeCommentFromList(commentId){
            this.totalComments = this.totalComments.filter(item => item.cid !== commentId);
            window.location.reload();
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
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg> {{countComments}}
                </button>


            </div>
            <div v-if="totalComments!= undefined &&totalComments.length>0" class="container-comments">
                <div class="scrollable-content ">
                <Comment 
                    v-for="(comment, index) in this.comments"
                    :key="index"
                    :cid="comment.id"
                    :userid="comment.uid"
                    :username="comment.user"
                    :text="comment.string"
                    :date="comment.date"
                    :isOwner="isOwnerComment(comment.uid)"
                    :phid="photoid"
                    @eliminateComment="removeCommentFromList(comment.cid)"
                />
                </div>
            </div>
            
            
        </div>
    </div>
</template>

<style>

.scrollable-content {
  height: 100%; /* Utilizza tutta l'altezza del container */
  overflow-y: auto; /* Aggiunge una scrollbar verticale se il contenuto è più alto del container */
}
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