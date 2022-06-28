package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

/////////////////////////////////////////////////////////////STRUCTS
type mbr = struct {
	Mbr_tamano         [100]byte
	Mbr_fecha_creacion [100]byte
	Mbr_dsk_signature  [100]byte
	Dsk_fit            [100]byte
	Mbr_partition      [4]partition
}

type partition struct {
	Part_status [100]byte
	Part_type   [100]byte
	Part_fit    [100]byte
	Part_start  [100]byte
	Part_size   [100]byte
	Part_name   [100]byte
}

type particion struct {
	Name string
	Path string
	Id   string
}

var particonesM []particion
var numeroP = 1
var letra = "a"

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
			var commandArray []string
			commandArray = append(commandArray, comando)
			commandArray = strings.Split(comando, " ")
			data := strings.ToLower(commandArray[0])
			if strings.Contains(data, "#") {
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
		fmt.Println("")
		fmt.Println("/////////////////////Creacion de Disco///////////////////////")
		crear_disco(commandArray)
	} else if data == "rmdisk" {
		fmt.Println("")
		fmt.Println("/////////////////////Eliminacion de Disco///////////////////////")
		borrar_disco(commandArray)
	} else if data == "fdisk" {
		//mostrar()
	} else if data == "rep" {
		fmt.Println("")
		fmt.Println("/////////////////////Reporte de Disco///////////////////////")
		crear_reporte(commandArray)
	} else if data == "mount" {
		fmt.Println("")
		fmt.Println("/////////////////////Montaje de Particiones///////////////////////")
		montar_particion(commandArray)
	} else if data == "exec" {
		fmt.Println("")
		fmt.Println("/////////////////////Ejecucion de Archivo///////////////////////")
		ejecutar_archivo(commandArray)
		fmt.Println("")
	} else if data == "pause" {
		fmt.Println("Presione Enter para continuar")
		reader := bufio.NewReader(os.Stdin)
		comando, _ := reader.ReadString('\n')
		if strings.Contains(comando, "\n") {
			fmt.Println("Continuando")
		}
	} else {
		fmt.Println("Comando ingresado no es valido")
	}
}

////////////////////////////////////////////////////////////////COMANDOS
func borrar_disco(commandArray []string) {
	path := ""
	// Lectura de parametros del comando
	for i := 0; i < len(commandArray); i++ {
		data := strings.ToLower(commandArray[i])
		if strings.Contains(data, "-path=") {
			path = strings.Replace(data, "-path=", "", 1)
			path = strings.Replace(path, "\"", "", 2)

		} else {
			fmt.Println("Parametro Incorrecto")
		}
	}

	//BUSQUEDA DE DISCO

	//LECTURA DE DISCO
	// Apertura de archivo
	// fmt.Println("PATHELIM " + path)
	err := os.Remove(path)
	if err != nil {
		fmt.Println("No existe el Archivo que desea Borrar")
	} else {
		fmt.Println("Archivo Borrado")
	}

}

func ejecutar_archivo(commandArray []string) {
	path := ""
	// Lectura de parametros del comando
	for i := 0; i < len(commandArray); i++ {
		data := commandArray[i]
		if strings.Contains(data, "-path=") {
			path = strings.Replace(data, "-path=", "", 1)
			path = strings.Replace(path, "\"", "", 2)

		} else {
			fmt.Println("Parametro Incorrecto")
		}
	}

	//RECORRER LINEAS DE ARCHIVO
	fmt.Println(path)
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println("No existe el Script")
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	for _, line := range lines {
		if len(line) > 1 {
			// fmt.Println("ComandoEx " + line)
			var commandArray []string
			commandArray = append(commandArray, line)
			commandArray = strings.Split(line, " ")
			data := strings.ToLower(commandArray[0])
			if strings.Contains(data, "#") {
				fmt.Println(line)
			} else if line != "" && line != "exit\n" {
				split_comando(line)
			}
		}

	}
}

