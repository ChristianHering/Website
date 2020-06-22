package utils

func runConfigSetup() error {
	err := setConfig()
	if err != nil {
		return err
	}
}
