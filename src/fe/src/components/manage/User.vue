<template>
  <div>
    <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createUserDialogVisible = true">新建用户</el-button>
    <el-table
        :data="users"
        border
        highlight-current-row
        style="width: 100%">
        <el-table-column
            prop="username"
            label="用户名">
        </el-table-column>
        <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
                <el-button @click="removeUser(scope.row)" type="danger" plain size="small">删除</el-button>
            </template>
        </el-table-column>
    </el-table>


    <el-dialog title="新建用户" :visible.sync="createUserDialogVisible">
        <el-form :model="createUserForm">
            <el-form-item label="用户名" :label-width="formLabelWidth">
                <el-input v-model="createUserForm.username" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="密码" :label-width="formLabelWidth">
                <el-input v-model="createUserForm.password" autocomplete="off" type="password"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="createUserDialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="createUser">确 定</el-button>
        </div>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios"
import Config from '../../config'

export default {
  name: 'User',
  props: {
  },
  data () {
      return {
          users: [],
          config: new Config(),
          createUserDialogVisible: false,
          formLabelWidth: '120px',
          createUserForm: {
              username: "",
              password: ""
          }
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
        that.createUserDialogVisible = false
        that.createUserForm = {
              username: "",
              password: ""
          }

        Promise.all([
            that.syncUsers(),
        ]).then(values => {
            that.$store.commit("setPageLoading", false)
        }).catch(errors => {
            that.$message.error("页面加载异常")
            console.error(errors)
            that.$store.commit("setPageLoading", false)
        })
    },
    syncUsers () {
    var that = this
    return axios.post(that.config.getAddress("LIST_USERS"))
                .then(response => {
                    that.users = response.data
                })
                .catch(error => {
                    console.error(error)
                    that.users = []
                    that.$message.error("获取数据异常")
                })
    },
    removeUser(user) {
        var that = this
        axios.post(that.config.getAddress("REMOVE_USER"), {username: user.username})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message.error(error.response.data.msg)
             })
    },
    createUser () {
        var that = this
        axios.post(that.config.getAddress("CREATE_USER"), {username: that.createUserForm.username, password: that.createUserForm.password})
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
