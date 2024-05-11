<script setup>
import { ExclamationTriangleIcon } from '@heroicons/vue/24/outline';
  defineProps({
    title: String,
    text: String,
  })
</script>

<template>
  <div class="alert-component">
      <ExclamationTriangleIcon class="icon-alert" />
      <div class="alert-content">
        <h4>{{ title }}</h4>
        <p>{{ text }}</p>
      </div>
  </div>

</template>

<style scoped>

.alert-component {
  background-color: #e2e8f039;
  border-radius: 8px;
  padding: 16px;
  display: flex;
  align-items: self-start;
  margin-bottom: 32px;
  gap: 12px;
  border-left: 4px solid #FB923C;
}

.alert-content {
  display: flex;
  flex-direction: column;
  flex: 1 0 0;
}

h4 {
  color: #fff;
  font-size: 14px;
  margin: 0px;
}

p{
  color: #E2E8F0;
  font-size: 14px;
  margin: 0px;
}

.icon-alert {
    height: 24px;
    color: #FB923C;
  }

</style>