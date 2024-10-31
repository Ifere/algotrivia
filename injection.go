package main

import (
	dbs "github.com/ifere/algotrivia/config/db"

	questiondelivery "github.com/ifere/algotrivia/src/questions/delivery/http"
	questionrepo "github.com/ifere/algotrivia/src/questions/repository"
	questionusecase "github.com/ifere/algotrivia/src/questions/usecase"
	
)


func InitializeQuestions(conn dbs.MongoDB) questiondelivery.QuestionDelivery {
	questionRepoInit := questionrepo.MongodbRepo{MongoDB: conn}


	questionUseCase := questionusecase.NewQuestionUseCase(questionRepoInit)
	return questiondelivery.NewQuestion(questionUseCase)
}

