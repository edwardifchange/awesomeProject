package json_demo

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	//ServerName  string `json:"name"` //可映射名字
	//ServerIp    string `json:"ip"`
	//ServerPoint int    `json:"point"`

	ServerName  string
	ServerIp    string
	ServerPoint int
}

//序列化结构体
func SerializeStruct() {
	server := new(Server)
	server.ServerName = "Demo-for-json-struct"
	server.ServerIp = "127.0.0.1"
	server.ServerPoint = 8080

	b, err := json.Marshal(server) //序列化成json 字节数组
	if err != nil {
		fmt.Println("Marshal error : ", err.Error())
		return
	}
	fmt.Println("Marshal json : ", string(b)) //将json字节数组转化为string
}

//序列化结map
func SerializeMap() {
	server := make(map[string]interface{})
	server["ServerName"] = "Demo-for-json-map"
	server["ServerIp"] = "192.168.31.100"
	server["ServerPoint"] = 9090

	b, err := json.Marshal(server) //序列化成json 字节数组
	if err != nil {
		fmt.Println("Marshal error : ", err.Error())
		return
	}
	fmt.Println("Marshal json : ", string(b)) //将json字节数组转化为string
}

//反序列化结构体
func DeSerializeStruct() {
	jsonString := `{"ServerIp":"192.168.31.100","ServerName":"Demo-for-json-map","ServerPoint":9090}`
	server := new(Server)
	err := json.Unmarshal([]byte(jsonString), &server)
	if err != nil{
		fmt.Println("Unmarshal error : ", err.Error())
		return
	}
	fmt.Println("Unmarshal struct : ", server)
}

//反序列化map
func DeSerializeMap() {
	jsonString := `{"ServerIp":"192.168.31.100","ServerName":"Demo-for-json-map","ServerPoint":9090}`
	server := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &server)
	if err != nil{
		fmt.Println("Unmarshal error : ", err.Error())
		return
	}
	fmt.Println("Unmarshal map : ", server)
}