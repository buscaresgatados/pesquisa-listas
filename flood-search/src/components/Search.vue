<script setup>
import { UInput } from 'nuxt-ui-vue';
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
  <p>Nome completo</p>
  <div class="cta">
    <UInput v-model="searchTerm" class="UInput" autocomplete="off" placeholder="Buscar por nome" />
    <button @click="onSearch">Buscar</button>
  </div>
</template>

<style scoped>
  p {
    margin-bottom: 2px;
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