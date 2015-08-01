(function  () {
	'use strict'

	angular.module('student')
		.controller('SignUpController', function  (Account, toaster, $state) {
			var self = this;
			self.message = "";
			self.checkUserName = function () {
				console.log("Function called");
				Account.getUserNames()
						.then(function (response) {

							if (!response.data) {
								self.enable = true;
							} else {

								var index = response.data.indexOf(self.username);

								if (index !== -1) {
									self.message = "Username taken";
									self.enable = false;
								} else {
									self.message = "";
									self.enable = true;
								}
							}

						})
			};

			self.submit =function  () {
				console.log(self);

				var groupsFinal = [];

				var groups = self.group;
				for (var prop in groups) {
					if (groups[prop]) {

						switch(prop) {
							case 'ele': prop = "11th"
										break;

							case 'twe': prop = "12th"
										break;
						}

						groupsFinal.push(prop);

					}
				}
				
				var data = {
					username: self.username,
					password: self.password,
					groups: groupsFinal
				};

				Account.signUp(data)
					.then(function  (response) {

						//Signed up now sign in

						Account.login({
							username: data.username,
							password: data.password
						}).then(function () {
							toaster.pop('success', 'Hurray', "Logged In");
							$state.go('exam');
						}).catch(function (reason) {
							toaster.pop('error', 'Error', reason.data.err);
						})
					})
					.catch(function  (reason) {
						console.log(reason);
						toaster.pop('error', 'Error', reason.data.err);
					});
			}
		});
})();