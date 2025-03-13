module.exports = {
  title: 'Elastic Crontab System',
  description: '一款简单易用的分布式定时任务管理系统',
  serviceWorker: true,
  markdown: {
    lineNumbers: true
  },
  themeConfig: {
    sidebarDepth: 3,
    nav: [
      { text: '项目介绍', link: '/introduction/architecture' },
      { text: '开发文档', link: '/developer/' },
      { text: '下载', link: 'https://github.com/betterde/ects/releases' },
      { text: 'Github', link: 'https://github.com/betterde/ects' },
    ],
    sidebar: {
      '/introduction/': [
        'architecture',
        'dependencies',
        'configuration',
        'services',
        'managerment',
        'more'
      ],
      '/developer/': [
        '',
        'api'
      ]
    }
  },
  base: '/ects/',
  dest: '../docs'
};
