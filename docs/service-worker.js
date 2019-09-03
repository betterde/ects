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
    "revision": "d5143f16abb4aa4614b35b9c6b524590"
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
    "url": "assets/js/10.1b389d6c.js",
    "revision": "9b8b4c6df7017a6c8ada3c6aac40f673"
  },
  {
    "url": "assets/js/11.fda4346e.js",
    "revision": "181930ac1198d2fed031747915b1bf48"
  },
  {
    "url": "assets/js/12.d8f832d5.js",
    "revision": "8e2fd6f948dee967851e2d292f07932e"
  },
  {
    "url": "assets/js/13.ca1e11ec.js",
    "revision": "d2a8f2917f48aad26adec4cf7e345dd7"
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
    "url": "assets/js/6.1dcb0c51.js",
    "revision": "faa51e8e1259cd9f76157d5743049e3c"
  },
  {
    "url": "assets/js/7.d14b781a.js",
    "revision": "fd709d05938d57086c0c1f28ce9fd243"
  },
  {
    "url": "assets/js/8.96975e9c.js",
    "revision": "d23f991797860895ccf61614551750e1"
  },
  {
    "url": "assets/js/9.e349a620.js",
    "revision": "dcf9c18d59bf9b11875545257e00a38c"
  },
  {
    "url": "assets/js/app.86143f7c.js",
    "revision": "abd9f7ee93181ae9e04a1666aa8dbc0c"
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
    "revision": "7d807a7ec253e63de84588f65a0759f1"
  },
  {
    "url": "developer/index.html",
    "revision": "7f4badb075bf91668be6418b5971768b"
  },
  {
    "url": "index.html",
    "revision": "d6e87d5557d5eb349d54a67a261f87bc"
  },
  {
    "url": "introduction/architecture.html",
    "revision": "fd55af5d4cc42fa39aefae3bd5aefc66"
  },
  {
    "url": "introduction/configuration.html",
    "revision": "acd8164971a63ca05c6630c0c11d28d4"
  },
  {
    "url": "introduction/dependencies.html",
    "revision": "a89dab88461e274cd941e77b69fe9c27"
  },
  {
    "url": "introduction/managerment.html",
    "revision": "eb86102cabe126fadb37195269a56a7e"
  },
  {
    "url": "introduction/more.html",
    "revision": "cf1e6767ab0deba195b0a7ff890b7086"
  },
  {
    "url": "introduction/services.html",
    "revision": "779c7d4a091dced37bc25b0424872759"
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
    "revision": "d883f3451c7606c7839500a2de826033"
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
