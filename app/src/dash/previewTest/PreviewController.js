/**
 * Created by akash on 27/7/15.
 */

/**
 * Controller for page to preview the test
 */
angular.module('question')
    .controller('PreviewController', ['User', 'toaster', '$state', '$stateParams', 'Test',
        function (User, toaster, $state, $stateParams, Test) {
            console.log('Preview Controller');

            var test,
                questions,
                ids,
                self = this;

            self.questions = [];

            if ($stateParams.testId) {
                //if test id exist then viewing an existing id
                console.log($stateParams.testId);
                var testId = $stateParams.testId;

                //get details for this id
                Test.getTestById(testId)
                    .then(function (response) {
                        self.name = response.data.test.name;
                        self.submitButtonText = "Edit test";
                        self.questions = response.data.questions;
                    })
                    .catch(function (reason) {
                        console.log('Why not here', reason);
                        toaster.pop('error', 'OOPS sorry', "not there exactly")
                    })
            } else {
                //new test
                test = User.getTest();
                questions = User.getQuestions();


                ids = test.ids;


                self.name = test.name;
                self.questions = [];
                self.submitButtonText = "Create Test";

                if (Array.isArray(questions)) {
                    ids.forEach(function (id) {
                        self.questions.push(questions[id]);
                        console.log('Q Q')
                    });
                }
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
                            $state.go('dash.view');
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