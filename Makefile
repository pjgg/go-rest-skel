VERSION = $(shell  if [ $TAG = "" ]; then echo "0.0.1"; else echo "$TAG"; fi)
SRC_PATH=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

version:
	echo $(VERSION)

clean: 
	rm -rf\
	 $(SRC_PATH)/dist\
	 $(SRC_PATH)/debug\
	 $(SRC_PATH)/godepgraph.png\
	 $(SRC_PATH)/*/cover.out\
	 $(SRC_PATH)/*/cover.html\
	 $(SRC_PATH)/pkg\
	 $(SRC_PATH)/src\