<script>
import ErrorMsg from '../components/ErrorMsg.vue';
export default {
    data: function() {
        return {
            errormsg: null,
            username: "",
            disabled: true,

        }
    },
    methods: {
        async login() {
            this.errormsg = null
            try {
                let response = await this.$axios.post("/session", {
                    username : this.username.trim()

                });
                localStorage.setItem('token', response.data.uid)
                this.$router.replace("/home")
                this.$emit('updatedLoggedChild',true)
            } catch (e) {
                this.errormsg = e.toString();
                console.log(ErrorMsg);
            }
        },
    },
    mounted() {
        if (localStorage.getItem('token')){
            this.$router.replace("/home") 
        }
    },
}
</script>

<template>
  
    <div class="container-fluid login">
        
        <div  class="row">
            <div  class="col">
                <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
            </div>
        </div>
        
        <!--<div class="row mt-5 mb-6">
            <div class="col text-center">
                <h1 class="head-title">Benvenuto su WASAPHOTO</h1>
            </div>
            </div>
        -->
        <div class="row h-100 w-100 m-0">
			
			<form @submit.prevent="login" class="d-flex flex-column align-items-center justify-content-center p-0">

				<div class="row mt-2 mb-3 border-bottom">
					<div class="col">
						<h2 class="login-title">Welcome to WASAPhoto</h2>
					</div>
				</div>
                <div class="form-floating mb-3">
                    <input type="text" class="form-control" id="floatingInput" v-model="username" maxlength="16" minlength="3" placeholder="Username" />
                    <label for="floatingInput">Username</label>
                </div>
                

				<div class="row mt-2 mb-5 ">
					<div class="col ">
						<button class="btn btn-dark" :disabled="username == null || username.length >16 || username.length <3 || username.trim().length<3"> 
						Registration/Login 
						</button>
					</div>
				</div>
			</form> 
		</div>
    </div>
 
</template>

<style>
.btn-dark {
  background-color: #6680B6; 
  padding: 10px 20px;
  border: none;
  margin-top: 10px;
}


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
    font-size: 40px;
    font-family:'Times New Roman';
}

.head-title {
    color: #3F749C;
}

</style>