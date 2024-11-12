// ~/Workspace/sqlcdemo/main.go

package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"moustacios.dev/sqlcdemo/moustacios"
)

func play() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", "sqlcdemo:secret@tcp(localhost:3306)/sqlcdemo?parseTime=true")
	if err != nil {
		return err
	}

	queries := moustacios.New(db)

	// create a comment
	_, err = queries.SaveComment(ctx, moustacios.SaveCommentParams{
		Email:       "someone@theinter.net",
		CommentText: "Un comentariu salvat din funcția run().",
	})
	if err != nil {
		return err
	}

	// create a second comment
	result, err := queries.SaveComment(ctx, moustacios.SaveCommentParams{
		Email:       "bot@theinter.net",
		CommentText: "Un alt comentariu salvat din funcția run().",
	})
	if err != nil {
		return err
	}

	// Marchează ca bot comment
	botCommentID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	_, err = queries.FlagBotComment(ctx, moustacios.FlagBotCommentParams{
		ID:             botCommentID,
		BotProbability: sql.NullInt16{Int16: 90, Valid: true},
	})
	log.Println("botCommentID", botCommentID)

	//insertedCommentID, err := result.LastInsertId()
	//if err != nil {
	//	return err
	//}
	//log.Println("insertedCommentID", insertedCommentID)

	// Listează toate comentariile
	comments, err := queries.ListComments(ctx)
	if err != nil {
		return err
	}
	log.Println(comments)

	// delete the comment we just inserted
	//_, err = queries.DeleteComment(ctx, insertedCommentID)
	//if err != nil {
	//	return err
	//}

	// Elimină comentariile marcate cu probabilitate de bot > 50
	_, err = queries.PurgeBotComments(ctx, sql.NullInt16{Int16: 50, Valid: true})
	if err != nil {
		return err
	}

	// Listează comentariile rămase
	comments, err = queries.ListComments(ctx)
	if err != nil {
		return err
	}
	log.Println(comments)

	// done
	return nil
}

func main() {
	if err := play(); err != nil {
		log.Fatal(err)
	}
}
