<script setup>
import ResultHeader from './ResultHeader.vue';
import ResultFooter from './ResultFooter/ResultFooter.vue'
import Result from './Results/Result.vue';

defineProps({
    search: String,
  })
</script>

<template>
  <ResultHeader :name="search"  />
  <Result />
  <ResultFooter />
</template>

<style scoped>
  .search-title{
    color: white;
  }
</style>
