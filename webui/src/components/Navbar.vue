<script>
export default {
  data(){
    return{
      textVar: "",
      iconProfile: "fa-regular",
    }
  },
  methods:{
    logout(){
      localStorage.removeItem('token')
      this.$emit('logoutNavbar',false)
    },
    goBackHome(){
      this.$emit('requestUpdateView',"/home")
    },
    searchFunc(){
      this.$emit('searchNavbar',this.textVar)
      this.textVar=""
    },
    myProfile(){
      this.$emit('requestUpdateView',"/users/"+localStorage.getItem('token'))
    },
    profileIconInactive(){
      this.iconProfile = "fa-regular"
    },
    profileIconActive(){
      this.iconProfile = "fa-solid"
    },
  },
}
</script>

<template>
  <nav class="navbar navbar-expand-lg bg-light d-flex justify-content-between sticky-top mb-3 my-nav bg-transparent">
      <div class="col-4">
          <a class="navbar-brand ms-2 d-flex" @click="goBackHome">
              <img class="brand-img" src="../assets/images/brand.png">
              <div>WASAPhoto</div>
          </a>
      </div>

      <!-- -->
      <div class="col-4">
          <form class="form-inline my-2 my-lg-0 d-flex justify-content-center m-auto">
              <input class="form-control mr-sm-2 w-50" v-model="textVar" type="search" placeholder="Search users">
              <button class="btn btn-light ms-2" type="submit" @click.prevent="searchFunc" style="display: none;">Search</button>
          </form>
      </div>

      <div class="col-4 d-flex justify-content-end">
          <button @click="myProfile" class="my-trnsp-btn me-2" type="button">
              <!--Profile-->
              <i :class="'my-nav-icon-profile me-1 w-100 h-100 '+iconProfile+ ' fa-user'" @mouseover="profileIconActive" @mouseout="profileIconInactive"></i>
          </button>

          <button @click="logout" class="my-trnsp-btn me-2" type="button">
              <!--Logout-->
              <i class="my-nav-icon-quit me-1 w-100 h-100 fa-solid fa-right-from-bracket"></i>
          </button>
      </div>
  </nav>
</template>

<style>

.my-nav {
  background: transparent;
}
.my-nav:hover{
  cursor: pointer;
}

.navbar-btn {
    background-color: rgba(235, 79, 79, 0.795);
    color: grey;
    border-color: white;
}

.navbar-elements {
    color: rgb(231, 152, 47);

}
.brand-img{
  height: 30px;
  width: 30px;
}

.my-nav-icon-profile{
  color: black;
}
.my-nav-icon-quit{
  color: black;
}
.my-nav-icon-quit:hover{
  color: var(--color-red-danger);
  transform: scale(1.2);
}
</style>
