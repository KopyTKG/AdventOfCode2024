package stream

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadBytes(f string) []byte {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	var data []byte

	reader := bufio.NewReader(file)
	buf := make([]byte, 1)
	for {
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
		}

		data = append(data, buf[0])
	}

	return data

}

func BtoSa(d []byte) []string {
	var lines []string
	line := ""
	for i := 0; i < len(d); i++ {
		item := d[i]

		if item != 10 {
			line += string(item)
		} else {
			lines = append(lines, line)
			line = ""
		}
	}
	return lines
}
