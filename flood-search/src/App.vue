<script setup>
import { RouterView } from 'vue-router';

</script>

<template>
  <div class="background">
    <div class="full-wrapper">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
  .background{
    width: 100vw;
    height: 100vh;
  }
  .full-wrapper{
    margin: 0 auto;
    max-width: 650px;
  }
</style>
