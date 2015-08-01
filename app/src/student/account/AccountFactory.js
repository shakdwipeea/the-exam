(function  () {
	'use strict'

	angular.module('student')
		.factory('Account', function  ($http) {
			var token = null;

			function signUp (data) {
				return $http.post('/studentSignup', data);
			}

			function login (data) {
				return $http.post('/studentLogin', data)
						.then(function (response) {
							console.log(response);
							token = response.data.token;
							return response;
						})
			}

			function usernames () {
				return $http.get('/usernames');
			}

			function getToken () {
				return token;
			}

			return {
				signUp: signUp,
				login: login,
				getUserNames: usernames,
				getToken: getToken
			}
		});
})();