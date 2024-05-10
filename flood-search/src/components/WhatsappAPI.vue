<script setup>


    defineProps({
    name: String,
    
  })

  

</script>

<template>
  <div class="whatsapp-wrapper">
    <img src="../assets/whatsapp.svg" />
    <h4 class="whatsapp-title">Receba atualizações sobre {{ name }}</h4>
    <p class="whatsapp-text">
      Enviaremos uma mensagem no WhatsApp caso o nome procurado seja adicionado em nossas listas.
    </p>

    <a target="_blank" :href="`https://wa.me/5554996016629?text=${name}%20-%20receber%20alertas`">
    <button>Receber atualizações</button>
    </a>

  </div>
</template>

<style scoped>

.whatsapp-wrapper{
  margin-top: 40px;
  margin-bottom: 40px;
}
  .whatsapp-title{
    font-size: 24px !important;
    font-weight: 700 !important;
    color: white;
    margin: 0;
  }
  .whatsapp-text{
    color: white;
    font-size: 16px !important;
    font-weight: 400 !important;
    line-height: 24px;
    margin-top: 8px;
    padding-bottom: 8px;
  }

</style>
