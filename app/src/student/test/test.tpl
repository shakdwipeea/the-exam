<div class="container">
    <div class="card">
        <form ng-submit="test.next()">
            <div class="card-content">
                <div class="card-title">Question</div>
                <p mathjax-bind="test.question.QuestionText"></p>

                <div class="input-field col s6"><p>
                    <input name="answer" ng-model="test.myAnswer" type="radio" id="option1" value="1"/>
                    <label for="option1" mathjax-bind="test.question.Option1"></label>
                </p></div>

                <div class="input-field col s6"><p>
                    <input name="answer" ng-model="test.myAnswer" type="radio" id="option2" value="2"/>
                    <label for="option2" mathjax-bind="test.question.Option2"></label>
                </p></div>

                <div class="input-field col s6"><p>
                    <input name="answer" ng-model="test.myAnswer" type="radio" id="option3" value="3"/>
                    <label for="option3" mathjax-bind="test.question.Option3"></label>
                </p></div>


                <div class="input-field col s6"><p>
                    <input name="answer" ng-model="test.myAnswer" type="radio" id="option4" value="4"/>
                    <label for="option4" mathjax-bind="test.question.Option4"></label>
                </p></div>
            </div>
            <div class="card-action">
                <button type="submit" class="waves-effect waves-light btn-flat">Next
                </button>
            </div>
        </form>
    </div>
</div>