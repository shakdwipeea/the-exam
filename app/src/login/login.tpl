<div class="row">

<div class="container intro">
    <div class="col-lg-offset-4 col-lg-4">
        <h2 class="h2">Login</h2>
    <form>
            <div class="form-group">
                <input placeholder="Username" class="form-control" id="username" required name="username" ng-model="login.user.username">
            </div>
            <div class="form-group">
                <input placeholder="Password" id="password" class="form-control" required type="password" name="password" ng-model="login.user.password">
            </div>
            <div class="form-group">
                <button  type="submit" class="btn btn-default" ng-click="login.doLogin()">
                   {{login.loginText}}
                </button>
                </div>
    </form>
        </div>
    <div class="col-lg-offset-2 col-lg-2">
        <toaster-container></toaster-container>
    </div>
</div></div>