import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import Auth from './utils/auth'

Vue.config.productionTip = false

Vue.use(ElementUI);

router.beforeEach((to, _from, next) => {
  store.commit("setRoutePath", to.path)
  store.commit("setPageLoading", false)
  store.commit("setPageLoading", true)

  // 检查用户是否登录，如果没有登陆则跳转登录页
  if (store.state.currentUser === "游客") {
    (new Auth()).getCurrentUsername((username) => {
      store.commit("setCurrentUser", username)
      next()
    })
  } else {
    next()
  }
})

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
