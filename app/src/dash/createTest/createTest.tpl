<!--
 -- TOdO the entire panel body should be clickable
-->
<div class="container">
    <div class="col-lg-10">
        <div data-ng-repeat="question in test.questions" class="panel panel-default">
            <div class="form-group">
                <div class="panel-body">
                    <input data-ng-model="test.checked[$index]" type="checkbox"> <span
                        mathjax-bind="question.QuestionText"></span>
                </div>
            </div>
        </div>
    </div>
    <div class="col-lg-2">
        <toaster-container></toaster-container>
    </div>
</div>
<nav class="navbar navbar-default navbar-fixed-bottom bottom-bar-container ">
    <div class="container bottom-bar">
        <div class="form-group ">
            <div class="col-lg-6"><input type="text" class="form-control" data-ng-model="test.name" required
                                         placeholder="Enter the name of test"></div>
            <div class="col-lg-4">
                <select class="form-control" data-ng-model="test.group">
                    <option value="Engineering">Engineering</option>
                    <option value="Medical">Medical</option>
                    <option value="11th">11th</option>
                    <option value="12th">12th</option>
                </select>
            </div>
            <div class="col-lg-2">
                <button type="submit" ng-click="test.preview()" class="btn btn-default">Preview</button>
            </div>
        </div>
    </div>
</nav>