include $(GOROOT)/src/Make.inc

TARG=alsa

CGOFILES=\
        alsa.go\

CGO_LDFLAGS=-lasound

include $(GOROOT)/src/Make.pkg

format:
	find . -type f -name '*.go' -exec gofmt -w {} \;

arch-install:
	mkdir -p "$(DESTDIR)$(pkgdir)"
	cp _obj/$(TARG).a "$(DESTDIR)$(pkgdir)"
