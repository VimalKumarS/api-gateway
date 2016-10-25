package service

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
    "fmt"

    "github.com/unrolled/render"

)

func gateWayHandler(formatter *render.Render, repo repository) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        url := strings.Split(req.URL.String(), "/")
        fmt.Println(url)
        fmt.Println(len(url))
        if len(url) <= 2 {
            formatter.JSON(w, http.StatusNotFound, "Not valid route")
            return
        }
        root := url[2]
        fmt.Println(root)
        value, err := repo.redisGetValue(root)
        fmt.Println(value)
        if err != nil {
            formatter.JSON(w, http.StatusNotFound, "Not valid route")
            return
        }

        service := ServiceWebClient{URL:value}
        addOn := strings.Join(url[2:], "/")
        response := service.SendCommand(req.Method, addOn, req.Body, req.Header)
        payload, _ := ioutil.ReadAll(response.Body)
        var res interface{}
        json.Unmarshal(payload, &res)
        formatter.JSON(w, response.StatusCode, res)
    }
}

func postAddServiceHandler(formatter *render.Render, repo repository) http.HandlerFunc  {
    return func(w http.ResponseWriter, req *http.Request) {
        var service Service
        payload, _ := ioutil.ReadAll(req.Body)
        err := json.Unmarshal(payload, &service)
        if err != nil || (service == Service{}) {
            formatter.JSON(w, http.StatusBadRequest, "Failed to parse service.")
            return
        }
        fmt.Println(service)
        err = repo.redisSetValue(service.Name, service.URL)
        fmt.Println(err)
        formatter.JSON(w, http.StatusOK, "Added service.")
    }
}
