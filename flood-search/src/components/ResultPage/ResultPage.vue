<script setup>
import ResultHeader from './ResultHeader.vue';
import ResultFooter from './ResultFooter/ResultFooter.vue'
import Result from './Results/Result.vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const term = route.params.term;

</script>

<template>

<div class="wrapper">
    <div class="full-wrapper">
        <ResultHeader :name="term"  />
        <Result />
        <ResultFooter />
    </div>
  </div>
  
</template>

<style scoped>
  .search-title{
    color: white;
  }
</style>
