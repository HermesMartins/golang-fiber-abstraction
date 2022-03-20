package contracts

type Fn func(params []string, body []byte) string

type IHttp interface {
	Get(url string, fn Fn)
	Listen(port string)
}
