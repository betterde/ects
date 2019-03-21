import Vue from 'vue'
import api from '../../apis'
import * as types from '../types'

export default {
  state: {
    profile: {
      id: '',
      name: '',
      email: global.JSON.parse(localStorage.getItem('profile')),
      avatar: '',
      group_id: 0,
      manager: false
    },
    access_token: localStorage.getItem('access_token'),
  },
  mutations: {
    SET_PROFILE: (state, data) => {
      // 如果数据无效则删除
      if (data === false){
        Vue.set(state, 'profile', {
          id: '',
          name: '',
          email: '',
          avatar: '',
          group_id: 0,
          manager: false
        });
        localStorage.removeItem('profile')
      } else {
        let profile = {
          id: data.id,
          name: data.name,
          email: data.email,
          avatar: data.avatar,
          group_id: data.group_id,
          manager: data.manager
        };
        Vue.set(state, 'profile', profile);
        localStorage.setItem('profile', global.JSON.stringify(profile))
      }
    },
    SET_ACCESS_TOKEN: (state, data) => {
      if (data === false) {
        Vue.set(state, 'expires_in', '');
        Vue.set(state, 'access_token', '');
        // 删除 localStorage 中的用户凭证信息
        localStorage.removeItem('expires_in');
        localStorage.removeItem('access_token');
      } else {
        Vue.set(state, 'access_token', data.access_token);
        // 写入 localStorage 防止刷新后用户凭证丢失
        localStorage.setItem('access_token', data.access_token);
      }
    }
  },
  actions: {
    signIn({ dispatch, commit }, params) {
      return new Promise((resolve, reject) => {
        api.account.signin(params).then(res => {
          commit(types.SET_ACCESS_TOKEN, res.data);
        }).catch(err => {
          commit(types.SET_ACCESS_TOKEN, false);
          return err.response.data;
        });
        resolve();
      });
    },
    fetchProfile({ commit }) {
      api.account.profile().then(res => {
        commit(types.SET_PROFILE, res.data)
      }).catch(err => {
        window.console.log(err);
        commit(types.SET_PROFILE, false)
      });
    }
  }
}

