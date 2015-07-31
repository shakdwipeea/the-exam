(function () {
	'use strict';

	angular.module('question',['ui.router', 'toaster', 'ngAnimate'])
			.config( function  ($stateProvider, $urlRouterProvider) {
				$urlRouterProvider.otherwise('/main');

				$stateProvider
						.state('home', {
							url: '/main',
							templateUrl: '/public/src/admin/home/home.tpl'
						})
						.state('main', {
							url:'/login',
							templateUrl: '/public/src/admin/login/login.tpl',
							controller:'LoginController as login'
						})

						.state('dash', {
							url: '/dash',
							templateUrl: '/public/src/admin/dash/dash.tpl',
							controller: 'DashController as dash'
						})
						.state('dash.welcome', {
							url:'/welcome',
							templateUrl: '/public/src/admin/dash/welcome/welcome.tpl',
							controller: 'WelcomeController as welcome'
						})
						.state('dash.add', {
							url: '/add',
							templateUrl: '/public/src/admin/dash/addQuestion/add.tpl',
							controller: 'AddController as add'
						})
						.state('dash.createTest', {
							url: '/test',
							templateUrl: '/public/src/admin/dash/createTest/createTest.tpl',
							controller: 'CreateTestController as test'
						})
						.state('dash.preview', {
							url: '/preview/:testId',
							templateUrl: '/public/src/admin/dash/previewTest/preview.tpl',
							controller: 'PreviewController as preview'
						})
						.state('dash.view', {
							url: '/view',
							templateUrl: '/public/src/admin/dash/viewTests/view.tpl',
							controller: 'ViewController as view'
						})
				;
			});

})();

/**

git rm -r --cached .

**/
