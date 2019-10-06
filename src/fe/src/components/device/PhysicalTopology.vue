<template>
  <div style="">
    <Datacenter :datacenter="datacenter" :key="dcKey"/>
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
          config: new Config(),
          dcKey: 0,
          // 单位都为绘图单位
          datacenter: {
            size: {
                  height: 150,
                  width: 150,
            },
            racks: [],
        },
      }
  },
  created () {
    var that = this
    that.initData()
  },
  components: {
      Datacenter
  },
  mounted () {
    
  },
  methods: {
    initData () {
        var that = this
        that.datacenter = {
            size: {
                  height: 150,
                  width: 150,
            },
            racks: [],
        }
        Promise.all([
            that.syncDatacenter()
        ]).then(values => {
            that.$store.commit("setPageLoading", false)
            that.dcKey += 1
        }).catch(errors => {
            that.$message.error("页面加载异常")
            console.error(errors)
            that.$store.commit("setPageLoading", false)
        })
    },
    syncDatacenter () {
        var that = this
        return axios.post(that.config.getAddress("GET_PHYSICAL_TOPOLOGY"), {})
                    .then(response => {
                        that.datacenter = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.$message.error("获取数据异常")
                    })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
