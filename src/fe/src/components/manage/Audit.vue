<template>
  <div>
    <el-table
        :data="records"
        border
        highlight-current-row
        style="width: 100%">
        <el-table-column
            prop="username"
            label="操作人">
        </el-table-column>
        <el-table-column
            prop="action"
            label="操作类型">
        </el-table-column>
        <el-table-column
            prop="action_time"
            label="操作时间">
        </el-table-column>
        <el-table-column
            prop="args"
            label="说明">
        </el-table-column>
    </el-table>

    <div class="block" style="float: right;">
        <el-pagination
            @current-change="handlePageChange"
            layout="prev, pager, next"
            :page-size="recordsPerPage"
            :total="totalCnt"
        ></el-pagination>
    </div>

  </div>
</template>

<script>
import axios from "axios"
import Config from '../../config'

export default {
  name: 'Audit',
  data () {
      return {
          config: new Config(),
          records: [],
          currentPage: 1,
          totalCnt: 1,
          recordsPerPage: 10,
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
        
        that.records = []
        that.totalCnt = 1

        Promise.all([
            that.syncAuditRecords(),
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
    syncAuditRecords () {
        var that = this
        return axios.post(that.config.getAddress("LIST_AUDIT_RECORDS"), {current_page: that.currentPage, records_per_page: that.recordsPerPage})
                    .then(response => {
                        that.records = response.data.records
                        that.totalCnt = response.data.total_cnt
                    })
                    .catch(error => {
                        console.error(error)
                        that.records = []
                        that.totalCnt = 1
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    handlePageChange (val) {
        this.currentPage = val
        this.syncAuditRecords()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
