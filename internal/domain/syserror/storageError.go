package syserror

import "fmt"

type StorageError struct {
	message string
	file    string
	path    string
}

func (sr StorageError) Error() string {
	return fmt.Sprintf("[Error Alert] \n message: %s \n path: %s \n file: %s \n", sr.message, sr.path, sr.file)
}

func StorageDirNotExist(file, path, message string) StorageError {
	return StorageError{
		message: fmt.Sprintf("error while read file, directory dont exist: %v", message),
		file:    file,
		path:    path,
	}
}

func StorageErrorReadFile(message, file, path string) StorageError {
	return StorageError{
		message: fmt.Sprintf("Error while open file: %s", message),
		file:    file,
		path:    path,
	}
}

func StorageErrorHandle(message, file, path string) StorageError {
	return StorageError{
		message: message,
		file:    file,
		path:    path,
	}
}
