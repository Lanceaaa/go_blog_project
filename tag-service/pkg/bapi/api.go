package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	APP_KEY    = "lance"
	APP_SECRET = "go-programming-tour-book"
)

type AccessToken struct {
	Token string `josn:"token"`
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	// body, err := a.httpGet(ctx, fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", APP_KEY, APP_SECRET))
	body, err := a.httpPost(ctx, "auth", url.Values{"app_key": {APP_KEY}, "app_secret": {APP_SECRET}})
	if err != nil {
		return "", err
	}

	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	// resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	url := fmt.Sprintf("%s/%s", a.URL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 创建设置当前跨度的信息和标签内容（传入上下文信息，以保证链路完整性）
	span, newCtx := opentracing.StartSpanFromContext(
		ctx, "HTTP GET："+a.URL,
		opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
	)
	span.SetTag("url", url)
	// 然后传入附带的信息，并把它设置到对应的链路信息上
	_ = opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	// 最后进行调用，并返回新的上下文
	req = req.WithContext(newCtx)
	client := http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	defer span.Finish()

	// 读取消息主体，在实际封装中可以将其抽离
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
	// resp, err := ctxhttp.Get(ctx, http.DefaultClient, fmt.Sprintf("%s/%s", a.URL, path))
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// return body, nil
}

func (a *API) httpPost(ctx context.Context, path string, data url.Values) ([]byte, error) {
	resp, err := http.PostForm(fmt.Sprintf("%s/%s", a.URL, path), data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		return nil, err
	}

	return body, nil
}
