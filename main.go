package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "log"
  "os"
  "regexp"
)

type UpsData struct {
  lines []string
}

func main() {
  upsData := UpsData{}
  scanner := bufio.NewScanner(os.Stdin)
  r := regexp.MustCompile(`^(\S+)(\s+)?\:\s+(.*)$`)
  // TODO: scan the data into a more robust go struct
  for scanner.Scan() {
    upsData.lines = append(upsData.lines, scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    log.Println(err)
  }

  m := map[string]interface{}{}

  for _, line := range upsData.lines {
    // fmt.Println(line)
    res := r.FindStringSubmatch(line)
    if len(res) > 2 {
      // fmt.Printf("%v\n", res)
      fmt.Printf("%s - %s\n", res[1], res[3])
      m[res[1]] = res[3]
    }
  }
  fmt.Printf("\n%v\n", m)
  jsonData, err := json.Marshal(m)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("\n%s\n", jsonData)
}
