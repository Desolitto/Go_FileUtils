.PHONY: myFind myFindA
myFind: 
	@echo "\033[38;5;212mRunning tests...\033[0m" && \
	go build -o myFind ./cmd/myFind && \
	echo "\033[38;5;212mRunning test 1...\033[0m" && \
	./myFind testdata/test01 && \
	echo "\033[38;5;212mRunning test 2 (sl)...\033[0m" && \
	./myFind -sl testdata/test01 && \
	echo "\033[38;5;212mRunning test 3 (d)...\033[0m" && \
	./myFind -d testdata/test01 && \
	echo "\033[38;5;212mRunning test 4 (f)...\033[0m" && \
	./myFind -f testdata/test01 && \
	echo "\033[38;5;212mRunning test 5 (f-ext)...\033[0m" && \
	./myFind -f -ext 'txt' testdata/test01 && \
	wait

myFindA:
	@echo "\033[38;5;122mRunning tests...\033[0m" && \
	go build -o myFind ./cmd/myFind && \
	echo "\033[38;5;122mRunning test 1...\033[0m" && \
	./myFind testdata/test01 && \
	echo "\033[38;5;122mRunning test 2 (sl)...\033[0m" && \
	./myFind -sl testdata/test01 && \
	echo "\033[38;5;122mRunning test 3 (d)...\033[0m" && \
	./myFind -d testdata/test01 && \
	echo "\033[38;5;122mRunning test 4 (f)...\033[0m" && \
	./myFind -f testdata/test01 && \
	echo "\033[38;5;122mRunning test 5 (f-ext)...\033[0m" && \
	./myFind -f -ext 'txt' testdata/test01 && \
	wait

.PHONY:myWc
myWc:
	@echo "\033[38;5;122mRunning tests...\033[0m" && \
	go build -o myWc ./cmd/myWc  && \
	echo "\033[38;5;122mRunning test 1...\033[0m" && \
	./myWc testdata/test02/text.txt testdata/test02/text1.txt && \
	echo "\033[38;5;122mRunning test 2 (w)...\033[0m" && \
	./myWc -w testdata/test02/text.txt testdata/test02/text1.txt && \
	echo "\033[38;5;122mRunning test 3 (m)...\033[0m" && \
	./myWc -m testdata/test02/text.txt testdata/test02/text1.txt && \
	echo "\033[38;5;122mRunning test 4 (l)...\033[0m" && \
	./myWc -l testdata/test02/text.txt testdata/test02/text1.txt && \
	wait

build:
	@echo "\033[38;5;122mBuilding files...\033[0m"
	go build -o myXargs ./cmd/myXargs
	go build -o myFind ./cmd/myFind
	go build -o myWc ./cmd/myWc


.PHONY:myXargs
myXargs: build
	@echo "\033[38;5;122mRunning tests...\033[0m" && \
	echo "\033[38;5;122mRunning test 1...\033[0m" && \
	./myFind -f -ext 'log' testdata/test03/logs | ./myXargs ./myWc 

.PHONY:myRotate
myRotate:
	@echo "\033[38;5;122mRunning tests...\033[0m" && \
	go build -o myRotate ./cmd/myRotate && \
	echo "\033[38;5;122mRunning test 1...\033[0m" && \
	./myRotate testdata/test04/logs/log.log && \
	echo "\033[38;5;122mRunning test 2...\033[0m" && \
	./myRotate -a testdata/archive testdata/test04/logs/log2.log testdata/test04/logs/log.log && \
	echo "\033[38;5;122mRunning error test 1 ...\033[0m" && \
	./myRotate  testdata/archive testdata/test04/logs/log2.log && \
	echo "\033[38;5;122mRunning error test 1 ...\033[0m" && \
	./myRotate testdata/test04/logs/log1.log && \
	wait



# Форматирование кода
.PHONY: fmt
fmt:
	@echo "Running gofmt..."
	@gofmt -w $$(find . -type f -name '*.go' | sed s/^\\.\\///)
	@echo "Gofmt completed."

clean:
	rm -rf myFind myXargs myWc myRotate 
	find . -type f -name '*.tag.gz' -exec rm -f {} +