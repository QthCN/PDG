<template>
  <div>
    <el-tabs type="border-card">
        <el-tab-pane label="机房">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createDataCenterDialogVisible = true">新建机房</el-button>
            <el-table
                :data="datacenters"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    prop="name"
                    label="机房名">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="100">
                    <template slot-scope="scope">
                        <el-button @click="removeDataCenter(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建机房" :visible.sync="createDataCenterDialogVisible">
                <el-form :model="createDataCenterForm">
                    <el-form-item label="机房名" :label-width="formLabelWidth">
                        <el-input v-model="createDataCenterForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createDataCenterDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createDataCenter">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="机柜">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createRackDialogVisible = true">新建机柜</el-button>
            <el-table
                :data="racks"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    prop="name"
                    label="机柜名">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="100">
                    <template slot-scope="scope">
                        <el-button @click="removeRack(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建机柜" :visible.sync="createRackDialogVisible">
                <el-form :model="createRackForm">
                    <el-form-item label="机柜名" :label-width="formLabelWidth">
                        <el-input v-model="createRackForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createRackDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createRack">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="物理服务器">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createServerDeviceDialogVisible = true">新建物理服务器</el-button>
            <el-table
                :data="serverDevices"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    fixed
                    prop="hostname"
                    label="主机名">
                </el-table-column>
                <el-table-column
                    prop="brand"
                    label="厂商">
                </el-table-column>
                <el-table-column
                    prop="model"
                    label="型号">
                </el-table-column>
                <el-table-column
                    prop="disk_capacity"
                    label="磁盘(TB)">
                </el-table-column>
                <el-table-column
                    prop="memory_capacity"
                    label="内存(GB)">
                </el-table-column>
                <el-table-column
                    prop="enable_time"
                    label="启用时间">
                </el-table-column>
                <el-table-column
                    prop="expire_time"
                    label="过保时间">
                </el-table-column>
                <el-table-column
                    prop="os"
                    label="操作系统">
                </el-table-column>
                <el-table-column
                    prop="comment"
                    label="备注">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="100">
                    <template slot-scope="scope">
                        <el-button @click="removeServerDevice(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建物理服务器" :visible.sync="createServerDeviceDialogVisible">
                <el-form :model="createServerDeviceForm">
                    <el-form-item label="主机名" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.hostname" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="厂商" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.brand" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="型号" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.model" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="磁盘容量(TB)" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.disk_capacity" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="内存容量(GB)" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.memory_capacity" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="启用时间" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.enable_time" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="过保时间" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.expire_time" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="操作系统" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.os" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="备注" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.comment" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createServerDeviceDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createServerDevice">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="存储设备">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createStorageDeviceDialogVisible = true">新建存储设备</el-button>
            <el-table
                :data="storageDevices"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    fixed
                    prop="name"
                    label="设备名">
                </el-table-column>
                <el-table-column
                    prop="brand"
                    label="厂商">
                </el-table-column>
                <el-table-column
                    prop="model"
                    label="型号">
                </el-table-column>
                <el-table-column
                    prop="enable_time"
                    label="启用时间">
                </el-table-column>
                <el-table-column
                    prop="expire_time"
                    label="过保时间">
                </el-table-column>
                <el-table-column
                    prop="comment"
                    label="备注">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="100">
                    <template slot-scope="scope">
                        <el-button @click="removeStorageDevice(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建存储设备" :visible.sync="createStorageDeviceDialogVisible">
                <el-form :model="createStorageDeviceForm">
                    <el-form-item label="设备名" :label-width="formLabelWidth">
                        <el-input v-model="createStorageDeviceForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="厂商" :label-width="formLabelWidth">
                        <el-input v-model="createStorageDeviceForm.brand" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="型号" :label-width="formLabelWidth">
                        <el-input v-model="createStorageDeviceForm.model" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="启用时间" :label-width="formLabelWidth">
                        <el-input v-model="createStorageDeviceForm.enable_time" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="过保时间" :label-width="formLabelWidth">
                        <el-input v-model="createStorageDeviceForm.expire_time" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="备注" :label-width="formLabelWidth">
                        <el-input v-model="createStorageDeviceForm.comment" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createStorageDeviceDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createStorageDevice">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="网络设备">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createNetworkDeviceDialogVisible = true">新建网络设备</el-button>
            <el-table
                :data="networkDevices"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    fixed
                    prop="name"
                    label="设备名">
                </el-table-column>
                <el-table-column
                    prop="brand"
                    label="厂商">
                </el-table-column>
                <el-table-column
                    prop="model"
                    label="型号">
                </el-table-column>
                <el-table-column
                    prop="enable_time"
                    label="启用时间">
                </el-table-column>
                <el-table-column
                    prop="expire_time"
                    label="过保时间">
                </el-table-column>
                <el-table-column
                    prop="comment"
                    label="备注">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="100">
                    <template slot-scope="scope">
                        <el-button @click="removeNetworkDevice(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建网络设备" :visible.sync="createNetworkDeviceDialogVisible">
                <el-form :model="createNetworkDeviceForm">
                    <el-form-item label="设备名" :label-width="formLabelWidth">
                        <el-input v-model="createNetworkDeviceForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="厂商" :label-width="formLabelWidth">
                        <el-input v-model="createNetworkDeviceForm.brand" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="型号" :label-width="formLabelWidth">
                        <el-input v-model="createNetworkDeviceForm.model" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="启用时间" :label-width="formLabelWidth">
                        <el-input v-model="createNetworkDeviceForm.enable_time" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="过保时间" :label-width="formLabelWidth">
                        <el-input v-model="createNetworkDeviceForm.expire_time" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="备注" :label-width="formLabelWidth">
                        <el-input v-model="createNetworkDeviceForm.comment" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createNetworkDeviceDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createNetworkDevice">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="其它设备">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createCommonDeviceDialogVisible = true">新建其它设备</el-button>
            <el-table
                :data="commonDevices"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    fixed
                    prop="name"
                    label="设备名">
                </el-table-column>
                <el-table-column
                    prop="brand"
                    label="厂商">
                </el-table-column>
                <el-table-column
                    prop="model"
                    label="型号">
                </el-table-column>
                <el-table-column
                    prop="enable_time"
                    label="启用时间">
                </el-table-column>
                <el-table-column
                    prop="expire_time"
                    label="过保时间">
                </el-table-column>
                <el-table-column
                    prop="comment"
                    label="备注">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="100">
                    <template slot-scope="scope">
                        <el-button @click="removeCommonDevice(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建其它设备" :visible.sync="createCommonDeviceDialogVisible">
                <el-form :model="createCommonDeviceForm">
                    <el-form-item label="设备名" :label-width="formLabelWidth">
                        <el-input v-model="createCommonDeviceForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="厂商" :label-width="formLabelWidth">
                        <el-input v-model="createCommonDeviceForm.brand" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="型号" :label-width="formLabelWidth">
                        <el-input v-model="createCommonDeviceForm.model" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="启用时间" :label-width="formLabelWidth">
                        <el-input v-model="createCommonDeviceForm.enable_time" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="过保时间" :label-width="formLabelWidth">
                        <el-input v-model="createCommonDeviceForm.expire_time" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="备注" :label-width="formLabelWidth">
                        <el-input v-model="createCommonDeviceForm.comment" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createCommonDeviceDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createCommonDevice">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import axios from "axios"
