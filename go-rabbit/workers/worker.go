package workers

// Worker defines the worker interface
type Worker interface {
	// Use allows a middleware like function to be passed to the worker
	Use(...WorkerMiddlewareFunc)
	// Health returns an error if connection to the queue is not healthy
	Health() error
	// Config changes the worker address
	Config(address string) error
	// Listen starts to listen with a given quantity of jobs, in a given queue (if queue does not exists it is created), and processes the body with the given function
	Listen(quantity int, queueName string, processBody func([]byte) error) (Channel, error)
	// Closes the worker and all of its connections
	Close() error
	// Send sends messages to a queue
	Send(queue string, messages ...[]byte) error
}

type WorkerMiddlewareFunc func(name string) WorkerMiddleware

type WorkerMiddleware interface {
	Start()
	Stop(error)
}

// Channel that can be closed in order to stop listening without having to close the whole worker
type Channel interface {
	Close() error
}

func BuildRabbitWorker(address string) (worker Worker, err error) {
	worker = new(rabbitWorker)
	err = worker.Config(address)
	return
}
