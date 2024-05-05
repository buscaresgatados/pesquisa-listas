<script setup>
import { UIcon } from 'nuxt-ui-vue';

</script>

<template>
  <div class="text-base-bold">Listas consultadas</div>
  <p>Em tempo real, atualizamos nosso banco de dados com os nomes de pessoas resgatadas nas enchentes, com base nas seguintes listas:</p>

  <div class="link-list">
  <a href="">Abrigados em Canoas 01 <UIcon name="i-heroicons-arrow-top-right-on-square-solid"/> </a>
  <a href="">Abrigados em Canoas 01</a>
  </div>

  <div class="text-base-bold">Alguma lista ficou de fora?</div>
  <p>Envie sua sugestão para <a href="mailto:buscadeabrigados@gmail.com">buscadeabrigados@gmail.com</a></p>
</template>

<style scoped>

.link-list {
  display: flex;
  gap: 16px;
  padding-bottom: 24px;
}

.link-list a {
  width: 100%;
}


h3 {
  font-size: 1rem;
  font-weight: bold;
  color: #EFFDF5;
}

p {
  color: #E5E5E5
}

</style>
