import { createRouter, createWebHistory } from 'vue-router'

const IndexPage = () => import('../pages/TopPage.vue')
const EateryPage = () => import('../pages/EateryList.vue')
const NotFound = () => import('../pages/NotFound.vue')
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
    path: '/:path(.*)*',
    name: 'NotFound',
    component: NotFound,
  },
  {
    path: '/eataly/:eateryid',
    name: 'RestaurantOverview',
    component: RestaurantPage,
  }
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
