package core

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func (o *Ocsa) RunServer() {
	listener := o.getListner()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting: %s\n", err)
			continue
		}

		go o.handleConnection(conn)
	}

}

func (o *Ocsa) handleConnection(conn net.Conn) {

	headerAvailable := false

	defer conn.Close()

	buf := make([]byte, 1024*1024*100)

	if o.verbose {
		log.Println("server: conn: opened")
	}

	fileHeader := OcsaHeader{}
	file := &os.File{}

	io.WriteString(conn, "<<<<START_HEADER>>>>")

	for {

		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		data := string(buf[:n])

		if !headerAvailable {
			isSet, errc, header := o.parseHeader(data)
			if errc {
				if o.verbose {
					log.Println("server: conn: header parsing error")
				}
				break
			}
			headerAvailable = isSet

			if headerAvailable {
				fileHeader = header

				_, err = io.WriteString(conn, "<<<<START_FILE>>>>")
				if err != nil {
					break
				}

				err = os.Remove(o.rootDir + fileHeader.FilePath)
				os.MkdirAll(filepath.Dir(o.rootDir+fileHeader.FilePath), os.ModePerm)

				if (err != nil) && (err.Error() != "remove "+o.rootDir+fileHeader.FilePath+": no such file or directory") {
					log.Println("server: conn: file remove error")
					break
				}

				f, err := os.OpenFile(o.rootDir+fileHeader.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

				if err != nil {
					log.Println("server: conn: file open error")
					break
				}

				file = f

			}

			continue

		} else {

			//fmt.Println("data: ", string(buf[:n]))

			if strings.HasPrefix(string(buf[:n]), "<<<<END_FILE>>>>") {
				//fmt.Println("end: ")
				file.Close()
				break
			}

			file.Write(buf[:n])

			//_, err = o.file.Write(buf[:n])
			//if err != nil {
			//	log.Println("server: conn: file write error")
			//	break

			//}
		}

	}

	log.Println("server: conn: closed")

}



func (o *Ocsa) parseHeader(headersText string) (bool, bool, OcsaHeader) {
	//fmt.Println("headersText: ", headersText)

	header := OcsaHeader{}

	headerLines := strings.Split(headersText, "\n")

	for _, line := range headerLines {
		if strings.HasPrefix(line, "<<<<END_HEADER>>>>") {
			break
		} else if strings.HasPrefix(line, "FILE_PATH:") {
			header.FilePath = strings.Trim(line, "FILE_PATH:")
		} else if strings.HasPrefix(line, "AUTH_TOKEN:") {
			header.Token = strings.Trim(line, "AUTH_TOKEN:")
		} else {
			return false, true, OcsaHeader{}
		}
	}

	if header.FilePath == "" || header.Token == "" {
		return false, false, OcsaHeader{}
	} else {
		//fmt.Println("fileHeaders: ", o.fileHeaders)
		return true, false, header
	}

}