func montar_particion(commandArray []string) {
	name := ""
	path := ""
	id := ""
	// Lectura de parametros del comando
	for i := 0; i < len(commandArray); i++ {
		data := strings.ToLower(commandArray[i])
		if strings.Contains(data, "-name=") {
			name = strings.Replace(data, "-name=", "", 1)
			name = strings.Replace(name, "\"", "", 2)

		} else if strings.Contains(data, "-path=") {
			path = strings.Replace(data, "-path=", "", 1)
			path = strings.Replace(path, "\"", "", 2)

		} else if strings.Contains(data, "#") {

		} else {
			fmt.Println("Parametro Incorrecto")
		}
	}

	id = generar_Id()
	particonesM = append(particonesM, particion{Name: name, Path: path, Id: id})
	fmt.Println("Particiones Montadas:")
	for i := 0; i < len(particonesM); i++ {
		fmt.Println("DISCO=" + particonesM[i].Path + " NOMBRE=" + particonesM[i].Name + " #ID=" + particonesM[i].Id)
	}

}
func crear_reporte(commandArray []string) {

	name := ""
	path := ""
	id := ""
	ruta := ""

	// Lectura de parametros del comando
	for i := 0; i < len(commandArray); i++ {
		data := strings.ToLower(commandArray[i])
		fmt.Println("DATA: " + data)
		if strings.Contains(data, "-name=") {
			name = strings.Replace(data, "-name=", "", 1)
			name = strings.Replace(name, "\"", "", 2)

		} else if strings.Contains(data, "-path=") {
			path = strings.Replace(data, "-path=", "", 1)
			path = strings.Replace(path, "\"", "", 2)

		} else if strings.Contains(data, "-id=") {
			id = strings.Replace(data, "-id=", "", 1)
			id = strings.Replace(id, "\"", "", 2)

		} else if strings.Contains(data, "-ruta=") {
			ruta = strings.Replace(data, "-ruta=", "", 1)
			ruta = strings.Replace(ruta, "\"", "", 2)
		} else if strings.Contains(data, "#") {

		} else {
			fmt.Println("Parametro Incorrecto")
		}
	}

	//BUSCAR DISCO
	pathDisco := "null"

	for i := 0; i < len(particonesM); i++ {
		if strings.Contains(particonesM[i].Id, id) {
			pathDisco = particonesM[i].Path
		}
	}
	fmt.Println("PATHD " + pathDisco)

	if strings.Contains(pathDisco, "null") {
		fmt.Println("No existe el Disco o la Particion")
	} else {

		//LECTURA DE DISCO
		// Apertura de archivo
		disco, err := os.OpenFile(pathDisco, os.O_RDWR, 0660)
		if err != nil {
			fmt.Println("No se pudo abrir el Archivo")
		}
		mbr := mbr{}
		sstruct := size_struct(pathDisco)
		fmt.Println(sstruct)
		// Lectrura de conjunto de bytes en archivo binario
		lectura := make([]byte, sstruct)
		_, err = disco.ReadAt(lectura, 0)
		if err != nil && err != io.EOF {
			fmt.Println("Error al Leer el Archivo")
		}

		// Conversion de bytes a struct
		mbr = bytes_to_struct(lectura)
		if err != nil {
			msg_error(err)
		}
		//fmt.Println(mbr)
		fmt.Println("Lectura de Disco:")
		fmt.Print(" Tamano: ")
		fmt.Println(string(mbr.Mbr_tamano[:]))
		fmt.Print(" Fecha: ")
		fmt.Println(string(mbr.Mbr_fecha_creacion[:]))
		fmt.Print(" DiskSIgnature: ")
		fmt.Println(string(mbr.Mbr_dsk_signature[:]))
		fmt.Print(" Fit: ")
		fmt.Println(string(mbr.Dsk_fit[:]))

		disco.Close()

		//LIMPIAR DATOS DE DISCO
		Tamano := string(bytes.Trim(mbr.Mbr_tamano[:], "\x00"))
		Fecha := string(bytes.Trim(mbr.Mbr_fecha_creacion[:], "\x00"))
		Signature := string(bytes.Trim(mbr.Mbr_dsk_signature[:], "\x00"))
		Fit := string(bytes.Trim(mbr.Dsk_fit[:], "\x00"))

		//

		if strings.Contains(name, "disk") {

			// Estructura inicial de Dotfile
			dotfile := ""
			////////////C
			dotfile += " digraph G{ \n\n"

			dotfile += "graph[\n"
			dotfile += "bgcolor = \"GREEN\" \n"
			dotfile += "]\n"

			dotfile += "subgraph cluster{\n label=\"MBR " + Fecha + "\" "
			dotfile += "nodo[shape=plaintext, label=<\n"
			dotfile += "<table bgcolor=\"white:lightblue\">"

			dotfile += "<tr>\n"
			dotfile += "<td>Nombre</td>\n"
			dotfile += "<td>Valor</td>\n"
			dotfile += "</tr>\n"
			dotfile += "<tr>\n"
			dotfile += "<td>mbr_tamano</td>\n"
			dotfile += "<td>"
			dotfile += Tamano
			dotfile += "</td>\n"
			dotfile += "</tr>\n"
			dotfile += "<tr>\n"
			dotfile += "<td>mbr_fecha_creacion</td>\n"
			dotfile += "<td>"
			dotfile += Fecha
			dotfile += "</td>\n"
			dotfile += "</tr>\n"
			dotfile += "<tr>\n"
			dotfile += "<td>mbr_disk_signature</td>\n"
			dotfile += "<td>"
			dotfile += Signature
			dotfile += "</td>\n"
			dotfile += "</tr>\n"
			dotfile += "<tr>\n"
			dotfile += "<td>disk_fit</td>\n"
			dotfile += "<td>"
			dotfile += Fit
			dotfile += "</td>\n"
			dotfile += "</tr>\n"
			dotfile += "<tr><td height='25' colspan='2'></td></tr>\n"

			contador := 1

			for i := 0; i < 4; i++ {

				dotfile += " <tr>\n"
				dotfile += " <td  bgcolor=\"green\"> part_status_" + strconv.Itoa(i+1) + "</td>"
				dotfile += " <td  bgcolor=\"green\">" + string(bytes.Trim(mbr.Mbr_partition[i].Part_status[:], "\x00")) + "</td>"
				dotfile += " </tr>\n"

				if string(mbr.Mbr_partition[i].Part_status[:]) != "-1" && string(mbr.Mbr_partition[i].Part_status[:]) != "1" {

					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_type_" + strconv.Itoa(contador) + "</td> <td>" + string(bytes.Trim(mbr.Mbr_partition[i].Part_type[:], "\x00")) + "</td></tr>\n"
					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_fi_" + strconv.Itoa(contador) + "</td> <td>" + string(bytes.Trim(mbr.Mbr_partition[i].Part_fit[:], "\x00")) + "</td></tr>\n"
					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_start_" + strconv.Itoa(contador) + "</td> <td>" + string(bytes.Trim(mbr.Mbr_partition[i].Part_start[:], "\x00")) + "</td></tr>\n"
					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_size_" + strconv.Itoa(contador) + "</td> <td>" + string(bytes.Trim(mbr.Mbr_partition[i].Part_size[:], "\x00")) + "</td></tr>\n"
					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_name_" + strconv.Itoa(contador) + "</td> <td>" + string(bytes.Trim(mbr.Mbr_partition[i].Part_name[:], "\x00")) + "</td></tr>\n"

				} else {

					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_type_" + strconv.Itoa(contador) + "</td> <td> -- </td></tr>\n"
					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_fi_" + strconv.Itoa(contador) + "</td> <td> -- </td></tr>\n"
					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_start_" + strconv.Itoa(contador) + "</td> <td> -- </td></tr>\n"
					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_size_" + strconv.Itoa(contador) + "</td> <td> -- </td></tr>\n"
					dotfile += "<tr> <td bgcolor=\"" + getColor_reporte(contador) + "\"> part_name_" + strconv.Itoa(contador) + "</td> <td> -- </td></tr>\n"
				}

				dotfile += "<tr><td height='25' colspan='2'></td></tr>\n"
				contador++

			}

			dotfile += "</table>\n"
			dotfile += ">]}\n"

			///////////
			// Cierre de Dotfile
			dotfile += "}\n"
			// Generacion de Dotfile y Imagen
			//Creacion de Directorio
			crearDirectorio(getDirectorio(path))
			generar_IMG(dotfile, path)
			disco.Close()
		} else {
			fmt.Println("No existe ese Reporte")
		}

	}

}

