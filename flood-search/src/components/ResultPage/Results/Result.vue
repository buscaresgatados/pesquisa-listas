<script setup>
import FoundResults from './FoundResults.vue';
import ResultService from '../../../service/ResultService';
import { onMounted, ref } from 'vue';

defineProps({
    name: String,
    required: true,
  })
  
  const searchResults = ref([]);
  const isValueReady = ref(false); 
  
  const fetchData = async () => {
  try {
    const response = await ResultService.search(this.name);
    searchResults.value = response.data;
    isValueReady.value = true;
  } catch (error) {
    console.error('Erro na busca:', error);
  }

  onMounted(fetchData)
}


</script>

<template>
  <div class="results-wrapper">
    <div class="results-wrapper-top">
      <p class="found-results">{{ searchResults.length }} resultados encontrados</p>
      <p class="last-update">Atualizado por último às {xx:xx} do dia {xx de Maio}</p>
    </div>
    <div v-if="isValueReady">
      <div v-for="result in searchResults" :key="result.listId">
        <FoundResults
          :Nome="result.Nome"
          :Idade="result.Idade"
          :listId="result.listId"
          :Abrigo="result.Abrigo"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
  .results-wrapper-top{
    display: flex;
    width: 100%;
    justify-content: space-between;
    margin-top: 40px;
  }
  .found-results {
    font-size: 16px;
    font-weight: 400;
    color: #CBD5E1;
  }
  .last-update{
    font-size: 12px;
    font-weight: 400;
    color: #CBD5E1;
  }
</style>
