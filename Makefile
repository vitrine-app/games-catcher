NAME=vitrine_games_catcher
GO=go
BINARY_NAME=games_catcher
SRCS=games_catcher.go \


all: $(NAME)

$(NAME):
	$(GO) build -o $(BINARY_NAME).so -buildmode=c-shared $(SRCS)
	rm $(BINARY_NAME).h

clean:
	$(GO) clean
	rm -f $(BINARY_NAME).so

deps:
	$(GO) get github.com/Henry-Sarabia/igdb
