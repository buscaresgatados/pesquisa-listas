<script setup>
import FoundResults from './FoundResults.vue';

defineProps({
  searchResults: Array
  })
</script>

<template>
  <div class="results-wrapper">
    <div class="results-wrapper-top">
      <p class="found-results">X resultados encontrados</p>
      <p class="last-update">Atualizado por último às {xx:xx} do dia {xx de Maio}</p>
    </div>
    <div>
      <div v-for="result in searchResults" :key="result.SheetId">
      <FoundResults
        :Nome="result.Nome"
        :Idade="result.Idade"
        :listId="result.SheetId"
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
