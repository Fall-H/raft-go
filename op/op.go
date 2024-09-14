package op

import "raft/model"

func Op(op string, number int32) {
	switch op {
	case "add":
		model.Number += number
		break
	case "sub":
		model.Number -= number
		break
	case "set":
		model.Number = number
		break
	}
}
