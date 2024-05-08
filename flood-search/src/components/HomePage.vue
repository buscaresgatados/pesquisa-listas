<script setup>
import LinkList from './LinkList.vue';
import Search from './Search.vue';
import Footer from './Footer.vue';
import Partners from './Partners.vue';

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
      <Partners />
      <Footer />
    </div>
    </div>
</template>

<style scoped>

</style>
