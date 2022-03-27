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
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var Logger *zap.Logger
var API_URL string

// =====================================================================================================================
//
// 	Funciones de configuración
//
// =====================================================================================================================

// tareas iniciales del sistema
func Configure() {
	Logger, _ = zap.NewProduction()

	// se leen las variables de entorno del archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	API_URL = mustGetenv("API_URL")
}
