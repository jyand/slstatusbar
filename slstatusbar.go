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
func PipeStatusBar(outof *exec.Cmd, into *exec.Cmd) {
        if output, err := outof.Output() ; err == nil {
                outof.Start()
                if input, er := into.StdinPipe() ; er == nil {
                        go func() {
                                defer input.Close()
                                io.WriteString(input, string(output))
                        }()
                        into.Run()
                }
        }
}

func main() {
        if PathCheck("lemonbar", "slstatus") {
                fmt.Println("no error reported")
        }
        sl := exec.Command("slstatus", "-s")
        lemon := exec.Command("lemonbar")
        PipeStatusBar(sl, lemon)
}

/*func ConfigValues(fpath string) [][]string {
        f, _ := os.Open(fpath)
        defer f.Close()
        jdec := json.NewDecoder(strings.NewReader(f))
}*/
