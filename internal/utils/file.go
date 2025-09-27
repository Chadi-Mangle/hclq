package utils

import "os"

func saveBytesToFile(data []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(data)

	return nil
}
