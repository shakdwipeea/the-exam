/**
 * Created by akash on 1/8/15.
 */
(function () {
    'use strict'

    angular.module('student')
        .controller('ResultController', function ($state, Account, Result) {

            if (!Account.getToken()) {
                $state.go('account.login');
                return;
            }

            var score = Result.evaluate();

            var self = this;
            self.total = score.total;
            self.scored = score.scored;
        });
})();