package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	id            int
	name          string
	class         string
	contactnumber string
}

func main() {

	fmt.Println(" \n*** Student Management Application ***\n ")
	var input string
	counter := 1
	ListOfStudent := []Student{}

Recovery:
	fmt.Println("\nEnter Your Choice : \n1: Add  \n2 : Edit \n3 : Delete \n4 : Print \n5 : Print All Records \nRest : Exit \n ")
	fmt.Scanln(&input)

	switch input {
	case "1":
		{
			var std Student
			std.id = counter

			fmt.Println("\nEnter Name : ")
			in := bufio.NewReader(os.Stdin)
			Spacewithname, _ := in.ReadString('\n')
			Spacewithname = strings.TrimSpace(Spacewithname)
			std.name = Spacewithname

			fmt.Println("\nEnter Class : ")
			fmt.Scanln(&std.class)

			fmt.Println("\nEnter Contact Number : ")
			fmt.Scanln(&std.contactnumber)

			ListOfStudent = append(ListOfStudent, std)
			counter++
			goto Recovery
		}
	case "2":
		{
			var Eid int
			var Newname string
			var Newclass string
			var Contactnumber string

			fmt.Println("\nEnter Your Id")
			fmt.Scanln(&Eid)
			if Eid <= len(ListOfStudent) {
				obj := ListOfStudent[Eid-1]
				var std = &Student{}

				(*std).id = obj.id

				fmt.Println("Previous Name : ", obj.name)
				fmt.Print("Enter New Name : ")
				in := bufio.NewReader(os.Stdin)
				Newname, _ = in.ReadString('\n')
				(*std).name = strings.TrimSpace(Newname)

				fmt.Println("Previous Class : ", obj.class)
				fmt.Print("Enter New Class : ")
				fmt.Scanf("%s\n", &Newclass)
				(*std).class = Newclass

				fmt.Println("Previous Contact Number : ", obj.contactnumber)
				fmt.Print("Enter New Contact number : ")
				fmt.Scanf("%s\n", &Contactnumber)
				(*std).contactnumber = Contactnumber

				ListOfStudent[Eid-1] = *std

			} else {
				fmt.Println("\nPlease Enter Valid Id")
			}
			goto Recovery
		}
	case "3":
		{
			var Eid int

			fmt.Println("\nEnter Your Id")
			fmt.Scanln(&Eid)

			if Eid < len(ListOfStudent) {
				ListOfStudent = append(ListOfStudent[:Eid], ListOfStudent[Eid+1:]...)
			} else if Eid == len(ListOfStudent) {
				ListOfStudent = append(ListOfStudent[:Eid-1])
			} else {
				fmt.Println("\nPlease Enter Valid Id")
			}

			goto Recovery
		}
	case "4":
		{
			var Eid int

			fmt.Println("\nEnter Your Id")
			fmt.Scanln(&Eid)
			fmt.Println("\nFormat : ID | Name | Class | Contact Number")
			if Eid <= len(ListOfStudent) {
				fmt.Println("\n", ListOfStudent[Eid-1])
			} else {
				fmt.Println("\nPlease Enter Valid Id")
			}

			goto Recovery
		}

	case "5":
		{
			fmt.Println("\nFormat : ID | Name | Class | Contact Number")
			fmt.Println("\n", ListOfStudent)
			goto Recovery
		}

	default:
		{
			fmt.Println("\n*** Happy GO Coding ***")
			break
		}
	}
}
