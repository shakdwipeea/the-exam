/**
 * Created by akash on 22/7/15.
 */

angular.module('question')
    .controller('DashController', function ($state, User) {
        console.log("dash Controller");
        var dash = this;
        dash.add = function () {
            console.log('Trying to go');
            $state.go('dash.add');
        }

        if(!User.isLoggedIn()) {
            $state.go('main');
        } else {
            $state.go('dash.welcome');
        }
     });