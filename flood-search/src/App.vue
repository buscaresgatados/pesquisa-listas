<script setup>
import LinkList from './components/LinkList.vue';
import Search from './components/Search.vue';

</script>

<template>
  <div class="wrapper">
    <div class="header">
      <h1 class="text-3xl-bold">Lista de Abrigados</h1>
      <p class="main-text">Pesquisa em compilado de listas de pessoas resgatadas nas enchentes do Rio Grande do Sul</p>
    </div>

    <Search />
    <LinkList />


  </div>
</template>

<style scoped>
.text-3xl-bold {
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
  display: flex;
  flex-direction: column;
  height: 100vh;
  width: 100vw;
  margin-top: 32px;
  
}

</style>
