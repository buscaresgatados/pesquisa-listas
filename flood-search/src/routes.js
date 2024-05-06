import { createRouter, createMemoryHistory } from 'vue-router';
import Homepage from './components/HomePage.vue'
import ResultPage from './components/ResultPage/ResultPage.vue';

const history = createMemoryHistory();

const routes = [
  {
    path: '/',
    component: Homepage,
  },
  {
    path: '/results/:term',
    name: 'results',
    component: () => ResultPage,
    props: true
  }
];

const router = createRouter({
  history,
  routes,
});

export default router;
