<script>
export default {
    name: 'ModalSettings',
    data() {
        return {
            newUsername: "",
        }
    },
    props: ["uid", "username"], 
    methods: {
        close() {
            this.$emit('close')
        },
        async setUsername() {        
            try {
                let response = await this.$axios.put("/users/" + this.$route.params.uid, {
                    username : this.newUsername.trim(), 
                })
                this.$emit('setUsername')
            } catch (e) {
                console.log(e)
                this.errormsg = e.toString();
            }
            this.newUsername = ""
            this.close() 
        },
        
    },
}
</script>

<template>
  <transition name="modal-fade">
    <div class="modal-backdrop">
      <div class="modal" role="dialog" aria-labelledby="modalTitle" aria-describedby="modalDescription">
        <header class="modal-header" id="modalTitle">
          <slot name="header">
            Set a new Username
          </slot>
          <button type="button" class="btn-close" @click="close" aria-label="Close modal"> x </button>
        </header>

        <section class="modal-body" id="modalDescription">
          
          <div class="form-floating mb-3">
                    <input type="text" class="form-control" id="floatingInput" v-model="newUsername" maxlength="16" minlength="3" placeholder="newUsername" />
                    <label for="floatingInput">New username</label>
                </div>
                
        </section>

        <footer class="modal-footer">
          <div class="container-fluid-justify-center">
          <button type="button" class="btn-green" @click="setUsername" >Save</button> 
 
          </div>
          
        </footer>
      </div>
    </div>
  </transition>
</template>

<style>
.modal-backdrop {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.3);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal {
  background: #FFFFFF;
  box-shadow: 2px 2px 20px 1px;
  overflow-x: auto;
  width: 50%;
  max-height: 30%;
  padding: 20px;
  display: flex;
  flex-direction: column;
  position: relative;
  top: 10%;
  transform: translateY(-50%);
}

.modal-header,
.modal-footer {
  padding: 15px;
  display: flex;
}

.modal-header {
  position: relative;
  border-bottom: 1px solid #eeeeee;
  color: #4AAE9B;
  justify-content: space-between;
}

.modal-footer {
  border-top: 1px solid #eeeeee;
  flex-direction: column;
  justify-content: flex-end;
}

.modal-body {
  position: relative;
  padding: 20px 10px;
  flex-grow: 1;
}

.btn-close {
  position: absolute;
  top: 0;
  right: 0;
  border: none;
  font-size: 20px;
  padding: 10px;
  cursor: pointer;
  font-weight: bold;
  color: #4AAE9B;
  background: transparent;
}

.btn-green {
  color: white;
  background: #4AAE9B;
  border: 1px solid #4AAE9B;
  border-radius: 2px;
}

.modal-fade-enter,
.modal-fade-leave-to {
  opacity: 0;
}

.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity .5s ease;
}
</style>
