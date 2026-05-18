package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql" // Turso (libSQL) ドライバ
)

type Idiom struct {
	ID      int    `json:"id"`
	Phrase  string `json:"phrase"`
	Reading string `json:"reading"`
	Meaning string `json:"meaning"`
	Usage   string `json:"usage"`
}

var db *sql.DB

func initDB() error {
	if db != nil {
		return nil
	}
	var err error

	// 1. Tursoの環境変数を取得
	dbUrl := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	if dbUrl == "" || authToken == "" {
		return fmt.Errorf("TURSO_DATABASE_URL または TURSO_AUTH_TOKEN が設定されていません")
	}

	// 2. Tursoに接続するための文字列を作成
	connStr := fmt.Sprintf("%s?authToken=%s", dbUrl, authToken)
	db, err = sql.Open("libsql", connStr)
	if err != nil {
		return err
	}
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("リクエストを受信しました！")

	// DBの初期化（Vercel用）
	if err := initDB(); err != nil {
		log.Printf("DB初期化エラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 環境変数から許可するドメインを取得（設定がない場合は開発用に * を使用）
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "*"
	}

	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// プリフライトリクエスト (OPTIONS) の対応
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// クエリパラメータからジャンルを取得し、テーブルを切り替える
	genre := r.URL.Query().Get("genre")
	if genre == "ことわざ" || genre == "koto" {
		tableName = "proverbs"
	}

	idioms := []Idiom{} // 空のスライスで初期化することで、データがない場合に null ではなく [] を返すようにします

	// 3. データベースから全件取得
	query := fmt.Sprintf("SELECT id, phrase, reading, meaning, usage FROM %s ORDER BY id ASC", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("データベースクエリのエラー: %v", err) // 詳細なエラーをGoのコンソールに表示
		http.Error(w, "DBエラーが発生しました", http.StatusInternalServerError)
		return
	}
	defer rows.Close() // 忘れずにrowsをクローズする

	for rows.Next() {
		var idiom Idiom
		err := rows.Scan(&idiom.ID, &idiom.Phrase, &idiom.Reading, &idiom.Meaning, &idiom.Usage)
		if err != nil {
			log.Printf("行スキャンエラー: %v", err)
			http.Error(w, "DBデータ読み込みエラー", http.StatusInternalServerError)
			return
		}
		idioms = append(idioms, idiom)
	}
	if err = rows.Err(); err != nil {
		log.Printf("行イテレーションエラー: %v", err)
		http.Error(w, "DBデータ処理エラー", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(idioms) // スライス全体をJSONとしてエンコード
}
