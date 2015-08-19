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
            console.log(self.question);
            if (!self.question.questionText || !self.question.option1 ||
                !self.question.option2 || !self.question.option3 ||
                !self.question.option4 || !self.question.correct) {
                toaster.pop('error', 'Error ocurred', 'Missing Values');
                return;
            } else if (!self.subjectSelect) {
                toaster.pop('error', 'Error ocurred', 'No Subjects');
                return;
            };

			self.addText = "Adding.............";

        	var promise = User.add(self.question, self.selectedTags, self.subjectSelect);
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
						toaster.pop('error', 'Error ocurred', 'Missing Values');
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


        self.addTag = function ($event) {
            if ($event.keyCode === 13) {

                var promise = User.newTag({
                    name: self.newTagText
                });

                if (promise instanceof Error) {
                    toaster.pop('error', 'Error', 'Could not add Tags');
                } else {
                    promise.then(function (response) {
                        getTags();
                        if (response.data.err === false) {
                            self.newTagText = "";
                            toaster.pop('success', 'Success', 'New tag added');
                        } else {
                            toaster.pop('error', 'Error', response.data.msg);
                        }
                    })
                        .catch(function (reason) {
                            console.log(reason);
                            toaster.pop('error', 'Error', response.data.msg);
                        })
                }
            }
        };

        getTags();

        function getTags () {
            User.getTags()
                .then(function (response) {
                    if (response.data.err === false) {
                        self.tags = response.data.tags;
                        var x = self.tags.length / 2;

                        self.tags1 = self.tags.slice(0 , x);
                        self.tags2 = self.tags.slice(x, x * 2);

                        console.log(response.data.tags);
                    } else {
                        toaster.pop('error', 'Colud not get Tags', response.data.msg);
                    }
                })
                .catch(function (reason) {
                    toaster.pop('error', 'Server No gud ', "Ouuta here");
                })
        }

        self.Select = function () {
            console.log(self.tagSelect);
            var k = [];
            for (var key in self.tagSelect) {
                if (self.tagSelect[key] === true) {
                    k.push(key);
                }
            }

            self.selectedTags = k;
        }


        //Subjects
        self.addSubject = function ($event) {
            if ($event.keyCode === 13) {

                var promise = User.newSubject({
                    name: self.newSubjectText
                });

                if (promise instanceof Error) {
                    toaster.pop('error', 'Error', 'Could not add Tags');
                } else {
                    promise.then(function (response) {
                        getSubject();
                        if (response.data.err === false) {
                            self.newTagText = "";
                            toaster.pop('success', 'Success', 'New tag added');
                        } else {
                            toaster.pop('error', 'Error', response.data.msg);
                        }
                    })
                        .catch(function (reason) {
                            console.log(reason);
                            toaster.pop('error', 'Error', response.data.msg);
                        })
                }
            }
        };

        getSubject();

        function getSubject () {
            User.getSubject()
                .then(function (response) {
                    if (response.data.err === false) {
                        self.subjects = response.data.subjects;

                        console.log(response.data.subjects);
                    } else {
                        toaster.pop('error', 'Colud not get Tags', response.data.msg);
                    }
                })
                .catch(function (reason) {
                    toaster.pop('error', 'Server No gud ', "Ouuta here");
                })
        }

        self.SelectSubject = function  () {
            console.log(self.subjectSelect);
        }


    });