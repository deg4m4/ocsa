package core

type Ocsa struct {
	tls bool
	tlsOptions struct{
		cert string
		key string
	}

	host string
	port int

	rootDir string

	verbose bool

}

type OcsaHeader struct {
	FilePath string
	Token string
}

func (o *Ocsa) SetHost(host string) {
	o.host = host
}

func (o *Ocsa) SetPort(port int) {
	o.port = port
}

func (o *Ocsa) SetTls(tls bool) {
	o.tls = tls
}

func (o *Ocsa) SetTlsOptions(cert string, key string) {
	o.tlsOptions.cert = cert
	o.tlsOptions.key = key
}

func (o *Ocsa) SetRootDir(rootDir string) {
	o.rootDir = rootDir
}

func (o *Ocsa) SetVerbose(verbose bool) {
	o.verbose = verbose
}
