/**
 * Created by akash on 26/7/15.
 */

angular.module('question')
    .controller('CreateTestController',
    ['$state', 'User', 'toaster', function ($state, User, toaster) {
        console.log("Create Test Controller");
        var self = this;

        var promise = User.getQuestions();

        if (typeof promise.then === "function") {
            promise
                .then(function (response) {
                    console.log('hi', response);
                    self.questions = response.data.questions;
                })
                .catch(function (reason) {
                    toaster.pop('error', 'Error ocurred', 'Oh boy!! ' + reason.data.err);
                    console.log('HOHO', reason);
                });
        } else {
            self.questions = promise;
        }


        self.preview = function () {
            var ids = [];
            //get name of test
            if (!self.name || self.checked.length === 0) {
                toaster.pop('error', 'No name', 'Specify a name for the test');
                return;
            }
            console.log('U should not be here');
            // get the checked questions
            console.log(self.checked);

            for (var questionIndex in self.checked) {
                if (self.checked.hasOwnProperty(questionIndex)) {
                    console.log(questionIndex);
                    if (self.checked[questionIndex]) {
                        ids.push(questionIndex);
                    }
                }
            }

            // store the checked questions in a factory
            User.setTest(self.name, self.group, ids);

            // redirect to preview page
            $state.go('dash.preview');
        }
    }]);