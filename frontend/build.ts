const { build } = require('esbuild');
const esbuildPluginTsc = require('esbuild-plugin-tsc');
const buildParams = {
  color: true,
  assetNames: 'assets/[name]-[hash]',
  outdir: 'build',

  entryPoints: [
    './frontend/modules/**/*.ts',
    './frontend/modules/**/*.tsx',
    './frontend/main.ts',
  ],
  bundle: true,
  minify: true,
  format: 'cjs',
  sourcemap: true,
  loader: {
    '.js': 'jsx',
    '.ts': 'tsx',
    '.json': 'json',
    '.png': 'file',
    '.jpeg': 'file',
    '.jpg': 'file',
    '.svg': 'file',
  },
  plugins: [
    esbuildPluginTsc({
      force: true,
    }),
  ],
};

build(buildParams)
  .then((res) => {
    console.log(res);
  })
  .catch((err) => {
    console.log(err);
  });
