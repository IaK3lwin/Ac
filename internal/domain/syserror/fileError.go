package syserror

import "fmt"

type FileError struct {
	message string
	path    string
}

func (fr FileError) Error() string {
	return fmt.Sprintf("Erro message: %v\n path: %v\n", fr.message, fr.path)
}

func FactoryFileError(message, path string) FileError {
	messageFormat := fmt.Sprintf("Message Error: %v", message)

	return FileError{
		message: messageFormat,
		path:    path,
	}
}
