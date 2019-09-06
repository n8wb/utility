package utils

import(
	log "github.com/sirupsen/logrus"
	"github.com/whiteblock/go.uuid"
	"runtime"
)

//GetUUIDString generates a new UUID
func GetUUIDString() string {
	uid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return uid.String()
}

func LogErrorN(err error,n int) error {
	if err == nil {
		return err // do nothing if the given err is nil
	}
	_, file, line, ok := runtime.Caller(n)
	if !ok {
		log.Error(err.Error())
	} else {
		log.WithFields(log.Fields{"file": file, "line": line}).Error(err.Error())
	}

	return err
}


// LogError takes in an error, logs that error and returns that error.
// Used to help reduce code clutter and unify the error handling in the code.
// Has no effect if err == nil
func LogError(err error) error {
	return LogErrorN(err,2)
}
