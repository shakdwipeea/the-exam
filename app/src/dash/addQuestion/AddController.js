/**
 * Created by akash on 22/7/15.
 */

/**
 * Created by akash on 22/7/15.
 */

angular.module('question')
    .controller('AddController', function (User, toaster) {
        console.log("add Controller");
        var self = this;
		self.question = {};
		self.addText = "Add";
        self.submit = function  () {
			self.addText = "Adding.............";

        	var promise = User.add(self.question);
			console.log(promise instanceof Error);

			if (!(promise instanceof Error)) {
				promise.then(function (response) {
					console.log(response);
					if (response.data.err == false) {
						console.log("Successfully added");
						self.question.questionText = "";
						self.question = {};
						toaster.pop('success', 'Success', 'Question Added');
					} else {
						console.log("error");
						toaster.pop('error', 'Error ocurred', 'OOps Try Again');
					}
					self.addText = "Add";
				})
				.catch(function (error) {
					self.question.questionText = "";
					self.question = {};
					console.log("Errror", error);
					toaster.pop('error', 'Error ocurred', 'OOps Try Again');
							self.addText = "Add";
				})
			} else {
				self.addText = "Add";
				toaster.pop('error', 'Error ocurred', 'Type the question');
			}
        }
    });