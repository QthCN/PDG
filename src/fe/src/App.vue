<template>
  <div id="app">
    <el-container style="height: 100%;">
      <el-header style="padding: 0px;">

        <el-menu mode="horizontal" background-color="#545c64" text-color="#ffffff" active-text-color="#ffffff" style="box-shadow: 0 12px 14px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04); z-index: 9999;">

          <div style="float: left; height: 60px; margin-left: 20px;" >
            <!-- <img style="display: inline-block; height: 40px; margin-top: 10px; margin-left: 20px;" src="./assets/logo.png"/> -->
            <i class="el-icon-s-promotion" style="height: 60px; font-size: 30px; vertical-align: middle; line-height: 60px; color: #ffffff;"></i>
          </div>
          <div style="float: left; height: 60px; margin-left: 10px;">
              <span style="vertical-align: middle; line-height: 60px; color: #ffffff;">管理平台</span>
          </div>

          <el-submenu index="登陆菜单" style="float: right;">
            <template slot="title">{{$store.state.currentUser}}</template>
            <el-menu-item index="注销登陆" @click="logout">注销登陆</el-menu-item>
          </el-submenu>
        </el-menu>

      </el-header>


      <el-container>

        <el-aside :width="asideWidth">
          
          <el-menu :default-active="indexSelected" router @select="changeIndex" style="height: 100%;" :collapse="isCollapse" :collapse-transition="false" :unique-opened="true">
            <el-menu-item index="/r/dashboard">
                <i class="el-icon-s-marketing"></i>
                <span slot="title">平台大盘</span>
            </el-menu-item>

            <el-submenu index="资源情况">
              <template slot="title">
                <i class="el-icon-location"></i>
                <span slot="title">资源情况</span>
              </template>
              <el-menu-item index="/r/device/p">物理拓扑</el-menu-item>
              <el-menu-item index="/r/device/n">网络拓扑</el-menu-item>
              <el-menu-item index="/r/device/l">资源拓扑</el-menu-item>
            </el-submenu>

            <el-submenu index="监控数据">
              <template slot="title">
                <i class="el-icon-s-order"></i>
                <span slot="title">监控数据</span>
              </template>
              <el-menu-item index="/r/monitor/m">监控大盘</el-menu-item>
              <el-menu-item index="/r/monitor/a">告警大盘</el-menu-item>
            </el-submenu>

            <el-submenu index="平台管理">
              <template slot="title">
                <i class="el-icon-menu"></i>
                <span slot="title">平台管理</span>
              </template>
              <el-menu-item index="/r/manage/d">资源管理</el-menu-item>
              <el-menu-item index="/r/manage/l">布线管理</el-menu-item>
              <el-menu-item index="/r/manage/m">监控管理</el-menu-item>
              <el-menu-item index="/r/manage/i">IP及网段管理</el-menu-item>
              <el-menu-item index="/r/manage/u">人员管理</el-menu-item>
              <el-menu-item index="/r/manage/a">操作审计</el-menu-item>
            </el-submenu>

            <div class="collapse-button" @click="collapse">
              <template v-if="isCollapse">
                <el-divider><i class="el-icon-s-unfold"></i></el-divider>
              </template>
              <template v-else>
                <el-divider><i class="el-icon-s-fold"></i></el-divider>
              </template>
            </div>
          </el-menu>

        </el-aside>

        <el-container>
          <el-main style="height: 100%;" v-loading="$store.state.pageLoading">
            <router-view/>
          </el-main>

          <el-footer>
            <el-divider><span style="color: #909399; font-size: 10px;">{{copyrightContent}}</span></el-divider>
          </el-footer>
        </el-container>
        
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isCollapse: false,
      asideWidth: "200px",
      indexSelected: "/r/device/p",
    }
  },
  computed: {
    copyrightContent () {
      var year = (new Date()).getFullYear()
      return `©2018-${year} 管理平台`
    }
  },
  mounted () {
    this.indexSelected = this.$store.state.routePath
  },
  methods: {
    changeIndex(key, keyPath) {
      this.indexSelected = key
    },
    collapse () {
      this.isCollapse = !this.isCollapse
      if (this.isCollapse) {
        this.asideWidth = "68px"
      } else {
        this.asideWidth = "200px"
      }
    },
    logout () {
      this.$store.commit("setCurrentUser", "游客")
      window.location.href = "/v1/ajax/auth/logout"
    }
  }
}
</script>

<style lang="scss">
html,
body,
#app {
  margin: 0px;
  padding: 0px;
  height: 100%;
}

.collapse-button {
  color: #909399; 
  font-size: 20px; 
  height: 56px; 
  line-height: 56px; 
  text-align: center;
  width: 100%;
}

.collapse-button i {
  color: #909399; 
}


</style>
