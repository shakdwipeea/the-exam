<div class="row">
    <form ng-submit="signup.submit()" class="col s12">
        <div class="row">
            <div class="input-field col s12">
                <span class="ak-error-notify">{{signup.message}}</span>
                <input ng-model="signup.username" ng-change="signup.checkUserName()" id="email" type="email"
                       class="validate">
                <label for="email" data-error="wrong" data-success="right">Email</label>

            </div>
            <div class="input-field col s12">
                <input ng-model="signup.password" id="password" type="password" class="validate">
                <label for="password">Password</label>
            </div>
            <div class="row">
                <div class="input-field col s6">
                    <p>
                        <input ng-model="signup.group.Engineering" type="checkbox" id="test5"/>
                        <label for="test5">Engineering</label>
                    </p>
                </div>
                <div class="input-field col s6"><p>
                    <input ng-model="signup.group.Medical" type="checkbox" id="test6"/>
                    <label for="test6">Medical</label>
                </p></div>
            </div>
            <div class="row">
                <div class="input-field col s6">
                    <p>
                        <input ng-model="signup.group.ele" type="checkbox" id="11th"/>
                        <label for="11th">11th</label>
                    </p>
                </div>
                <div class="input-field col s6"><p>
                    <input ng-model="signup.group.twe" type="checkbox" id="12th"/>
                    <label for="12th">12th</label>
                </p></div>
            </div>
            <div class=" col s12">
                <button ng-disabled="!signup.enable" class="waves-effect waves-light btn" type="submit" name="action">
                    Sign Up
                    <i class="material-icons">send</i>
                </button>
            </div>
        </div>
    </form>
</div>