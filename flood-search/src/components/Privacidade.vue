<script setup>
import Footer from './Footer.vue';
import Partners from './Partners.vue';
import { ChevronLeftIcon } from '@heroicons/vue/16/solid';

</script>

<template>
  <div class="wrapper">
    <div class="full-wrapper">
      <RouterLink to="/" class="button-voltar"><ChevronLeftIcon class="icon-button"/><p>Voltar</p></RouterLink>
      <div class="header">
        <h1 class="text-3xl-bold">Política de Privacidade</h1>
        <p>A coleta e tratamento de dados pessoais realizada por esta software, incluindo nome e CPF, dentre outros dados, para a finalidade de auxiliar nas buscas de vítimas de tragédias, está em conformidade com a Lei Geral de Proteção de Dados (LGPD), Lei nº 13.709/2018, que estabelece as bases legais para o tratamento de dados pessoais. Segundo o Art. 7º, VII, da LGPD, o tratamento de dados é permitido para a proteção da vida ou da incolumidade física do titular ou de terceiro. Portanto, o tratamento de dados pessoais realizado por este software está alinhado com os princípios e bases legais da LGPD, uma vez que em situações de desastres naturais, como a enchente do Rio Grande do Sul, a utilização desses dados torna-se essencial para a localização e o resgate de indivíduos afetados, justificando-se pela necessidade de salvaguardar a vida humana.</p>
      </div>  
      <Partners />
      <Footer />
    </div>
    </div>
</template>

<style scoped>

</style>
