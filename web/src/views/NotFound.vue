<template>
  <div class="page-error page-error-404 layout-full">
    <div class="page animsition vertical-align text-center" data-animsition-in="fade-in" data-animsition-out="fade-out">
      <div class="page-content">
        <header>
          <h1 class="code animation-slide-top">404</h1>
          <p>The page does not exist，redirect to {{ access_token ? 'Index' : 'Sign In' }} page after <span v-html="second"></span> seconds</p>
        </header>
      </div>
    </div>
    <el-row type="flex" align="bottom" class="footer-row">
      <el-col :span="24">
        <footer class="copyright page-copyright-inverse text-center">
          <p>Copyright © 2017-2019 Betterde Inc. All Rights Reserved.</p>
        </footer>
      </el-col>
    </el-row>
  </div>
</template>
<script>
  import router from '../router'
  import { mapState } from 'vuex'

  export default{
    data () {
      return {
        second: 3,
        styles: {
          height: window.screen.height - 120 + 'px'
        }
      }
    },
    computed: {
      ...mapState ({
        access_token: state => state.account.access_token
      })
    },
    watch: {
      second (value) {
        if (value === 0) {
          clearInterval();
          router.push({
            path: this.access_token ? '/' : '/signin'
          });
        }
      }
    },
    mounted () {
      this.$store.commit("SET_LAYOUT_CURRENT", 'guest');
      /**
       * 设置定时器
       */
      setInterval(() => {
        this.second -= 1;
      }, 1000);
    }
  }
</script>
<style lang="scss">
  body {
    background: #f1f4f5 !important;
    font-family: Roboto, sans-serif;
    font-size: 14px;
    line-height: 1.57142857;
    color: #76838f;
    margin: 0;
    display: block;
  }
  .page-error {
    height: 100%;
    text-align: center;
  }
  .layout-full {
    .page {
      height: 80%;
      margin: 0 !important;
      max-width: none;
      padding: 0;
      background: #f1f4f5;
      position: relative;
      min-height: -webkit-calc(100% - 44px);
      min-height: calc(100% - 44px);
      display: flex;
      display: -webkit-flex;
      align-items: center;
    }
    .page-content {
      left: 50%;
      margin-left: -580px;
      position: absolute;
    }
  }

  .code{
    font-size: 10em;
    font-weight: 400
  }
  .page-error header p {
    margin-bottom: 30px;
    font-size: 30px;
    text-transform: uppercase;
  }
  p{
    margin: 0 0 11px;
  }
  .page-error .error-advise {
    margin-bottom: 25px;
    color: #A9AFB5
  }
  .btn-round{
    border-radius: 1000px
  }
  .btn{
    padding: 6px 15px;
    font-size: 14px;
    line-height: 1.57142857;
    display: inline-block;
    margin-bottom: 0;
    text-align: center;
    white-space: nowrap;
    vertical-align: middle;
  }
  .btn-primary{
    color: #fff;
    background-color: #62a8ea;
    border-color: #62a8ea
  }
  a{
    text-decoration: none
  }
  h1{
    margin-top: 22px;
    margin-bottom: 11px;
  }
  .page-copyright{
    margin-top: 60px;
    color: #37474f;
    font-size: 12px;
    letter-spacing: 1px
  }
  @-webkit-keyframes slide-top {
    0% {
      -webkit-transform:translate3d(0, -100%, 0);
    }
    100% {
      -webkit-transform:translate3d(0, 0, 0);
    }
  }
  [class*=animation-] {
    -webkit-animation-duration:.5s;
    -webkit-animation-timing-function:ease-out;
    -webkit-animation-fill-mode:both;
  }
  .animation-slide-top {
    -webkit-animation-name:slide-top;
  }
</style>
