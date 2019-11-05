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
    "revision": "54eba0e2abfc3a1844153c4b8ffbdd65"
  },
  {
    "url": "architecture.png",
    "revision": "46d3e361c7c3ce6509950651d8b47a8a"
  },
  {
    "url": "assets/css/0.styles.4b878e04.css",
    "revision": "ffd12f6bec255f7eeba8a188db8025dd"
  },
  {
    "url": "assets/img/search.83621669.svg",
    "revision": "83621669651b9a3d4bf64d1a670ad856"
  },
  {
    "url": "assets/js/10.a9b3e2ed.js",
    "revision": "603f6fb3b08bff08d96e4bffe657e7bd"
  },
  {
    "url": "assets/js/11.8d62d449.js",
    "revision": "efb6d66c15fff4a0d5d26a4d64bbbe9c"
  },
  {
    "url": "assets/js/12.f9d3cea2.js",
    "revision": "1ac4034ddf97e84d590391f9de9a2713"
  },
  {
    "url": "assets/js/13.de2731d4.js",
    "revision": "d2a8f2917f48aad26adec4cf7e345dd7"
  },
  {
    "url": "assets/js/2.f8f0e8d9.js",
    "revision": "7de301f1c7cc3c474642ba67365ad9aa"
  },
  {
    "url": "assets/js/3.5a2a6f56.js",
    "revision": "49db9933116b5e32c8cc8f2ac51b5c87"
  },
  {
    "url": "assets/js/4.fc5cf636.js",
    "revision": "7260de8b78c377582959bd6fe190a758"
  },
  {
    "url": "assets/js/5.3002656a.js",
    "revision": "e440725f2a35b8c7c333881b0ae2a3f1"
  },
  {
    "url": "assets/js/6.fe2e5ce2.js",
    "revision": "9cf8c1968da15f97dfcd04c9c7847381"
  },
  {
    "url": "assets/js/7.fbbe362c.js",
    "revision": "15c59ae657978a8fa5c3d1d33217eff5"
  },
  {
    "url": "assets/js/8.50df79d7.js",
    "revision": "3b068b99057a151f6403e6cabfefdf0c"
  },
  {
    "url": "assets/js/9.4a0eec1b.js",
    "revision": "b6f8560da2a27caf982d2262af6fe608"
  },
  {
    "url": "assets/js/app.dc03944a.js",
    "revision": "f150bce080750363100ff75d859895ee"
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
    "revision": "2b8b1e1c7eeb35ae261c0e88ae1e1514"
  },
  {
    "url": "developer/index.html",
    "revision": "6feed1f2e2832b3ddac616848486a4a6"
  },
  {
    "url": "index.html",
    "revision": "12457a488f0304fe73b333516dffafab"
  },
  {
    "url": "introduction/architecture.html",
    "revision": "1f0d92340e3abeeca2cb69477a39d7f3"
  },
  {
    "url": "introduction/configuration.html",
    "revision": "f4c00f19fd21192fad2cfd2968884281"
  },
  {
    "url": "introduction/dependencies.html",
    "revision": "2f5a69ead0ef992597c2e4ac54be0504"
  },
  {
    "url": "introduction/managerment.html",
    "revision": "8790c82ece0da1f10c1a398d2bdb0c83"
  },
  {
    "url": "introduction/more.html",
    "revision": "64d21c4622e6aec1a05af645aa14c951"
  },
  {
    "url": "introduction/services.html",
    "revision": "c2591d479c3eba20ce21afcf37ac2bc1"
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
    "revision": "fa2ff94516c3dfab1265fc92a1fa9885"
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
