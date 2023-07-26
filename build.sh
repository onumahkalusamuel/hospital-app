#!/bin/sh
cd web
yarn build
cd ../
go build -o hcms.exe main.go