package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "sync"
)

type Comment struct {
    PostID int    `json:"postId"`
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Email  string `json:"email"`
    Body   string `json:"body"`
}

func main() {
    var comments []Comment
    var wg sync.WaitGroup

    for i := 1; i <= 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/comments", i)
            resp, err := http.Get(url)
            if err != nil {
                fmt.Printf("请求地址失败，失败原因: %s\n", err)
                return
            }
            defer resp.Body.Close()

            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                fmt.Printf("获取body信息失败，失败原因: %s\n", err)
                return
            }

            var pageComments []Comment
            err = json.Unmarshal(body, &pageComments)
            if err != nil {
                fmt.Printf("解析json信息失败，失败原因: %s\n", url, err)
                return
            }

            comments = append(comments, pageComments...)
        }(i)
    }

    wg.Wait()

    if len(comments) == 0 {
        fmt.Println("没有找到资料")
        return
    }

    var emails []string
    for _, comment := range comments {
        emails = append(emails, comment.Email)
    }

    file, err := os.Create("emails.txt")
    if err != nil {
        fmt.Println("创建emails.txt文件失败，失败原因:", err)
        return
    }
    defer file.Close()

    for _, email := range emails {
        fmt.Fprintln(file, email)
    }

    fmt.Println("Emails内容已保存到emails.txt")
}
