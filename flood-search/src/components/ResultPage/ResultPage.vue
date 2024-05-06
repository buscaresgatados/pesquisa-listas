<script setup>
import ResultHeader from './ResultHeader.vue';
import ResultFooter from './ResultFooter/ResultFooter.vue'
import Result from './Results/Result.vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const term = route.params.term;

</script>

<template>
  <ResultHeader :name="term"  />
  <Result />
  <ResultFooter />
</template>

<style scoped>
  .search-title{
    color: white;
  }
</style>
