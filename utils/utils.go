package utils

func RunUtilSetup() error {
	err := setupConfig()
	if err != nil {
		return err
	}

	err := setupSQL()
	if err != nil {
		return err
	}

	return nil
}
