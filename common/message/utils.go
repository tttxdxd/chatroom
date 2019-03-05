package message

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

var (
	data   []byte
	length []byte
)

func init() {
	data = make([]byte, 4096)
	length = make([]byte, 4)
}

func ReadMsg(conn net.Conn) (msg Msg, err error) {

	//读取数据长度
	n, err := conn.Read(data[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Read(data[:4]) error:", err)
		return
	}
	msgLen := binary.BigEndian.Uint32(data[:4])

	//根据长度读取数据
	n, err = conn.Read(data[:msgLen])
	if uint32(n) != msgLen || err != nil {
		fmt.Println("conn.Read(data[:msgLen]) error:", err)
		return
	}

	//获取数据
	err = json.Unmarshal(data[:msgLen], &msg)
	if err != nil {
		fmt.Println("Unserializer(data[:msgLen]) error:", err)
		return
	}
	return
}

func WriteMsg(conn net.Conn, msg Msg) (err error) {
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("Marshal(msg) error:", err)
		return
	}
	msglen := uint32(len(data))
	binary.BigEndian.PutUint32(length, msglen)

	n, err := conn.Write(length)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(length) error:", err)
		return
	}

	n, err = conn.Write(data[:msglen])
	if uint32(n) != msglen || err != nil {
		fmt.Println("conn.Write(data[:msglen]) error:", err)
		return
	}
	return
}

func ReadResponse(conn net.Conn) (msg Response, err error) {

	//读取数据长度
	n, err := conn.Read(data[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Read(data[:4]) error:", err)
		return
	}
	msgLen := binary.BigEndian.Uint32(data[:4])

	//根据长度读取数据
	n, err = conn.Read(data[:msgLen])
	if uint32(n) != msgLen || err != nil {
		fmt.Println("conn.Read(data[:msgLen]) error:", err)
		return
	}

	//获取数据
	err = json.Unmarshal(data[:msgLen], &msg)
	if err != nil {
		fmt.Println("Unserializer(data[:msgLen]) error:", err)
		return
	}
	return
}

func WriteResponse(conn net.Conn, msg Response) (err error) {
	data, err = Serializer(msg)
	if err != nil {
		fmt.Println("Serializer(msg) error:", err)
		return
	}
	msglen := uint32(len(data))
	binary.BigEndian.PutUint32(length, msglen)

	n, err := conn.Write(length)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(length) error:", err)
		return
	}

	n, err = conn.Write(data[:msglen])
	if uint32(n) != msglen || err != nil {
		fmt.Println("conn.Write(data[:msglen]) error:", err)
		return
	}
	return
}
