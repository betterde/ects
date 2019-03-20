import Vue from 'vue'
import store from '../store'
import Router from 'vue-router'

Vue.use(Router);

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'dashboard',
      meta: {
        requiresAuth: true
      },
      component: () => import('../views/Dashboard.vue')
    },
    {
      path: '/signin',
      name: 'signin',
      meta: {
        requiresAuth: false
      },
      component: () => import('../views/SignIn.vue')
    },
    {
      // 会匹配所有路径
      path: '*'
    }
  ]
});

/**
 * 路由拦截
 */
router.beforeEach((to, from, next) => {
  /**
   * 判断前往的路由是否需要身份验证
   */
  if (to.matched.some(record => record.meta.requiresAuth)) {
    const access_token = store.state.account.access_token;
    //通过access_token判断用户是否已经登录
    if (!access_token) {
      next({
        path: '/signin'
      })
    } else {
      next()
    }
  } else {
    next()
  }
});

export default router;