func getColor_reporte(numero int) string {

	switch numero {

	case 1: // TONO VERDE
		return "#ADFF2F:#FFFAFA"

	case 2: // TONO NARANJA
		return "#FFA500:#FFFAFA"

	case 3: // TONO GRIS
		return "#008B8B:#FFFAFA"

	case 4: // TONO ROJO
		return "#FF4500:#FFFAFA"

	case 5: // TONO DE verde obscuro
		return "#008000:#FFFAFA"

	case 6: // TONO DE ANARANJADO
		return "#DAA520:#FFFAFA"

	case 7: // TONO DE CELESTE
		return "white:lightblue"

	case 8: //VERDE SIN DEGRADAR
		return "GREEN"

	default:
		return "#ADFF2F:#FFFAFA"
	}
}
func crear_disco(commandArray []string) {
	tamano := 0
	dimensional := ""
	fit := "FF"
	path := ""
	fit2 := "F"
	tamano_archivo := 0
	limite := 0
	bloque := make([]byte, 1024)
	//MBR
	mbr := mbr{}

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
				copy(mbr.Mbr_tamano[:], "0")
				msg_error(err)
			} else {
				copy(mbr.Mbr_tamano[:], strtam)
			}
		} else if strings.Contains(data, "-unit=") {
			dimensional = strings.Replace(data, "-unit=", "", 1)
			dimensional = strings.Replace(dimensional, "\"", "", 2)
		} else if strings.Contains(data, "-fit=") {
			fit = strings.Replace(data, "-fit=", "", 1)
			fit = strings.Replace(fit, "\"", "", 2)
			fmt.Println("FIT: " + fit)
			if strings.Contains(fit, "BF") {
				fit2 = "B"
				copy(mbr.Dsk_fit[:], fit2)
			} else if strings.Contains(fit, "FF") {
				fit2 = "F"
				copy(mbr.Dsk_fit[:], fit2)
			} else if strings.Contains(fit, "WF") {
				fit2 = "W"
				copy(mbr.Dsk_fit[:], fit2)
			} else {
				copy(mbr.Dsk_fit[:], fit2)
			}

		} else if strings.Contains(data, "-path=") {
			path = strings.Replace(data, "-path=", "", 1)
			path = strings.Replace(path, "\"", "", 2)
		} else if strings.Contains(data, "#") {

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
			// msg_error(err)
		}
		limite++
	}
	disco.Close()

	//Escritura MBR
	copy(mbr.Mbr_fecha_creacion[:], obtenerFecha())
	copy(mbr.Mbr_dsk_signature[:], strconv.Itoa(obtenerNumeroRandom()))

	for i := 0; i < 4; i++ {
		copy(mbr.Mbr_partition[i].Part_status[:], "0")
		copy(mbr.Mbr_partition[i].Part_type[:], "0")
		copy(mbr.Mbr_partition[i].Part_size[:], "0")
		copy(mbr.Mbr_partition[i].Part_start[:], "-1")
		copy(mbr.Mbr_partition[i].Part_fit[:], "0")
		copy(mbr.Mbr_partition[i].Part_name[:], "")
	}

	disco2, err := os.OpenFile(path, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("No existe el Disco")
	}

	mbrbyte := struct_to_bytes(mbr)
	// newpos, err := disco2.Seek(0, os.SEEK_SET)
	// if err != nil {
	// 	fmt.Println("Error al Posicionarse")
	// }
	_, err = disco2.WriteAt(mbrbyte, 0)
	if err != nil {
		fmt.Println("Error al Escribir en el disco")
	}

	disco2.Close()
	leerdisco(path)

}

