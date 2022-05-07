package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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
			if comando != "" && comando != "exit\n" {
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
		commandArray = append(commandArray, comando)
	} else {
		commandArray = strings.Split(comando, " ")
	}
	ejecucion_comando(commandArray)
}

func ejecucion_comando(commandArray []string) {
	data := strings.ToLower(commandArray[0])
	if data == "mkdisk" {
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
		} else if strings.Contains(data, "-dimensional=") {
			dimensional = strings.Replace(data, "-dimensional=", "", 1)
			dimensional = strings.Replace(dimensional, "\"", "", 2)
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

	// Creacion, escritura y cierre de archivo
	disco, err := os.Create("Disco_Ejemplo.dk")
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
}

//////////////////////////////////////////////////////////////////////

func struct_to_bytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil && err != io.EOF {
		msg_error(err)
	}
	return buf.Bytes()
}

func bytes_to_struct(s []byte) ejemplo {
	p := ejemplo{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		msg_error(err)
	}
	return p
}

func size_struct() int {
	disco, err := ioutil.ReadFile("Disco_Ejemplo.dk")
	if err != nil {
		msg_error(err)
	}
	ejm := bytes_to_struct(disco)
	ejm2 := struct_to_bytes(ejm)
	return len(ejm2)
}
