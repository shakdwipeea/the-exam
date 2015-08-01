/**
 * Created by akash on 1/8/15.
 */

(function () {
    'use strict'

    angular.module('student')
        .controller('TestController', function ($stateParams, Test, $state, Account, toaster) {
            var self = this;
            var curId = $stateParams.id;

            if (!Account.getToken()) {
                $state.go('account.login');
                return;
            }

            self.question = Test.getQuestion(curId);
            console.log('is', self.question, curId);

            self.next = function () {
                console.log(self);
                var valid = Test.answer(self.question, self.myAnswer);

                if (valid) {
                    var nextId = Test.next();
                    if (nextId) {
                        $state.go('test', {
                            id: nextId
                        })
                    } else {
                        console.log('Complete');
                        toaster.pop('success', 'Success', 'Test Complete');
                        $state.go('result');
                    }
                } else {
                    toaster.pop('error', 'Error', 'Its not working');
                }
            }
        });
})();