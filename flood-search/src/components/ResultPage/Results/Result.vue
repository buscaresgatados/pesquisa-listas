<script setup>
import { ref, computed } from 'vue';


defineProps({
  searchResults: Array
  })

  const logResults = computed(() => {
  searchResults.value.forEach(result => {
    console.log(result);
  });
});
</script>

<template>
  <div class="results-wrapper">
    <div class="results-wrapper-top">
      <p class="found-results">{{ searchResults.length }} resultados encontrados</p>
      <!-- We'll add it once we have the endpoint -->
      <!-- <p class="last-update">Atualizado às {xx:xx} do dia {xx de Maio}</p> -->
    </div>
    <div>
      <div v-for="result in searchResults" :key="result.SheetId">
      <FoundResults
        :Nome="result.Nome"
        :Idade="result.Idade"
        :listId="result.SheetId"
        :Abrigo="result.Abrigo"
        :Url="result.URL"
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
    color: white;
  }
  .last-update{
    font-size: 13px;
    font-weight: 400;
    color: white;
    max-width: 160px
  }
</style>
