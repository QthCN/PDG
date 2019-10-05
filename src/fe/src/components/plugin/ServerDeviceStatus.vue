<template>
  <div style="width: 100%; height: 100%;">
      <el-row v-loading="loading">
          <el-col :span="12">
            <el-row style="margin-right: 5px;">
                <el-tabs type="border-card">
                    <el-tab-pane label="基本信息">
                    </el-tab-pane>
                    <el-tab-pane label="位置信息">
                    </el-tab-pane>
                </el-tabs>
            </el-row>
          </el-col>

          <el-col :span="12">
              <el-row>
                <el-tabs type="border-card">
                    <el-tab-pane label="实时告警">
                    </el-tab-pane>
                    <el-tab-pane label="监控数据">
                    </el-tab-pane>
                </el-tabs>
            </el-row>
          </el-col>
      </el-row>
  </div>
</template>

<script>
import axios from "axios"

import Config from '../../config'



export default {
  name: 'ServerDeviceStatus',
  props: ['uuid'],
  data () {
      return {
          loading: true,
      }
  },
  created () {

  },
  components: {
  },
  mounted () {
    var that = this
    that.initData()
  },
  watch: {
      uuid: function(val) {
          this.initData()
      }
  },
  methods: {
    initData () {
        var that = this
        that.loading = true
        if (that.uuid == "") {
            return
        }
        that.loading = false
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
