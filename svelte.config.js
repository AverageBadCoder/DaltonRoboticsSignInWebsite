import adapter from '@sveltejs/adapter-static';
import sveltePreprocess from 'svelte-preprocess';

const dev = process.argv.includes('dev');

const config = {
  preprocess: sveltePreprocess(),
  kit: {
    adapter: adapter({
      pages: 'build',
      assets: 'build',
      fallback: 'index.html'
    }),
    paths: {
      base: dev ? '' : '/DaltonRoboticsSignInWebsite'
    },
    prerender: {
      handleHttpError: 'warn'
    }
  }
};

export default config;
