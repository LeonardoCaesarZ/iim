package auth

func Init() error {
	if err := readPriKeyIntoMemroy(); err != nil {
		return err
	}

	return nil
}
