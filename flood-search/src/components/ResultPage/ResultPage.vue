<script setup>
import ResultHeader from './ResultHeader.vue';
import ResultFooter from './ResultFooter/ResultFooter.vue'
import ResultService from '../../service/ResultService';
import Result from './Results/Result.vue';
import { useRoute } from 'vue-router';
import { ref, onMounted } from 'vue';

const route = useRoute();
const term = route.params.term;
const searchResults = ref([])
const search = () => {
  ResultService.search(term)
    .then(response => {
      searchResults.value = response
    })
    .catch(error => {
      console.error('Erro na busca:', error);
    });
  };
onMounted(search);

</script>

<template>

<div class="wrapper">
    <div class="full-wrapper"> 
        <ResultHeader :name="term"  />
        <Result :searchResults="searchResults"/>
        <ResultFooter />
    </div>
  </div>
  
</template>

<style scoped>
  .search-title{
    color: white;
  }
</style>
