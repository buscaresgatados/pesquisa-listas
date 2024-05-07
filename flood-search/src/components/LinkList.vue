<script setup>
import { ArrowTopRightOnSquareIcon, DocumentIcon, ExclamationCircleIcon } from '@heroicons/vue/24/outline';
import LinkComponent from './LinkComponent.vue';
</script>

<template>
  <div class="wrapper-list">

      <div class="section">
        <DocumentIcon class="icon-decoration"/>
        <div class="text-base-bold">Listas consultadas</div>
        <p>Em tempo real, atualizamos nosso banco de dados com os nomes de pessoas resgatadas nas enchentes, com base nas seguintes listas:</p>

      <div class="link-list">
          <div class="link-list--col">
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1Pd8NVuEtnR7-IlLF7cJ3XY7yVSoJMgY47-eepe2BBXo/edit#gid=156688726" name="RESGATADOS (Cruzeiro do Sul)"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1yuzazWMydzJKUoBnElV1YTxSKLJsT4fSVHfyJBjLlAY/edit#gid=0" name="SESC PROTÁSIO"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1hKJVs-RLiSUpx-1Rd9wS1k8RLqxkPWK4hNob-t8v2Ko/edit#gid=0" name="LISTA RESGATADOS ELDORADO"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1T_yd-M6BG1qYdQKeMo2U_AffqRCxkExqpB39iQXig5s/htmlview" name="ULBRA - CANOAS PRÉDIO 11"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1-1q4c8Ns6M9noCEhQqBE6gy3FWUv-VQgeUO9c7szGIM/htmlview" name="ABRIGADOS EM CANOAS 01"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1zt_yrzvU2nmihyZG7rR67iqlkwlcDz3LjSb1UBynQ3c/edit#gid=0" name="SAO LEOPOLDO - LISTA ALOJ..."/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1f5gofOOv4EFYWhVqwPWbgF2M-7uHrJrCMiP7Ug4y6lQ/htmlview#gid=1619683963" name="Lista Desabrigados" />
            <LinkComponent url="https://www.instagram.com/stories/highlights/18055125328538288/" name="Abrigados Estrela" tag title="NOVA"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1pwL_AAO-7H7-cXujB_aFhkMLHbiH7xqH/edit#gid=874938892" name="ABRIGADOS LAR MÃE ESTHER"/>
          </div>

          <div class="link-list--col">
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1bNw-t0RUE-AP-2quCU80w5meGYVtD7WjjuESb7tXxTo/edit#gid=0" name="Abrigados Lajeado"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1-xhmS1VQ95LFI05XG8o9JO3mPk8KxQtrxAZe4GNYO3I/htmlview?fbclid=PAZXh0bgNhZW0CMTEAAaa8hrVtObdF2fVAKqXLDmpSNIxwrdxrYeR67gCKkLLjjmmYX_uyO3mwB5Q_aem_AYGIoQtWukDR9BFdDp29TRY_mG75KmR5EbnNgeh08Il9XLI1z2UPRwhYD94Gttg1zSkakpKoGViZoK5vtA7pT8Iz" name="Resgatados Pontal"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1RGRoIzSFQaaJF1xZsJhQsMJxXnXWzfZfas29T_PefmY/edit#gid=904509716" name="ALOJAMENTOS DA PMNH"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1Gf78W5yY0Yiljg-E0rYqbRjxYmBPcG2BtfpGwFk-K5M/edit#gid=0" name="LISTA GERAL EM ATUALIZ..."/>
            <LinkComponent url="https://docs.google.com/spreadsheets/u/0/d/1frgtJ9eK05OqsyLwOBiZ2Q6E7e4_pWyrb7fJioqfEMs/htmlview" name="RESGATADOS - Enchente RS"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1hxzHYE4UR1YbcH3ZQoPfcQTgPbRm5T6lShkDwGDoeXA/edit#gid=0" name="To Salvo Vale dos Sinos" />
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1IjefUU8g5OaHxMfcvqc0ewyQDn5ZVLBE/edit#gid=747398008" name="ABRIGADOS EM GUAÍBA"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/1IVtSmKRFynQH9I9Cox93YxZe0uwKfjx_CYFzKE96its/edit#gid=2046654110" name="ABRIGO COELHÃO" tag title="NOVA"/>
            <LinkComponent url="https://docs.google.com/spreadsheets/d/16X-68-x7My4u0WEfscL7t4YYw_Ebeco6gaLhE80Q8Wc/edit#gid=0" name="Viaduto da Dom Pedro" tag title="NOVA"/>
          </div>
        </div>
      </div>

      <div class="section">
        <ExclamationCircleIcon class="icon-decoration" />
        <div class="text-base-bold">Alguma lista ficou de fora?</div>
        <p> Estamos trabalhando para adicionar essas listas ao banco de dados: </p>

          <div class="link-list">
            <div class="link-list--col">
              <LinkComponent url="https://docs.google.com/spreadsheets/d/1xaEPlk8JonATIOAvQEc0Dev-QVAzx2AwUzLHBhbA3rI/edit#gid=0" name="Eldorado - Guaiba" tag title="EM PROGRESSO"/>
              <LinkComponent url="https://docs.google.com/spreadsheets/d/1Pd8NVuEtnR7-IlLF7cJ3XY7yVSoJMgY47-eepe2BBXo/edit#gid=156688726" name="RESGATADOS DO VALE DO TAQUARI"/>

            </div>

            <div class="link-list--col">
              <LinkComponent url="https://docs.google.com/spreadsheets/d/1wvtgK7ZO9KuJsFDI9syyPWmEyqYoKw2PKssmgfo_jCU/edit#gid=446068527" name="LOCALIZAÇÃO DESABRIGADOS CANOAS"/>
            </div>
          </div>


        <p>Envie sua sugestão para <a href="mailto:buscaresgatados@gmail.com">buscaresgatados@gmail.com</a></p>
      </div>


      <Footer></Footer>
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
}


</style>
