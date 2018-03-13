package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)



func main(){
	scanner := bufio.NewScanner(os.Stdin)

	connection,err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("Connection au serveur échouée")
	} else{
		writer := bufio.NewWriter(connection)

    //saisie de la requête
    var scanner2 string
    if scanner.Scan(){
      scanner2 = scanner.Text()
    }
    _,_ = writer.WriteString(scanner2+"\n")
    writer.Flush()
		}

}
