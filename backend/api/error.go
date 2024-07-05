package api

type Error struct {
	Code    int
	Func    string
	Err     error
	Message string
	Details []any
}

func (e *Error) Error() string {

	if e.Err == nil {
		return e.Message
	}

	return e.Err.Error()
}

func (e *Error) Fields() map[string]interface{} {

	fields := make(map[string]interface{})
	fields["err"] = e.Error()

	return fields
}
