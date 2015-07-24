(function () {
	'use strict';

	angular.module('question',['ui.router', 'toaster', 'ngAnimate'])
			.config( function  ($stateProvider, $urlRouterProvider) {
				$urlRouterProvider.otherwise('/main');

				$stateProvider
						.state('home', {
							url: '/main',
							templateUrl: '/public/src/home/home.tpl'
						})
						.state('main', {
							url:'/login',
							templateUrl: '/public/src/login/login.tpl',
							controller:'LoginController as login'
						})

						.state('dash', {
							url: '/dash',
							templateUrl: '/public/src/dash/dash.tpl',
							controller: 'DashController as dash'
						})
						.state('dash.welcome', {
							url:'/welcome',
							templateUrl: '/public/src/dash/welcome/welcome.tpl',
							controller: 'WelcomeController as welcome'
						})
						.state('dash.add', {
							url: '/add',
							templateUrl: '/public/src/dash/addQuestion/add.tpl',
							controller: 'AddController as add'
						})
				;
			});

})();

/**

git rm -r --cached .

**/
