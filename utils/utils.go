package utils

//Initialization for our utility functions
func RunUtilSetup() error {
	err := setupConfig()
	if err != nil {
		return err
	}

	errChan := make(chan error)
	go setupSQL(errChan)
	err = <-errChan
	if err != nil {
		return err
	}

	setupAuth()

	return nil
}
