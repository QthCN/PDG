import Vue from 'vue'
import Router from 'vue-router'
import User from './components/manage/User.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/r/manage/u',
      name: 'User',
      component: User
    }
  ]
})
