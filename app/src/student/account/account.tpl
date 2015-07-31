<div class="mdl-tabs mdl-js-tabs mdl-js-ripple-effect">
    <div class="mdl-tabs__tab-bar">
        <a href="#starks-panel" class="mdl-tabs__tab is-active">Log In</a>
        <a href="#lannisters-panel" class="mdl-tabs__tab">Sign Up</a>
    </div>

    <div class="mdl-tabs__panel is-active" id="starks-panel">
        <div class="mdl-grid">
            <div class="mdl-cell mdl-cell-4-col"></div>
            <div class="mdl-cell mdl-cell--4-col">
                <form>
                    <div class="mdl-textfield mdl-js-textfield textfield-demo">
                        <input class="mdl-textfield__input" type="text" id="username"/>
                        <label class="mdl-textfield__label" for="username">Username</label>
                    </div>
                    <div class="mdl-textfield mdl-js-textfield textfield-demo">
                        <input class="mdl-textfield__input" type="text" id="password"/>
                        <label class="mdl-textfield__label" for="password">password</label>
                    </div>

                    <div>  <!-- Accent-colored raised button with ripple -->
                        <button type="submit"
                                class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent">
                            Log In
                        </button>
                    </div>
                </form>
            </div>
        </div>
    </div>
    <div class="mdl-tabs__panel" id="lannisters-panel">
        <div class="mdl-grid">
            <div class="mdl-cell mdl-cell-4-col"></div>
            <div class="mdl-cell mdl-cell--4-col">
                <form>
                    <div class="mdl-textfield mdl-js-textfield textfield-demo">
                        <input class="mdl-textfield__input" type="text" id="username"/>
                        <label class="mdl-textfield__label" for="username">Username</label>
                    </div>
                    <div class="mdl-textfield mdl-js-textfield textfield-demo">
                        <input class="mdl-textfield__input" type="text" id="password"/>
                        <label class="mdl-textfield__label" for="password">password</label>
                    </div>

                    <div>
                        <label class="mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect" for="checkbox-1">
                            <input type="checkbox" id="checkbox-1" class="mdl-checkbox__input"/>
                            <span class="mdl-checkbox__label">Engineering</span>
                        </label>
                        <label class="mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect" for="checkbox-2">
                            <input type="checkbox" id="checkbox-2" class="mdl-checkbox__input"/>
                            <span class="mdl-checkbox__label">Medical</span>
                        </label>
                        <label class="mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect" for="checkbox-3">
                            <input type="checkbox" id="checkbox-3" class="mdl-checkbox__input"/>
                            <span class="mdl-checkbox__label">11th</span>
                        </label>
                        <label class="mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect" for="checkbox-4">
                            <input type="checkbox" id="checkbox-4" class="mdl-checkbox__input"/>
                            <span class="mdl-checkbox__label">12th</span>
                        </label>
                    </div>

                    <div>  <!-- Accent-colored raised button with ripple -->
                        <button type="submit"
                                class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent">
                            Log In
                        </button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>