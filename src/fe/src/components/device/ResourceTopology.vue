<template>
  <div style="">
        <el-input
            placeholder="输入关键字进行设备过滤"
            v-model="filterText">
        </el-input>
        <br/>
        <br/>
        <el-tree
            :data="items"
            :props="defaultProps"
            :filter-node-method="filterNode"
            @node-click="itemClick"
            ref="tree">
            <span slot-scope="{ node, data }">
                <template v-if="data.device_type === 'DATACENTER'">
                    <i class="el-icon-office-building"></i>&nbsp;&nbsp;<span>{{node.label}}</span>
                </template>
                <template v-else-if="data.device_type === 'RACK'">
                    <i class="el-icon-c-scale-to-original"></i>&nbsp;&nbsp;<span>{{node.label}}</span>
                </template>
                <template v-else-if="data.device_type === 'SERVER'">
                    <i class="el-icon-refrigerator"></i>&nbsp;&nbsp;<span>{{node.label}}</span>
                </template>
                <template v-else-if="data.device_type === 'NETWORK'">
                    <i class="el-icon-cpu"></i>&nbsp;&nbsp;<span>{{node.label}}</span>
                </template>
                <template v-else-if="data.device_type === 'STORAGE'">
                    <i class="el-icon-coin"></i>&nbsp;&nbsp;<span>{{node.label}}</span>
                </template>
                <template v-else>
                    <i class="el-icon-mobile"></i>&nbsp;&nbsp;<span>{{node.label}}</span>
                </template>
            </span>
        </el-tree>


    <el-drawer
        title="设备信息"
        size="60%"
        :visible.sync="serverDeviceStatusDialogVisible"
        :direction="'rtl'">
        <DeviceStatus :uuid="deviceUUID" :device-type="deviceType"></DeviceStatus>
    </el-drawer>
  </div>
</template>

<script>
import axios from "axios"

import Config from '../../config'
import DeviceStatus from '../plugin/DeviceStatus.vue'


export default {
  name: 'ResourceTopology',
  data () {
    return {
        config: new Config(),
        items: [],
        filterText: '',
        defaultProps: {
            children: 'children',
            label: 'label'
        },
        serverDeviceStatusDialogVisible: false,
        deviceUUID: "",
        deviceType: "",
    }
  },
  created () {
    var that = this
    that.initData()
  },
  components: {
      DeviceStatus
  },
  mounted () {
    
  },
  watch: {
    filterText(val) {
        this.$refs.tree.filter(val);
    }
  },
  methods: {
    initData () {
        var that = this

        that.filterText = ''
        that.items = []
        
        Promise.all([
            that.syncItems()
        ]).then(values => {
            that.$store.commit("setPageLoading", false)
        }).catch(errors => {
            that.$message.error("页面加载异常")
            console.error(errors)
            that.$store.commit("setPageLoading", false)
        })
    },
    filterNode(value, data) {
        if (!value) return true;
        return data.label.indexOf(value) !== -1;
    },
    syncItems() {
        var that = this
        return axios.post(that.config.getAddress("GET_RESOURCE_TOPOLOGY"), {})
                    .then(response => {
                        console.log(JSON.stringify(response.data))
                        that.items = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.$message.error("获取数据异常")
                    })
    },
    itemClick(data, node, h) {
        if (data.device_type == "SERVER" || data.device_type == "NETWORK" || data.device_type == "STORAGE" || data.device_type == "OTHER" ) {
            this.deviceUUID = data.uuid
            this.deviceType = data.device_type
            this.serverDeviceStatusDialogVisible = true
        }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
