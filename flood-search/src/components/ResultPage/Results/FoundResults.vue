<script setup>
import { DocumentMagnifyingGlassIcon } from '@heroicons/vue/24/solid'
    defineProps({
    Nome: String,
    Idade: Number,
    Abrigo: String,
    Observacao: String,
    listId: String,
    Url: String,
  })
</script>

<template>
  <div class="found-results-wrapper">
    <div >
      <h3 class="found-results">
        {{ Nome }}
      </h3>
      <div class="inner-result-container">
        <p v-if="Idade">Idade: {{ Idade }}</p>
        <p>Abrigo: {{ Abrigo }}<br></p>
        <p v-if="Observacao">Obs: {{ Observacao }}</p>
        <a v-if="listId" style="padding: 12px;" class="link" target="_blank" :href="`https://docs.google.com/spreadsheets/u/0/d/${listId}`" >
          <div class="list-name">
            <DocumentMagnifyingGlassIcon style="width: 24px; height: 24px;" />
          </div>
        </a>
        <a v-if="Url" style="padding: 12px;" class="link" target="_blank" :href="Url" >
          <div class="list-name">
            <DocumentMagnifyingGlassIcon style="width: 24px; height: 24px;" />
          </div>
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped>

  .list-name{
    font-size: 14px;
    text-decoration: none;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
    max-width: 120px;
    border-bottom: 1px solid #00DC82 !important;    
  }

  .link {
    text-decoration: none;
    text-transform: uppercase;
    display:flex;
    gap: 4px;
  }

  .found-results-wrapper{
    box-sizing: border-box;
    padding: 16px;
    border-bottom: 0.1px solid rgba(255, 255, 255,0.3);
  }
  .inner-result-container{
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .found-results{
    margin: 0;
  }

  .found-results-wrapper:hover {
    background-color: rgba(255, 255, 255, 0.3);
    border-radius: 12px;
  }
</style>