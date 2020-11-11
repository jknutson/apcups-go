package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "regexp"
)

type UpsData struct {
  lines []string
}

func main() {
  data := UpsData{}
  scanner := bufio.NewScanner(os.Stdin)
  r := regexp.MustCompile(`^(\S+)(\s+)?\:\s+(.*)$`)
  // TODO: scan the data into a more robust go struct
  for scanner.Scan() {
    data.lines = append(data.lines, scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    log.Println(err)
  }

  for _, line := range data.lines {
    // fmt.Println(line)
    res := r.FindAllStringSubmatch(line, -1)
    if len(res) > 0 {
      fmt.Printf("%v\n", res)
    }
  }
}
