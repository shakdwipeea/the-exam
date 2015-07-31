/**
 * Created by akash on 22/7/15.
 */

angular.module('question')
    .controller('DashController', function ($state, User) {
        console.log("dash Controller");

        if (!User.isLoggedIn()) {
            $state.go('main');
            return
        }

        $state.go('dash.welcome');

        var dash = this;
        dash.add = function () {
            console.log('Trying to go');
            $state.go('dash.add');
        };

        dash.questionTemp = false;
        dash.testTemp = false;

     });