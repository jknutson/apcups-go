package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "log"
  "os"
  "regexp"
  "time"
)

/*
type metricData struct {
	Series []struct {
    MetricName string     `json:"metric"`
    Points     [][]string `json:"points"`
    Tags       []string   `json:"tags"`
  } `json:"series"`
}
*/

type metricSeries struct {
	MetricName string     `json:"metric"`
	Points     [][]string `json:"points"`
	Tags       []string   `json:"tags"`
	MetricType string     `json:"type"`
}

type metricData struct {
	Series []metricSeries `json:"series"`
}

type UpsData struct {
  lines []string
}

func main() {
  /*
  datadogAPIKey := os.Getenv("DD_API_KEY")
  datadogAPIUrl := fmt.Sprintf("https://api.datadoghq.com/api/v1/series?api_key=%s", datadogAPIKey)
  */

  upsData := UpsData{}
  scanner := bufio.NewScanner(os.Stdin)
  // r := regexp.MustCompile(`^(\S+)(\s+)?\:\s+(.*)$`)
  r := regexp.MustCompile(`^(\S+)(\s+)?\:\s+([\d+\.\d+])`)
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
      // fmt.Printf("%s - %s\n", res[1], res[3])
      m[res[1]] = res[3]
    }
  }

  hostname, err := os.Hostname()
  if err != nil {
    log.Fatal(err)
  }
  tags := []string{fmt.Sprintf("device:%s", hostname)}
  metricPoints := []string{fmt.Sprintf("%v", time.Now().Unix()), m["LOADPCT"].(string)}
  var series []metricSeries
  series = append(series, metricSeries{MetricName: "apcupsd.loadpct", Points: [][]string{metricPoints}, Tags: tags, MetricType: "gauge"})
  payload := metricData{Series: series}
  jsonPayload, err := json.Marshal(payload)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("\n%s\n\n", jsonPayload)


  /*
  req, err := http.NewRequest("POST", datadogAPIUrl, bytes.NewBuffer(payload))
  if err != nil {
    log.Fatal(err)
  }
  */
}
