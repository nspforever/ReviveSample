package main

import (
    "github.com/kolo/xmlrpc"
    "fmt"
    "log"
)

type BannerInfo struct {
    BannerId int `xmlrpc:"bannerId"`
    CampaignId int `xmlrpc:"campaignId"`
    BannerName string `xmlrpc:"bannerName"`
    StorageType string `xmlrpc:"storageType"`
    ImageURL string `xmlrpc:"imageURL"`
    HtmlTemplate string `xmlrpc:"htmlTemplate"`
    Width int `xmlrpc:"width"`
    Height int `xmlrpc:"height"`
    Weight int `xmlrpc:"Weight"`
    Target string `xmlrpc:"target"`
    Url string `xmlrpc:"url"`
    BannerText string `xmlrpc:"bannerText"`
    Status int `xmlrpc:"status"`
    Adserver string `xmlrpc:"adserver"`
    Transparent int `xmlrpc:"transparent"`
    Capping int `xmlrpc:"capping"`
    SessionCapping int `xmlrpc:"sessionCapping"`
    Block int `xmlrpc:"block"`
    Comments string `xmlrpc:"comments"`
    Alt string `xmlrpc:"alt"`
    Filename string `xmlrpc:"filename"`
    Append string `xmlrpc:"append"`
    Prepend string `xmlrpc:"prepend"`
}

func main() {

    client := GetXmlRpcClient("https://revive-gudu.c9users.io/revive/www/api/v2/xmlrpc/index.php")
    sessionId := GetSessionId(client);
    bannerInfo := GetBannerInfo(client, sessionId, 1);
    fmt.Printf("banner status: %d", bannerInfo.Status)
    if bannerInfo.Status == 0 {
        bannerInfo.Status = 1
    } else {
        bannerInfo.Status = 0
    }
    ret := ModifyBanner(client, sessionId, bannerInfo);
    fmt.Printf("Modify result: %t", ret)
}

func GetXmlRpcClient(serverUrl string) (*xmlrpc.Client) {
    client, err := xmlrpc.NewClient(serverUrl, nil)
    if err != nil {
        log.Fatal(err)
    }
    return client;
}

func HelloBach(rpcClient *xmlrpc.Client) string{
    //client, _ := xmlrpc.NewClient("https://revive-gudu.c9users.io/revive/www/api/v2/xmlrpc/index.php", nil)
    greeting := new (string);
    rpcClient.Call("ox.helloBach", nil, &greeting)
    fmt.Printf("Greeting: %s\n", *greeting)
    return *greeting
}

func GetSessionId(rpcClient *xmlrpc.Client) string {
    //client, _ := xmlrpc.NewClient("https://revive-gudu.c9users.io/revive/www/api/v2/xmlrpc/index.php", nil)
    sessionId := new (string);
    var interfaceSlice []interface{} = make([]interface{}, 2)
    interfaceSlice[0] = "admin"
    interfaceSlice[1] = "123"

    err := rpcClient.Call("ox.logon", interfaceSlice, &sessionId)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("SessionId: %s\n", *sessionId)
    return *sessionId
}

func GetBannerInfo(rpcClient *xmlrpc.Client, sessionId string, bannerId int) (BannerInfo) {
    bannerInfo := BannerInfo{};
    var interfaceSlice []interface{} = make([]interface{}, 2)
    interfaceSlice[0] = sessionId
    interfaceSlice[1] = bannerId

    err := rpcClient.Call("ox.getBanner", interfaceSlice, &bannerInfo)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Status: %d\n", bannerInfo.Status)
    return bannerInfo;
}

func ModifyBanner(rpcClient *xmlrpc.Client, sessionId string, bannerInfo BannerInfo) (bool) {
    var interfaceSlice []interface{} = make([]interface{}, 2)
    interfaceSlice[0] = sessionId
    interfaceSlice[1] = bannerInfo
    result := false
    err := rpcClient.Call("ox.modifyBanner", interfaceSlice, &result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Result: %t\n", result)
    return result;
}

