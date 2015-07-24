<div class="row">

<div class="col-lg-offset-1 col-lg-4"><form class="form-horizontal">
    <div class="form-group ">
        <label for="question" >Question</label>
        <textarea ng-model="add.question.questionText" id="question" class="form-control" rows="3" placeholder="Type your question"></textarea>
    </div>

    <div class="form-group">
        <label for="option1" >Option1</label>
        <input ng-model="add.question.option1" id="option1" class="form-control" placeholder="Option1" />
    </div>

    <div class="form-group">
        <label for="option2" >Option2</label>
        <input ng-model="add.question.option2" id="option2" class="form-control" placeholder="Option2" />
    </div>

    <div class="form-group">
        <label for="option3" >Option3</label>
        <input ng-model="add.question.option3" id="option3" class="form-control" placeholder="Option3" />
    </div>

    <div class="form-group">
        <label for="option4" >Option4</label>
        <input ng-model="add.question.option4" id="option4" class="form-control" placeholder="Option4" />
    </div>

    <div class="form-group">
        <button ng-click="add.submit()" class="btn btn-primary btn-block" type="submit">{{add.addText}}</button>
    </div>
</form></div>

    <div class="col-lg-offset-1 col-lg-4">
        <toaster-container></toaster-container>
        <div class="row">Preview</div>
        <div class="row">
            <div class="panel">
                <div mathjax-bind="add.question.questionText"></div>
            </div>
            <div class="row"></div>
            <ul class="list-group">
                <li class="list-group-item">1.{{add.question.option1}}</li>
                <li class="list-group-item">2.{{add.question.option2}}</li>
                <li class="list-group-item">3.{{add.question.option3}}</li>
                <li class="list-group-item">4.{{add.question.option4}}</li>
            </ul>
        </div>
        <div class="row">
            <div class="form-group">

                <h3><label for="tags">Tags:</label> <div class="label label-primary">New</div></h3>
                <input id="label" type="text" class="form-control">
            </div>
        </div>
    </div>
</div>