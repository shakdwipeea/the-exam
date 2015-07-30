/**
 * Created by akash on 27/7/15.
 */
angular.module('question')
    .controller('ViewController', ['Test', 'toaster', function (Test, toaster) {
        var self = this;

        Test.getTest()
            .then(function (response) {
                console.log(response);
                self.tests = response.data.tests;
            })
            .catch(function (reason) {
                console.log(reason);
                var msg = reason.data.err || reason.statusText;
                toaster.pop('error', 'Not happening', msg);
            });

        self.EnableTest = function ($index, testId, enable) {
            console.log(testId, enable);
            //Call the api to enabled/disable the test
            Test.enableTest(testId, enable)
                .then(function (response) {
                    console.log(response);
                    //on success give a proper growl
                    var message = enable == true ? "disabled" : "enabled";
                    toaster.pop('success', 'Test ' + message);

                    // update text in the current scope
                    self.tests[$index].Enable = !enable;
                })
                .catch(function (reason) {
                    console.log(reason);
                    toaster.pop('error', 'Error', reason.statusText);
                })
        }
    }]);