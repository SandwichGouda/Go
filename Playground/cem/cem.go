package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// app := "echo"

	// arg0 := "-e"
	// arg1 := "Hello world"
	// arg2 := "\n\tfrom"
	// arg3 := "golang"

	// 	curl 'https://cemantix.certitudes.org/score?n=1087' -H 'origin: https://cemantix.certitudes.org' --data-raw 'word=négatif' \
	//   -H 'accept: */*' \
	//   -H 'accept-language: en-GB,en;q=0.9' \
	//   -H 'content-type: application/x-www-form-urlencoded' \
	//   -b '_ga=GA1.1.735944376.1740169438; FCCDCF=%5Bnull%2Cnull%2Cnull%2C%5B%22CQNLZQAQNLZQAEsACBENBeFgAAAAAEPgAAqIAAANVSD2F2K2kKFkPCmwXYAYBCujYAAhQgAAAkCBMACgAUgQAgFJIAgCIFAAAAAAAAAQEiCQAAQABAAAIACgAAAAAAIAAAAAAAQQAABAAIAAAAAAAAEAQAAIAAQAAAAIAABEhAAAQQAEAAAAAAAQAA%22%2C%222~~dv.70.89.93.108.122.149.184.196.236.259.311.313.323.358.415.442.486.494.495.540.574.609.864.981.1029.1048.1051.1095.1097.1126.1205.1276.1301.1365.1415.1449.1514.1570.1577.1598.1651.1716.1735.1753.1765.1870.1878.1889.1958.1960.2072.2253.2299.2373.2415.2506.2526.2531.2568.2571.2575.2624.2677.2778%22%2C%22DD19B502-7EA2-40C8-8436-6E2C6B9C8589%22%5D%5D; __eoi=ID=8b44fa51eef4c0ec:T=1740169444:RT=1740169444:S=AA-Afja3i3r4Yh1tlcoGF3-P7FDR; _ga_7E8QW0JJHF=GS1.1.1740169438.1.0.1740169447.0.0.0' \
	//   -H 'dnt: 1' \
	//    \
	//   -H 'priority: u=1, i' \
	//   -H 'referer: https://cemantix.certitudes.org/' \
	//   -H 'sec-ch-ua: "Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'sec-ch-ua-platform: "Windows"' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-site: same-origin' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36' \

	// tail := "tail"
	// cem := "cem.go"
	// opt := "-n 3"

	// cmd := exec.Command(tail, cem, opt)
	// stdout, err := cmd.Output()

	curl := "curl"
	link := "'https://cemantix.certitudes.org/score?n=1087'"
	origin := "-H 'origin: https://cemantix.certitudes.org'"
	data := "--data-raw 'word=négatif'"

	// tail := "tail"
	// cem := "cem.go"
	// opt := "-n 3"

	// cmd := exec.Command(tail, cem, opt)
	// stdout, err := cmd.Output()

	cmd := exec.Command(curl, origin, link, data)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
