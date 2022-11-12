package generator

/* A listener is a mothership */
type LISTENER interface{
  StartListener()
}

type TcpListener struct{}
func (tl *TcpListener) StartListener(){}

type UdpListener struct{}
func (ul *UdpListener) StartListener(){}

type DnsListener struct{}
func (dl *DnsListener) StartListener(){}

type TlsListener struct{}
func (Tls *TlsListener) StartListener(){}

type HttpListener struct{}
func (httpl *HttpListener) StartListener(){}

type HttpsListener struct{}
func (httpsl *HttpsListener) StartListener(){}
