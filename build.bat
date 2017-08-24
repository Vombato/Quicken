@echo off

SET APP_NAME=Quicken Project Generator
SET APP_PATH=%~dp0

cd /d %~dp0
ECHO
SET NAMESPACE=github.com/matteojoliveau/quicken
SET MODULES=cmd modules recipe utils
SET WORKDIR=%APP_PATH%build\go\src
ECHO %WORKDIR%

if not exist %WORKDIR% mkdir "%WORKDIR%\%NAMESPACE%"

setlocal ENABLEDELAYEDEXPANSION

cd %APP_PATH%

SET GOPATH=%APP_PATH%build\go

goto :run

:run
ECHO --------------------------------------------------------------
ECHO Copying project modules...
For %%a in (%MODULES%) do robocopy  %%a "%WORKDIR%\%NAMESPACE%\%%a" /MIR
 rem xcopy /s /y %%a "%WORKDIR%\%NAMESPACE%\%%a"

IF "%1" NEQ "-d" (
    ECHO Getting dependencies...
    go get ./...
) ELSE ECHO DEBUG MODE: Skipping dependencies download
ECHO --------------------------------------------------------------
echo Building project...
go build -o build/qk.exe
ECHO Done!
goto :end
:end
endlocal
REM pause
