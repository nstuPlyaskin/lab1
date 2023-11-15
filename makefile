SILENT = @
FILES = main.go
TARGET = bin/lab1

.PHONY: clean

TARGET: 
	$(SILENT)go run $(FILES)

clean:
	rm -r $(TARGET)
