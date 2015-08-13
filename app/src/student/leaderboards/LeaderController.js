(function () {
    'use strict'

    angular.module('student')
        .controller('LeaderController', function ($stateParams, Leader) {
          var self = this;

          Leader.getLeaders($stateParams.id)
            .then(function (response) {
              console.log(response);
              self.results = response.data.results;

              if(self.results) {
                self.results.sort(function (a, b) {
                  return b.Score - a.Score;
                })
              }

            })
            .catch(function (reason) {
              console.log(reason);
            })
        })
})();
