package main

import (
	"bufio"

	"fmt"
	"io"
	"os"

	"os/exec"
	"strings"
)

type Ipmac struct {
	ip     string
	mac    string
	vendor string
}

func main() {

	imacs, err := getips()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("IP                      mac                       vendor")
	for _, v := range imacs {

		fmt.Printf("%s        %s       %s\n", v.ip, v.mac, v.vendor)
	}

}

func getips() ([]Ipmac, error) {
	cmd := exec.Command("arp", "-a")
	re, err := cmd.Output()

	ipmac := make([]Ipmac, 0)

	if err != nil {
		fmt.Println(err)
		return ipmac, err
	}

	res := string(re)
	lines := strings.Split(res, "\n")

	for _, v := range lines {
		if len(v) > 0 {
			v = strings.TrimLeft(v, " ")
			v = strings.TrimRight(v, " ")
			out := strings.Split(v, " ")

			ip := out[0]
			var mac string

			for _, v := range out[1:] {
				if strings.TrimSpace(v) != "" {
					mac = v
					break
				}
			}
			if len(out) > 1 {
				if len(mac) < 17 {

					continue
				}

				ipmac = append(ipmac, Ipmac{ip, mac, getVendor(mac)})
			}
		}
	}

	return ipmac, nil

}

func getVendor(mac string) string {

	fi, err := os.Open("oui.txt")
	if err != nil {
		return ""
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		lines, _, err := br.ReadLine()
		if err == io.EOF {
			return ""
		}
		if len(lines) > 0 {
			linestr := strings.TrimSpace(string(lines))
			linestrs := strings.Split(linestr, "(hex)")

			if strings.TrimRight(linestrs[0], " ") == strings.ToUpper(mac[:8]) {
				linestrs[1] = strings.TrimLeft(linestrs[1], "	")
				return strings.TrimLeft(linestrs[1], " ")
			}
		}
	}
}
