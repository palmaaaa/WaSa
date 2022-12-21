<script>
export default {
  data(){
    return{
      textVar: "",
    }
  },
  methods:{
    logout(){
      localStorage.removeItem('token')
      this.$emit('logoutNavbar',false)
    },
    goBackHome(){
      //console.log("anything?")
      this.$emit('requestUpdateView',"/home")
    },
    searchFunc(){
      //console.log("hit off the search bruv")
      this.$emit('searchNavbar',this.textVar)
      this.textVar=""
    },
    myProfile(){
      this.$emit('requestUpdateView',"/users/"+localStorage.getItem('token'))
    },
  },
}
</script>

<template>
  <nav class="navbar navbar-expand-lg bg-light d-flex justify-content-between sticky-top mb-3 my-nav bg-transparent">
    <a class="navbar-brand ms-2 d-flex" @click="goBackHome">
      <img class="brand-img" src="../assets/images/brand.png">
      <div class="my-brand-name">WASAPhoto</div>
    </a>

    <form class="form-inline my-2 my-lg-0 d-flex justify-content-center m-auto">
      <input class="form-control mr-sm-2" v-model="textVar" type="search" placeholder="Search users" >
      <button class="btn btn-light btn-outline-success ms-2" type="submit" @click.prevent="searchFunc">Search</button>
    </form>
    <button @click="myProfile" class="btn btn-light btn-outline-danger me-2" type="button">Profile</button>
    <button @click="logout" class="btn btn-light btn-outline-danger me-2" type="button">Logout</button>
  </nav>

</template>

<style>
.my-brand-name{
  color: white;
}

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
</style>
