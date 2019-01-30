package main

import (
	"time"
	"strings"
	"fmt"
	"bufio"
	"os"
	"io"
)

type Reader interface {
	Read(chan string)
}

type Writer interface {
	Write(chan string)
}

type ReadFromFile struct {
	path string
}

type LogProcess struct {
	reader Reader
	writer Writer
	rc     chan string
	wc     chan string
}

type WriteToInfluxDB struct {
	dsn string
}

func (r *ReadFromFile) Read(rc chan string) {

	f, err := os.Open(r.path)
	defer f.Close()
	if err != nil {
		panic(fmt.Sprintf("open file fail,err:%s", err.Error()))
	}

	f.Seek(0, io.SeekEnd)
	reader := bufio.NewReader(f)
	for {
		message, err := reader.ReadBytes(byte('\n'))
		if err == io.EOF {
			fmt.Println(fmt.Sprintf("read file fail eof"))
			continue
		} else if err != nil {
			fmt.Println(fmt.Sprintf("read file fail,err:%s", err.Error()))
			continue
		}
		rc <- string(message)
	}
}

func (l *LogProcess) Parse() {
	l.wc <- strings.ToUpper(<-l.rc)
}

func (w *WriteToInfluxDB) Write(wc chan string) {
	fmt.Println(<-wc)
}

func main() {

	r := &ReadFromFile{
		path: "access.log",
	}

	w := &WriteToInfluxDB{
		dsn: "",
	}

	logProcess := &LogProcess{
		reader: r,
		writer: w,
		rc:     make(chan string),
		wc:     make(chan string),
	}

	go logProcess.reader.Read(logProcess.rc)
	go logProcess.Parse()
	go logProcess.writer.Write(logProcess.wc)

	time.Sleep(time.Second)
}
