// +build !unit

package acceptance

import "github.com/DATA-DOG/godog"

type personCrudTestContext struct {
	personContext *testContext
}

func newPersonCRUDTestContext(context *testContext) *personCrudTestContext {
	return &personCrudTestContext{
		personContext: context,
	}
}

func (context *personCrudTestContext) resetPersonDb(interface{}, error) {
	context.personContext.resetContext(nil)
	// TODO reset test resources
}

func (context *personCrudTestContext) aPersonWithName(personName string) error {
	context.personContext.userName = personName
	return nil
}

func (context *personCrudTestContext) iRequestToPingAPerson() error {
	req, _ := newHTTPRequest("GET", "http://"+getAddr()+"/v1/person/ping", nil, nil)
	req.Header.Add("content-type", context.personContext.requestContentType)
	resp, err := makeHTTPQuery(req)

	context.personContext.responseCode = resp.StatusCode
	context.personContext.body = resp.Body
	context.personContext.resp = resp

	return err
}

func PersonCrudFeatureContext(personCrudSuit *godog.Suite, context *testContext) {

	personContext := newPersonCRUDTestContext(context)
	personCrudSuit.AfterScenario(personContext.resetPersonDb)

	personCrudSuit.Step(`^a person with name "([^"]*)"$`, personContext.aPersonWithName)
	personCrudSuit.Step(`^I request to ping a person$`, personContext.iRequestToPingAPerson)

}
