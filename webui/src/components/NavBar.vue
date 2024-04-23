<script>
export default {
    data() {
        return {
          textVar: "",  
        }
      
    },
    methods:{
        logout() {
            localStorage.removeItem('token')
            this.$emit('logoutNavBar', false)
        },
        searchFunc() {
            this.$emit('searchNavBar', this.textVar)
            this.textVar = ""
        },
        goBackHome(){
            this.$emit('requestUpdateView',"/home")
        },
        getMyProfile() {
            this.$emit('requestUpdateView', "/users/" + localStorage.getItem('token'))
        }
    },
}
</script>

<template>

    <nav class="navbar navbar-expand-lg d-flex justify-content-between sticky-top mb-3 my-nav" style="background-color: #e3f2fd;">
      <div class="col-4">
          <a class="navbar-brand ms-2 d-flex" @click="goBackHome" style="background-color: #e3f2fd;" type="button">
              <img class="home-img" src="../assets/images/home.png">
              <div>WASAPhoto</div>
              
          </a>
        
      </div>

      <!-- -->
      <div class="col-4">
          <form class="form-inline my-2 my-lg-0 d-flex justify-content-center m-auto">
              <input class="form-control mr-sm-2 w-50" v-model="textVar" type="search" placeholder="Search user">
              <button class="btn btn-light ms-2" type="submit" @click.prevent="searchFunc" style="display: none;">Search</button>
          </form>
      </div>
    
    
      <div class="col-4 d-flex justify-content-end">
            
            

          <button @click="getMyProfile" class="my-trnsp-btn me-2" type="button">
              <!--Profile-->
       
              <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
              
          </button>

          <button @click="logout" class="my-trnsp-btn me-2" type="button">
              <!--Logout-->
              
              <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
          </button>
      </div>
    </nav>
</template>

<style>
.navbar {
  border-bottom: 1px solid #dee2e6; /* Colore del bordo e spessore */
}
.my-trnsp-btn {
    background-color: transparent; 
    border-radius: 10px;
}

.feather {
  width: 30px; /* Larghezza desiderata */
  height: 30px; /* Altezza desiderata */
}
</style>