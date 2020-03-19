NAME := hidemaru_ide

DIST_DIR := $(APPDATA)\Hidemaruo

.PHONY: dist
dist:
	cp -pr Hidemaru $(DIST_DIR)
