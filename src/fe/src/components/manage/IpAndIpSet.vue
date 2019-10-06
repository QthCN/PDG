<template>
  <div>
    <el-tabs type="border-card">
        <el-tab-pane label="IP地址">
            <el-table
                :data="ips"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    prop="ip_address"
                    label="IP">
                </el-table-column>
                <el-table-column
                    prop="role"
                    label="类型">
                </el-table-column>
                <el-table-column
                    prop="ipset_id"
                    label="所属网段">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="200">
                    <template slot-scope="scope">
                        <el-button @click="showTargetDevice(scope.row)" type="danger" plain size="small">查看绑定设备</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-tab-pane>

        <el-tab-pane label="网段">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createIpsetDialogVisible = true">新建网段</el-button>
            <el-table
                :data="ipsets"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    prop="cidr"
                    label="CIDR">
                </el-table-column>
                <el-table-column
                    prop="comment"
                    label="备注">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="300">
                    <template slot-scope="scope">
                        <el-button @click="showAvailableIP(scope.row)" type="primary" plain size="small">查询可用IP</el-button>
                        <el-button @click="removeIPSet(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-tab-pane>
    </el-tabs>

    <el-dialog title="新建网段" :visible.sync="createIpsetDialogVisible">
        <el-form :model="createIpsetForm">
            <el-form-item label="CIDR" :label-width="formLabelWidth">
                <el-input v-model="createIpsetForm.cidr" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="备注" :label-width="formLabelWidth">
                <el-input v-model="createIpsetForm.comment" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="createIpsetDialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="createIPSet">确 定</el-button>
        </div>
    </el-dialog>

  </div>
</template>

<script>
import axios from "axios"
import Config from '../../config'

export default {
  name: 'IpAndIpSet',
  data () {
      return {
          config: new Config(),
          formLabelWidth: '120px',
          ips: [],
          ipsets: [],
          createIpsetDialogVisible: false,
          createIpsetForm: {
              cidr: "",
              comment: ""
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

        that.ips = []
        that.ipsets = []
        that.createIpsetDialogVisible = false
        that.createIpsetForm = {
            cidr: "",
            comment: ""
        }
        
        Promise.all([
            that.syncIPS(),
            that.syncIPSets()
        ]).then(values => {
            that.$store.commit("setPageLoading", false)
        }).catch(errors => {
            that.$message.error("页面加载异常")
            console.error(errors)
            that.$store.commit("setPageLoading", false)
        })
    },
    syncIPS () {
        var that = this
        return axios.post(that.config.getAddress("LIST_IPS"))
                    .then(response => {
                        that.ips = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.ips = []
                        that.$message.error("获取数据异常")
                    })
    },
    syncIPSets () {
        var that = this
        return axios.post(that.config.getAddress("LIST_IPSETS"))
                    .then(response => {
                        that.ipsets = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.ipsets = []
                        that.$message.error("获取数据异常")
                    })
    },
    showTargetDevice (ip) {

    },
    showAvailableIP (ipSet) {

    },
    removeIPSet (ipset) {
        var that = this
        return axios.post(that.config.getAddress("DELETE_IPSET"), {uuid: ipset.uuid})
                    .then(response => {
                        that.initData()
                    })
                    .catch(error => {
                        console.error(error)
                        that.$message.error("获取数据异常")
                    })
    },
    createIPSet () {
        var that = this
        return axios.post(that.config.getAddress("CREATE_IPSET"), {cidr: that.createIpsetForm.cidr, comment: that.createIpsetForm.comment})
                    .then(response => {
                        that.initData()
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
