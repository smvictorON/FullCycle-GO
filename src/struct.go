// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type Cliente struct {
// 	Nome  string
// 	Email string
// 	CPF   int
// }

// func (c Cliente) Exibe() {
// 	fmt.Println("EXIBINDO", c.Nome)
// }

// type ClienteInternacional struct {
// 	Cliente
// 	Pais string `json:"pais"`
// }

// func main() {
// 	cliente := Cliente{
// 		Nome:  "Victor",
// 		Email: "victorsm9@hotmail.com",
// 		CPF:   39093749803,
// 	}

// 	fmt.Println(cliente)

// 	cliente2 := Cliente{"teste", "teste@gmail.com", 111111111111}

// 	fmt.Printf("Nome: %s\nEmail: %s\nCPF: %d\n\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

// 	cliente3 := ClienteInternacional{
// 		Cliente: Cliente{
// 			Nome:  "Inter",
// 			Email: "inter@gmail.com",
// 			CPF:   111111,
// 		},
// 		Pais: "EUA",
// 	}

// 	fmt.Printf("Nome: %s\nEmail: %s\nCPF: %d\nPaís %s\n\n", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)

// 	cliente.Exibe()
// 	cliente2.Exibe()
// 	cliente3.Exibe()

// 	clienteJson, err := json.Marshal(cliente3)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	fmt.Println(string(clienteJson))

// 	jsoncliente4 := `{"Nome":"Inter","Email":"inter@gmail.com","CPF":111111,"pais":"EUA"}`
// 	cliente4 := ClienteInternacional{}

// 	json.Unmarshal([]byte(jsoncliente4), &cliente4)

// 	fmt.Println(cliente4.Nome)

// }
