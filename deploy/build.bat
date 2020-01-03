set project_root=%~dp0..

pushd %project_root%
	
	set GOARCH=amd64
	set GOOS=windows
	go build -o %project_root%\artifacts\windows\penv.exe .\src\penv || exit
		
	set GOOS=linux
	go build -o %project_root%\artifacts\linux\penv .\src\penv || exit
	
	set GOOS=darwin
	go build -o %project_root%\artifacts\darwin\penv .\src\penv || exit
		
popd