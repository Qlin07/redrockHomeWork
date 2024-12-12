package function

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dsn = "root:@tcp(127.0.0.1:3306)/messageboard?charset=utf8mb4&parseTime=True&loc=Local"

func SignUp() {
	db, err := sql.Open("mysql", dsn)
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for {
		fmt.Printf("please enter your username:\n")
		var name string
		fmt.Scanln(&name)
		// 准备查询语句
		query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)`
		// 执行查询
		var exists bool
		err = db.QueryRow(query, name).Scan(&exists)
		if err != nil {
			log.Fatal(err)
		}
		// 处理结果
		if exists {
			fmt.Printf("The name has contained,please try other names.\n")
		} else {
			var pwd, ispwd string
			for {
				fmt.Println("please enter your password:")
				fmt.Scanln(&pwd)
				fmt.Println("please enter your password again:")
				fmt.Scanln(&ispwd)
				if ispwd == pwd {
					fmt.Println("account signup success")
					break
				} else {
					fmt.Println("your password is different,please try again")
				}
			}
			result, err := db.Exec("insert into users(username,password) values(?,?)", name, pwd)
			if err != nil {
				log.Println(err)
				return
			}
			result.LastInsertId()
			result.RowsAffected()
			break
		}
	}

}
func SignIn() {
	db, err := sql.Open("mysql", dsn)
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	fmt.Println("please enter your username:")
	var username string
	fmt.Scanln(&username)
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)`
	var exists bool
	err = db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		//从数据库查询用户密码
		rows, err := db.Query("SELECT password from users where username = ?", username)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var password string
		for rows.Next() {
			err = rows.Scan(&password)
			if err != nil {
				log.Fatal(err)
			}
		}
		//用户输入密码
		fmt.Println("please enter your password:")
		for {
			var pwd string
			fmt.Scanln(&pwd)
			//对比密码
			if pwd == password {
				fmt.Println("account sign in success")
				break
			} else {
				fmt.Println("incorrect password,please try again")
			}
		}
	} else {
		fmt.Printf("not found such account,do you want to signUp?(y/n)\n")
		var answer string
		fmt.Scanln(&answer)
		if answer == "y" {
			SignUp()
		} else {
			return
		}
	}
}
func UpLoad() {
	
}
func GetMessage() {

}
func del() {

}
