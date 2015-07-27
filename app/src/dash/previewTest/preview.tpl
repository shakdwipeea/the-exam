<div class="row">
    <div class="test-name col-lg-8">{{preview.name}}</div>
    <div data-ng-click="preview.createTest()" class="btn btn-primary create-test col-lg-4">
        {{preview.submitButtonText}}
    </div>
</div>

<div class="row">
    <div class="container">
        <div class="col-lg-12">
            <div data-ng-repeat="question in preview.questions" class="panel panel-default">
                <div class="panel-body">
                    <div class="list-group">
                        <div class="list-group-item" mathjax-bind="question.QuestionText"></div>
                        <div class="list-group-item" mathjax-bind="question.Option1"></div>
                        <div class="list-group-item" mathjax-bind="question.Option2"></div>
                        <div class="list-group-item" mathjax-bind="question.Option3"></div>
                        <div class="list-group-item" mathjax-bind="question.Option4"></div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>