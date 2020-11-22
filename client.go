package main

import (
	"fmt"
	"net/rpc"
)

type Data struct {
	Student string
	Subject string
	Grade   float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1) Agregar calificacion de alumno")
		fmt.Println("2) Obtener promedio de alumno")
		fmt.Println("3) Obtener promedio de todos los alumnos")
		fmt.Println("4) Obtener promedio por materia")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var data Data
			fmt.Print("Name: ")
			fmt.Scanln(&data.Student)
			fmt.Print("Materia: ")
			fmt.Scanln(&data.Subject)
			fmt.Print("Calificacion: ")
			fmt.Scanln(&data.Grade)


			var result string
			err = c.Call("Server.AddGrade", data, &result)
			if err != nil {
				fmt.Println(err)
			} 
		case 2: 
			var data Data
			fmt.Print("Name: ")
			fmt.Scanln(&data.Student)
			var result float64
			err = c.Call("Server.StudentAVG", data, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("el promedio de ", data.Student, " es: ", result)
			}
		case 3:
			var data Data
			var result float64
			err = c.Call("Server.GeneralAVG", data, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("el promedio de es: ", result)
			}
		case 4: 
			var data Data
			fmt.Print("Materia: ")
			fmt.Scanln(&data.Subject)
			var result float64
			err = c.Call("Server.SubjectAVG", data, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("el promedio de ", data.Subject, " es: ", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}