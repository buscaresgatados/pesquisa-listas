<script setup>
import LinkList from '../../LinkList.vue';
import { ChevronLeftIcon } from '@heroicons/vue/16/solid';



</script>

<template>
  <div class="searched-wrapper">
    <a href='HomePage' class="button-voltar"><ChevronLeftIcon class="icon-button"/><p>Voltar</p></a>
    <h4 class="searched-title">Não encontrou quem procurava?</h4>
    <p class="searched-text">Continuamos atualizando nosso banco de dados à medida que novos nomes são adicionados às listas consultadas. Tente novamente mais tarde.</p>
  </div>
  <LinkList />
</template>

<style scoped>


.button-voltar {
    display: flex;
    align-items: center;
    text-decoration: none;
  }

  .button-voltar p {
    font-weight: bold;

  }


.searched-wrapper{
  margin-top: 40px;
  margin-bottom: 40px;
  border-bottom: 0.1px solid rgba(255, 255, 255,0.3);
}
  .searched-title{
    font-size: 24px !important;
    font-weight: 700 !important;
    color: white;
    margin: 0;
  }
  .searched-text{
    color: white;
    font-size: 16px !important;
    font-weight: 400 !important;
    line-height: 24px;
    padding-bottom: 24px;
  }
</style>
