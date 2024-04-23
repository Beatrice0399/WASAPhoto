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

            countFollower: 0,
            countFollowing: 0,
            countPhoto: 0,

            newName: "",
            showPanel: false,
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
        togglePanel() {
            this.showPanel = true;
        },
        closePanel() { 
            this.showPanel = false;
        },
        async follow(){
            try{
                if (this.isFollowed) {
                    await this.$axios.delete("/users/" + this.$route.params.uid + "/followers/" + localStorage.getItem('token'))
                    this.countFollower -= 1

                } else {
                    await this.$axios.put("/users/" + this.$route.params.uid + "/followers/" + localStorage.getItem('token'));
                    this.countFollower += 1
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
                if (this.isBanned && this.isFollowed) {
                    this.countFollower -= 1
                    this.isFollowed = false;
                }
            } catch (e) {
                this.errormsg = e.toString();
            }
            
        },

        async getInfo() {

            try {
                let response = await this.$axios.get("/users/"+this.$route.params.uid);
    
                this.username = response.data.name
                this.follower = response.data.followers
                this.following = response.data.following
                this.countFollower = response.data.followers != null ? response.data.followers.length: 0
                this.countFollowing = response.data.following != null ? response.data.following.length : 0
                this.photo = response.data.photos
                this.countPhoto = response.data.photos != null ? response.data.photos.length : 0

                if (response.data.followers != null) {
                    for (let i = 0; i < response.data.followers.length; i++) {
                        if (response.data.followers[i].uid == localStorage.getItem('token')) {
                            this.isFollowed = true;
                        }
                    }
                }
                
            } catch (e) {
                this.errormsg = e.toString();
                this.isBanned = true;
            }
            
        },
        async setUsername() {
            
            try {
                let response = await this.$axios.put("/users/" + this.$route.params.uid, {
                    username : this.newName.trim(),
                    
                })
                this.showPanel = false
            } catch (e) {
                console.log(e)
                this.errormsg = e.toString();
            }
            this.closePanel()
            this.getInfo()

        },
        async addPhoto() {

        },
        
    },

    async mounted() {
        await this.getInfo()
    },
}

</script>

<template>
   <div class="container-fluid">
        <div class="row">
            <div class="col-12 d-flex justify-content-center">
                <div class="card w-50 container-fluid">

                    <div class="row">
                        <div class="col">
                            <div class="card-body d-flex justify-content-between align-items-center">
                                <h5 class="card-title p-0 me-auto mt-auto">{{this.$route.params.uid}} @{{this.username}}
                                <button v-if="isOwner" @click="togglePanel" class="my-trnsp-btn me-2" type="button" style="border: null">
                                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#edit-3"/></svg>
                                    <div v-if="showPanel" class="pannello">
                                    	
                                        <div class="pannello-contenuto">
                                            <input type="text" class="form-control" v-model="newName" maxlength="16" minlength="3" placeholder="New username" style="margin-bottom: 10px;"/>
                                            <button @click="setUsername" class="btn" :disabled="newName == null || newName.length >16 || newName.length <3 || newName.trim().length<3">Save</button>
                                            <button @click="closePanel" class="btn">Cancel</button>
                                        </div>
                                        
                                    </div>
                                </button>
                               
                                </h5>                  
                            </div>
                            <p> Follower: {{countFollower}}</p>
                            <p> Following: {{countFollowing}}</p>
                            <p> Photo: {{countPhoto}}</p>
                            <button v-if="isOwner" @click="addPhoto" class="btn ms-3" style="background-color: rgb(114, 152, 174); margin-top: 10px; margin-bottom: 10px;">Add photo</button> 
                            <button v-if="!isOwner && !isBanned" @click="follow" class="btn ms-3" style="background-color: rgb(114, 152, 174); margin-top: 10px; margin-bottom: 10px;">
                                    {{isFollowed ? "Unfollow" : "Follow"}}
                                </button>

                                <button v-if="!isOwner" @click="ban" class="btn ms-3" style="background-color: rgb(114, 152, 174); margin-top: 10px; margin-bottom: 10px;">
                                    {{isBanned ? "Unban" : "Ban"}}
                                </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
   </div>
</template>

<style>

.my-trnsp-btn {
    border: none;
}

.pannello {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* Sfondo semitrasparente */
  display: flex;
  align-items: center;
  justify-content: center;
}

.pannello-contenuto {
  background-color: white;
  padding: 20px;
  border-radius: 5px;
}

textarea {
  width: 100%;
  height: 100px;
  margin-bottom: 10px;
}

</style>