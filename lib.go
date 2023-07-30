package oslib

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strconv"
)

func PrintErrIfNotNil(err error) {
	if err != nil {
		log.Println(err)
	}
}

func PanicErrIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfNotNil(err interface{}, msg string, args ...interface{}) {
	if err != nil {
		panic(fmt.Sprintf(msg, args...))
	}
}

func MustExec(f func() *exec.Cmd) []byte {
	cmd := f()
	log.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(string(out[:]))
	}
	// log.Println(string(out[:]))
	return out
}

func Pwd() string {
	cmd := exec.Command("pwd")
	cmd.Env = os.Environ()
	ret, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(ret[:])
}

func Cp(from string, to string) {
	args := []string{
		from,
		to,
	}
	cmd := exec.Command("cp", args...)
	cmd.Env = os.Environ()
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}

func Cd(dirname string) {
	log.Println("cd", dirname)
	err := os.Chdir(dirname)
	if err != nil {
		log.Fatal(err)
	}
}

func Mkdir(dirname string) {
	newpath := filepath.Join(".", dirname)
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func MustParseUint64(str string) (ret uint64) {
	ret, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		debug.PrintStack()
		log.Fatalln(err)
	}
	return
}

func MustParseInt(str string) (ret int) {
	ret, err := strconv.Atoi(str)
	if err != nil {
		debug.PrintStack()
		log.Fatalln(err)
	}
	return
}

func HexToHexstring(input string) string {
	del := "\\x"
	ret := ""
	for idx, c := range input {
		if idx%2 == 0 {
			ret = ret + del

		}
		ret = ret + string(c)
	}
	return ret
}

func ToLE16(n int) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(n))
	return b
}
