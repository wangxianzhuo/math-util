package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Conf 配置
type Conf struct {
	Port string `json:"port"`
}

// StartServer 启动一个http服务器
func StartServer(conf Conf) {
	// http.HandleFunc("/det/value", GetDeterminantValue)
	// http.HandleFunc("/det/expanse", GetDeterminantExpanse)
	// http.HandleFunc("/util/inversionnumber", CalculateInversionNumber)
	// http.HandleFunc("/equations/solves", SolveEquations)

	// log.Printf("Listening %v\n", conf.Port)
	// err := http.ListenAndServe(conf.Port, nil) //设置监听的端口
	// if err != nil {
	// 	log.Fatalf("Listen %v error: %v", conf.Port, err)
	// }
}

// GetDeterminantExpanse 求行列式展开形式
func GetDeterminantExpanse(w http.ResponseWriter, r *http.Request) {

}

// CalculateInversionNumber 计算逆序数
func CalculateInversionNumber(w http.ResponseWriter, r *http.Request) {

}

// SolveEquations 方程组求解，默认克拉默法则
func SolveEquations(w http.ResponseWriter, r *http.Request) {

}

type ErrMessage struct {
	Status  int
	Message string
	Err     error
}

func errHandler(err ErrMessage, w http.ResponseWriter) {
	log.Printf("%v: %v", err.Message, err.Err)
	http.Error(w, err.Message, err.Status)
}

type CalcResult struct {
	Value interface{} `json:"value"`
}

func encodeCalcResult(cr CalcResult) ([]byte, error) {
	return json.Marshal(cr)
}
