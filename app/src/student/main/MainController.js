(function () {
    'use strict'

    angular.module('student')
        .controller('MainController', function ($rootScope) {
            this.progress = false;
            this.Message = $rootScope.Message;
        })
})();
