<script setup>
import router from '../routes'
import { ref, computed, onUnmounted, onMounted } from 'vue';

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
const mediaQuery = window.matchMedia('(min-width : 480px)');
const isScreenLarge = ref(mediaQuery.matches)

const update = (event) => (isScreenLarge.value = event.matches);
onMounted(() => mediaQuery.addEventListener("change", update));
onUnmounted(() => mediaQuery.removeEventListener("change", update));

const logoLarge = new URL('../assets/Algolia-logo-white.svg', import.meta.url).href
const logoSmall = new URL('../assets/Algolia-mark-white.svg', import.meta.url).href
const algoliaLogo = computed(() => isScreenLarge.value ? logoLarge : logoSmall);
</script>

<template>
  <div class="search-wrap">
    <p>Nome</p>
    <div class="cta">
      <div class="input-wrapper">
        <input class="input-style" v-model="searchTerm" @keyup.enter="handleKeyPress" autocomplete="off" placeholder="Buscar por nome" />
        <span class="inline-algolia">Powered by 
          <a target="_blank" href="https://www.algolia.com/?utm_medium=AOS-referral">
            <img class="algolia-logo" :src="algoliaLogo"/>
          </a>
      </span>
      </div>
    </input>
      <button @click="onSearch">Buscar</button>
    </div>
    <span class="error-message" v-if="showError">Busca inválida</span>
  </div>
</template>

<style scoped>

  .input-style {
    padding: 12px 26% 12px 8px;
    background-color: transparent ;
    border-radius: 8px;
    border-color: #00DC82;
    border-style: solid;
    box-sizing:border-box;
    display: flex;
    color: white;
    width:100%;
  }

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
    @media (max-width: 380px) {
      font-size: 0;
    }
  }

  .algolia-logo {
    position:relative;
    top: 3px;
    width: 60px;
    height: 13.66px;
    color: #E5E5E5;
    @media (max-width: 480px) {
      width: 13.66px;
    }
  }

  .error-message{
    display: block;
  }

  .search-wrap {
    padding: 48px 0px 80px 0px;
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