import Application from '@ember/application';
import Resolver from 'ember-resolver';
import loadInitializers from 'ember-load-initializers';
import config from 'wildfire/config/environment';

export default class App extends Application {
  /* eslint-disable */
  modulePrefix = config.modulePrefix; // eslint-disable-line
  podModulePrefix = config.podModulePrefix; // eslint-disable-line
  Resolver = Resolver;
  /* eslint-enable */
}

loadInitializers(App, config.modulePrefix);
