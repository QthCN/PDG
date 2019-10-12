<template>
  <div style="width: 100%; height: 100%;">
      <el-row v-loading="loading">
          <el-col :span="12">
            <el-row style="margin-right: 5px;">
                <el-tabs type="border-card">
                    <el-tab-pane label="基本信息">
                    </el-tab-pane>
                </el-tabs>
            </el-row>
          </el-col>

          <el-col :span="12">
              <el-row>
                <el-tabs type="border-card">
                    <el-tab-pane label="监控数据">
                      <el-row>
                        <LinePic :title="'CPU监控数据'" :records="cpuMonitorRecords"></LinePic>
                      </el-row>
                        <LinePic :title="'内存监控数据'" :records="memoryMonitorRecords"></LinePic>
                      <el-row>
                        
                      </el-row>
                    </el-tab-pane>
                    <el-tab-pane label="实时告警">
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
import LinePic from './LinePic.vue'



export default {
  name: 'DeviceStatus',
  props: ['uuid', 'deviceType'],
  data () {
      return {
          config: new Config(),
          loading: true,
          // 监控项信息
          monitorItems: [],
          // 基本监控数据类型: CPU/MEMORY
          cpuMonitorRecords: [],
          memoryMonitorRecords: [],
      }
  },
  created () {
    var that = this
    that.initData()
  },
  components: {
    LinePic
  },
  mounted () {

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

        Promise.all([
            that.syncMonitorItems()
        ]).then(values => {
            Promise.all([
                that.syncCPUMonitorRecords(),
                that.syncMemoryMonitorRecords()
            ]).then(values => {
                that.loading = false
            }).catch(errors => {
                that.$message({
                    type: 'error',
                    message: "数据加载异常",
                    offset: 200,
                })
                console.error(errors)
            })
        }).catch(errors => {
            that.$message({
                type: 'error',
                message: "数据加载异常",
                offset: 200,
            })
            console.error(errors)
        })
    },
    syncMonitorItems () {
        var that = this
        return axios.post(that.config.getAddress("LIST_MONITOR_ITEMS"))
                    .then(response => {
                        that.monitorItems = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.monitorItems = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    AddMinutesToDate(date, minutes) {
      return new Date(date.getTime() + minutes*60000);
    },
    DateFormat(date){
      var days = date.getDate();
      days = days < 10 ? '0' + days : days;
      var year = date.getFullYear();
      var month = (date.getMonth()+1);
      month = month < 10 ? '0' + month : month;
      var hours = date.getHours();
      hours = hours < 10 ? '0' + hours : hours;
      var minutes = date.getMinutes();
      minutes = minutes < 10 ? '0' + minutes : minutes;
      var strTime = year + '-' + month + '-' + days + ' ' + hours + ':' + minutes + ':00'
      return strTime;
    },
    getMonitorItemByName (itemName) {
      for (var item of this.monitorItems) {
        if (item.name === itemName) {
          return item
        }
      }
      return null
    },
    syncCPUMonitorRecords () {
      var that = this
      var monitorItem = that.getMonitorItemByName("CPU使用率")
      if (monitorItem === null) {
        that.cpuMonitorRecords = []
        return 
      }
      return axios.post(that.config.getAddress("GET_MONITOR_HISTORY_RECORDS"), {
        device_uuid: that.uuid,
        item_id: monitorItem.id,
        query_begin_date: that.DateFormat(that.AddMinutesToDate(new Date(), -30)),
        query_end_date: that.DateFormat(that.AddMinutesToDate(new Date(), 0)),
      })
                  .then(response => {
                      that.cpuMonitorRecords = response.data
                  })
                  .catch(error => {
                      console.error(error)
                      that.$message({
                          type: 'error',
                          message: error.response.data.msg,
                          offset: 200,
                      })
                  })
    },
    syncMemoryMonitorRecords () {
      var that = this
      var monitorItem = that.getMonitorItemByName("内存使用率")
      if (monitorItem === null) {
        that.memoryMonitorRecords = []
        return 
      }
      return axios.post(that.config.getAddress("GET_MONITOR_HISTORY_RECORDS"), {
        device_uuid: that.uuid,
        item_id: monitorItem.id,
        query_begin_date: that.DateFormat(that.AddMinutesToDate(new Date(), -30)),
        query_end_date: that.DateFormat(that.AddMinutesToDate(new Date(), 0)),
      })
                  .then(response => {
                      that.memoryMonitorRecords = response.data
                  })
                  .catch(error => {
                      console.error(error)
                      that.$message({
                          type: 'error',
                          message: error.response.data.msg,
                          offset: 200,
                      })
                  })
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
