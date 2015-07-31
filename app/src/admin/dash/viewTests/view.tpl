<style>
    .enable-btn {
        margin-top: -1%;
    }
</style>
<div class="container tests">
    <div data-ng-repeat="test in view.tests" class="panel panel-default">
        <div class="panel-body" ui-sref="dash.preview({testId: test.Id})">
            Name: {{test.Name}}
            <span class="label label-default pull-right"><h4>{{test.Enable === true && "Enabled" ||
                "Disabled"}}</h4></span>
        </div>
        <div class="panel-footer">
            Group: {{test.Group}}
            <span class="btn btn-default pull-right enable-btn"
                  data-ng-click="view.EnableTest($index, test.Id, test.Enable)">{{test.Enable !== true && "Enable" || "Disable"}}</span>
        </div>
    </div>
    <toaster-container></toaster-container>
</div>