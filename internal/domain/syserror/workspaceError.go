package syserror

import "fmt"

type WorkspaceError struct {
	message string
	path    string
}

func (wr WorkspaceError) Error() string {
	return fmt.Sprintf("[Error workspace found!] \n message: %s \n path: %s \n", wr.message, wr.path)
}

func WsActionsNotFound(actionName, message, path string) WorkspaceError {
	return WorkspaceError{
		message: message,
		path:    path,
	}
}

func FactoryWsError(message, path string) WorkspaceError {
	messageFormat := fmt.Sprintf("Error: %s", message)

	return WorkspaceError{
		message: messageFormat,
		path:    path,
	}
}
