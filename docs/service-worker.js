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
    "revision": "2e0eba35e8a930f1759c3cb79dbebc09"
  },
  {
    "url": "architecture.png",
    "revision": "46d3e361c7c3ce6509950651d8b47a8a"
  },
  {
    "url": "assets/css/0.styles.66cb41a9.css",
    "revision": "71c0212b0f65d04cbdcf821388a5dad8"
  },
  {
    "url": "assets/img/search.83621669.svg",
    "revision": "83621669651b9a3d4bf64d1a670ad856"
  },
  {
    "url": "assets/js/10.0cd891f6.js",
    "revision": "75169689de2658541b79c730bc1964a3"
  },
  {
    "url": "assets/js/11.c1439af8.js",
    "revision": "e9b0177ec513fbe1117cffeff01942d0"
  },
  {
    "url": "assets/js/12.93e1ec7f.js",
    "revision": "efac44b52682855112482c4ed404d0a9"
  },
  {
    "url": "assets/js/13.ac6343ad.js",
    "revision": "54abe441186bcc8b9e6150a83672bcfc"
  },
  {
    "url": "assets/js/2.890c2903.js",
    "revision": "f036d80cd49aebad3c482662897c60e4"
  },
  {
    "url": "assets/js/3.b04656e7.js",
    "revision": "c7702f2cfc687f260fb943698c41f666"
  },
  {
    "url": "assets/js/4.b4c8d6b9.js",
    "revision": "f938e355b2db5d7d12f7c30f0f7012aa"
  },
  {
    "url": "assets/js/5.ada6e86d.js",
    "revision": "be6a6623bbf878c942633128b5c6d8a1"
  },
  {
    "url": "assets/js/6.8e603eb7.js",
    "revision": "73adfe36ae980ed0bab5d43bbf0db46f"
  },
  {
    "url": "assets/js/7.0d65a2e7.js",
    "revision": "455f13fec87b5790f3dd4e107fffa84f"
  },
  {
    "url": "assets/js/8.7e31ce14.js",
    "revision": "c5f5349846f878ac3d7a3d2406ee93b2"
  },
  {
    "url": "assets/js/9.73afb34e.js",
    "revision": "3cc80cee5ae1db51a80eba8552d8c6dc"
  },
  {
    "url": "assets/js/app.87168bd6.js",
    "revision": "ae6da47850a68740552a3f62d177d49f"
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
    "revision": "562d10173c2ccd4cf70eb99960d16fe7"
  },
  {
    "url": "developer/index.html",
    "revision": "eccbdb02345170fe321227481164ccc7"
  },
  {
    "url": "index.html",
    "revision": "54fe55168f73f2ad8bdcc63e23b0807e"
  },
  {
    "url": "introduction/architecture.html",
    "revision": "ae68546ec86d7a1e75a50936b0d32990"
  },
  {
    "url": "introduction/configuration.html",
    "revision": "986cf29b57531b7639fc496cbea2e3b5"
  },
  {
    "url": "introduction/dependencies.html",
    "revision": "a4557e5a680dd0684e0b8e9773a3dc20"
  },
  {
    "url": "introduction/managerment.html",
    "revision": "834dcba13fdeca87c917fb2e2e80b4cb"
  },
  {
    "url": "introduction/more.html",
    "revision": "5776e3df7f4074eba516aa7fd74fcaf4"
  },
  {
    "url": "introduction/services.html",
    "revision": "9de182b5e057c7cbb09025a11167d7bb"
  },
  {
    "url": "jetbrains.svg",
    "revision": "856cab912ded136de723e526c8ab5f04"
  },
  {
    "url": "log.png",
    "revision": "8666569baeaabdcdae6b4e1f45babd35"
  },
  {
    "url": "logo_mini.png",
    "revision": "2ba17246e1d910de78474dc2374dc96a"
  },
  {
    "url": "logo.png",
    "revision": "395c87f17ea4aed9c29247b57e5e34b6"
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
    "revision": "fd382c3b2e6887a33507c15e0f63addf"
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
