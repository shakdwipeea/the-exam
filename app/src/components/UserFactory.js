/**
 * Created by akash on 22/7/15.
 */
angular.module('question')
.constant('Host', {
        add: 'http://127.0.0.1:3000'
    })
.factory('User', function ($http, Host) {
        var token = null,
            subject = null,
            username = null;

        var login = function (userData) {
            console.log(userData);
           return $http.post(Host.add + '/login', userData)
                .then(function (response) {
                    console.log(response);
                    token = response.data.token;
                    subject = response.data.teacher.Subject;
                    username = response.data.teacher.Username;
                    return response;
                })
                .catch(function (error) {
                    console.log(error);
                    return error;
                });
        };

        var addQuestion = function  (questionData) {
            console.log('Q in addQues factory',questionData);

            if (token && subject && username) {
                var requestBody = {
                    questionText: questionData.questionText,
                    option1: questionData.option1,
                    option2: questionData.option2,
                    option3: questionData.option3,
                    option4: questionData.option4,
                    token: token,
                    subject: subject
                };

                console.log(requestBody);
                return $http.post(Host.add + '/add_question', requestBody);
            } else {
                console.log('This hould not be happening');
                return new Error("Go to hell");
            }

        }

        return {
            login: login,
            add: addQuestion
        }
});