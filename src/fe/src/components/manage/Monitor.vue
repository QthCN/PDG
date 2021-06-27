<template>
  <div>
        <el-tabs type="border-card">
                <el-tab-pane label="监控项管理">
                    <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createMonitorItemDialogShow = true">新建监控项</el-button>
                    <el-table
                        :data="monitorItems"
                        border
                        highlight-current-row
                        style="width: 100%">
                        <el-table-column
                            prop="name"
                            label="名称">
                        </el-table-column>
                        <el-table-column
                            prop="dc_type"
                            label="数据收集模块">
                        </el-table-column>
                        <el-table-column
                            fixed="right"
                            label="操作"
                            width="400">
                            <template slot-scope="scope">
                                <el-button @click="configDC(scope.row)" type="primary" plain size="small">配置数据收集模块</el-button>
                                <el-button @click="bindDevice(scope.row)" type="primary" plain size="small">关联设备</el-button>
                                <el-button @click="removeMonitorItem(scope.row)" type="danger" plain size="small">删除</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>

                <el-tab-pane label="告警事件管理">
                    <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createAlertItemDialogShow = true">新建告警项</el-button>
                    <el-table
                        :data="alerts"
                        border
                        highlight-current-row
                        style="width: 100%">
                        <el-table-column
                            prop="item_name"
                            label="名称">
                        </el-table-column>
                        <el-table-column
                            prop="alert_type"
                            label="告警事件源">
                        </el-table-column>
                        <el-table-column
                            prop="event_id"
                            label="事件ID">
                        </el-table-column>
                        <el-table-column
                            fixed="right"
                            label="操作"
                            width="400">
                            <template slot-scope="scope">
                                <el-button @click="removeAlertItem(scope.row)" type="danger" plain size="small">删除</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>

                <el-tab-pane label="监控服务管理">
                    <el-collapse>
                        <el-collapse-item title="ZABBIX" name="ZABBIX">
                            <el-form :model="updateMonitorBackendZabbixForm">
                                <el-form-item label="Zabbix地址" :label-width="formLabelWidth">
                                    <el-input v-model="updateMonitorBackendZabbixForm.address" autocomplete="off"></el-input>
                                </el-form-item>
                                <el-form-item label="用户名" :label-width="formLabelWidth">
                                    <el-input v-model="updateMonitorBackendZabbixForm.username" autocomplete="off"></el-input>
                                </el-form-item>
                                <el-form-item label="密码" :label-width="formLabelWidth">
                                    <el-input 
                                    type="password" v-model="updateMonitorBackendZabbixForm.password" autocomplete="off"></el-input>
                                </el-form-item>
                            </el-form>

                                <el-button type="primary" plain style="float: right;" @click="updateMonitorBackendZabbixCfg">更 新</el-button>
                                <br/><br/>
                        </el-collapse-item>
                    </el-collapse>
                </el-tab-pane>
        </el-tabs>

    <el-dialog title="新建告警事件" :visible.sync="createAlertItemDialogShow">
        <el-form :model="createAlertItemForm">
            <el-form-item label="名称" :label-width="formLabelWidth">
                <el-input v-model="createAlertItemForm.item_name" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="事件源" :label-width="formLabelWidth">
                <el-select v-model="createAlertItemForm.alert_type" placeholder="请选择">
                    <el-option value="FAKE">FAKE</el-option>
                    <el-option value="ZABBIX">ZABBIX</el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="事件ID" :label-width="formLabelWidth">
                <el-input v-model="createAlertItemForm.event_id" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="createAlertItemDialogShow = false">取 消</el-button>
            <el-button type="primary" @click="createAlertItem">确 定</el-button>
        </div>
    </el-dialog>

    <el-dialog title="新建监控项" :visible.sync="createMonitorItemDialogShow">
        <el-form :model="createMonitorItemForm">
            <el-form-item label="名称" :label-width="formLabelWidth">
                <el-input v-model="createMonitorItemForm.name" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="createMonitorItemDialogShow = false">取 消</el-button>
            <el-button type="primary" @click="createMonitorItem">确 定</el-button>
        </div>
    </el-dialog>

    <el-dialog title="配置数据收集模块" :visible.sync="editDCDialogShow">
        <h3>监控项 - {{editDCItem.name}}
        <el-popover
                title="标题"
                width="200"
                trigger="hover"
                content="${ip} - 设备IP"
                style="float: right;"
            >
            <el-button type="primary" plain slot="reference">变量信息</el-button>
        </el-popover>
        </h3>
        <el-form :model="editDCItem">
            <el-form-item label="监控服务类型" :label-width="formLabelWidth">
                <el-select v-model="editDCItem.dc_type" placeholder="请选择">
                    <el-option value="FAKE">FAKE</el-option>
                    <el-option value="ZABBIX">ZABBIX</el-option>
                </el-select>
            </el-form-item>

            <template v-if="editDCItem.dc_type === 'FAKE'">
                <el-form-item label="Fake监控项名" :label-width="formLabelWidth">
                    <el-input v-model="editDCItem.dc_fake_cfg.item_name" autocomplete="off"></el-input>
                </el-form-item>

                <el-form-item label="Host IP" :label-width="formLabelWidth">
                    <el-input v-model="editDCItem.dc_fake_cfg.host_ip" autocomplete="off"></el-input>
                </el-form-item>
            </template>

            <template v-if="editDCItem.dc_type === 'ZABBIX'">
            </template>

        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="editDCDialogShow = false">取 消</el-button>
            <el-button type="primary" @click="doConfigDC">确 定</el-button>
        </div>
    </el-dialog>

    <el-dialog title="绑定设备" :visible.sync="bindDeviceDialogShow">

        <el-transfer
                filterable
                filter-placeholder="输入关键字进行过滤"
                :titles="titles"
                :data="allDevices"
                v-model="bindDevices"
            >
        </el-transfer>

        <div slot="footer" class="dialog-footer">
            <el-button @click="bindDeviceDialogShow = false">取 消</el-button>
            <el-button type="primary" @click="doBindDevice">确 定</el-button>
        </div>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios"
