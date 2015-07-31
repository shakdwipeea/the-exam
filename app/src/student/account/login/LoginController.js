/**
 * Created by akash on 31/7/15.
 */
(function () {
    'use strict'

    angular.module('student')
        .controller('LoginController', function (Account, toaster, $state) {
            var self = this;

            self.submit = function () {
                Account.login({
                    username: self.username,
                    password: self.password
                }).then(function () {
                    toaster.pop('success', 'Hurray', 'Logged In');
                    $state.go('exam');
                }).catch(function (reason) {
                    toaster.pop('error', 'Error', reason.data.err);
                })
            }
        });
})();