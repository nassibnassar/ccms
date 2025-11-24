package harvest

import (
	"fmt"

	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/cmd/ccd/marcxml"
	"github.com/nassibnassar/goharvest/oai"
)

func Harvest() {
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

		//Verb:  "ListRecords",
		//From:  "2025-10-14T22:52:14Z",
		//Until: "2025-10-14T22:52:16Z",

		Verb:           "GetRecord",
		Identifier:     "b6c6160c-6bbb-41cc-9e07-690049d7d537",
		MetadataPrefix: "marcxml",
	}
	//fmt.Printf("%#v\n\n", rq)
	//fmt.Println("harvesting...")
	var c int
	//rq.HarvestRecords(func(record *oai.Record) {
	rq.Harvest(func(rs *oai.Response) {
		//fmt.Printf("identifier: %s\n", record.Header.Identifier)
		//fmt.Printf("datestamp: %s\n", record.Header.DateStamp)
		//fmt.Printf("setspec: %v\n", record.Header.SetSpec)
		//fmt.Printf("status: %s\n", record.Header.Status)
		//fmt.Printf("about: %s\n", record.About.Body)
		//fmt.Printf("metadata: %s\n", record.Metadata.Body)

		//fmt.Printf("%s %s\n", record.Header.Identifier, record.Header.DateStamp)

		fmt.Printf("=======================================================================\n")
		fmt.Printf("%v\n", c)
		c++
		fmt.Printf("=======================================================================\n")
		fmt.Printf("%#v\n", rs)
		fmt.Printf("=======================================================================\n")
		fmt.Printf("retrieved = %s\n", rs.ResponseDate)
		fmt.Printf("identifier = %s\n", rs.GetRecord.Record.Header.Identifier)
		fmt.Printf("data = %s\n", rs.GetRecord.Record.Metadata.Body)
		m, err := marcxml.Unmarshal(rs.GetRecord.Record.Metadata.Body)
		if err != nil {
			panic(err)
		}
		title245, _ := m.Lookup("245", " ", " ", "a")
		fmt.Printf("title245 = %s\n", title245)
	})
	// TODO add to ccd.conf:
	// [oai]
	// base_url = https://cclp-okapi.reshare-dev.indexdata.com/_/invoke/tenant/cclp/reservoir/oai
	fmt.Printf("=======================================================================\n")
	fmt.Printf("harvest exiting\n")
	fmt.Printf("=======================================================================\n")
}
