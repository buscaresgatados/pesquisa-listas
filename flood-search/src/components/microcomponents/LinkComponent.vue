<script setup>
import { ArrowTopRightOnSquareIcon } from '@heroicons/vue/24/outline';

  defineProps({
    url: String,
    listname: String,
    tag: Boolean,
    title: String,
  })

</script>

<template>

  <a :href="url" target="_blank"> {{ listname }} <ArrowTopRightOnSquareIcon class="icon-size" /></a>

</template>


<style scoped>

a {
  width: 100%;
  display: flex;
  gap: 4px;
  line-height: 32px;
  align-items: center;
}


</style>