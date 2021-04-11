package services

import (
	"github.com.br/JoaoVictorPereiraTeixeira/domain"
	"github.com.br/JoaoVictorPereiraTeixeira/framework/utils"
	"github.com/streadway/amqp"
)

type JobWorkerResult struct {
	Job     domain.Job
	Message *amqp.Delivery
	Error   error
}

func JobWorker(messageChannel chan amqp.Delivery, returnChan chan JobWorkerResult, JobService JobService, workerID int) {

	for message := range messageChannel {
		err := utils.IsJson(string(message.Body))

		if err != nil {
			returnChan <- returnJobResult(domain.Job{}, message, err)
			continue
		}
	}
}

func returnJobResult(job domain.Job, message amqp.Delivery, err error) JobWorkerResult {
	result := JobWorkerResult{
		Job:     job,
		Message: &message,
		Error:   err,
	}

	return result
}