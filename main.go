package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/////////////////////////////////////////////////////////////STRUCTS

///////////////////////////////////////////////////////////////
func main() {
	analizar()
}

func msg_error(err error) {
	fmt.Println("Error: ", err)
}

func analizar() {
	finalizar := false
	fmt.Println("PROYECTO 2 ARCHIVOS EXT2 201602813(exit para salir...)")
	reader := bufio.NewReader(os.Stdin)
	//  Ciclo para lectura de multiples comandos
	for !finalizar {
		fmt.Print("> ")
		comando, _ := reader.ReadString('\n')
		if strings.Contains(comando, "exit") {
			finalizar = true
		} else {
			if strings.Contains(comando, "#") {
				fmt.Println(comando)
			} else if comando != "" && comando != "exit\n" {
				split_comando(comando)
			}
		}
	}
}

func split_comando(comando string) {
	var commandArray []string
	comando = strings.Replace(comando, "\n", "", 1)
	comando = strings.Replace(comando, "\r", "", 1)
	if strings.Contains(comando, "mostrar") {

	} else {
		commandArray = strings.Split(comando, " ")
	}
	ejecucion_comando(commandArray)
}

func ejecucion_comando(commandArray []string) {
	data := strings.ToLower(commandArray[0])
	if data == "mkdisk" {
		fmt.Println("Creacion de Disco")
		crear_disco(commandArray)
	} else if data == "rmdisk" {
		//escribir(commandArray)
	} else if data == "fdisk" {
		//mostrar()
	} else {
		fmt.Println("Comando ingresado no es valido")
	}
}

////////////////////////////////////////////////////////////////COMANDOS
func crear_disco(commandArray []string) {
	tamano := 0
	dimensional := ""
	fit := "FF"
	path := ""
	tamano_archivo := 0
	limite := 0
	bloque := make([]byte, 1024)

	// Lectura de parametros del comando
	for i := 0; i < len(commandArray); i++ {
		data := strings.ToLower(commandArray[i])
		if strings.Contains(data, "-size=") {
			strtam := strings.Replace(data, "-size=", "", 1)
			strtam = strings.Replace(strtam, "\"", "", 2)
			strtam = strings.Replace(strtam, "\r", "", 1)
			tamano2, err := strconv.Atoi(strtam)
			tamano = tamano2
			if err != nil {
				msg_error(err)
			}
		} else if strings.Contains(data, "-unit=") {
			dimensional = strings.Replace(data, "-unit=", "", 1)
			dimensional = strings.Replace(dimensional, "\"", "", 2)
		} else if strings.Contains(data, "-fit=") {
			fit = strings.Replace(data, "-fit=", "", 1)
			fit = strings.Replace(fit, "\"", "", 2)
		} else if strings.Contains(data, "-path=") {
			path = strings.Replace(data, "-path=", "", 1)
			path = strings.Replace(path, "\"", "", 2)
		} else {
			fmt.Println("Parametro Incorrecto")
		}
	}

	// Calculo de tamaño del archivo
	if strings.Contains(dimensional, "k") {
		tamano_archivo = tamano
	} else if strings.Contains(dimensional, "m") {
		tamano_archivo = tamano * 1024
	} else if strings.Contains(dimensional, "g") {
		tamano_archivo = tamano * 1024 * 1024
	}

	// Preparacion del bloque a escribir en archivo
	for j := 0; j < 1024; j++ {
		bloque[j] = 0
	}

	//Creacion de Directorio
	crearDirectorio(getDirectorio(path))
	// Creacion, escritura y cierre de archivo
	disco, err := os.Create(path)
	if err != nil {
		msg_error(err)
	}
	for limite < tamano_archivo {
		_, err := disco.Write(bloque)
		if err != nil {
			msg_error(err)
		}
		limite++
	}
	disco.Close()

	fmt.Print("Creacion de Disco:")
	fmt.Print(" Tamaño: ")
	fmt.Print(tamano)
	fmt.Print(" Dimensional: ")
	fmt.Println(dimensional)
	fmt.Print(" Fit: ")
	fmt.Println(fit)
	fmt.Print(" Path: ")
	fmt.Println(path)
	fmt.Print(" PathDir: ")
	fmt.Println(getDirectorio(path))
	fmt.Print(" PathArch: ")
	fmt.Println(getArchivo(path))
}

//////////////////////////////////////////////////////////////////////
func getDirectorio(direccion string) string {
	var aux string = filepath.Dir(direccion)
	return aux
}

func getArchivo(direccion string) string {
	var aux string = filepath.Base(direccion)
	return aux
}

func crearDirectorio(direccion string) {
	if _, err := os.Stat(direccion); os.IsNotExist(err) {
		err = os.MkdirAll(direccion, os.ModePerm)
		if err != nil {
			fmt.Println("No se pudo crear el Directorio")
		}
	}
}

func struct_to_bytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil && err != io.EOF {
		msg_error(err)
	}
	return buf.Bytes()
}

// func bytes_to_struct(s []byte) ejemplo {
// 	p := ejemplo{}
// 	dec := gob.NewDecoder(bytes.NewReader(s))
// 	err := dec.Decode(&p)
// 	if err != nil && err != io.EOF {
// 		msg_error(err)
// 	}
// 	return p
// }

// func size_struct() int {
// 	disco, err := ioutil.ReadFile("Disco_Ejemplo.dk")
// 	if err != nil {
// 		msg_error(err)
// 	}
// 	ejm := bytes_to_struct(disco)
// 	ejm2 := struct_to_bytes(ejm)
// 	return len(ejm2)
// }
