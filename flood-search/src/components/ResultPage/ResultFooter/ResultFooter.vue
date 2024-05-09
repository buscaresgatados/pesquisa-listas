<script setup>
import LinkList from '../../LinkList.vue';

defineProps({
    name: String,
    showName: Boolean
  })

</script>

<template>
  <div class="searched-wrapper">
    <h4 class="searched-title">Não encontrou quem procurava?</h4>
    <p class="searched-text">
      Isso não quer dizer que a pessoa não foi resgatada! Isso pode significar que ela ainda não foi cadastrada e/ou a planilha ainda não foi importada.
      Estamos trabalhando continuamente para atualizar nossa base de dados.
    </p>
  </div>
  <WhatsappAPI  v-if="showName" :name="name" />
  <LinkList />
</template>

<style scoped>

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
