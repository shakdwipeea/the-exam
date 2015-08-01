/**
 * Created by akash on 1/8/15.
 */
(function () {
    'use strict'

    angular.module('student')
        .factory('Result', function ($http, Account) {
            var _questions = [],
                test_id = "";

            function save(score) {
                var response = [];
                response = _questions.map(function (q) {
                    return {
                        id: q.Id,
                        answer: q.answer
                    }
                });

                $http.post("/saveResult", {
                    token: Account.getToken(),
                    response: response,
                    score: score + "",
                    test_id: test_id
                }).then(function (response) {
                    console.log(response);
                }).catch(function (reason) {
                    console.log(reason);
                    throw new Error("Could not save Result");
                })
            }

            return {
                setId: function (id) {
                    test_id = id;
                },

                setQuestions: function (questions) {
                    _questions = questions;
                },

                evaluate: function () {
                    console.log(_questions);
                    var totalScore = 0,
                        yourScore = 0;

                    _questions.forEach(function (question) {
                        totalScore += 4;
                        if (question.answer == question.Correct) {
                            yourScore += 4;
                        }
                    });

                    save(yourScore);

                    return {
                        total: totalScore,
                        scored: yourScore
                    }

                }
            }
        });
})();