/*
 * Copyright (c) 2024 Michael Plunkett (https://github.com/michplunkett)
 * All rights reserved.
 * Used to parse Slack messages.
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

type Message struct {
	TimeStamp       string
	UserID          string         `json:"user"`
	TS              string         `json:"ts"`
	Type            string         `json:"type"`
	ClientMessageID string         `json:"client_msg_id"`
	Text            string         `json:"text,omitempty"`
	UserProfile     *UserProfile   `json:"user_profile,omitempty"`
	Attachments     *[]*Attachment `json:"attachments,omitempty"`
	Files           *[]File        `json:"files,omitempty"`
	IsUpload        bool           `json:"upload,omitempty"`
}

type UserProfile struct {
	ProfileImage string `json:"image_72,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	RealName     string `json:"real_name,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Attachment struct {
	Text        string `json:"text"`
	Fallback    string `json:"fallback"`
	FromURL     string `json:"from_url"`
	ServiceName string `json:"service_name"`
	ID          int    `json:"id"`
	OriginalURL string `json:"original_url"`
}

type File struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	MimeType     string `json:"mimetype"`
	FileType     string `json:"pretty_type"`
	IsExternal   bool   `json:"is_external"`
	IsPublic     bool   `json:"is_public"`
	DownloadLink string `json:"url_private_download"`
	Height       int    `json:"original_w"`
	Width        int    `json:"original_h"`
}

type CSVRecord struct {
	TimeStamp   string
	UserID      string
	UserName    string
	RealName    string
	MessageType string
	Text        string
	Attachments []string
	Files       []string
}

func main() {
	files, err := filepath.Glob("./SlackMessages/*.json")
	if err != nil {
		fmt.Println(err)
	}

	messages := make([]*Message, 0)

	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error occured while opening %s: %+v\n", file, err)
			continue
		}

		fileMessages := make([]*Message, 0)
		err = json.Unmarshal(bytes, &fileMessages)
		if err != nil {
			fmt.Printf("Error occured while parsing JSON from %s: %+v\n", file, err)
			continue
		}

		messages = append(messages, fileMessages...)
	}

	csvRecords := make([]*CSVRecord, 0)
	for _, msg := range messages {
		record := &CSVRecord{
			UserID:      msg.UserID,
			UserName:    "",
			RealName:    "",
			MessageType: msg.Type,
			Attachments: make([]string, 0),
			Files:       make([]string, 0),
		}

		timeStampSplit := strings.Split(msg.TS, ".")
		seconds, err := strconv.ParseInt(timeStampSplit[0], 10, 64)
		if err != nil {
			fmt.Printf("Error occured while parsing seconds: %+v\n", err)
			continue
		}

		nanoseconds, err := strconv.ParseInt(timeStampSplit[1], 10, 64)
		if err != nil {
			fmt.Printf("Error occured while parsing nanoseconds: %+v\n", err)
			continue
		}

		record.TimeStamp = time.Unix(seconds, nanoseconds).Format(time.RFC3339)

		if msg.UserProfile != nil {
			record.UserName = msg.UserProfile.Name
			record.RealName = msg.UserProfile.RealName
		}

		if msg.Text != "" {
			record.Text = msg.Text
		}

		if msg.Attachments != nil {
			for _, attachment := range *msg.Attachments {
				record.Attachments = append(record.Attachments, attachment.OriginalURL)
			}
		}

		if msg.Files != nil {
			for _, file := range *msg.Files {
				record.Files = append(record.Files, file.DownloadLink)
			}
		}

		csvRecords = append(csvRecords, record)
	}

	csvFile, err := os.Create("slack_records.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	_ = gocsv.MarshalFile(csvRecords, csvFile)
}
