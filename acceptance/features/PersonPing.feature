Feature: ping person acceptanceTest
  
Scenario: As  developer I would like to ping a person
  Given a person with name "Pablo"
  And a content-type "application/json"   
  When I request to ping a person
  Then response code must be 200
  Then response must be a content type "application/json; charset=utf-8"

Scenario: As  developer I would like to ping a person
  Given a person with name "Pablo"
  When I request to ping a person
  Then response code must be 415
 