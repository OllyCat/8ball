build:
	go build -ldflags "-s -w" -o release/linux/amd64/8ball
	GOOS=windows go build -ldflags "-s -w" -o release/windows/amd64/8ball.exe
	GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o release/windows/386/8ball.windows.386.exe
	#GOOS=android go build -ldflags "-s -w" -o release/android/go8ball.apk
	#GOOS=darwin go build -ldflags "-s -w" -o release/macosx/8ball
