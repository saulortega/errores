package errores

import (
	"errors"
	"fmt"
)

type Error struct {
	código  int
	mensaje string
	err     error
}

func Nuevo(err error) *Error {
	return &Error{
		err: err,
	}
}

func (O *Error) ConCódigo(código int) *Error {
	O.código = código
	return O
}

func (O *Error) ConMensaje(mensaje string) *Error {
	O.mensaje = mensaje
	return O
}

func (O Error) String() string {
	if len(O.mensaje) > 0 {
		return fmt.Sprintf("%s: %s", O.mensaje, O.err.Error())
	}

	return O.err.Error()
}

func (O Error) Error() string {
	return O.mensaje
}

func (O Error) Código() int {
	return O.código
}

func (O Error) Mensaje() string {
	return O.mensaje
}

func (O Error) ErrorOriginal() error {
	return O.err
}

func (O Error) Unwrap() error {
	return O.err
}

func ExtraerCódigoMensajeErrorOriginal(err error) (int, string, error, bool) {
	var errOriginal = err
	var Err = &Error{}
	var vld = false
	var msj string
	var cdg = 500

	var transformado = errors.As(err, &Err)
	if !transformado {
		transformado = errors.As(err, Err)
	}

	if transformado {
		errOriginal = Err.ErrorOriginal()
		msj = Err.Mensaje()
		cdg = Err.Código()

		if cdg < 100 || cdg >= 600 {
			cdg = 500
		}

		vld = true
	}

	return cdg, msj, errOriginal, vld
}
