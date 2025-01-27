package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	defer db.Close()
	require.NoError(t, err)
	clientID := 1
	// напиши тест здесь
	cl, err := selectClient(db, clientID)
	require.NoError(t, err)
	if assert.NotEmpty(t, cl.ID) {
		require.Equal(t, cl.ID, clientID)
	}
	require.NotEmpty(t, cl.FIO)
	require.NotEmpty(t, cl.Login)
	require.NotEmpty(t, cl.Birthday)
	require.NotEmpty(t, cl.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	defer db.Close()
	require.NoError(t, err)
	clientID := -1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)
	if assert.Error(t, err) {
		require.Equal(t, err, sql.ErrNoRows)
	}
	require.Empty(t, cl.FIO)
	require.Empty(t, cl.Login)
	require.Empty(t, cl.Birthday)
	require.Empty(t, cl.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	defer db.Close()
	require.NoError(t, err)

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)
	// напиши тест здесь
	testCl, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	require.Equal(t, cl.ID, testCl.ID)
	require.Equal(t, cl.FIO, testCl.FIO)
	require.Equal(t, cl.Login, testCl.Login)
	require.Equal(t, cl.Birthday, testCl.Birthday)
	require.Equal(t, cl.Email, testCl.Email)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	defer db.Close()
	require.NoError(t, err)

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь

	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	testCl, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	require.Equal(t, cl.ID, testCl.ID)

	err = deleteClient(db, testCl.ID)
	require.NoError(t, err)

	_, err = selectClient(db, testCl.ID)
	if assert.Error(t, err) {
		require.Equal(t, err, sql.ErrNoRows)
	}
}
