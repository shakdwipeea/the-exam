<div class="container">
    <div data-ng-repeat="test in exam.tests" class="card">
        <div class="card-content">
            <span class="card-title left-align" style="color: red">Group : {{test.Group}}</span>
            <span class="card-title right" style="color: blue">Subject : {{test.Subject}}</span>

            <div class="center-align">Test Name: {{test.Name}}</div>
        </div>
        <div class="card-action">
            <button data-ng-click="exam.getTestDetail(test.Id)" class="waves-effect waves-light btn-flat">Take Test
            </button>
        </div>
    </div>
</div>