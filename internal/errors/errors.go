package errors

type Error struct {
  Message string
}

func New(msg string) Error {
  return Error{Message: msg}
}
