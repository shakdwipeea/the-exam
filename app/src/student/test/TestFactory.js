/**
 * Created by akash on 1/8/15.
 */
(function () {
    'use strict'

    angular.module('student')
        .factory('Test', function (Result) {
            var _questions = [],
                _current = 0;

            return {

                setQuestions: function (questions) {
                    console.log('et que', questions)
                    _questions = questions;
                    _current = 0;
                },

                getQuestion: function (id) {
                    console.log('Getting quetion', _questions);
                    var q = {};

                    for (var i = 0; i < _questions.length; i++) {
                        if (_questions[i].Id == id) {
                            q = _questions[i];
                            break;
                        }
                    }

                    return q;
                },

                answer: function (question, answer) {
                    var index = _questions.indexOf(question);

                    if (index === -1) {
                        return false;
                    }

                    _questions[index].answer = answer;
                    return true;

                },

                next: function () {
                    if (_current == _questions.length - 1) {
                        Result.setQuestions(_questions);
                        return false;
                    } else {
                        _current++;
                        return _questions[_current].Id;
                    }
                }
            }
        });
})();