/**
 * Created by akash on 1/8/15.
 */
(function () {
    'use strict';

    angular.module('student')
        .controller('ExamController', function (Exam, toaster) {
            var self = this;

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
                    })
                    .catch(function (reason) {
                        console.log(reason);
                        toaster.pop('error', 'Error Occured', reason.data.err)
                    })
            }
        });
})();