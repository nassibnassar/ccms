package harvest

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/cmd/ccd/marcxml"
	"github.com/indexdata/ccms/internal/global"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nassibnassar/goharvest/oai"
)

func Harvest(dp *pgxpool.Pool) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("harvest aborted: %v", r)
		}
	}()
	/*
		(&oai.Request{
			BaseURL: "http://services.kb.nl/mdo/oai", Set: "DTS", MetadataPrefix: "dcx",
			From: "2012-09-06T014:00:00.000Z",
		}).HarvestRecords(func(record *oai.Record) {
			fmt.Printf("%s\n\n", record.Metadata.Body[0:500])
		})
	*/

	//(&oai.Request{
	//        BaseURL: "https://cclp-okapi.reshare-dev.indexdata.com/_/invoke/tenant/cclp/reservoir/oai",

	//        //Verb:  "ListRecords",
	//        //From:  "2025-10-14T22:52:14Z",
	//        //Until: "2025-10-14T22:52:16Z",

	//        Verb:           "GetRecord",
	//        Identifier:     "b6c6160c-6bbb-41cc-9e07-690049d7d537",
	//        MetadataPrefix: "marcxml",
	//}).HarvestRecords(func(record *oai.Record) {
	//        //fmt.Printf("identifier: %s\n", record.Header.Identifier)
	//        //fmt.Printf("datestamp: %s\n", record.Header.DateStamp)
	//        //fmt.Printf("setspec: %v\n", record.Header.SetSpec)
	//        //fmt.Printf("status: %s\n", record.Header.Status)
	//        //fmt.Printf("about: %s\n", record.About.Body)
	//        //fmt.Printf("metadata: %s\n", record.Metadata.Body)

	//        fmt.Printf("%s %s\n", record.Header.Identifier, record.Header.DateStamp)
	//})

	// The following OAI-PMH verbs are supported by the Reservoir: ListIdentifiers, ListRecords,
	// GetRecord, Identify.
	rq := &oai.Request{
		BaseURL: "https://cclp-okapi.reshare-dev.indexdata.com/_/invoke/tenant/cclp/reservoir/oai",

		Verb:  "ListRecords",
		From:  "2024-10-14T22:52:14Z",
		Until: "2026-10-14T22:52:16Z",

		//Verb:           "GetRecord",
		//Identifier:     "b6c6160c-6bbb-41cc-9e07-690049d7d537",
		//MetadataPrefix: "marcxml",
	}
	//fmt.Printf("%#v\n\n", rq)
	//fmt.Println("harvesting...")
	//var c int
	rq.HarvestRecords(func(record *oai.Record) {
		//rq.Harvest(func(rs *oai.Response) {
		//fmt.Printf("identifier: %s\n", record.Header.Identifier)
		//fmt.Printf("datestamp: %s\n", record.Header.DateStamp)
		//fmt.Printf("setspec: %v\n", record.Header.SetSpec)
		//fmt.Printf("status: %s\n", record.Header.Status)
		//fmt.Printf("about: %s\n", record.About.Body)
		//fmt.Printf("metadata: %s\n", record.Metadata.Body)

		//fmt.Printf("%s %s\n", record.Header.Identifier, record.Header.DateStamp)

		/*
			fmt.Printf("=======================================================================\n")
			fmt.Printf("%v\n", c)
			c++
			fmt.Printf("=======================================================================\n")
			fmt.Printf("%#v\n", rs)
			fmt.Printf("=======================================================================\n")
			fmt.Printf("retrieved = %s\n", rs.ResponseDate)
			fmt.Printf("identifier = %s\n", rs.GetRecord.Record.Header.Identifier)
			fmt.Printf("data = %s\n", rs.GetRecord.Record.Metadata.Body)
		*/
		m, err := marcxml.Unmarshal(record.Metadata.Body)
		//m, err := marcxml.Unmarshal(rs.GetRecord.Record.Metadata.Body)
		if err != nil {
			panic(err)
		}
		metadata := strings.TrimSpace(string(record.Metadata.Body))
		identifier := strings.TrimPrefix(record.Header.Identifier, "oai:")
		dateStamp := record.Header.DateStamp
		author100a := strings.Join(m.Lookup("100", "a"), "\n")
		title245a := m.Lookup("245", "a")
		title245b := m.Lookup("245", "b")
		title245 := strings.Join(append(title245a, title245b...), "\n")
		//fmt.Printf("%s [%s] %s\t%s\n", dateStamp, identifier, author100a, title245)
		_ = m
		//c++
		//if c%10000 == 0 {
		//        log.Info("harvested %d records", c)
		//}

		var author, title *string
		if author100a != "" {
			author = &author100a
		}
		if title245 != "" {
			title = &title245
		}

		tx, err := dp.Begin(context.TODO())
		if err != nil {
			panic(err)
		}
		defer tx.Rollback(context.TODO())

		var id int64
		q := "insert into ccms.md (identifier, retrieved, date_stamp, data) " +
			"values ($1, $2, $3, $4) on conflict do nothing returning id"
		err = tx.QueryRow(context.TODO(), q, identifier, time.Now(), dateStamp, metadata).Scan(&id)
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Info("conflict: skipping record %s", identifier)
		case err != nil:
			panic(fmt.Sprintf("writing to table "+global.SystemSchema+".md: %v", err))
		default:
		}

		if id != 0 {
			q = "insert into ccms.attr (id, author, title) " +
				"values ($1, $2, $3) on conflict do nothing"
			if _, err = tx.Exec(context.TODO(), q, id, author, title); err != nil {
				panic(fmt.Sprintf("writing to table "+global.SystemSchema+".attr: %v", err))
			}

			q = "insert into ccms.reserve (id) " +
				"values ($1) on conflict do nothing"
			if _, err = tx.Exec(context.TODO(), q, id); err != nil {
				panic(fmt.Sprintf("writing to table "+global.SystemSchema+".reserve: %v", err))
			}

			log.Info("(%d) %s %s / %s", id, identifier, author100a, title245)
		}

		if err = tx.Commit(context.TODO()); err != nil {
			panic(fmt.Sprintf("writing harvested data: committing changes: %v", err))
		}

		//os.Exit(0)
	})
	// TODO add to ccd.conf:
	// [oai]
	// base_url = https://cclp-okapi.reshare-dev.indexdata.com/_/invoke/tenant/cclp/reservoir/oai
	fmt.Printf("=======================================================================\n")
	fmt.Printf("harvest exiting\n")
	fmt.Printf("=======================================================================\n")
}
