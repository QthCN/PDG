import Vue from 'vue'
import Router from 'vue-router'
import User from './components/manage/User.vue'
import Device from './components/manage/Device.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/r/manage/u',
      name: 'User',
      component: User
    },
    {
      path: '/r/manage/d',
      name: 'Device',
      component: Device
    }
  ]
})
