<script setup>
import router from '../routes'
import { ref } from 'vue';

const searchTerm = ref('');
const showError = ref(false);
const onSearch = () => {
  if(searchTerm.value) {
    showError.value=false;
    router.push({ name: 'results', params: { term: searchTerm.value } });
  } else {
    showError.value = true;
  }
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
      <div class="input-wrapper">
        <input class="input-style" v-model="searchTerm" @keyup.enter="handleKeyPress" autocomplete="off" placeholder="Buscar por nome" />
        <span class="inline-algolia">Powered by <a href="https://www.algolia.com/?utm_medium=AOS-referral"><img class="algolia-logo" src="../assets/Algolia-logo-white.svg"/></a></span>
      </div>
    </input>
      <button @click="onSearch">Buscar</button>
    </div>
    <span class="error-message" v-if="showError">Busca inválida</span>
  </div>
</template>

<style scoped>
  .input-wrapper { 
    position: relative;
    width: 97%; 
  }

  .inline-algolia {
    font-size: 12px;
    position: absolute;
    top: 13px;
    right: 10px;
    flex-direction: row;
    text-decoration: none;
    opacity: 0.66;
    a {
      display: inline
    }
    /* I'm going to do this and you can't stop me, CSS. ~ Igor */ 
    @media (max-width: 480px) {
      font-size: 0;
    }
  }

  .algolia-logo {
    position:relative;
    top: 3px;
    width: 60px;
    color: #E5E5E5;
  }
  .error-message{
    display: block;
}

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