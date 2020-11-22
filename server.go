package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Server struct {
	subjects map[string]map[string]float64
	students map[string]map[string]float64
}

type Data struct {
	Student string
	Subject string
	Grade   float64
}

func (this *Server) AddGrade(data Data, reply *string) error {

	if _, ok := this.students[data.Student]; !ok {
		grade := make(map[string]float64)
		grade[data.Subject] = data.Grade
		this.students[data.Student] = grade
	} else {
		if _, ok := this.students[data.Student][data.Subject]; ok {
			return errors.New("Ya existe una calificacion para este alumno")
		}

		this.students[data.Student][data.Subject] = data.Grade
	}

	if _, ok := this.subjects[data.Subject]; !ok {
		grade := make(map[string]float64)
		grade[data.Student] = data.Grade
		this.subjects[data.Subject] = grade
	} else {
		this.subjects[data.Subject][data.Student] = data.Grade
	}

	*reply = "Done"
	return nil
}

func (this *Server) StudentAVG(data Data, reply *float64) error {
	avg := GetStudentAVG(data.Student, this)
	*reply = avg
	return nil
}

func (this *Server) GeneralAVG(data Data, reply *float64) error {

	avg := float64(0)
	for student := range this.students {
		avg += GetStudentAVG(student, this)
	}
	avg = avg / float64(len(this.students))
	*reply = avg
	return nil
}

func (this *Server) SubjectAVG(data Data, reply *float64) error {
	avg := float64(0)
	for student := range this.subjects[data.Subject] {
		avg += this.subjects[data.Subject][student]
	}
	avg = avg / float64(len(this.subjects[data.Subject]))
	*reply = avg
	return nil
}

func GetStudentAVG(student string, this *Server) float64 {
	avg := float64(0)
	for subject := range this.students[student] {
		avg += this.students[student][subject]
	}
	avg = avg / float64(len(this.students[student]))
	return avg
}

func server() {
	ser := new(Server)
	ser.subjects = make(map[string]map[string]float64)
	ser.students = make(map[string]map[string]float64)
	rpc.Register(ser)
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
