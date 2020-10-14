module.exports = {
  title: 'Gommitizen',
  tagline: 'A git utility written in go',
  url: 'https://gommitizen.epsxy.xyz',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  favicon: 'img/favicon.ico',
  organizationName: 'epsxy', // Usually your GitHub org/user name.
  projectName: 'gommitizen', // Usually your repo name.
  themeConfig: {
    navbar: {
      title: 'Gommitizen',
      logo: {
        alt: 'Gommitizen Logo',
        src: 'img/logo.png',
      },
      items: [
        {
          to: 'docs/introduction',
          activeBasePath: 'docs',
          label: 'Docs',
          position: 'left',
        },
        {
          href: 'https://github.com/epsxy/gommitizen',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Introduction',
              to: 'docs/introduction',
            },
            {
              label: 'Quickstart',
              to: 'docs/installation/quickstart/',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/epsxy/gommitizen',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Gommitizen. Built with Docusaurus.`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl:
            'https://github.com/epsxy/gommitizen/edit/master/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
