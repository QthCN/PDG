import Vue from 'vue'
import Router from 'vue-router'
import User from './components/manage/User.vue'
import Connection from './components/manage/Connection.vue'
import Device from './components/manage/Device.vue'
import PhysicalTopology from './components/device/PhysicalTopology.vue'
import IpAndIpSet from './components/manage/IpAndIpSet.vue'
import ResourceTopology from './components/device/ResourceTopology.vue'
import NetworkTopology from './components/device/NetworkTopology.vue'



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
    },
    {
      path: '/r/manage/l',
      name: 'Connection',
      component: Connection
    },
    {
      path: '/r/device/p',
      name: 'PhysicalTopology',
      component: PhysicalTopology
    },
    {
      path: '/r/device/l',
      name: 'ResourceTopology',
      component: ResourceTopology
    },
    {
      path: '/r/device/n',
      name: 'NetworkTopology',
      component: NetworkTopology
    },
    {
      path: '/r/manage/i',
      name: 'IpAndIpSet',
      component: IpAndIpSet
    }
  ]
})
