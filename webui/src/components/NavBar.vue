<script>
export default {
    data() {
        return {
          text: "",  
          iconProfile: "fa-regular",
        }
      
    },
    methods:{
        logout() {
            localStorage.removeItem('token')
            this.$emit('logoutNavBar', false)
        },
    },
}
</script>

<template>

    <nav class="navbar navbar-expand-lg bg-light d-flex justify-content-between sticky-top mb-3 my-nav bg-transparent">

        <div class="col-4">
            <a class="navbar-brand ms-2 d-flex" @click="goBackHome">
              <img class="brand-img" src="../assets/images/favicon.ico">
              <div>WASAPhoto</div>
          </a>
            <button @click="myProfile" class="my-trnsp-btn me-" type="button" >
                <!--Profile-->
                <i :class="'my-nav-icon-profile me-1 w-100 h-100 '+iconProfile+ ' fa-user'" @mouseover="profileIconActive" @mouseout="profileIconInactive"></i>
            </button>

            <button @click="logout" class="my-trnsp-btn me-2" type="button">
                <!--Logout-->
                <i class="my-nav-icon-quit me-1 w-100 h-100 fa-solid fa-right-from-bracket"></i>
            </button>
        </div>

        <div class="col-4">
            <form class="form-inline my-2 my-lg-0 d-flex justify-content-center m-auto">             
                <input class="form-control mr-sm-2 w-50" v-model="textVar" type="search" placeholder="Search users">
                <button class="btn btn-light ms-2" type="submit" @click.prevent="searchFunc" style="display: none;">Search</button>
            </form>
        </div>
    </nav>
</template>

<style>
.my-nav {
    color: blue;
}

</style>