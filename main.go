package main

import "fmt"

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
)

func main() {
	var command string
	userList := []string{"user1_password1", "user1_password1"}
	productList := make([]string, 0, 10)

	_ = productList
	for command != exit {
		fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			message := ""
			if !string_in_list(command, userList) {
				userList = append(userList, command)
				message = fmt.Sprintf("Пользователь %s успешно добавлен", command)
			} else {
				message = fmt.Sprintf("Пользователь %s уже существует", command)
			}
			fmt.Println(message)

			fmt.Println(userList)
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command)

			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магази")

				} else {
					fmt.Println("Вы не зарегистрированны")
				}

			}
		}
	}
}

func string_in_list(get_str string, list []string) bool {
	for _, str := range list {
		if str == get_str {
			return true
		}
	}
	return false
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
// test comment
