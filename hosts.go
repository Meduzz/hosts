package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const helpText = `Usage: hosts
  + hostname [ip] [file] - adds the hostname and ip (default 127.0.0.1) to file (default /etc/hosts).
  - hostname [file] - removes the hostname from the file (default /etc/hosts).
  help - displayes this message.
`

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) <= 1 {
		fmt.Print(helpText)
		os.Exit(0)
	}

	cmd := args[0]
	hostname := args[1]

	if cmd != "+" && cmd != "-" {
		fmt.Print(helpText)
		os.Exit(0)
	}

	ip := "127.0.0.1"
	file := "/etc/hosts"

	if len(args) == 3 {
		if cmd == "+" {
			ip = args[2]
		} else {
			file = args[2]
		}
	}

	if len(args) == 4 {
		file = args[3]
	}

	fileContent, err := readFile(file)

	if err != nil {
		fmt.Printf("An error occured: %s.\n", err)
		os.Exit(1)
	}

	fmt.Printf("Data in: %s %s %s %s\n", cmd, hostname, ip, file)
	fmt.Println(string(fileContent))

	if cmd == "+" {
		err = setHostname(fileContent, hostname, ip, file)
	} else {
		err = dropHostname(fileContent, hostname, file)
	}

	if err != nil {
		fmt.Printf("An error occured storing changes: %s.\n", err)
		os.Exit(1)
	}
}

func readFile(file string) ([]byte, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0755)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	return ioutil.ReadAll(f)
}

func setHostname(data []byte, hostname, ip, file string) error {
	if strings.Contains(string(data), hostname) {
		// find it and replace it

		writeBuff := make([]string, 0)
		hasChanged := false
		seen := false

		rows := toRows(string(data))
		for _, row := range rows {
			if row != "" && !strings.HasPrefix(strings.TrimSpace(row), "#") {
				fileIp, fileHost := toCells(row)
				if fileHost == hostname && fileIp != ip {
					fmt.Printf("Changing %s to %s for %s.\n", fileIp, ip, fileHost)
					writeBuff = append(writeBuff, fmt.Sprintf("%s %s", ip, fileHost))
					hasChanged = true
					seen = true
				} else {
					writeBuff = append(writeBuff, row)
					seen = true
				}
			} else {
				if row != "" {
					writeBuff = append(writeBuff, row)
				}
			}
		}

		if hasChanged {
			return store(strings.Join(writeBuff, "\n"), file)
		} else {
			if !seen {
				return appendHost(hostname, ip, file)
			}
			return nil
		}
	} else {
		// add it
		fmt.Printf("Adding %s %s to %s.\n", ip, hostname, file)
		return appendHost(hostname, ip, file)
	}
}

func dropHostname(data []byte, hostname, file string) error {
	if strings.Contains(string(data), hostname) {
		// find it and remove it

		writeBuff := make([]string, 0)
		hasChanged := false

		rows := toRows(string(data))
		for _, row := range rows {
			if row != "" && !strings.HasPrefix(strings.TrimSpace(row), "#") {
				_, fileHost := toCells(row)
				if fileHost != hostname {
					writeBuff = append(writeBuff, row)
				} else {
					fmt.Printf("Removing row for %s.\n", hostname)
					hasChanged = true
				}
			} else {
				if row != "" {
					writeBuff = append(writeBuff, row)
				}
			}
		}

		if hasChanged {
			return store(strings.Join(writeBuff, "\n"), file)
		}
	}

	return nil
}

func appendHost(hostname, ip, file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0755)

	if err != nil {
		fmt.Printf("There was an error opening file for appending (%s).", err)
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("\n%s %s", ip, hostname))

	return err
}

func store(data string, file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0755)

	if err != nil {
		fmt.Printf("There was an error opening file for appending (%s).", err)
	}

	defer f.Close()

	err = f.Truncate(0)

	if err != nil {
		fmt.Printf("There was an error truncating file: %s.", err)
	}

	_, err = f.WriteString(fmt.Sprintf("%s\n", data))

	return err
}

func toRows(data string) []string {
	return strings.Split(data, "\n")
}

func toCells(data string) (string, string) {
	for strings.Contains(data, "\t") || strings.Contains(data, "  ") {
		data = strings.Replace(data, "\t", " ", -1)
		data = strings.Replace(data, "  ", " ", -1)
	}

	cells := strings.Split(data, " ")
	return cells[0], cells[1]
}
