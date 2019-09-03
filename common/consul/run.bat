::consul agent -server -ui -data-dir=./tmp -node=consul-1 -bind=127.0.0.1 
::consul agent -dev -config-dir . -data-dir=./tmp
consul agent -server -ui -bootstrap -data-dir=./tmp -node=consul-1 -client=0.0.0.0 -bind=127.0.0.1 -config-dir .