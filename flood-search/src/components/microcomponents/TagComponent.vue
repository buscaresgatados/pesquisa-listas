<script setup>
  defineProps({
    title: String
    })
</script>

<template>
  <div class="tag">{{ title }}</div>
</template>

<style scoped>
  .tag {
    text-transform: uppercase;
  }
</style>