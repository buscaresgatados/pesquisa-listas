<script setup>
import { ref } from 'vue';
import ResultService from '../service/ResultService';
import router from '../routes'

const searchTerm = ref('');
const searchResults = ref([]); 

const onSearch = () => {
  const term = searchTerm.value;
  router.push({ name: 'results', params: { term } });
  ResultService.search(term)
    .then(response => {
      searchResults.value = response.data;
    })
    .catch(error => {
      console.error('Erro na busca:', error);
    });
  }
</script>

<template>
  <div class="search-wrap">
    <p>Nome</p>
    <div class="cta">
      <input class="input-style" v-model="searchTerm" autocomplete="off" placeholder="Buscar por nome" />
      <button @click="onSearch">Buscar</button>
    </div>
  </div>
</template>

<style scoped>

  .search-wrap {
    padding: 32px 0px 48px 0px;
  }

  p {
    margin-bottom: 4px;
    margin-left: 4px;

  }

  .cta {
    display: flex;
    gap: 16px;
    width: 100%;
  }

  .relative {
    width: 100%;
  }
</style>