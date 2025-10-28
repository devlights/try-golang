package main

/*
#include <stdio.h>
#include <string.h>

extern void go_func(const char *s, size_t n);

void c_func() {
	char s[] = "helloworld";
	size_t s_size = sizeof(s);
	go_func(s, s_size);
}

void c_func2(const char *s, size_t n) {
	char buf[n];
	{
		memcpy(buf, s, n);
		buf[n-1] = '\0';
	}

	printf("[C ] %s\n", buf);
}
*/
import "C"
