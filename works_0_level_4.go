package main

import "fmt"

func main() {

	var users = []string{"Matthew", "Sarah", "Augustus", "Heidi", "test", "Emilie", "Peter", "Giana", "Adriano", "Elizabeth"}

	/* e,E 给一个
	   i,I 给两个
	   o,O   三个
	   u,U   四个
	*/

	var sum = 0

	for i := len(users) - 1; i >= 0; i-- {

		var num int = 0

		for k := len(users[i]) - 1; k >= 0; k-- {

			switch users[i][k] {

			case 'e':
				num += 1
			case 'E':
				num += 1
			case 'i':
				num += 2
			case 'I':
				num += 2
			case 'o':
				num += 3
			case 'O':
				num += 3
			case 'U':
				num += 4
			case 'u':
				num += 4

			}

		}

		sum += num

		fmt.Println(users[i], " 获得了金币 ", num)

	}

	fmt.Println("共发出金币数 ", sum)

}
