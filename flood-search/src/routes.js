import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    component: () => import('./components/HomePage.vue')
  },
  {
    path: '/results/:term',
    name: 'results',
    component: () => import('./components/ResultPage/ResultPage.vue'),
    props: true
  }
];

const router = new VueRouter({
  routes
});

new Vue({
  router,
  el: '#app'
});