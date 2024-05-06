<script setup>
import { ArrowTopRightOnSquareIcon } from '@heroicons/vue/24/solid'
    defineProps({
    name: String,
    age: Number,
    shelter: String,
    url: String,
    listName: String,
  })
</script>

<template>
  <div class="found-results-wrapper">
    <div >
      <h3 class="found-results">
        {{ name }}
      </h3>
      <div class="inner-result-container">
        <p>{{ age }} anos, em {{ shelter }}</p>
        <a class="link" :href="url"> {{ listName }} <ArrowTopRightOnSquareIcon class="icon-size" /></a>
      </div>
    </div>
  </div>
</template>

<style scoped>
  .icon-size{
    width: 16px;
  }
  .link{
    text-transform: uppercase;
    text-decoration: none;
    border-bottom: 1px solid #00DC82 !important;
  }
  .found-results-wrapper{
    padding: 8px;
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
    background-color: rgba(255, 255, 255,0.3);
    border-radius: 12px;
  }
</style>
