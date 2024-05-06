<script setup>
import { ArrowTopRightOnSquareIcon, DocumentIcon, ExclamationCircleIcon } from '@heroicons/vue/24/outline';

</script>

<template>
  <div class="wrapper-list">

      <div class="section">
        <DocumentIcon class="icon-decoration"/>
        <div class="text-base-bold">Listas consultadas</div>
        <p>Em tempo real, atualizamos nosso banco de dados com os nomes de pessoas resgatadas nas enchentes, com base nas seguintes listas:</p>

        <div class="link-list">
          <div class="link-list--col">
            <a href="https://docs.google.com/spreadsheets/d/1f5gofOOv4EFYWhVqwPWbgF2M-7uHrJrCMiP7Ug4y6lQ/htmlview#gid=1619683963">Lista Desabrigados<ArrowTopRightOnSquareIcon class="icon-size" /></a>
            <a href="https://docs.google.com/spreadsheets/d/1hxzHYE4UR1YbcH3ZQoPfcQTgPbRm5T6lShkDwGDoeXA/edit#gid=0">To Salvo Vale dos Sinos <ArrowTopRightOnSquareIcon class="icon-size" /> </a>
          </div>
          <div class="link-list--col">
            <a href="https://docs.google.com/spreadsheets/d/1f5gofOOv4EFYWhVqwPWbgF2M-7uHrJrCMiP7Ug4y6lQ/htmlview#gid=1619683963">Lista Desabrigados<ArrowTopRightOnSquareIcon class="icon-size" /></a>
            <a href="https://docs.google.com/spreadsheets/d/1hxzHYE4UR1YbcH3ZQoPfcQTgPbRm5T6lShkDwGDoeXA/edit#gid=0">To Salvo Vale dos Sinos <ArrowTopRightOnSquareIcon class="icon-size" /> </a>
          </div>
        </div>
      </div>

      <div class="section">
        <ExclamationCircleIcon class="icon-decoration" />
        <div class="text-base-bold">Alguma lista ficou de fora?</div>
        <p>Envie sua sugestão para <a href="mailto:buscadeabrigados@gmail.com">buscadeabrigados@gmail.com</a></p>
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

.link-list a {
  width: 100%;
  display: flex;
  gap: 4px;
  line-height: 32px;
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
