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
	"io/ioutil"
	"net/http"
	"strings"
	sync "sync"

	"go.uber.org/zap"
)

type ResumenCompras struct {
	Total         float64            `json:"total"`
	ComprasPorTDC map[string]float64 `json:"comprasPorTDC"`
	NoCompraron   int64              `json:"nocompraron"`
	CompraMasAlta float64            `json:"compraMasAlta"`
}

/**
*
* peticion unica al api, se retorna por los canales de acumulacion y errores
*
 */
func peticionApiCompras(idx int, fecha string, ch chan<- []Compra, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	// contruir la url con los parametros
	mUrl := []string{API_URL, "/compras/", fecha}
	url := strings.Join(mUrl, "")
	Logger.Info("API", zap.Int("idx", idx), zap.String("url", url))

	// realizar la peticion
	resp, err := http.Get(url)
	if err != nil {
		// Logger.Error("error de conexion", zap.Error(err))
		errChan <- err
		return
	}
	defer resp.Body.Close()

	// leer la respuesta
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errChan <- err
		return
	}

	// deserealizar
	var body []Compra
	if err := json.Unmarshal(msg, &body); err != nil {
		errChan <- err
		return
	}

	// reportar a los canales
	ch <- body
	errChan <- nil
}

/**
*
* Acumulador de la informacion de las compras
*
 */
func acumularCompras(idx int, resumenCompras *ResumenCompras, comprasChan <-chan []Compra, wg *sync.WaitGroup) {
	for compras := range comprasChan {
		for _, compra := range compras {
			// noCompraron: numero de registros donde no hubo compra
			if !compra.GetCompro() {
				resumenCompras.NoCompraron++
				continue
			}

			// Total: la suma del monto de todas las ventas de todos los días indicados
			resumenCompras.Total += compra.GetMonto()

			// comprasPorTDC: un map que tenga como llave el tipo de tdc y como valor la suma del monto correspondiente a esa tipo de tarjeta de esos días.
			if _, existe := resumenCompras.ComprasPorTDC[compra.Tdc]; !existe {
				resumenCompras.ComprasPorTDC[compra.Tdc] = 0.0
			}
			resumenCompras.ComprasPorTDC[compra.Tdc] += compra.GetMonto()

			// compraMasAlta: la compra mas alta en el periodo.
			if resumenCompras.CompraMasAlta < compra.GetMonto() {
				resumenCompras.CompraMasAlta = compra.GetMonto()
			}
		}

		wg.Done()
	}
}

/**
*
* Acumulador de errores
*
 */
func erroresCompras(idx int, erroresAcumulados *[]error, errChan <-chan error, wg *sync.WaitGroup) {
	for err := range errChan {
		(*erroresAcumulados) = append((*erroresAcumulados), err)
		wg.Done()
	}
}
