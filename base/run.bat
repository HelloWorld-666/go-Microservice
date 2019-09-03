@echo off
set CONSUL_CONFIG_ADDRESS=127.0.0.1:8500
base.exe --registry=consul -- registry_address=127.0.0.1:8500