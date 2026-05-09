package hanlder

const HelloServiceName = "hanlder/HelloService"

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	//err类型的返回值方法 Http的返回可以通过修改指针来进行操作
	*reply = "Hello, " + request
	return nil
}
