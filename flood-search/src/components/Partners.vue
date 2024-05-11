<script setup>
import { ArrowTopRightOnSquareIcon, DocumentIcon, ExclamationCircleIcon } from '@heroicons/vue/24/outline';
import LinkComponent from './LinkComponent.vue';
</script>

<template>
  <div class="wrapper-list">

      <div class="section">
        <p>Estamos trabalhando em parceria com as seguintes iniciativas:</p>

      <div class="link-list">
          <div class="link-list--col">
            <LinkComponent url="https://estou-salvo-rs.bubbleapps.io/" name="Estou Salvo RS" />
            <LinkComponent url="https://www.tosalvo.ong.br/" name="Tô Salvo ONG" />
            <LinkComponent url="https://petmapa.com.br/" name="PetMapa" />
         
          </div>


          <div class="link-list--col">
            <LinkComponent url="https://salvars.com.br/" name="Salva RS" />
            <LinkComponent url="https://pessoaperdida.com.br/" name="Pessoa Perdida"/>
            <LinkComponent url="https://sosrs.net.br/" name="SOS RS"/>
 
          </div>
        </div>
      </div>

  </div>
</template>

<style scoped>

.section {
  padding: 16px 0px;
}

.link-list {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 24px;
}


h3 {
  font-size: 1rem;
  font-weight: bold;
  color: #EFFDF5;
}

p {
  color: #E5E5E5
}

.wrapper-list {
  padding-top: 0px;
  padding-bottom: 40px;
}

@media (max-width: 480px) 
  {
    .link-list {
      flex-direction: column;
      gap: 0px;
    }

  }


.link-list--col {
  display: flex;
  flex-direction: column;
  flex-basis: auto;
}


</style>
