package errores

import (
	"errors"
	"testing"
)

type testNuevo struct {
	parámetros testNuevoParámetros
}

type testNuevoParámetros struct {
	err      error
	código   int
	mensaje  string
	errTexto string
}

func casoNuevo(err error, cdg int, msj string, errTexto string) testNuevo {
	return testNuevo{
		parámetros: testNuevoParámetros{
			err:      err,
			código:   cdg,
			mensaje:  msj,
			errTexto: errTexto,
		},
	}
}

func TestNuevo(t *testing.T) {
	pruebas := []testNuevo{
		casoNuevo(errors.New("abc"), 400, "pruebas 1", "pruebas 1: abc"),
		casoNuevo(errors.New("xyz"), 404, "pruebas 2", "pruebas 2: xyz"),
		casoNuevo(errors.New("qwe"), 500, "pruebas 3", "pruebas 3: qwe"),
		casoNuevo(errors.New("rty"), 0, "pruebas 4", "pruebas 4: rty"),
		casoNuevo(errors.New("uio"), 401, "", "uio"),
		casoNuevo(errors.New("asd"), 0, "", "asd"),
	}

	for _, prueba := range pruebas {
		Err := Nuevo(prueba.parámetros.err).ConCódigo(prueba.parámetros.código).ConMensaje(prueba.parámetros.mensaje)
		var Err2 = Error{}

		if Err.Código() != prueba.parámetros.código {
			t.Errorf("Err.Código() debe ser «%d»; se obtuvo «%d»",
				prueba.parámetros.código,
				Err.Código(),
			)
		} else if Err.Mensaje() != prueba.parámetros.mensaje {
			t.Errorf("Err.Mensaje() debe ser «%s»; se obtuvo «%s»",
				prueba.parámetros.mensaje,
				Err.Mensaje(),
			)
		} else if Err.Error() != prueba.parámetros.mensaje {
			t.Errorf("Err.Error() debe ser «%s»; se obtuvo «%s»",
				prueba.parámetros.mensaje,
				Err.Error(),
			)
		} else if Err.ErrorOriginal().Error() != prueba.parámetros.err.Error() {
			t.Errorf("Err.ErrorOriginal() debe ser «%s»; se obtuvo «%s»",
				prueba.parámetros.err.Error(),
				Err.ErrorOriginal().Error(),
			)
		} else if Err.String() != prueba.parámetros.errTexto {
			t.Errorf("Err.String() debe ser «%s»; se obtuvo «%s»",
				prueba.parámetros.errTexto,
				Err.String(),
			)
		} else if Err.Unwrap() != prueba.parámetros.err {
			t.Errorf("Err.Unwrap() debe ser «%s»; se obtuvo «%s»",
				prueba.parámetros.err,
				Err.Unwrap(),
			)
		} else if !errors.Is(Err, prueba.parámetros.err) {
			t.Errorf("errors.Is falló")
		} else if !errors.As(error(*Err), &Err2) {
			t.Errorf("errors.As falló")
		}

		var cdg, msj, errOriginal, válido = ExtraerCódigoMensajeErrorOriginal(Err)
		var cdgEsperado = prueba.parámetros.código
		if cdgEsperado < 200 || cdgEsperado >= 600 {
			cdgEsperado = 500
		}

		if !válido {
			t.Errorf("ExtraerCódigoMensajeErrorOriginal falló [1]: %d, %s, %s, %t", cdg, msj, errOriginal, válido)
		} else if cdg != cdgEsperado {
			t.Errorf("ExtraerCódigoMensajeErrorOriginal falló [2]: %d, %s, %s, %t", cdg, msj, errOriginal, válido)
		} else if msj != prueba.parámetros.mensaje {
			t.Errorf("ExtraerCódigoMensajeErrorOriginal falló [3]: %d, %s, %s, %t", cdg, msj, errOriginal, válido)
		} else if errOriginal != prueba.parámetros.err {
			t.Errorf("ExtraerCódigoMensajeErrorOriginal falló [4]: %d, %s, %s, %t", cdg, msj, errOriginal, válido)
		}

		cdg, msj, errOriginal, válido = ExtraerCódigoMensajeErrorOriginal(prueba.parámetros.err)
		if válido {
			t.Errorf("ExtraerCódigoMensajeErrorOriginal falló [5]: %d, %s, %s, %t", cdg, msj, errOriginal, válido)
		} else if cdg != 500 {
			t.Errorf("ExtraerCódigoMensajeErrorOriginal falló [6]: %d, %s, %s, %t", cdg, msj, errOriginal, válido)
		} else if msj != "" {
			t.Errorf("ExtraerCódigoMensajeErrorOriginal falló [7]: %d, %s, %s, %t", cdg, msj, errOriginal, válido)
		} else if errOriginal != prueba.parámetros.err {
			t.Errorf("ExtraerCódigoMensajeErrorOriginal falló [8]: %d, %s, %s, %t", cdg, msj, errOriginal, válido)
		}
	}
}
