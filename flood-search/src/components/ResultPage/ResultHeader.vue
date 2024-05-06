<script setup>

    defineProps({
    name: String,
  })
</script>

<template>
  <div>
    <h5 class="result-header__title ">Buscando por</h5>
    <h2 class="result-header__name">{{ name }}</h2>
  </div>
  
</template>

<style scoped>

  .result-header__title{
    color: white;
    font-size: 16px;
    line-height: 24px;
    font-weight: 400;
    margin:0;
  }
  .result-header__name{
    color: #00DC82;
    font-weight: 800;
    font-size: 30px;
    margin: 0;
  }
</style>