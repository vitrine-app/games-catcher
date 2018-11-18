NAME=vitrine_games_catcher
GO=go
BINARY_NAME=games_catcher
SRCS=main.go \
	 db.go \
	 db_games.go \
	 db_companies.go \
	 db_series.go \
	 db_genres.go \
	 get_game.go \
	 get_company.go \
	 get_series.go \
	 get_genre.go \


all: $(NAME)

$(NAME):
	$(GO) build -o $(BINARY_NAME).so -buildmode=c-shared $(SRCS)
	rm $(BINARY_NAME).h

binary:
	$(GO) build -o $(BINARY_NAME) $(SRCS)

clean:
	$(GO) clean
	rm -f $(BINARY_NAME).so

deps:
	$(GO) get github.com/Henry-Sarabia/igdb
	$(GO) get github.com/go-sql-driver/mysql
