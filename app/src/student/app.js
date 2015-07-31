(function () {
    'use strict';

    angular.module('student', ['ui.router', 'toaster', 'ngAnimate'])
        .config(function ($stateProvider, $urlRouterProvider) {
            $urlRouterProvider.otherwise('/account');

            var dir = '/public/src/student/';

            $stateProvider

                .state('account', {
                    url: '/account',
                    templateUrl: dir + 'account/account.tpl'
                })

            ;
        });

})();

/**

 git rm -r --cached .

 **/
