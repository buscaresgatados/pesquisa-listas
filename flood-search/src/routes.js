import { createRouter, createWebHistory  } from 'vue-router';
import App from './App.vue'; // Assuming component location
import ResultPage from './components/ResultPage/ResultPage.vue'; // Assuming component location

const history = createWebHistory();

const routes = [
  {
    path: '/',
    component: App,
  },
  {
    path: '/results/:term',
    name: 'results',
    component: ResultPage,
    props: true,
  },
];

const router = createRouter({
  history,
  routes,
});

export default router;
