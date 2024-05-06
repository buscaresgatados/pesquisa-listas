<script setup>
import LinkList from './LinkList.vue';
import Search from './Search.vue';

</script>

<template>
  <div class="wrapper">
    <div class="full-wrapper">
      <div class="header">
        <h1 class="text-3xl-bold">Lista de Abrigados</h1>
        <p>Pesquisa em compilado de listas de pessoas resgatadas nas enchentes do Rio Grande do Sul</p>
      </div>  
      <Search />
      <LinkList />
      
      <!-- <ResultPage search="John Doe"/> -->
    </div>
    </div>
</template>

<style scoped>


</style>
