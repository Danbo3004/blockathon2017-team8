import Vue from 'vue';
import Vuetify from 'vuetify';
import VueRouter from 'vue-router';
import App from './App.vue';
import HomePage from './pages/HomePage.vue';
import TransferPage from './pages/TransferPage.vue';
import RegisterPage from './pages/RegisterPage.vue';
import ListCompanyPage from './pages/ListCompanyPage.vue';
import BuyTokensPage from './pages/BuyTokensPage.vue';

const routes = [
  {
    path: '/home',
    name: 'home',
    component: HomePage,
  },
  {
    path: '/transfer',
    name: 'transfer',
    component: TransferPage,
  },
  {
    path: '/list_company',
    name: 'list_company',
    component: ListCompanyPage,
  },
  {
    path: '/buy_tokens/:id',
    name: 'buy_tokens',
    component: BuyTokensPage,
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterPage,
  },
  {
    path: '*',
    redirect: {
      name: 'home',
    },
  },
];
const router = new VueRouter({
  routes,
  root: '/home',
});

Vue.use(Vuetify);
Vue.use(VueRouter);

new Vue({ // eslint-disable-line no-new
  el: '#app',
  router,
  render: h => h(App),
});
