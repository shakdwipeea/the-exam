!function () {
    "use strict";
    angular.module("student", ["ui.router", "toaster", "ngAnimate"]).config(["$stateProvider", "$urlRouterProvider", function (t, e) {
        e.otherwise("/account");
        var n = "/public/src/student/";
        t.state("account", {
            url: "/account",
            templateUrl: n + "account/account.tpl",
            controller: "AccountController as account"
        }).state("account.login", {
            url: "/login",
            templateUrl: n + "account/login/login.tpl",
            controller: "LoginController as login"
        }).state("account.signup", {
            url: "/signup",
            templateUrl: n + "account/signup/signup.tpl",
            controller: "SignUpController as signup"
        }).state("exam", {
            url: "/exam",
            templateUrl: n + "exam/exam.tpl",
            controller: "ExamController as exam"
        }).state("test", {
            url: "/test/:id",
            templateUrl: n + "test/test.tpl",
            controller: "TestController as test"
        }).state("result", {
            url: "/result",
            templateUrl: n + "result/result.tpl",
            controller: "ResultController as result"
        })
    }])
}(), function () {
    "use strict";
    angular.module("student").controller("ExamController", ["Exam", "toaster", "Account", "$state", function (t, e, n, o) {
        var r = this;
        return n.getToken() ? (t.getExams().then(function (t) {
            console.log(t), r.tests = t.data.tests
        })["catch"](function (t) {
            console.log(t), e.pop("error", "Error Occured", t.data.err)
        }), void(r.getTestDetail = function (n) {
            t.getCompleteTest(n).then(function (t) {
                console.log(t), o.go("test", {id: t.data.questions[0].Id})
            })["catch"](function (t) {
                console.log(t), e.pop("error", "Error Occured", t.data.err)
            })
        })) : void o.go("account.login")
    }])
}(), function () {
    "use strict";
    angular.module("student").factory("Exam", ["$http", "Account", "Test", "Result", function (t, e, n, o) {
        var r = [];
        return {
            getExams: function () {
                return t({method: "GET", url: "/getTests", params: {token: e.getToken()}}).then(function (t) {
                    return r = t.data.tests, t
                })
            }, getCompleteTest: function (r) {
                return o.setId(r), t({
                    method: "GET",
                    url: "/secure/test/" + r,
                    params: {token: e.getToken()}
                }).then(function (t) {
                    return n.setQuestions(t.data.questions), t
                })
            }
        }
    }])
}(), function () {
    "use strict";
    angular.module("student").controller("MainController", function () {
        this.progress = !1
    })
}(), function () {
    "use strict";
    angular.module("student").controller("ResultController", ["$state", "Account", "Result", function (t, e, n) {
        if (!e.getToken())return void t.go("account.login");
        var o = n.evaluate(), r = this;
        r.total = o.total, r.scored = o.scored
    }])
}(), function () {
    "use strict";
    angular.module("student").factory("Result", ["$http", "Account", function (t, e) {
        function n(n) {
            var s = [];
            s = o.map(function (t) {
                return {id: t.Id, answer: t.answer}
            }), t.post("/saveResult", {token: e.getToken(), response: s, score: n + "", test_id: r}).then(function (t) {
                console.log(t)
            })["catch"](function (t) {
                throw console.log(t), new Error("Could not save Result")
            })
        }

        var o = [], r = "";
        return {
            setId: function (t) {
                r = t
            }, setQuestions: function (t) {
                o = t
            }, evaluate: function () {
                console.log(o);
                var t = 0, e = 0;
                return o.forEach(function (n) {
                    t += 4, n.answer == n.Correct && (e += 4)
                }), n(e), {total: t, scored: e}
            }
        }
    }])
}(), function () {
    angular.module("student").directive("mathjaxBind", function () {
        return {
            restrict: "A", controller: ["$scope", "$element", "$attrs", function (t, e, n) {
                t.$watch(n.mathjaxBind, function (t) {
                    e.html(t), MathJax.Hub.Queue(["Typeset", MathJax.Hub, e[0]]), t || e.html("")
                })
            }]
        }
    })
}(), function () {
    "use strict";
    angular.module("student").controller("TestController", ["$stateParams", "Test", "$state", "Account", "toaster", function (t, e, n, o, r) {
        var s = this, u = t.id;
        return o.getToken() ? (s.question = e.getQuestion(u), console.log("is", s.question, u), void(s.next = function () {
            console.log(s);
            var t = e.answer(s.question, s.myAnswer);
            if (t) {
                var o = e.next();
                o ? n.go("test", {id: o}) : (console.log("Complete"), r.pop("success", "Success", "Test Complete"), n.go("result"))
            } else r.pop("error", "Error", "Its not working")
        })) : void n.go("account.login")
    }])
}(), function () {
    "use strict";
    angular.module("student").factory("Test", ["Result", function (t) {
        var e = [], n = 0;
        return {
            setQuestions: function (t) {
                console.log("et que", t), e = t, n = 0
            }, getQuestion: function (t) {
                console.log("Getting quetion", e);
                for (var n = {}, o = 0; o < e.length; o++)if (e[o].Id == t) {
                    n = e[o];
                    break
                }
                return n
            }, answer: function (t, n) {
                var o = e.indexOf(t);
                return -1 === o ? !1 : (e[o].answer = n, !0)
            }, next: function () {
                return n == e.length - 1 ? (t.setQuestions(e), !1) : (n++, e[n].Id)
            }
        }
    }])
}(), function () {
    "use strict";
    angular.module("student").controller("AccountController", function () {
    })
}(), function () {
    "use strict";
    angular.module("student").factory("Account", ["$http", function (t) {
        function e(e) {
            return t.post("/studentSignup", e)
        }

        function n(e) {
            return t.post("/studentLogin", e).then(function (t) {
                return console.log(t), s = t.data.token, t
            })
        }

        function o() {
            return t.get("/usernames")
        }

        function r() {
            return s
        }

        var s = null;
        return {signUp: e, login: n, getUserNames: o, getToken: r}
    }])
}(), function () {
    "use strict";
    angular.module("student").controller("SignUpController", ["Account", "toaster", "$state", function (t, e, n) {
        var o = this;
        o.message = "", o.checkUserName = function () {
            console.log("Function called"), t.getUserNames().then(function (t) {
                if (t.data) {
                    var e = t.data.indexOf(o.username);
                    -1 !== e ? (o.message = "Username taken", o.enable = !1) : (o.message = "", o.enable = !0)
                } else o.enable = !0
            })
        }, o.submit = function () {
            console.log(o);
            var r = [], s = o.group;
            for (var u in s)if (s[u]) {
                switch (u) {
                    case"ele":
                        u = "11th";
                        break;
                    case"twe":
                        u = "12th"
                }
                r.push(u)
            }
            var a = {username: o.username, password: o.password, groups: r};
            t.signUp(a).then(function (o) {
                t.login({username: a.username, password: a.password}).then(function () {
                    e.pop("success", "Hurray", "Logged In"), n.go("exam")
                })["catch"](function (t) {
                    e.pop("error", "Error", t.data.err)
                })
            })["catch"](function (t) {
                console.log(t), e.pop("error", "Error", t.data.err)
            })
        }
    }])
}(), function () {
    "use strict";
    angular.module("student").controller("LoginController", ["Account", "toaster", "$state", function (t, e, n) {
        var o = this;
        o.submit = function () {
            t.login({username: o.username, password: o.password}).then(function () {
                e.pop("success", "Hurray", "Logged In"), n.go("exam")
            })["catch"](function (t) {
                e.pop("error", "Error", t.data.err)
            })
        }
    }])
}();