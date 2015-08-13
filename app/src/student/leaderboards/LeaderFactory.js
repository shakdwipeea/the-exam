(function () {
    'use strict'

    angular.module('student')
        .factory('Leader', function ($http) {
            return {
              getLeaders: function (testId) {
                return $http.get('/leaderboards/' + testId);
              }
            }
        })
})();
