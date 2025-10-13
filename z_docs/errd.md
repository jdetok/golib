# GETERR GO PACKAGE
- standard error handling pkg derived from applog in go-api-jdeko.me
# Err struct & BuildErr func
```go
type Err struct {
	Func string
	Msg  string
}
func (e *Err) BuildErr(err error) error {
	startEnd := "************"
	return fmt.Errorf("%s\n** ERROR OCCURED IN %s\n** MSG: %s\n** ERR MSG FROM FUNC: %w\n%s",
		startEnd, e.Func, e.Msg, err, startEnd)
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
# Example Console Output
```
************
** ERROR OCCURED IN github.com/jdetok/golib/getenv.GetEnvStr
** MSG: key 'TEST' not found in .env
** ERR MSG FROM FUNC: &{%!e(string=key 'TEST' not found in .env)}
************
```
# Example Implementation
```go
func GetEnvStr(key string) (string, error) {
	e := geterr.InitErr()
	val, ok := os.LookupEnv(key)
	if !ok {
		e.Msg = fmt.Sprintf("key '%s' not found in .env", key)
		return "", e.BuildErr(errors.New(e.Msg))
	}
	return val, nil
}
```