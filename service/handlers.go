package service

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
    "github.com/unrolled/render"

)

func gateWayHandler(formatter *render.Render, repo repository) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        url := strings.Split(req.URL.String(), "/")
        if len(url) <= 2 {
            formatter.JSON(w, http.StatusNotFound, "Not valid route")
            return
        }

        root := url[1]
        value, err := repo.redisGetValue(root)
        if err != nil {
            formatter.JSON(w, http.StatusNotFound, "Not valid route")
            return
        }

        service := ServiceWebClient{URL:value}
        response := service.SendCommand(req.Method, value, req.Body)
        formatter.JSON(w, response.StatusCode, response.Body)
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
        repo.redisSetValue(service.Name, service.URL)
        formatter.JSON(w, http.StatusOK, "Added service.")
    }
}
