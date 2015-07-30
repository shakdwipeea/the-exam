<div class="container tests">
    <div data-ng-repeat="test in view.tests" ui-sref="dash.preview({testId: test.Id})" class="panel panel-default">
        <div class="panel-body">
            Name: {{test.Name}}
            <span class="label label-default pull-right"><h4>{{test.Enable === "true" && "Enabled" || "Disabled"}}</h4></span>
        </div>
        <div class="panel-footer">
            Group: {{test.Group}}
        </div>
    </div>
</div>