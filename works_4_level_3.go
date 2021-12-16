package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func WriteIn(m map[string]string) {

	var path = "user.data.json"

	newData, err := json.Marshal(&m)

	if err != nil {
		fmt.Println("写入错误")
		return
	}

	ioutil.WriteFile(path, newData, 0777)

}

func getData(m *map[string]string) {

	data, err := ioutil.ReadFile("user.data.json")

	if err != nil {
		fmt.Println("读取失败,可能是找不到文件", err)
		return
	} else {
		fmt.Println("读取成功,正在解析")
	}

	json.Unmarshal([]byte(string(data)), &m)

}

func main() {

	fmt.Println("正在读入文件")

	var users = make(map[string]string)

	getData(&users)

	var n int
	var user_Name string
	var password string

	for {

		fmt.Println("==========================\n登陆系统\n")
		fmt.Println("1.登录    2.注册    3.修改密码    4.退出\n==========================")

		fmt.Scanln(&n)

		if n == 4 {
			return
		} //exit

		if n == 3 {

			fmt.Println("请输入账号")

			fmt.Scanln(&user_Name)

			PW, exist := users[user_Name]

			if !exist {
				fmt.Println("账号不存在")
				continue
			} else {

				fmt.Println("请输入密码(区别大小写)")
				fmt.Scanln(&password)

				if PW == password {
					fmt.Println("登陆成功\n请输入新的密码")

					fmt.Scanln(&password)

					users[user_Name] = password

					WriteIn(users)

					fmt.Println("修改成功")

				}

			}

		}

		if n == 1 {

			fmt.Println("请输入账号")

			fmt.Scanln(&user_Name)

			PW, exist := users[user_Name]

			if !exist {
				fmt.Println("账号不存在")
				continue
			} else {

				fmt.Println("请输入密码(区别大小写)")
				fmt.Scanln(&password)

				if PW == password {
					fmt.Println("登陆成功")
					return
				}

			}

		}

		if n == 2 {

			fmt.Println("请输入账号")

			fmt.Scanln(&user_Name)

			_, exist := users[user_Name]

			if exist {

				fmt.Println("账号已经存在")
				continue

			} else {

				fmt.Println("请输入密码(区别大小写)")
				fmt.Scanln(&password)

				users[user_Name] = password

				WriteIn(users)

				fmt.Println("注册成功,并存入user.data.json中")

				return

			}

		}

	}
}