import Config from '../../config'

export default {
  name: 'Device',
  data () {
      return {
          config: new Config(),
          formLabelWidth: '120px',

          createDataCenterDialogVisible: false,
          datacenters: [],
          createDataCenterForm: {
              name: "",
          },
          
          createRackDialogVisible: false,
          racks: [],
          createRackForm: {
              name: "",
          },

          createServerDeviceDialogVisible: false,
          serverDevices: [],
          createServerDeviceForm: {
              brand: "",
              model: "",
              disk_capacity: 0,
              memory_capacity: 0,
              hostname: "",
              enable_time: "",
              expire_time: "",
              os: "",
              comment: "",
          },

          createStorageDeviceDialogVisible: false,
          storageDevices: [],
          createStorageDeviceForm: {
              brand: "",
              model: "",
              name: "",
              enable_time: "",
              expire_time: "",
              comment: "",
          },

          createNetworkDeviceDialogVisible: false,
          networkDevices: [],
          createNetworkDeviceForm: {
              brand: "",
              model: "",
              name: "",
              enable_time: "",
              expire_time: "",
              comment: "",
          },

          createCommonDeviceDialogVisible: false,
          commonDevices: [],
          createCommonDeviceForm: {
              brand: "",
              model: "",
              name: "",
              enable_time: "",
              expire_time: "",
              comment: "",
          },
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

        that.createDataCenterDialogVisible = false
        that.datacenters = []
        that.createDataCenterForm = {
              name: "",
        }

        that.createRackDialogVisible = false
        that.racks = []
        that.createRackForm = {
              name: "",
        }

        that.createServerDeviceDialogVisible = false
        that.serverDevices = []
        that.createServerDeviceForm = {
              brand: "",
              model: "",
              disk_capacity: 0,
              memory_capacity: 0,
              hostname: "",
              enable_time: "",
              expire_time: "",
              os: "",
              comment: "",
        }

        that.createStorageDeviceDialogVisible = false
        that.storageDevices = []
        that.createStorageDeviceForm = {
              brand: "",
              model: "",
              name: "",
              enable_time: "",
              expire_time: "",
              comment: "",
        }

        that.createNetworkDeviceDialogVisible = false
        that.networkDevices = []
        that.createNetworkDeviceForm = {
              brand: "",
              model: "",
              name: "",
              enable_time: "",
              expire_time: "",
              comment: "",
        }

        that.createCommonDeviceDialogVisible = false
        that.commonDevices = []
        that.createCommonDeviceForm = {
              brand: "",
              model: "",
              name: "",
              enable_time: "",
              expire_time: "",
              comment: "",
        }
        
        Promise.all([
            that.syncDataCenters(),
            that.syncRacks(),
            that.syncServerDevices(),
            that.syncStorageDevices(),
            that.syncNetworkDevices(),
            that.syncCommonDevices()
        ]).then(values => {
            that.$store.commit("setPageLoading", false)
        }).catch(errors => {
            that.$message.error("页面加载异常")
            console.error(errors)
            that.$store.commit("setPageLoading", false)
        })
    },
    syncDataCenters () {
        var that = this
        return axios.post(that.config.getAddress("LIST_DATACENTERS"))
                    .then(response => {
                        that.datacenters = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.datacenters = []
                        that.$message.error("获取数据异常")
                    })
    },
    removeDataCenter (datacenter) {
        var that = this
        axios.post(that.config.getAddress("DELETE_DATACENTER"), {uuid: datacenter.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    createDataCenter () {
        var that = this
        axios.post(that.config.getAddress("CREATE_DATACENTER"), {name: that.createDataCenterForm.name})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    syncRacks () {
        var that = this
        return axios.post(that.config.getAddress("LIST_RACKS"))
                    .then(response => {
                        that.racks = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.racks = []
                        that.$message.error("获取数据异常")
                    })
    },
    removeRack (rack) {
        var that = this
        axios.post(that.config.getAddress("DELETE_RACK"), {uuid: rack.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    createRack () {
        var that = this
        axios.post(that.config.getAddress("CREATE_RACK"), {name: that.createRackForm.name})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    syncServerDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_SERVERS"))
                    .then(response => {
                        that.serverDevices = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.serverDevices = []
                        that.$message.error("获取数据异常")
                    })
    },
    removeServerDevice (server) {
        var that = this
        axios.post(that.config.getAddress("DELETE_SERVER"), {uuid: server.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    createServerDevice () {
        var that = this
        axios.post(that.config.getAddress("CREATE_SERVER"), {
            brand: that.createServerDeviceForm.brand,
            model: that.createServerDeviceForm.model,
            disk_capacity: parseInt(that.createServerDeviceForm.disk_capacity),
            memory_capacity: parseInt(that.createServerDeviceForm.memory_capacity),
            hostname: that.createServerDeviceForm.hostname,
            enable_time: that.createServerDeviceForm.enable_time,
            expire_time: that.createServerDeviceForm.expire_time,
            os: that.createServerDeviceForm.os,
            comment: that.createServerDeviceForm.comment,
        })
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    syncStorageDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_STORAGE_DEVICES"))
                    .then(response => {
                        that.storageDevices = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.storageDevices = []
                        that.$message.error("获取数据异常")
                    })
    },
    removeStorageDevice (device) {
        var that = this
        axios.post(that.config.getAddress("DELETE_STORAGE_DEVICE"), {uuid: device.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    createStorageDevice () {
        var that = this
        axios.post(that.config.getAddress("CREATE_STORAGE_DEVICE"), {
            brand: that.createStorageDeviceForm.brand,
            model: that.createStorageDeviceForm.model,
            name: that.createStorageDeviceForm.name,
            enable_time: that.createStorageDeviceForm.enable_time,
            expire_time: that.createStorageDeviceForm.expire_time,
            comment: that.createStorageDeviceForm.comment,
        })
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    syncNetworkDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_NETWORK_DEVICES"))
                    .then(response => {
                        that.networkDevices = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.networkDevices = []
                        that.$message.error("获取数据异常")
                    })
    },
    removeNetworkDevice (device) {
        var that = this
        axios.post(that.config.getAddress("DELETE_NETWORK_DEVICE"), {uuid: device.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    createNetworkDevice () {
        var that = this
        axios.post(that.config.getAddress("CREATE_NETWORK_DEVICE"), {
            brand: that.createNetworkDeviceForm.brand,
            model: that.createNetworkDeviceForm.model,
            name: that.createNetworkDeviceForm.name,
            enable_time: that.createNetworkDeviceForm.enable_time,
            expire_time: that.createNetworkDeviceForm.expire_time,
            comment: that.createNetworkDeviceForm.comment,
        })
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    syncCommonDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_COMMON_DEVICES"))
                    .then(response => {
                        that.commonDevices = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.commonDevices = []
                        that.$message.error("获取数据异常")
                    })
    },
    removeCommonDevice (device) {
        var that = this
        axios.post(that.config.getAddress("DELETE_COMMON_DEVICE"), {uuid: device.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    createCommonDevice () {
        var that = this
        axios.post(that.config.getAddress("CREATE_COMMON_DEVICE"), {
            brand: that.createCommonDeviceForm.brand,
            model: that.createCommonDeviceForm.model,
            name: that.createCommonDeviceForm.name,
            enable_time: that.createCommonDeviceForm.enable_time,
            expire_time: that.createCommonDeviceForm.expire_time,
            comment: that.createCommonDeviceForm.comment,
        })
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    }
  }
}

</script>