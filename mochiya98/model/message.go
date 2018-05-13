package model

import (
	"database/sql"
)

// Message はメッセージの構造体です
type Message struct {
	ID       int64  `json:"id"`
	Body     string `json:"body"`
	Msgtype	 int64	`json:"type"`
	UserName string `json:"username"`
}

// MessagesAll は全てのメッセージを返します
func MessagesAll(db *sql.DB) ([]*Message, error) {

	// Tutorial 1-1. ユーザー名を表示しよう
	rows, err := db.Query(`select id, body, msgtype, username from message`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms []*Message
	for rows.Next() {
		m := &Message{}
		// Tutorial 1-1. ユーザー名を表示しよう
		if err := rows.Scan(&m.ID, &m.Body, &m.Msgtype, &m.UserName); err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// MessageByID は指定されたIDのメッセージを1つ返します
func MessageByID(db *sql.DB, id string) (*Message, error) {
	m := &Message{}

	// Tutorial 1-1. ユーザー名を表示しよう
	if err := db.QueryRow(`select id, body, msgtype, username from message where id = ?`, id).Scan(&m.ID, &m.Body, &m.Msgtype, &m.UserName); err != nil {
		return nil, err
	}

	return m, nil
}

// Insert はmessageテーブルに新規データを1件追加します
func (m *Message) Insert(db *sql.DB) (*Message, error) {
	// Tutorial 1-2. ユーザー名を追加しよう
	res, err := db.Exec(`insert into message (body, username, msgtype) values (?, ?, ?)`, m.Body, m.UserName, m.Msgtype)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Message{
		ID:   id,
		Body: m.Body,
		Msgtype: m.Msgtype,
		UserName: m.UserName,
	}, nil
}

// Update は...
func (m *Message) Update(db *sql.DB) (*Message, error) {
	_, err := db.Exec(`update message set body=?, username=? where id = ?`, m.Body, m.UserName, m.ID)
	if err != nil {
		return nil, err
	}

	return &Message{
		ID:       m.ID,
		Body:     m.Body,
		UserName: m.UserName,
	}, nil
}

// Mission 1-2. メッセージを削除しよう

// Delete は...
func (m *Message) Delete(db *sql.DB) (*Message, error) {
	_, err := db.Exec(`delete from message where id = ?`, m.ID)
	if err != nil {
		return nil, err
	}

	return &Message{
		ID:       m.ID,
		Body:     m.Body,
		UserName: m.UserName,
	}, nil
}
