<script setup>
import router from '../routes'
import { ref } from 'vue';

const searchTerm = ref('');
const onSearch = () => {
  router.push({ name: 'results', params: { term: searchTerm.value } });
}

const handleKeyPress = (event) => {
  if (event.key === 'Enter') {
    onSearch();
  }
}
</script>

<template>
  <div class="search-wrap">
    <p>Nome</p>
    <div class="cta">
      <input class="input-style" v-model="searchTerm" @keyup.enter="handleKeyPress" autocomplete="off" placeholder="Buscar por nome" />
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