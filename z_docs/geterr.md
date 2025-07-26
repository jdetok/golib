# GETERR GO PACKAGE
- standard error handling pkg derived from applog in go-api-jdeko.me
# Err struct & BuildErr func
```go
type Err struct {
	Func string
	Msg  string
}
func (e *Err) BuildEr(err error) error {
	return fmt.Errorf("** ERROR IN %s\n-- ***MSG: %s\n ****SOURCE FUNC ERR: %e",
		e.Func, e.Msg)
}
```
# InitErr func
the InitErr func returns an Err type with the Func value as the caller function's name
### IMPORTANT: must be Caller(1) - Caller(0) will display the InitErr func name itself 
```go
func InitErr() Err {
	var e Err
	pc, _, _, _ := runtime.Caller(1)
	e.Func = runtime.FuncForPC(pc).Name()
	return e
}
```