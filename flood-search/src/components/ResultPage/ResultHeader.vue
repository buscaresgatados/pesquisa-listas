<script setup>
import { ChevronLeftIcon } from '@heroicons/vue/16/solid';
import { RouterLink } from 'vue-router';

    defineProps({
    name: String,
  })
</script>

<template>
  <div>
    <RouterLink to="/" class="button-voltar"><ChevronLeftIcon class="icon-button"/><p>Voltar</p></RouterLink>
    <h5 class="result-header__title ">Buscando por</h5>
    <h2 class="result-header__name">{{ name }}</h2>
  </div>
  
</template>

<style scoped>

  .button-voltar {
    display: flex;
    align-items: center;
    text-decoration: none;
    padding-bottom: 16px;
    color: #CBD5E1;
  }

  .button-voltar p {
    font-weight: bold;

  }

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