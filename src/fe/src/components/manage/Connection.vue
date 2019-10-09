<template>
  <div>
    <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createConnectionDialogVisible = true">新增布线</el-button>
    <el-table
        :data="connections"
        border
        highlight-current-row
        style="width: 100%">
        <el-table-column
            prop="source_device_type"
            label="A端设备类型">
        </el-table-column>
        <el-table-column
            prop="source_device_name"
            label="A端设备">
        </el-table-column>
        <el-table-column
            prop="source_port"
            label="A端设备接线口">
        </el-table-column>
        <el-table-column
            prop="destination_device_type"
            label="B端设备类型">
        </el-table-column>
        <el-table-column
            prop="destination_device_name"
            label="B端设备">
        </el-table-column>
        <el-table-column
            prop="destination_port"
            label="B端设备接线口">
        </el-table-column>
        <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
                <el-button @click="removeConnection(scope.row)" type="danger" plain size="small">删除</el-button>
            </template>
        </el-table-column>
    </el-table>

    <el-dialog title="新增布线" :visible.sync="createConnectionDialogVisible">
        <el-form :model="createConnectionForm">
            <el-form-item label="A端设备" :label-width="formLabelWidth">
                <el-select v-model="createConnectionForm.sourceId" placeholder="请选择">
                    <el-option
                        v-for="item in devices"
                        :key="item.uuid"
                        :label="item.name"
                        :value="item.uuid">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="A端设备接线口" :label-width="formLabelWidth">
                <el-input v-model="createConnectionForm.sourcePort" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="B端设备" :label-width="formLabelWidth">
                <el-select filterable v-model="createConnectionForm.destinationId" placeholder="请选择">
                    <el-option
                        v-for="item in devices"
                        :key="item.uuid"
                        :label="item.name"
                        :value="item.uuid">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="B端设备接线口" :label-width="formLabelWidth">
                <el-input v-model="createConnectionForm.destinationPort" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="createConnectionDialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="createConnection">确 定</el-button>
        </div>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios"
import Config from '../../config'

export default {
  name: 'Connection',
  data () {
      return {
          servers: [],
          networks: [],
          storages: [],
          others: [],
          connections: [],
          config: new Config(),
          formLabelWidth: '120px',
          createConnectionDialogVisible: false,
          createConnectionForm: {
              sourceId: "",
              sourcePort: "",
              destinationId: "",
              destinationPort: "",
          },
      }
  },
  computed: {
      devices: function() {
          var records = []

          for (var server of this.servers) {
              records.push({
                  uuid: server.uuid,
                  name: server.hostname,
                  deviceType: "SERVER"
              })
          }

          for (var device of this.networks) {
              records.push({
                  uuid: device.uuid,
                  name: device.name,
                  deviceType: "NETWORK"
              })
          }

          for (var device of this.storages) {
              records.push({
                  uuid: device.uuid,
                  name: device.name,
                  deviceType: "STORAGE"
              })
          }

          for (var device of this.others) {
              records.push({
                  uuid: device.uuid,
                  name: device.name,
                  deviceType: "COMMON"
              })
          }

          return records
      }
  },
  created () {

  },
  mounted () {
    var that = this
    that.initData()
  },
  methods: {
    initData () {
        var that = this
        that.$store.commit("setPageLoading", true)

        that.connections = []
        that.createConnectionDialogVisible = false

        Promise.all([
            that.syncConnections(),
            that.syncServerDevices(),
            that.syncStorageDevices(),
            that.syncNetworkDevices(),
            that.syncCommonDevices()
        ]).then(values => {
            that.$store.commit("setPageLoading", false)
        }).catch(errors => {
            that.$message({
                type: 'error',
                message: "页面加载异常",
                offset: 200,
            })
            console.error(errors)
            that.$store.commit("setPageLoading", false)
        })
    },
    syncServerDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_SERVERS"))
                    .then(response => {
                        that.servers = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.servers = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    syncStorageDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_STORAGE_DEVICES"))
                    .then(response => {
                        that.storages = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.storages = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    syncNetworkDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_NETWORK_DEVICES"))
                    .then(response => {
                        that.networks = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.networks = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    syncCommonDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_COMMON_DEVICES"))
                    .then(response => {
                        that.others = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.others = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    syncConnections () {
        var that = this
        return axios.post(that.config.getAddress("LIST_CONNECTIONS"))
                    .then(response => {
                        that.connections = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.connections = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    removeConnection(connection) {
        var that = this
        axios.post(that.config.getAddress("DELETE_CONNECTION"), {uuid: connection.uuid})
             .then(response => {
                 that.initData()
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
    getDeviceByUUID(uuid) {
        for (var device of this.devices) {
            if (device.uuid === uuid) {
                return device
            }
        }
        return null
    },
    createConnection () {
        var that = this
        var sourceDevice = that.getDeviceByUUID(that.createConnectionForm.sourceId)
        var destinationDevice = that.getDeviceByUUID(that.createConnectionForm.destinationId)
        axios.post(that.config.getAddress("CREATE_CONNECTION"), {
            source_id: sourceDevice.uuid,
            source_port: that.createConnectionForm.sourcePort,
            source_device_type: sourceDevice.deviceType,
            source_device_name: sourceDevice.name,
            destination_id: destinationDevice.uuid,
            destination_port: that.createConnectionForm.destinationPort,
            destination_device_type: destinationDevice.deviceType,
            destination_device_name: destinationDevice.name,
        })
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    }
    
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
