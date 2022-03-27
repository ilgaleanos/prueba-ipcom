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
	"encoding/json"
	sync "sync"
	"time"

	"go.uber.org/zap"
)

/**
*
* Generar las n-1 fechas adicionales dada una fecha inicial y n
*
 */
func calcularFechas(parametros *ResumenParams) []string {
	var dias []string
	dias = append(dias, parametros.FechaInicial.Format("2006-01-02"))

	for diasAdicionales := 1; diasAdicionales < parametros.Dias; diasAdicionales++ {
		dias = append(dias, (parametros.FechaInicial.Add(
			time.Hour * 24 * time.Duration(diasAdicionales),
		)).Format("2006-01-02"))
	}

	return dias
}

/**
*
* Generador del resumen consumiendo el api concurrentemente
*
 */
func ResumenFlujo(parametros *ResumenParams) ([]byte, error) {
	// calcular las fechas de consultas
	fechas := calcularFechas(parametros)
	Logger.Info("fechas", zap.Reflect("slice", fechas))

	// crear el grupo
	var wg sync.WaitGroup
	wg.Add(3 * parametros.Dias)

	// crear los canalales
	var errores []error
	errChan := make(chan error, parametros.Dias)
	defer close(errChan)

	var resumenCompras ResumenCompras
	resumenCompras.ComprasPorTDC = make(map[string]float64)

	comprasChan := make(chan []Compra, parametros.Dias)
	defer close(comprasChan)

	// disparar las peticiones concurrentemente
	for idx, fecha := range fechas {
		go peticionApiCompras(idx, fecha, comprasChan, errChan, &wg)
		go acumularCompras(idx, &resumenCompras, comprasChan, &wg)
		go erroresCompras(idx, &errores, errChan, &wg)
	}
	wg.Wait()

	// verificar que no haya errores en la consulta
	for _, err := range errores {
		if err != nil {
			return nil, err
		}
	}

	// entregar las estadísticas
	return json.Marshal(resumenCompras)
}
