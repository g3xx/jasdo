package msg

type Msg struct {
	Status string
	Text   string
	Data   []interface{}
	Total  int64
}

func Pesan(s string, m string) Msg {
	msg := Msg{s, m, nil, 0}
	return msg
}

func DbPesan(t []interface{}, i int64) Msg {
	msg := Msg{"Success", "", t, i}
	return msg
}
