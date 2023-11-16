SILENT = @
FILES = main.go
TARGET = bin/lab1

.PHONY: clean

TARGET: 
	$(SILENT)go build $(FILES)

clean:
	rm -r $(TARGET)
