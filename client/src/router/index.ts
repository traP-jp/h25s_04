import { createRouter, createWebHistory } from 'vue-router'

const IndexPage = () => import('../pages/TopPage.vue')
const NotFound = () => import('../pages/NotFound.vue')

const routes = [
  {
    path: '/',
    name: 'Index',
    component: IndexPage,
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
