<div class="intro">
<div class="row">
    <div class="col-lg-2">
        <nav class="nav" role="navigation">
            <ul class="nav nav-pills nav-stacked">
                <li role="presentation" ng-class="{ 'active': select == 'question' }" ng-click="select = 'question'"><a
                        ui-sref="dash.add"> Add A Question</a></li>
                <li role="presentation" ng-class="{ 'active': select == 'test' }" ng-click="select = 'test'"><a
                        ui-sref="dash.createTest">New Test</a> </a></li>
                <li role="presentation"><a href="#"> View Tests</a></li>
            </ul>
        </nav>
    </div>

<div class="col-lg-10">
    <div class="row">
        <div ui-view></div>
    </div>
</div>
</div></div>