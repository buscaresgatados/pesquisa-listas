<script setup>
  import { UInput } from 'nuxt-ui-vue';
</script>

<template>
  <p>Nome completo</p>
  <div class="cta">
    <UInput class="UInput" autocomplete="off" placeholder="Buscar por nome" />
    <button>Buscar</button>
  </div>
</template>

<style scoped>
  p {
    margin-bottom: 2px;
  }

  .cta {
    display: flex;
    gap: 16px;
    width: 100R%;
  }

  .relative {
    width: 100%;
  }
</style>