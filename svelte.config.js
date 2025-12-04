import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';

const dev = process.env.NODE_ENV === 'development';

export default {
  preprocess: preprocess(),
  kit: {
    adapter: adapter(),
    // ensure the base path used on GH Pages matches your repo name
    paths: {
      base: dev ? '' : '/DaltonRoboticsSignInWebsite'
    },
    appDir: '_app'
  }
};
