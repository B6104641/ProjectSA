import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateUser from './components/Users';
import Show from './components/Show';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/user', CreateUser);
    router.registerRoute('/result', Show);
  },
});