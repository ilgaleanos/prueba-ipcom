/*
 * Copyright 2022 Leito. All Rights Reserved.
 * <p>
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * <p>
 * http://www.apache.org/licenses/LICENSE-2.0
 * <p>
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package src

import (
	"errors"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

/**
*
* funcion de health check
*
 */

var alive = []byte(`API-IPCOM`)

func Alive(ctx *fasthttp.RequestCtx) {
	ctx.SetBody(alive)
}

/**
*
* parametros aceptados por el servicio
*
 */
type ResumenParams struct {
	FechaInicial time.Time `json:"fechaInicial"`
	Dias         int       `json:"dias"`
}

/**
*
* Validar las entradas al controlador
*
 */
func resumenValidar(fechaInicial string, dias string) (ResumenParams, error) {
	var err error
	var body ResumenParams

	// validar, se excluyen valor de dias inválidos y
	// mayores de 15 dias [ mas puede ser  y no hay indicaciones, deslimitarlo es muy fácil, así que se deja por defecto]
	if body.Dias, err = strconv.Atoi(dias); err != nil || body.Dias < 1 || body.Dias > 15 {
		return body, errors.New("días inválidos: `" + dias + "`")
	}

	if body.FechaInicial, err = time.Parse("2006-01-02", fechaInicial); err != nil {
		return body, errors.New("fecha inválida: `" + fechaInicial + "`")
	}

	return body, nil
}

/**
*
* Controlador de la peticion y respuesta del generador de estadisticas
*
 */
func ResumenControlador(ctx *fasthttp.RequestCtx) {
	// verificar entradas
	fechaInicialIn := ctx.UserValue("fechaInicial").(string)
	diasIn := string(ctx.QueryArgs().Peek("dias"))
	Logger.Info("ResumenControlador", zap.String("diasIn", diasIn), zap.String("fechaInicialIn", fechaInicialIn))

	// manejamos mala parametrizacion
	if body, err := resumenValidar(fechaInicialIn, diasIn); err != nil {
		ctx.SetStatusCode(400)
		ctx.Response.SetBody([]byte(err.Error()))
		return

	} else {
		// manejamos la logica del endpoint
		estadisticas, err := ResumenFlujo(&body)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetContentType("application/json")
		ctx.Response.SetBody(estadisticas)
	}
}
