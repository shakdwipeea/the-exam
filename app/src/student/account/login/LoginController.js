/**
 * Created by akash on 31/7/15.
 */
(function () {
    'use strict'

    angular.module('student')
        .controller('LoginController', function (Account, $state, $rootScope) {
            var self = this;

            self.submit = function () {
                Account.login({
                    username: self.username,
                    password: self.password
                }).then(function () {
                    $state.go('exam');
                }).catch(function (reason) {
                    console.log(reason);
                    self.Message = "Incorrect Username or Password";
                })
            }
        });
})();