func leerdisco(path string) {

	// Apertura de archivo
	disco, err := os.OpenFile(path, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("No se pudo abrir el Archivo")
	}
	mbr := mbr{}
	sstruct := size_struct(path)
	fmt.Println(sstruct)
	// Lectrura de conjunto de bytes en archivo binario
	lectura := make([]byte, sstruct)
	_, err = disco.ReadAt(lectura, 0)
	if err != nil && err != io.EOF {
		fmt.Println("Error al Leer el Archivo")
	}

	// Conversion de bytes a struct
	mbr = bytes_to_struct(lectura)
	if err != nil {
		msg_error(err)
	}
	//fmt.Println(mbr)
	fmt.Println("Lectura de Disco:")
	fmt.Print(" Tamano: ")
	fmt.Println(string(mbr.Mbr_tamano[:]))
	fmt.Print(" Fecha: ")
	fmt.Println(string(mbr.Mbr_fecha_creacion[:]))
	fmt.Print(" DiskSIgnature: ")
	fmt.Println(string(mbr.Mbr_dsk_signature[:]))
	fmt.Print(" FIt: ")
	fmt.Println(string(mbr.Dsk_fit[:]))

	disco.Close()
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
	if err := os.MkdirAll(direccion, os.ModePerm); err != nil {
		log.Fatal(err)
	}

}

func obtenerFecha() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	// fmt.Println("Fecha: " + fecha)
	return fecha
}

func obtenerNumeroRandom() int {
	numero := rand.Intn(200)
	// fmt.Println("NumeroR: ", numero)
	return numero
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

func bytes_to_struct(s []byte) mbr {
	p := mbr{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		// msg_error(err)
		fmt.Println("Error al pasar Bytes a Struct")
	}
	return p
}

// Generacion de Dotfile y Imagen
func generar_IMG(dotinfo string, nombre string) {
	// Creacion de Archivo Dot
	nombreAR := strings.Split(nombre, ".")
	fmt.Println("nombrePath: " + nombreAR[0])
	file, err := os.Create(nombreAR[0] + ".dot")
	if err != nil {
		fmt.Println("Erro al crear el .dot")
		file.Close()
	}
	file.WriteString(dotinfo)
	file.Close()
	// Ejecucion de Comando para creacion de Imagen
	cmd := exec.Command("dot", "-Tjpg", nombreAR[0]+".dot", "-o", nombreAR[0]+".jpg")
	err = cmd.Run()
	if err != nil {
		fmt.Println("No se pudo crear el Reporte")

	}
}

func generar_Id() string {
	id := "13" + strconv.Itoa(numeroP) + letra
	numeroP++
	letraux := []rune(letra)
	char := string(letraux[0] + 1)
	letra = string(char)
	fmt.Println("idSALE: " + id)
	return id
}

func size_struct(path string) int {
	disco, err := ioutil.ReadFile(path)
	if err != nil {
		msg_error(err)
	}
	ejm := bytes_to_struct(disco)
	ejm2 := struct_to_bytes(ejm)
	return len(ejm2)
}
