<script setup>
import { RouterLink } from 'vue-router';


</script>

<template>

  <div class="footer">
      <div class="footer-section">
        <div class="algolia">
          <a href="https://www.algolia.com/">Busca inteligente por <img src="../assets/Algolia-logo-white.svg"/></a>
        </div>
        <div class="Políticas">

          <RouterLink to=./Privacidade><p>Política de Privacidade</p></RouterLink>

        </div>
      </div>
      <div> 
        <b>Equipe</b>: Gabi Branco, Bibiana Branco, Igor Carrasco, Amanda Medeiros, Douglas Eloy, 
        Adriano Viana, Pedro Vanzella, Lucas Trajano, Henrique Barbosa, Andressa Quines, Maicon Machado, 
        Victor Guimarães, Fernando Martinkovics, Luísa Giusti, Helena Martinelli Serra, Elisa Albuquerque, 
        Carol Goyer.
      </div>
  </div>


</template>


<style scoped>

.footer-section {
  display: flex;
  align-items: center;
  justify-content: space-between;
  
}

.footer img {
  margin: 32px 0px;
  height: 16px;
  }

.algolia a {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #E5E5E5;
  text-decoration: none;
}

.footer {
  display: flex;
  flex-direction: column;
  border-top: 1px solid #1E293B;
}
</style>
