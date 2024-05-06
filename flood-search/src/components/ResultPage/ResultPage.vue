<script setup>
import ResultHeader from './ResultHeader.vue';
import ResultFooter from './ResultFooter/ResultFooter.vue'
import ResultService from '../../service/ResultService';
import Result from './Results/Result.vue';
import { useRoute } from 'vue-router';
import { ref, onMounted } from 'vue';

const route = useRoute();
const term = route.params.term;
const searchResults = ref([]);
const isLoading = ref(false)

const search = async () => {
  isLoading.value = true; 
  try {
      const response = await ResultService.search(term);
      searchResults.value = response;
  } catch (error) {
      console.error('Erro na busca:', error);
  } finally {
      isLoading.value = false;
  }
}; 
onMounted(search);

</script>

<template>

<div class="wrapper">
    <div class="full-wrapper"> 
        <ResultHeader :name="term"  />
        <div>
          <div v-if="isLoading" class="lds-ring"><div></div><div></div><div></div><div></div></div>
        </div>
        <Result v-if="!isLoading" :searchResults="searchResults"/>
        <ResultFooter />
    </div>
  </div>
  
</template>

<style scoped>
  .search-title{
    color: white;
  }
  .lds-ring {
  color: #00DC82
}
.lds-ring,
.lds-ring div {
  box-sizing: border-box;
}
.lds-ring {
  display: inline-block;
  position: relative;
  width: 80px;
  height: 80px;
}
.lds-ring div {
  box-sizing: border-box;
  display: block;
  position: absolute;
  width: 64px;
  height: 64px;
  margin: 8px;
  border: 8px solid currentColor;
  border-radius: 50%;
  animation: lds-ring 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
  border-color: currentColor transparent transparent transparent;
}
.lds-ring div:nth-child(1) {
  animation-delay: -0.45s;
}
.lds-ring div:nth-child(2) {
  animation-delay: -0.3s;
}
.lds-ring div:nth-child(3) {
  animation-delay: -0.15s;
}
@keyframes lds-ring {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
