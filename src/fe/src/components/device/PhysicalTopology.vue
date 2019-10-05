<template>
  <div style="">
    <Datacenter :datacenter="datacenter"/>
  </div>
</template>

<script>
import axios from "axios"
import * as BABYLON from 'babylonjs'

import Config from '../../config'
import Datacenter from '../plugin/Datacenter.vue'


export default {
  name: 'PhysicalTopology',
  data () {
      return {
          // 单位都为绘图单位
          datacenter: {
              size: {
                  height: 150,
                  width: 150,
              },
              racks: [
                  {
                      name: "rack0",
                      x: 12,
                      z: 12,
                      u: 37,
                      servers: [
                          {
                              uuid: "a",
                              name: "server0",
                              begU: 0,
                              sizeU: 2,
                              type: "SERVER",
                              status: "BAD",
                          },
                          {
                              uuid: "a",
                              name: "server1",
                              begU: 12,
                              sizeU: 4,
                              type: "SERVER",
                              status: "BAD",
                          },
                          {
                              uuid: "a",
                              name: "server3",
                              begU: 35,
                              sizeU: 2,
                              type: "SERVER",
                              status: "GOOD",
                          }
                      ]
                  },
                  {
                      name: "rack1",
                      x: 22,
                      z: 22,
                      u: 37,
                      servers: [
                          {
                              uuid: "a",
                              name: "server0",
                              begU: 0,
                              sizeU: 2,
                              type: "SERVER",
                              status: "GOOD",
                          },
                          {
                              uuid: "a",
                              name: "server0",
                              begU: 2,
                              sizeU: 2,
                              type: "SERVER",
                              status: "BAD",
                          },
                          {
                              uuid: "a",
                              name: "server1",
                              begU: 10,
                              sizeU: 4,
                              type: "SERVER",
                              status: "GOOD",
                          },
                          {
                              uuid: "a",
                              name: "server3",
                              begU: 30,
                              sizeU: 2,
                              type: "SERVER",
                              status: "GOOD",
                          }
                      ]
                  },
                  {
                      name: "rack2",
                      x: 22,
                      z: 34,
                      u: 37,
                      servers: [
                          
                      ]
                  },
                  {
                      name: "rack3",
                      x: 0,
                      z: 0,
                      u: 42,
                      servers: [
                          {
                              uuid: "a",
                              name: "server0",
                              begU: 0,
                              sizeU: 2,
                              type: "SERVER",
                              status: "GOOD",
                          },
                          {
                              uuid: "a",
                              name: "server0",
                              begU: 2,
                              sizeU: 2,
                              type: "SERVER",
                              status: "BAD",
                          },
                          {
                              uuid: "a",
                              name: "server1",
                              begU: 10,
                              sizeU: 4,
                              type: "SERVER",
                              status: "GOOD",
                          },
                          {
                              uuid: "a",
                              name: "server3",
                              begU: 30,
                              sizeU: 2,
                              type: "SERVER",
                              status: "GOOD",
                          }
                      ]
                  },
                  {
                      name: "rack4",
                      x: -22,
                      z: 0,
                      u: 42,
                      servers: [
                          {
                              name: "server0",
                              begU: 0,
                              sizeU: 2,
                              type: "SERVER",
                              status: "GOOD",
                          },
                          {
                              name: "server0",
                              begU: 2,
                              sizeU: 2,
                              type: "SERVER",
                              status: "BAD",
                          },
                          {
                              name: "server1",
                              begU: 10,
                              sizeU: 4,
                              type: "SERVER",
                              status: "GOOD",
                          },
                          {
                              name: "server3",
                              begU: 36,
                              sizeU: 4,
                              type: "SERVER",
                              status: "GOOD",
                          }
                      ]
                  }
              ]
          },
      }
  },
  created () {

  },
  components: {
      Datacenter
  },
  mounted () {
    var that = this
    that.initData()
  },
  methods: {
    initData () {
        var that = this
        
        Promise.all([

        ]).then(values => {
            that.$store.commit("setPageLoading", false)
        }).catch(errors => {
            that.$message.error("页面加载异常")
            console.error(errors)
            that.$store.commit("setPageLoading", false)
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
