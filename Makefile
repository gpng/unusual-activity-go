GOCMD=go
GOBUILD=$(GOCMD) build

INPUTFILE=cmd/main.go
OUTPUTFILE=unusual_activity
OUTPUTFOLDER=build

# compile binaries
compile:
	@echo "=============building binary============="
	$(GOBUILD) -ldflags="-s -w" -o $(OUTPUTFOLDER)/$(OUTPUTFILE) $(INPUTFILE)
	@echo "=============building exe============="
	GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o $(OUTPUTFOLDER)/$(OUTPUTFILE).exe $(INPUTFILE)
	@echo "=============copy data folders to build============="
	cp -R data build/data