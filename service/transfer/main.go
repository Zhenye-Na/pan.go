package main

import (
    "bufio"
    "encoding/json"
    "log"
    "os"
    "pan.go/config"
    dblayer "pan.go/db"
    "pan.go/mq"
    "pan.go/store/oss"
)

// ProcessTransfer : 处理文件转移
func ProcessTransfer(msg []byte) bool {
    log.Println(string(msg))

    pubData := mq.TransferData{}
    err := json.Unmarshal(msg, &pubData)
    if err != nil {
        log.Println(err.Error())
        return false
    }

    fin, err := os.Open(pubData.CurLocation)
    if err != nil {
        log.Println(err.Error())
        return false
    }

    err = oss.Bucket().PutObject(
        pubData.DestLocation,
        bufio.NewReader(fin))
    if err != nil {
        log.Println(err.Error())
        return false
    }

    _ = dblayer.UpdateFileLocation(
        pubData.FileHash,
        pubData.DestLocation)
    return true
}

func main() {
    if !config.AsyncTransferEnable {
        log.Println("异步转移文件功能目前被禁用，请检查相关配置")
        return
    }
    log.Println("文件转移服务启动中，开始监听转移任务队列...")
    mq.StartConsume(
        config.TransOSSQueueName,
        "transfer_oss",
        ProcessTransfer)
}
