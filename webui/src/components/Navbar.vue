<script>
export default {
}
</script>

<template>

<nav class="my-nav navbar navbar-expand-lg rounded">
  <div class="container-fluid">
    <a class="navbar-brand" href="#/home">WASAPhoto</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">

      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <a class="nav-link navbar-elements" aria-current="page" href="#">Profile</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Link</a>
        </li>
        
		<li class="nav-item">
<form class="d-flex" role="search">
        <input class="form-control me-2" type="search" placeholder="Search users" aria-label="Search">
        <button class="btn btn-dark btn-outline-success navbar-btn" type="submit">Search</button>
      </form>
		</li>
      </ul>
      
    </div>
  </div>
</nav>


</template>

<style>
.my-nav {
    background-color: rgba(155, 155, 155, 0.911);
}

.navbar-btn {
    background-color: rgba(235, 79, 79, 0.795);
    color: grey;
    border-color: white;
}

.navbar-elements {
    color: rgb(231, 152, 47);

}
</style>
