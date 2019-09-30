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
            <template slot="title">游客</template>
            <el-menu-item index="登陆按钮">登陆</el-menu-item>
          </el-submenu>
        </el-menu>

      </el-header>


      <el-container>

        <el-aside :width="asideWidth">
          
          <el-menu style="height: 100%;" :collapse="isCollapse" :collapse-transition="false">
            <el-menu-item index="机柜拓扑">
              <i class="el-icon-location"></i>
              <span slot="title">机柜拓扑</span>
            </el-menu-item>
            <el-menu-item index="资源拓扑">
              <i class="el-icon-menu"></i>
              <span slot="title">资源拓扑</span>
            </el-menu-item>

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


        <el-main style="height: 100%;">main</el-main>

      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isCollapse: false,
      asideWidth: "200px"
    }
  },
  methods: {
    collapse () {
      this.isCollapse = !this.isCollapse
      if (this.isCollapse) {
        this.asideWidth = "68px"
      } else {
        this.asideWidth = "200px"
      }
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
  position: absolute;
  bottom: 5px;
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
