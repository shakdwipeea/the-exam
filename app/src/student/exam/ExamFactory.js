/**
 * Created by akash on 1/8/15.
 */
(function () {
    'use strict';

    angular.module('student')
        .factory('Exam', function ($http, Account, Test) {
            var exams = [],
                questions = [];

            return {
                getExams: function () {
                    return $http({
                        method: 'GET',
                        url: '/getTests',
                        params: {
                            token: Account.getToken()
                        }
                    }).then(function (response) {
                        exams = response.data.tests;
                        return response;
                    })
                },

                getCompleteTest: function (id) {
                    return $http({
                        method: 'GET',
                        url: '/secure/test/' + id,
                        params: {
                            token: Account.getToken()
                        }
                    }).then(function (response) {
                        Test.setQuestions(response.data.questions);
                        return response;
                    })
                }
            }
        });
})();