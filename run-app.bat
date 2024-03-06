@echo off
REM Запустить фронтенд
cd "frontend"
start npm start

REM Ожидание запуска фронтенда (предполагается, что порт 3000 используется)
timeout /t 10 /nobreak

REM Запустить Go-файл
cd "%~dp0\backend"
start go run main.go