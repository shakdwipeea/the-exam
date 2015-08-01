 <div class="row">
    <form ng-submit="login.submit()" class="col s12">
      <div class="row">
        <div class="input-field col s12">
          <input data-ng-model="login.username" id="email" type="email" class="validate">
          <label for="email" data-error="wrong" data-success="right">Email</label>
        </div>
        <div class="input-field col s12">
          <input data-ng-model="login.password" id="password" type="password" class="validate">
          <label for="password" >Password</label>
        </div>
        <div class=" col s12">
          <button class="waves-effect waves-light btn" type="submit" name="action">Log In
            <i class="material-icons">send</i>
         </button>
        </div>
      </div>
    </form>
  </div>
        