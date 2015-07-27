/**
 * Created by akash on 27/7/15.
 */

/**
 * Controller for page to preview the test
 */
angular.module('question')
    .controller('PreviewController', ['User', function (User) {
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

            /**
             * todo implement
             */
            User.createTest();
            throw new Error("Not implemented");

        }
    }]);