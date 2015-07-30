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
            })
    }]);