<script>

export default {
    
    data: function() {
        return {
            errormsg: "",
            username: "",
            follower: [],
            following: [],
            photos: [],
            isBanned: false,
            isFollowed: false,

            countFollower: 0,
            countFollowing: 0,
            countPhoto: 0,

            newName: "",
            showPanel: false,
        }
    },

    watch:{
        currentPath(newuid,olduid){
            if (newuid !== olduid && (this.$route.path !== '/home' && this.$route.path !== '/search')) {  
                this.getInfo()
                window.location.reload();
            }
        },
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
            this.showPanel = false
            // window.location.reload();
            // this.$router.replace("/users/" + localStorage.getItem('token')) 
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
                    this.getInfo()
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
                if (response.status === 206){
                    this.isBanned = true
                    return
                }
                this.username = response.data.name
                this.follower = response.data.followers
                this.following = response.data.following
                this.countFollower = response.data.followers != null ? response.data.followers.length: 0
                this.countFollowing = response.data.following != null ? response.data.following.length : 0
                this.photos = response.data.photos
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
            let fileInput = document.getElementById('fileUploader');
            const file = fileInput.files[0];
            const reader = new FileReader();

            reader.readAsArrayBuffer(file);
            reader.onload = async () => {
                let arrayBuffer = reader.result;
                let uint8Array = new Uint8Array(arrayBuffer);
                let response = await this.$axios.post("/users/"+this.$route.params.uid+"/photos", uint8Array, {
                    headers: {
                        'Content-Type': file.type
                    },
                });
                this.photos.unshift(response.data);
                this.countPhoto += 1;
            };
            window.location.reload();
        },
        removePhoto(phid){
			this.photos = this.photos.filter(item => item.phid !== phid)
            window.location.reload();
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
                            <div class="card-body">
                                <h5 class="card-title p-0 me-auto mt-auto">{{this.$route.params.uid}} @{{this.username}}
                                <button v-if="isOwner" @click="togglePanel" class="my-trnsp-btn me-2" type="button" style="border: null">
                                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#edit-3"/></svg>
                                    <div v-if="showPanel" class="pannello" style="position: absolute;">
                                        <div class="pannello-contenuto ">
                                            <input type="text" class="form-control" v-model="newName" maxlength="16" minlength="3" placeholder="New username" style="margin-bottom: 10px;"/>
                                            <button @click="setUsername" class="btn" :disabled="newName == null || newName.length >16 || newName.length <3 || newName.trim().length<3">Save</button>
                                            <button @click="showPanel = false" class="btn">Cancel</button>
                                        </div>
                                        
                                    </div>
                                </button>
                               
                                </h5>                  
                            </div>
                            <p> Follower: {{countFollower}}</p>
                            <p> Following: {{countFollowing}}</p>
                            <p> Photo: {{countPhoto}}</p>
                            <input id="fileUploader" type="file" class="profile-file-upload" @change="addPhoto" accept=".jpg, .png">
                            <label v-if="isOwner" class="btn my-btn-add-photo ms-2" for="fileUploader" style="background-color: rgb(114, 152, 174); margin-top: 10px; margin-bottom: 10px;"> Add photo</label>
    

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
         <div class="row">

            <div class="container-fluid mt-3">

                <div class="row ">
                    <div class="col-12 d-flex justify-content-center">
                        <h2>Photos</h2>
                    </div>
                </div>
                <hr>
            

                <div class="row">
                    <div class="col">
                        <div v-if="!isBanned && countPhoto>0">
                            <Photo v-for="(photo, index) in photos"
                            :key="index"
                            :phid="photo.phid"
                            :uid="this.$route.params.uid"
                            :username="photo.username"
                            :path="photo.path"
                            :comments="photo.comments"
                            :likes="photo.likes"
                            :date="photo.date"
                            @deletePhoto="removePhoto"
                            />
                        </div>
                        <div v-else>
                            <h3> No post</h3>
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
.profile-file-upload{
    display: none;
}

</style>