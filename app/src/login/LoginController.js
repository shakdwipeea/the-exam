/**
 * Created by akash on 22/7/15.
 */

angular.module('question')
    .controller('LoginController', function (User, $state, toaster) {
        console.log("main Controller");

        var self = this;
        self.loginText = "Submit";
        self.doLogin = function () {
            self.loginText = "Please Wait......."
            console.log("Username & Password", self.user);
            User.login(self.user)
                .then(function (response) {
                    console.log(response.data.msg);
                    if (response.data.err == false) {
                        console.log('Toast');
                        toaster.pop('success', 'Success', 'Verified');
                        $state.go('dash');
                    } else {
                       console.log("Nope", response);
                        toaster.pop('error', 'Error ocurred', 'OOps Try Again');
                        self.loginText = "Submit";
                    }

                })
                .catch(function (error) {
                    console.log(error.data.msg);
                    toaster.pop('error', 'Error', 'Damn Boy Damn!!');
                    self.loginText = "Submit";
                });
        }
    });