package main

import (
        "os/exec"
        "io"
        //"os"
        //"encoding/json"
        //"strings"
        "fmt"
)

func PathCheck(bin ...string) bool {
        valid := false
        for _, k := range bin {
                if dir, err := exec.LookPath(k) ; err == nil {
                        valid = true
                        fmt.Println(dir)
                }
        }
        return valid
}

// output of first cmd is input to last cmd
func main() {
        if PathCheck("lemonbar", "slstatus") {
                fmt.Println("no error reported")
        }
        cmd := exec.Command("lemonbar")
        cmd.Start()
        cmd.Wait()
        stdin, _ := cmd.StdinPipe()
        go func() {
                defer stdin.Close()
                sl := exec.Command("slstatus", "-s")
                sl.Start()
                stdout, _ := sl.Output()
                io.WriteString(stdin, string(stdout))
        }()
}

/*func ConfigValues(fpath string) [][]string {
        f, _ := os.Open(fpath)
        defer f.Close()
        jdec := json.NewDecoder(strings.NewReader(f))
}*/
