/**
 * Created by akash on 27/7/15.
 */

angular.module('question')
    .factory('Test', ['$http', 'User', 'Host', function ($http, User, Host) {
        var tests = [];

        return {
            getTest: function () {
                return $http({
                    method: 'GET',
                    url: Host.add + '/secure/test',
                    params: {
                        token: User.getToken()
                    }
                }).then(function (response) {
                    if (!response.data.err) {
                        tests = response.data.tests;
                        return response;
                    }
                }).catch(function (reason) {
                    console.log(reason);
                    return reason;
                })
            },

            getTestById: function (id) {
                return $http({
                    method: 'GET',
                    url: Host.add + '/secure/test/' + id,
                    params: {
                        token: User.getToken()
                    }
                }).then(function (response) {
                    console.log(response);
                    return response;
                }).catch(function (reason) {
                    console.log(reason);
                    return reason;
                })
            },

            enableTest: function (id, enable) {
                return $http({
                    method: 'GET',
                    url: Host.add + '/secure/enable/' + id,
                    params: {
                        token: User.getToken(),
                        enable: enable
                    }
                });
            }
        }
    }]);