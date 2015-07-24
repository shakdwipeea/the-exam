(function () {
	'use strict';

	angular.module('question',['ui.router', 'toaster', 'ngAnimate'])
			.config( function  ($stateProvider, $urlRouterProvider) {
				$urlRouterProvider.otherwise('/login');

				$stateProvider
						.state('main', {
							url:'/login',
							templateUrl:'src/login/login.tpl',
							controller:'LoginController as login'
						})

						.state('dash', {
							url: '/dash',
							templateUrl: 'src/dash/dash.tpl',
							controller: 'DashController as dash'
						})
						.state('dash.welcome', {
							url:'/welcome',
							templateUrl:'src/dash/welcome/welcome.tpl',
							controller: 'WelcomeController as welcome'
						})
						.state('dash.add', {
							url: '/add',
							templateUrl: 'src/dash/addQuestion/add.tpl',
							controller: 'AddController as add'
						})
				;
			});

})();

/**

git rm -r --cached .

**/
