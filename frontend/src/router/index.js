import { createRouter, createWebHashHistory } from 'vue-router';

import Explore from '../components/Explore.vue';
import User from '../components/User.vue';
import Question from '../components/Question.vue';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import Pay from '../views/pay.vue';
import Submit from '../views/submit-question.vue';
import Answerer from '../components/Answerer.vue';
import AdminLogin from '../components/AdminLogin.vue';
import AdminQuestion from '../components/AdminQuestion.vue';
import Income from '../components/Income.vue';

const routes = [
  {
    path: '/',
    name: 'Explore',
    component: Explore,
    meta: {
      public: true,
    },
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: {
      public: true,
    },
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: {
      public: true,
    },
  },
  {
    path: '/user',
    name: 'User',
    component: User,
  },
  {
    path: '/question',
    name: 'Question',
    component: Question,
  },
  {
    path: '/answerer',
    name: 'Answerer',
    component: Answerer,
  },
  {
    path: '/pay',
    name: 'Pay',
    component: Pay,
  },
  {
    path: '/submit',
    name: 'Submit',
    component: Submit,
  },
  {
    path: '/admin',
    name: 'AdminLogin',
    component: AdminLogin,
    meta: {
      public: true,
    },
  },
  {
    path: '/admin/question',
    name: 'AdminQuestion',
    component: AdminQuestion,
    meta: {
      public: true,
    },
  },
  {
    path: '/income',
    name: 'Income',
    component: Income,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
router.beforeEach((to, from, next) => {
  if (!to.matched.some((record) => record.meta.public)) {
    if (window.localStorage.getItem('token') == null) {
      next({ name: 'Login' });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
