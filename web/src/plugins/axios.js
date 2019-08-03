"use strict";

import Vue from 'vue';
import axios from 'axios'
import store from '../store'
import router from '../router'

// Full config:  https://github.com/axios/axios#request-config
// axios.defaults.baseURL = process.env.baseURL || process.env.apiUrl || '';
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

let config = {
  baseURL: process.env.baseURL || process.env.apiUrl || "",
  timeout: 60 * 1000, // Timeout
  withCredentials: true, // Check cross-site Access-Control
};

const request = axios.create(config);

request.interceptors.request.use(config => {
  for(let key in config.params) {
    if (config.params.hasOwnProperty(key) && (config.params[key] === "" || config.params[key] === null)) {
      delete config.params[key];
    }
  }

  if (store.state.account.access_token) {
    config.headers.Authorization = `Bearer ${store.state.account.access_token}`;
  }
  return config;
}, error => {
  return Promise.reject(error);
});

// Add a response interceptor
request.interceptors.response.use(
  response => {
    return response.data;
  },
  error => {
    if (error.response.status === 401) {
      store.commit('SET_PROFILE', false);
      store.commit('SET_ACCESS_TOKEN', false);
      router.push("/signin");
    }
    return Promise.reject(error.response.data);
  }
);

Plugin.install = function(Vue) {
  Vue.axios = request;
  window.axios = request;
  Object.defineProperties(Vue.prototype, {
    axios: {
      get() {
        return request;
      }
    },
    $axios: {
      get() {
        return request;
      }
    },
  });
};

Vue.use(Plugin);

export default Plugin;
