/**
 * Created by akash on 1/8/15.
 */
(function () {
    'use strict';

    angular.module('student')
        .controller('ExamController', function (Exam, toaster, Account, $state) {
            var self = this;

            if (!Account.getToken()) {
                $state.go('account.login');
                return;
            }

            Exam.getExams()
                .then(function (response) {
                    console.log(response);
                    self.tests = response.data.tests;
                })
                .catch(function (reason) {
                    console.log(reason);
                    toaster.pop('error', 'Error Occured', reason.data.err);
                });

            self.getTestDetail = function (id) {
                Exam.getCompleteTest(id)
                    .then(function (response) {
                        console.log(response);
                        $state.go('test', {
                            id: response.data.questions[0].Id
                        })
                    })
                    .catch(function (reason) {
                        console.log(reason);
                        toaster.pop('error', 'Error Occured', reason.data.err)
                    })
            }
        });
})();