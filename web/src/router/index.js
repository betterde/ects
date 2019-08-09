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
      component: () => import('../views/Dashboard.vue'),
    },
    {
      path: "/pipeline",
      name: "pipeline",
      meta: {
        requiresAuth: true
      },
      component: () => import('../views/Pipeline.vue'),
    },
    {
      path: '/task',
      name: 'task',
      meta: {
        requiresAuth: true
      },
      component: () => import('../views/Task.vue')
    },
    {
      path: '/node',
      name: 'node',
      meta: {
        requiresAuth: true
      },
      component: () => import('../views/Node.vue')
    },
    {
      path: '/user',
      name: 'user',
      meta: {
        requiresAuth: true
      },
      component: () => import('../views/User.vue')
    },
    {
      path: '/log',
      name: 'log',
      meta: {
        requiresAuth: true
      },
      component: () => import('../views/Log.vue')
    },
    {
      path: '/setting',
      name: 'setting',
      meta: {
        requiresAuth: true
      },
      component: () => import('../views/Setting.vue')
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
      path: '*',
      name: "notfound",
      meta: {
        requiresAuth: false
      },
      component: () => import('../views/NotFound.vue')
    }
  ]
});

/**
 * Routing to intercept
 */
router.beforeEach((to, from, next) => {
  // 如果跳转到NotFound页面则提前设置视图的Layout 为 guest
  if (to.name === 'notfound') {
    store.commit('SET_LAYOUT_CURRENT', 'guest');
  }

  // 如果从NotFound 页面返回，并且需要认证的话，则设置视图的Layout 为 backend
  if (from.name === 'notfound' && to.meta.requiresAuth === true) {
    store.commit('SET_LAYOUT_CURRENT', 'backend')
  }

  if (store.state.account.access_token && to.path === '/signin') {
    next({
      path: "/"
    })
  }

  /**
   * Determine if auth is required
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