import Config from '../../config'

export default {
  name: 'Monitor',
  data () {
      return {
          config: new Config(),
          titles: ["所有设备", "已绑定设备"],
          devices: [],
          alerts: [],
          monitorItems: [],
          createMonitorItemDialogShow: false,
          createAlertItemDialogShow: false,
          formLabelWidth: '120px',
          createMonitorItemForm: {
              name: "",
          },
          editDCDialogShow: false,
          editDCItem: {
              id: 0,
              name: "",
              is_internal: 0,

              dc_type: "",
              dc_fake_cfg: {
                  item_name: "",
                  host_ip: "",
              },
          },
          bindDeviceDialogShow: false,
          bindDeviceMonitorItemId: 0,
          bindDeviceMonitorItemName: "",
          bindDevices: [],
          monitorBackendCfgs: [],
          updateMonitorBackendZabbixForm: {
              address: "",
              username: "",
              password: "",
          },
          createAlertItemForm: {
              item_name: "",
              alert_type: "",
              event_id: ""
          }
      }
  },
  computed: {
      allDevices: function () {
          var records = []
          for (var device of this.devices) {
              records.push({
                  label: device.name,
                  key: device.uuid,
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

        that.monitorItems = []
        that.createMonitorItemDialogShow = false
        that.editDCDialogShow = false
        that.bindDeviceDialogShow = false
        that.createAlertItemDialogShow = false

        Promise.all([
            that.syncAlertItems(),
            that.syncMonitorItems(),
            that.syncMonitorBackendCfgs(),
            that.syncDevices()
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
    syncMonitorBackendCfgs () {
        var that = this
        return axios.post(that.config.getAddress("LIST_MONITOR_BACKEND_CFGS"))
                    .then(response => {
                        that.monitorBackendCfgs = response.data
                        for (var cfg of that.monitorBackendCfgs) {
                            if (cfg.name === "ZABBIX") {
                                that.updateMonitorBackendZabbixForm = cfg.cfg_zabbix
                            } else {

                            }
                        }
                    })
                    .catch(error => {
                        console.error(error)
                        that.monitorBackendCfgs = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    syncAlertItems () {
        var that = this
        return axios.post(that.config.getAddress("LIST_ALERT_ITEMS"))
                    .then(response => {
                        that.alerts = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.alerts = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    syncDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_DEVICES"))
                    .then(response => {
                        that.devices = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.devices = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
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
    removeMonitorItem (monitorItem) {
        var that = this
        axios.post(that.config.getAddress("DELETE_MONITOR_ITEM"), {id: monitorItem.id})
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
    createMonitorItem () {
        var that = this
        axios.post(that.config.getAddress("CREATE_MONITOR_ITEM"), that.createMonitorItemForm)
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
    createAlertItem () {
        var that = this
        axios.post(that.config.getAddress("CREATE_ALERT_ITEM"), that.createAlertItemForm)
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
    configDC (monitorItem) {
        var that = this
        that.editDCDialogShow = true
        Promise.all([
            that.getMonitorItem(monitorItem.id)
        ]).then(values => {
            that.editDCDialogShow = true
        }).catch(errors => {
            that.$message({
                type: 'error',
                message: "页面加载异常",
                offset: 200,
            })
            console.error(errors)
        })
    },
    getMonitorItem (id) {
        var that = this
        return axios.post(that.config.getAddress("GET_MONITOR_ITEM"), {id: id})
                    .then(response => {
                        that.editDCItem = response.data
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
    doConfigDC () {
        var that = this
        axios.post(that.config.getAddress("UPDATE_MONITOR_ITEM_DC_CFG"), {
            id: that.editDCItem.id,
            dc_type: that.editDCItem.dc_type,
            dc_fake_cfg_item_name: that.editDCItem.dc_fake_cfg.item_name,
            dc_fake_cfg_host_ip: that.editDCItem.dc_fake_cfg.host_ip,
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
    },
    listMonitorItemReleatedDevices (itemId) {
        var that = this
        that.bindDevices = []
        return axios.post(that.config.getAddress("LIST_MONITOR_ITEM_RELEATED_DEVICES"), {id: itemId})
                    .then(response => {
                        that.bindDevices = []
                        for (var bd of response.data) {
                            that.bindDevices.push(bd.device_uuid)
                        }
                    })
                    .catch(error => {
                        console.error(error)
                        that.bindDevices = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    bindDevice(monitorItem) {
        var that = this
        that.bindDeviceMonitorItemId = monitorItem.id
        that.bindDeviceMonitorItemName = monitorItem.name
        Promise.all([
            that.listMonitorItemReleatedDevices(monitorItem.id)
        ]).then(values => {
            that.bindDeviceDialogShow = true
        }).catch(errors => {
            that.$message({
                type: 'error',
                message: "页面加载异常",
                offset: 200,
            })
            console.error(errors)
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
    doBindDeviceAndMonitorItem (itemId, itemName, deviceUUID, deviceType, deviceName) {
        var that = this
        axios.post(that.config.getAddress("BIND_MONITOR_ITEM_AND_DEVICE"), {
            item_id: itemId,
            item_name: itemName,
            device_uuid: deviceUUID,
            device_type: deviceType,
            device_name: deviceName,
        })
             .then(response => {
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
    doBindDevice () {
        var that = this
        var records = []
        for (var toBindDeviceUUID of that.bindDevices) {
            var toBindDevice = that.getDeviceByUUID(toBindDeviceUUID)
            console.log(JSON.stringify(toBindDevice))
            records.push(that.doBindDeviceAndMonitorItem(that.bindDeviceMonitorItemId, that.bindDeviceMonitorItemName, toBindDevice.uuid, toBindDevice.name, toBindDevice.device_type))
        }

        Promise.all(records)
        .then(values => {
            that.initData()
        }).catch(errors => {
            that.$message({
                type: 'error',
                message: "操作遇到异常",
                offset: 200,
            })
            console.error(errors)
        })
    },
    updateMonitorBackendZabbixCfg () {
        var that = this
        axios.post(that.config.getAddress("UPDATE_BACKEND_CFG"), {
            backend_name: "ZABBIX",
            cfg: JSON.stringify(that.updateMonitorBackendZabbixForm),
        })
             .then(response => {
                that.$message({
                    type: 'success',
                    message: "操作成功",
                    offset: 200,
                })
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
    removeAlertItem (alertItem) {
        var that = this
        axios.post(that.config.getAddress("DELETE_ALERT_ITEM"), {id: alertItem.id})
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
