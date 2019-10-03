import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    // 当前用户信息
    currentUser: "游客",
    // 当前用户想访问的地址
    routePath: "",
    // 页面加载动画开关
    pageLoading: false,
  },
  mutations: {
    // 设置当前用户
    setCurrentUser(state, username) {
      state.currentUser = username
    },
    // 设置访问地址
    setRoutePath(state, url) {
      state.routePath = url
    },
    // 设置页面加载开关
    setPageLoading(state, loading) {
      state.pageLoading = loading
    }
  },
  actions: {

  }
})
