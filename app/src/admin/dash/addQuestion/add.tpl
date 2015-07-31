<div class="row">

    <div class="col-lg-offset-1 col-lg-4">
        <form class="form-horizontal">
            <div class="form-group ">
                <label for="question">
                    Question <a target="_blank" href="https://en.wikipedia.org/wiki/Help:Displaying_a_formula#Functions.2C_symbols.2C_special_characters">Help for typing Math</a>
                </label>
                <textarea ng-model="add.question.questionText" id="question" class="form-control" rows="3"
                          placeholder="Type your question"></textarea>
            </div>

            <div class="form-group">
                <label for="option1">Option1</label>
                <input ng-model="add.question.option1" id="option1" class="form-control" placeholder="Option1"/>
            </div>

            <div class="form-group">
                <label for="option2">Option2</label>
                <input ng-model="add.question.option2" id="option2" class="form-control" placeholder="Option2"/>
            </div>

            <div class="form-group">
                <label for="option3">Option3</label>
                <input ng-model="add.question.option3" id="option3" class="form-control" placeholder="Option3"/>
            </div>

            <div class="form-group">
                <label for="option4">Option4</label>
                <input ng-model="add.question.option4" id="option4" class="form-control" placeholder="Option4"/>
            </div>

            <div class="form-group">
                <label for="correct">Correct</label>
                <select id="correct" data-ng-model="add.question.correct" class="form-control">
                    <option value="1">1</option>
                    <option value="2">2</option>
                    <option value="3">3</option>
                    <option value="4">4</option>
                </select>
            </div>

            <div class="form-group">
                <button ng-click="add.submit()" class="btn btn-primary btn-block" type="submit">{{add.addText}}</button>
            </div>
        </form>
    </div>

    <div class="col-lg-offset-1 col-lg-4">
        <toaster-container></toaster-container>
        <div class="row">Preview</div>
        <div class="row">
            <div class="panel">
                <div mathjax-bind="add.question.questionText"></div>
            </div>
            <div class="row"></div>
            <ul class="list-group">
                <li class="list-group-item" mathjax-bind="add.question.option1">Option1</li>
                <li class="list-group-item" mathjax-bind="add.question.option2">Option2</li>
                <li class="list-group-item" mathjax-bind="add.question.option3">Option3</li>
                <li class="list-group-item" mathjax-bind="add.question.option4">Option3</li>
            </ul>
        </div>
        <div class="row">
            <div class="form-group">

                <h4>
                    <label for="tags">Tags:</label>
                    <span data-ng-repeat="tag in add.selectedTags" class="label label-default tag">{{tag}}</span>
                </h4>
                <input ng-keydown="add.addTag($event)" id="label" type="text" data-ng-model="add.newTagText" placeholder="Press enter to add a tag"
                       class="form-control">
            </div>
            <div class="row">
                <div class="form-group col-lg-6">
                    <div class="checkbox"   data-ng-repeat="tag in add.tags1">
                        <label>
                            <input data-ng-model="add.tagSelect[tag.Name]" ng-change="add.Select()" type="checkbox" >{{tag.Name}}
                        </label>
                    </div><!-- /input-group -->
                </div>
                <div class="form-group col-lg-6">
                    <div class="checkbox" data-ng-repeat="tag in add.tags2">
                        <label>
                            <input data-ng-model="add.tagSelect[tag.Name]" ng-change="add.Select()" type="checkbox" >{{tag.Name}}
                        </label>
                    </div><!-- /input-group -->
                </div>
            </div>
        </div>
    </div>
</div>