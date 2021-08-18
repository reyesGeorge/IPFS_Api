import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import IPFSAPI from '../views/IPFSAPI.vue'


Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/ipfs',
    name: 'IPFS',
    component: IPFSAPI
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
  }
]

const router = new VueRouter({
  routes
})

export default router
