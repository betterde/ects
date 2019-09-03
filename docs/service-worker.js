/**
 * Welcome to your Workbox-powered service worker!
 *
 * You'll need to register this file in your web app and you should
 * disable HTTP caching for this file too.
 * See https://goo.gl/nhQhGp
 *
 * The rest of the code is auto-generated. Please don't update this file
 * directly; instead, make changes to your Workbox build configuration
 * and re-run your build process.
 * See https://goo.gl/2aRDsh
 */

importScripts("https://storage.googleapis.com/workbox-cdn/releases/3.6.3/workbox-sw.js");

/**
 * The workboxSW.precacheAndRoute() method efficiently caches and responds to
 * requests for URLs in the manifest.
 * See https://goo.gl/S9QRab
 */
self.__precacheManifest = [
  {
    "url": "404.html",
    "revision": "4b89abb16662aa761bbecd73793d0dd4"
  },
  {
    "url": "architecture.png",
    "revision": "46d3e361c7c3ce6509950651d8b47a8a"
  },
  {
    "url": "assets/css/0.styles.048efcc2.css",
    "revision": "ceb572b3be3b081cbdb5f77418b4233b"
  },
  {
    "url": "assets/img/search.83621669.svg",
    "revision": "83621669651b9a3d4bf64d1a670ad856"
  },
  {
    "url": "assets/js/10.d1ba5c10.js",
    "revision": "879aaf34c42b6dee3f31e99d078eeb07"
  },
  {
    "url": "assets/js/11.1f45e276.js",
    "revision": "f0145fde537e5f2ad46f5815f0dc5f9d"
  },
  {
    "url": "assets/js/12.0beeb557.js",
    "revision": "09e7db5654b5051aac5a344b4eb0ac93"
  },
  {
    "url": "assets/js/13.171fce05.js",
    "revision": "a2629a505a726bc6969e1daa3e2610cf"
  },
  {
    "url": "assets/js/14.3ab2d6bc.js",
    "revision": "8441b3af0bb50f99fd58d2eb1c9f7a4a"
  },
  {
    "url": "assets/js/2.037b87e7.js",
    "revision": "cd879372c12642b707c9833d22b6efb9"
  },
  {
    "url": "assets/js/3.88e3f723.js",
    "revision": "b88c0b110abb249f5f23c3ec873b8eed"
  },
  {
    "url": "assets/js/4.efeec8cc.js",
    "revision": "fa0bcfd9fcbbd08465d67b3572ee2157"
  },
  {
    "url": "assets/js/5.338e9a1b.js",
    "revision": "075524e10e28b68cf07fe4c058e4eee5"
  },
  {
    "url": "assets/js/6.a252cc0f.js",
    "revision": "6eb915ba239dd8939e5e29cfeeff279f"
  },
  {
    "url": "assets/js/7.79b060f8.js",
    "revision": "877e751db679dfad9e72264aeced9bdd"
  },
  {
    "url": "assets/js/8.3ae275e8.js",
    "revision": "588283b7f3979a554c24d46ba30835c4"
  },
  {
    "url": "assets/js/9.c620737e.js",
    "revision": "348e1eb2de4d9b3a7245dcc385232449"
  },
  {
    "url": "assets/js/app.2749cedd.js",
    "revision": "8e540eb226891571c55b8a1e670c701c"
  },
  {
    "url": "configuration/db_alert.png",
    "revision": "7fc0ec38d14252764726ffc641549394"
  },
  {
    "url": "configuration/db.png",
    "revision": "139b4ac15a6a989a91712c47f79aed24"
  },
  {
    "url": "configuration/etcd.png",
    "revision": "7e84327cdf9bcf53a6e4934f7b022543"
  },
  {
    "url": "configuration/finish.png",
    "revision": "8f50abd02118b8080b6f3adcfa5433c2"
  },
  {
    "url": "configuration/jwt.png",
    "revision": "dec7a3c0d86cf80ef017e3c8f0d6974e"
  },
  {
    "url": "configuration/user.png",
    "revision": "7e6bf47e3c92e7951be6244cd5dfb7d4"
  },
  {
    "url": "cover.png",
    "revision": "7df9c35096d9cf836d215ba6c87e2f5e"
  },
  {
    "url": "developer/api.html",
    "revision": "a338d5311e437058e42038c9a6431d2c"
  },
  {
    "url": "developer/index.html",
    "revision": "87b5ce1648fb741c81aecc298c793607"
  },
  {
    "url": "index.html",
    "revision": "f89b8693750825a46720d1e90a242632"
  },
  {
    "url": "introduction/architecture.html",
    "revision": "effd5c82075bda9af75948b770186d84"
  },
  {
    "url": "introduction/configuration.html",
    "revision": "a42fca7797be0e3335e83f3f4003390c"
  },
  {
    "url": "introduction/dependencies.html",
    "revision": "9daca5224a9a6f225dd99f15f198a0b4"
  },
  {
    "url": "introduction/index.html",
    "revision": "d7027894e644d57d43500eeb9c0c337d"
  },
  {
    "url": "introduction/managerment.html",
    "revision": "4089720eb6d05a6c97eb55ad21a20075"
  },
  {
    "url": "introduction/more.html",
    "revision": "fe8ce28d003007844a5d8354d86951d9"
  },
  {
    "url": "introduction/services.html",
    "revision": "c56d116c3c54de796ec9e16e663d9220"
  },
  {
    "url": "log.png",
    "revision": "8666569baeaabdcdae6b4e1f45babd35"
  },
  {
    "url": "pipeline/bind_node.png",
    "revision": "b0a4a0fa89cf3502bc88f987e33afa13"
  },
  {
    "url": "pipeline/create_pipeline.png",
    "revision": "c696565e0826b747f7491766e96d3d33"
  },
  {
    "url": "pipeline/list.png",
    "revision": "d1a18fabdcc32494166787e00e6080a9"
  },
  {
    "url": "pipeline/node.png",
    "revision": "0b6f24960eaec065be77e567c478b45b"
  },
  {
    "url": "pipelines_relations.png",
    "revision": "7e4bd9b87e84717428227a0608548994"
  },
  {
    "url": "qrcode/alipay.jpg",
    "revision": "b2904660c399d465bdc6a92b794d1f4e"
  },
  {
    "url": "qrcode/wechatpay.jpg",
    "revision": "712ac818204c90c3ed2fdf4ea59f4442"
  },
  {
    "url": "task/create_task.png",
    "revision": "a678dbc168eff07bc6df7d6551d75aaf"
  },
  {
    "url": "task/list.png",
    "revision": "4ee2150ed5a03bddbce972092a6e0273"
  },
  {
    "url": "user.png",
    "revision": "3f76344994a3a3be584306127c90fde2"
  },
  {
    "url": "version/index.html",
    "revision": "15a562c7b683ce0aaba4ea2cace3b2c2"
  }
].concat(self.__precacheManifest || []);
workbox.precaching.suppressWarnings();
workbox.precaching.precacheAndRoute(self.__precacheManifest, {});
addEventListener('message', event => {
  const replyPort = event.ports[0]
  const message = event.data
  if (replyPort && message && message.type === 'skip-waiting') {
    event.waitUntil(
      self.skipWaiting().then(
        () => replyPort.postMessage({ error: null }),
        error => replyPort.postMessage({ error })
      )
    )
  }
})
