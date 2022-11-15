package main

import "fmt"

type userData struct {
	userId   int
	userName string
}

func valueReceiver(usr userData) {
	usr.userId = 111
	fmt.Println("user detail in valueReceiver() : ", usr)
}

func pointerReceiver(usr *userData) {
	usr.userId = 222
	fmt.Println("user detail in valueReceiver() : ", usr)
}

func main() {
	user1 := userData{
		userId:   1001,
		userName: "test-user-1",
	}
	user2 := userData{
		userId:   1002,
		userName: "test-user-2",
	}

	valueReceiver(user1)
	fmt.Println("user1 details after call to valueReceiver() : ", user1)
	pointerReceiver(&user2)
	fmt.Println("user2 details after call to pointerReceiver() : ", user2)
}
