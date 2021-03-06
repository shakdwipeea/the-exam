(function () {
    'use strict';

    angular.module('student', ['ui.router', 'toaster', 'ngAnimate'])
        .config(function ($stateProvider, $urlRouterProvider) {
            $urlRouterProvider.otherwise('/account');

            var dir = '/public/src/student/';

            $stateProvider

                .state('account', {
                    url: '/account',
                    templateUrl: dir + 'account/account.tpl',
                    controller: 'AccountController as account'
                })

                .state('account.login', {
                    url: '/login',
                    templateUrl: dir + 'account/login/login.tpl',
                    controller: 'LoginController as login'
                })

                .state('account.signup', {
                    url: '/signup',
                    templateUrl: dir + 'account/signup/signup.tpl',
                    controller: 'SignUpController as signup'
                })

                .state('exam', {
                    url: '/exam',
                    templateUrl: dir + 'exam/exam.tpl',
                    controller: 'ExamController as exam'
                })

                .state('test', {
                    url: '/test/:id',
                    templateUrl: dir + 'test/test.tpl',
                    controller: 'TestController as test'
                })

                .state('result', {
                    url: '/result',
                    templateUrl: dir + 'result/result.tpl',
                    controller: 'ResultController as result'
                })

                .state('leaderboards', {
                  url: '/leader/:id',
                  templateUrl: dir + 'leaderboards/leaderboards.tpl',
                  controller: 'LeaderController as leader'
                })

            ;
        });

})();

/**

 git rm -r --cached .

 **/
