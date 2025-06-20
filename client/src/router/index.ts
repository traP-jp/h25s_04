import { createRouter, createWebHistory } from 'vue-router'

const IndexPage = () => import('../pages/TopPage.vue')
const EateryPage = () => import('../pages/EateryList.vue')
const NotFound = () => import('../pages/NotFound.vue')
const Addstore = () => import('../pages/AddStore.vue')
const RestaurantPage = () => import('../pages/RestaurantOverview.vue')

const routes = [
  {
    path: '/',
    name: 'Index',
    component: IndexPage,
  },
  {
    path: '/eatery',
    name: 'Eatery',
    component: EateryPage,
  },

  {
    path: '/eatery/new',
    name: 'NewEatery',
    component: Addstore,
  },
  {
    path: '/eatery/:eateryId',
    name: 'RestaurantOverview',
    component: RestaurantPage,
  },
  {
    path: '/:path(.*)*',
    name: 'NotFound',
    component: NotFound,
  },
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
