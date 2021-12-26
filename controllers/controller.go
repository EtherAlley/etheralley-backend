package controllers

type IController interface {
	Start() error
	Exit()
}

func StartController(controller IController) {
	// start is intended to be a blocking call
	// if Exit() is called, we have caught a panic
	// if start returns, one of our controllers is no longer active and thus we should force a panic
	defer controller.Exit()
	err := controller.Start()
	panic(err)
}
