<div class="container">
  <ul class="collection with-header">
    <li class="collection-header"> LeaderBoards </li>
    <li class="collection-item" ng-repeat="result in leader.results">
      {{result.Username}} <div class="secondary-content"> {{result.Score}} </div>
    </li>
  </ul>
</div>
