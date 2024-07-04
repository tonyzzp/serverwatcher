package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func dumpIps() {
	log.Println("----------- dumpIps")
	inters, e := net.Interfaces()
	if e != nil {
		log.Println(e)
		return
	}
	for _, inter := range inters {
		log.Println("------------")
		log.Println(inter.Name, inter.HardwareAddr)
		addrs, e := inter.Addrs()
		if e != nil {
			log.Println(e)
			continue
		}
		for _, addr := range addrs {
			log.Println(addr)
		}
	}
}

func dumpIpconfig() {
	log.Println("----------- dumpIpconfig")
	var ipconfig = ""
	if runtime.GOOS == "windows" {
		ipconfig = "ipconfig"
	} else {
		ipconfig = "ifconfig"
	}

	cmd := exec.Command(ipconfig)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	e := cmd.Start()
	if e != nil {
		log.Println(e)
		return
	}
	e = cmd.Wait()
	if e != nil {
		log.Println(e)
		return
	}
	var state = cmd.ProcessState
	log.Println("exit", state.ExitCode())
}

func ping() bool {
	log.Println("----------- ping")
	resp, e := http.Get("https://www.baidu.com")
	if e != nil {
		log.Println(e)
		return false
	}
	log.Println("status", resp.StatusCode)
	return resp.StatusCode == 200
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(os.Stdout)
	log.Println("--------------------------------------------")
	dumpIpconfig()
	dumpIps()
	ping()
}
