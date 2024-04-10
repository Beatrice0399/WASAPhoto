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
    <nav class="navbar  navbar-expand-lg bg-dark d-flex justify-content-between sticky-top mb-3 my-nav bg-dark">
        <div class="col-4 d-flex justify-content-end">
          <button @click="myProfile" class="my-trnsp-btn me-2" type="button">
              <!--Profile-->
              <i :class="'my-nav-icon-profile me-1 w-100 h-100 '+iconProfile+ ' fa-user'" @mouseover="profileIconActive" @mouseout="profileIconInactive"></i>
          </button>

          <button @click="logout" class="my-trnsp-btn me-2" type="button">
              <!--Logout-->
              <i class="my-nav-icon-quit me-1 w-100 h-100 fa-solid fa-right-from-bracket"></i>
          </button>
    </nav>
</template>
