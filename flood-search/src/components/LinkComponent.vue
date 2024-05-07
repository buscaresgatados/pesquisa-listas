<script setup>
import { ArrowTopRightOnSquareIcon } from '@heroicons/vue/24/outline';
import TagComponent from './microcomponents/TagComponent.vue'
  defineProps({
    url: String,
    name: String,
    tag: Boolean,
    title: String
  })
</script>

<template>
  <div class="link-component">
    <a :href="url" target="_blank">{{  name  }}<ArrowTopRightOnSquareIcon class="icon-size" /></a>
    <TagComponent v-if="tag" :title="title" />
  </div>

</template>

<style scoped>
  .link-component{
    display: flex;
    gap: 4px;
    align-items: center;
    padding: 2px 0px;
  }

  a {
    display:flex;
    gap: 4px;
    color: #E5E5E5;
  }

</style>