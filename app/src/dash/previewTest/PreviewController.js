/**
 * Created by akash on 27/7/15.
 */

/**
 * Controller for page to preview the test
 */
angular.module('question')
    .controller('PreviewController', ['User', 'toaster', '$state', function (User, toaster, $state) {
        console.log('Preview Controller');
        var test = User.getTest();
        var questions = User.getQuestions();

        var ids = test.ids;

        var self = this;
        self.name = test.name;
        self.questions = [];
        self.submitButtonText = "Create Test";

        if (Array.isArray(questions)) {
            ids.forEach(function (id) {
                self.questions.push(questions[id]);
                console.log('Q Q')
            });
        }


        self.createTest = function () {
            self.submitButtonText = "Creating.........";
            var prom = User.createTest();

            if (typeof prom.then === "function") {
                prom
                    .then(function (response) {
                        console.log(response);
                        if (!response.data.err) {
                            self.submitButtonText = "Done";
                            toaster.pop('success', 'Success', 'Test created');
                        } else {
                            toaster.pop('error', 'Error', 'Sth happened');
                        }
                        $state.go('dash.welcome');
                    })
                    .catch(function (reason) {
                        console.log('Error hogaya', reason);
                        self.submitButtonText = "Create test";
                        if (reason.data.err) {
                            toaster.pop('error', 'Error', reason.data.err);
                        } else {
                            toaster.pop('error', 'Error', reason.statusText);
                        }
                    })
            }

        }
    }]);