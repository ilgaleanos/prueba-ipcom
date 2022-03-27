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
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

type Usuario struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

type Organizacion struct {
	Nombre   string    `json:"organization"`
	Usuarios []Usuario `json:"users"`
}

/**
*
* localizar o agregar una organizacion
*
 */
func existeOrganizacion(organizaciones *[]Organizacion, nombre string) int {
	lenOrganizaciones := len(*organizaciones)
	for indice := 0; indice < lenOrganizaciones; indice++ {
		if (*organizaciones)[indice].Nombre == nombre {
			return indice
		}
	}
	return -1
}

/**
*
* localizar o agregar un usuario
*
 */
func existeUsuario(usuarios *[]Usuario, username string) int {
	lenUsuarios := len(*usuarios)
	for indice := 0; indice < lenUsuarios; indice++ {
		if (*usuarios)[indice].Username == username {
			return indice
		}
	}
	return -1
}

/**
*
* agregar un rol, sin repeticion
*
 */
func agregarRol(organizaciones *[]Organizacion, indiceOrganizacion int, indiceUsuarios int, rol string) {
	for _, rolUsuario := range (*organizaciones)[indiceOrganizacion].Usuarios[indiceUsuarios].Roles {
		if rolUsuario == rol {
			return
		}
	}

	(*organizaciones)[indiceOrganizacion].Usuarios[indiceUsuarios].Roles = append(
		(*organizaciones)[indiceOrganizacion].Usuarios[indiceUsuarios].Roles,
		rol,
	)
}

/**
*
* Organizar dentro de la estructura una linea del archivo
*
 */
func agrupadorLinea2Jerarquia(organizaciones *[]Organizacion, linea string) {
	partes := strings.Split(linea, ",")
	// realizar trim de cada celda
	for index := 0; index < len(partes); index++ {
		partes[index] = strings.Trim(partes[index], " ")
	}

	// omitir encabezado
	if partes[2] == "rol" {
		return
	}

	// localizar o agregar una organizacion
	var indiceOrganizacion int
	if indiceOrganizacion = existeOrganizacion(organizaciones, partes[0]); indiceOrganizacion == -1 {
		(*organizaciones) = append((*organizaciones), Organizacion{Nombre: partes[0]})
		indiceOrganizacion = len(*organizaciones) - 1
	}

	// localizar o agregar un usuario
	var indiceUsuario int
	if indiceUsuario = existeUsuario(&(*organizaciones)[indiceOrganizacion].Usuarios, partes[1]); indiceUsuario == -1 {
		(*organizaciones)[indiceOrganizacion].Usuarios = append(
			(*organizaciones)[indiceOrganizacion].Usuarios,
			Usuario{Username: partes[1], Roles: []string{partes[2]}},
		)
		indiceUsuario = len((*organizaciones)[indiceOrganizacion].Usuarios) - 1
	}

	// agregar un rol, sin repeticion
	agregarRol(organizaciones, indiceOrganizacion, indiceUsuario, partes[2])
}

/**
*
* Leer el archivo y generar la jerarquia
*
 */
func AgrupadorFlujo(ruta string) ([]byte, error) {
	// leer fichero
	file, err := os.Open(ruta)
	if err != nil {
		Logger.Error("no existe: `" + ruta + "`")
	}
	defer file.Close()

	// agrupar roles por usuarios y usuarios por organizacion
	var organizaciones []Organizacion

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// ejecutar la jerarquización por cada línea
		agrupadorLinea2Jerarquia(&organizaciones, scanner.Text())
	}

	// error al leer la línea
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// generar respuesta
	return json.Marshal(organizaciones)
}
