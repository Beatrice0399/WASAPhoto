<script>
export default {
    props: ["uid", "username"],
    methods: {
        selectedUser() {
            this.$emit('selectedUser', this.uid)
        }
    },
}
</script>

<template>
   <div class="row mb-2 mt-2">
        <div class="col d-flex justify-content-center">
            <div class="user-mini-card card bg-transparent border-start">
                <div class="card-body">
                    
                    <h5 @click="selectedUser" class="user-mini-card-title d-flex justify-content-center " > 
                      <img class="profile-img" src="../assets/images/profile.png">  @{{username }}
                    </h5>
                </div>
            </div>
        </div>
    </div>
  
</template>

<style>

.profile-img {

  margin-right: 20px; /* Aggiunge un margine destro tra l'immagine e l'username */

}

.user{
    color: var(--color-text);
}

.user-mini-card{
    color: black;
    width: 25%;
    border: none;
    border-color: black;
}

.user-mini-card-title:hover{
    cursor: pointer;
    font-weight: bold;
}
</style>
