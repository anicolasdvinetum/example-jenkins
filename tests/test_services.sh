#!/bin/bash

echo "Esperando a que los servicios arranquen..."
sleep 5

echo "Testeando servicio en Go..."
curl -f http://go-service:9090 || exit 1
echo "Go OK"

echo "Testeando servicio en Python..."
curl -f http://python-service:5000 || exit 1
echo "Python OK"

echo "Tests completados correctamente."