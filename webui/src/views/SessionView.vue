<script>
import ErrorMsg from '../components/ErrorMsg.vue';
export default {
  components: { ErrorMsg },
    data: function() {
        return {
            errorMsg: null,
            idSession: "",
            disabled: true,
        }
    },
    methods: {
        async login() {
            this.errorMsg = null
            try {
                let response = await this.$axios.post("/session", {
                    uid : this.idSession.trim()

                });
                localStorage.setItem("token", response.data.uid);
                this.$router.replace("/home")
                this.$emit('updatedLoggedChild',true)
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
    },
    mounted() {
        if (localStorage.getItem("token")){
            this.$router.replace("/home") 
        }
    },
}
</script>

<template>
  
    <div class="container-fluid h-100 m-0 p-0 login">
        
        <div  class="row">
            <div  class="col">
                <ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>
            </div>
        </div>
        
        <div class="row mt-5 mb-6">
            <div class="col text-center">
                <h1 class="head-title">Benvenuto su WASAPHOTO</h1>
            </div>
            </div>
        
        <div class="row h-100 w-100 m-0">
			
			<form @submit.prevent="login" class="d-flex flex-column align-items-center justify-content-center p-0">

				<div class="row mt-2 mb-3 border-bottom">
					<div class="col">
						<h2 class="login-title">Login</h2>
					</div>
				</div>
                <div class="form-floating mb-3">
                    <input type="email" class="form-control" id="floatingInput" placeholder="name@example.com">
                    <label for="floatingInput">Username</label>
                </div>

				<div class="row mt-2 mb-5 ">
					<div class="col ">
						<button class="btn btn-dark" :disabled="identifier == null || identifier.length >16 || identifier.length <3 || identifier.trim().length<3"> 
						Registration/Login 
						</button>
					</div>
				</div>
			</form>
		</div>
    </div>
 
</template>

<style>

.login {
    background-image: url("../assets/images/sfondo.jpg");
    position: relative;
    height: 100vh;
    background-size: cover;
    width: 100vw; 
    
  }

  .absolute-center {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }

.login-title {
    color: #3F749C;
    margin-top: 0px
}

.head-title {
    color: #3F749C;
}

</style>