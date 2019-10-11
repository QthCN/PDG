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
            prop="role"
            label="角色">
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
            <el-form-item label="角色" :label-width="formLabelWidth">
                <el-select v-model="createUserForm.role" placeholder="请选择">
                    <el-option value="普通用户">普通用户</el-option>
                    <el-option value="管理员">管理员</el-option>
                </el-select>
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
  data () {
      return {
          users: [],
          config: new Config(),
          createUserDialogVisible: false,
          formLabelWidth: '120px',
          createUserForm: {
              username: "",
              password: "",
              role: "普通用户"
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
              password: "",
              role: "普通用户"
          }

        Promise.all([
            that.syncUsers(),
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
    syncUsers () {
        var that = this
        return axios.post(that.config.getAddress("LIST_USERS"))
                    .then(response => {
                        that.users = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.users = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
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
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    createUser () {
        var that = this
        axios.post(that.config.getAddress("CREATE_USER"), {username: that.createUserForm.username, password: that.createUserForm.password, role: that.createUserForm.role})
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
