<script setup>
import Search from './components/Search.vue';

</script>

<template>
  <div class="wrapper">
    <div class="header">
      <h1 class="title">Lista de Abrigados</h1>
      <p class="main-text">Pesquisa em compilado de listas de pessoas resgatadas nas enchentes do Rio Grande do Sul</p>
      <Search />
    </div>

  </div>
</template>

<style scoped>
.title {
  color: white;
  font-weight: 600;
  font-size: 30px;
  line-height: 36px;
  margin-top: 0;
  margin-left:0;
}
.main-text{
  color: white;
}
.wrapper {
  background-color: #111827;
  height: 100vh;
  width: 100vw;
  
}
</style>
