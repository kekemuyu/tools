package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Hosts struct {
	Domain    string
	Ip        string
	Isupdated bool
}

func getGitHubCom(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(doc, err)
	re := doc.Find("section.panel").Eq(0).Find("table").Find("tr").Find("td").Find("ul.comma-separated").Find("li").Text()
	fmt.Println(re)
	return re
}

//"https://fastly.net.ipaddress.com/github.global.ssl.fastly.net"
//"https://github.com.ipaddress.com"
func getGitHubGlobalSslFastlyNet(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(doc, err)
	re := doc.Find("section.panel").Eq(0).Find("table").Find("tr").Find("td").Find("ul.comma-separated").Find("li").Text()
	fmt.Println(re)
	return re
}

func updateHosts(hostsDir string, hosts []*Hosts) error {
	if len(hosts) <= 0 {
		return errors.New("len hosts <=0")
	}

	fin, err := os.OpenFile(hostsDir, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("open hosts fail")
		return err
	}
	defer fin.Close()

	rwio := bufio.NewReader(fin)

	seekp := int64(0) //文件当前位置
	lineLen := int64(0)

	for {
		linebs, _, err := rwio.ReadLine()
		if err == io.EOF {
			break
		}
		lineLen = int64(len(linebs) + 1)
		for _, v := range hosts {
			if strings.Contains(string(linebs), v.Domain) {
				fmt.Println("edit:", v)
				intext := v.Domain + "   " + v.Ip + "\n"
				fin.WriteAt([]byte(intext), seekp)
				lineLen = int64(len([]byte(intext)))
				v.Isupdated = true
				break
			}
		}

		seekp += lineLen

	}

	//append
	for _, v := range hosts {
		if v.Isupdated == false {
			fmt.Println("append:", v)
			intext := "\n" + v.Domain + "   " + v.Ip + "\n"
			fin.WriteAt([]byte(intext), seekp)
			lineLen = int64(len([]byte(intext)))
			seekp += lineLen
		}
	}
	return nil
}

func main() {
	fmt.Println("可能会损坏hosts文件，请先备份，继续请按回车")
	key := ""
	fmt.Scanln(&key)
	fmt.Println(key)
	ostype := runtime.GOOS
	fmt.Println("ostype:", ostype)

	hosts := make([]*Hosts, 0)

	ip := getGitHubCom("https://github.com.ipaddress.com")
	if ip != "" {
		hosts = append(hosts, &Hosts{"github.com", ip, false})
	}
	ip = getGitHubGlobalSslFastlyNet("https://fastly.net.ipaddress.com/github.global.ssl.fastly.net")
	if ip != "" {
		hosts = append(hosts, &Hosts{"github.global.ssl.fastly.net", ip, false})
	}
	for _, v := range hosts {
		fmt.Println("hosts:", v.Domain, v.Ip, v.Isupdated)
	}

	hostdir := "test.txt"
	if ostype == "windows" {
		hostdir = "C:/Windows/System32/drivers/etc/hosts"
	} else if ostype == "linux" {
		hostdir = "/etc/hosts"
	}
	updateHosts(hostdir, hosts)

	var cmddns, arg string
	if ostype == "windows" {
		cmddns = "ipconfig"
		arg = "/flushdns"
	} else if ostype == "linux" {
	}
	cmd := exec.Command(cmddns, arg)
	_, err := cmd.Output()
	if err == nil {
		fmt.Println("dns update success!")
	}

}
