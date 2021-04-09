# Swagger NiFi


## Generate Go Client

Download the release of nifi `https://github.com/apache/nifi/releases`

unzip / untar the file

`cd /{nifi-release}/nifi-nar-bundles/nifi-framework-bundle/nifi-framework/nifi-web/nifi-web-api`

mvn clean package

`cd target/swagger-ui`

you should find the `swagger.json` file

mv the file to the root of this project

`swagger generate client`

Most likely some errors occur when generating:

```
- "getState" is defined 4 times
- "updateRunStatus" is defined 5 times
- "clearState" is defined 3 times
- "createControllerService" is defined 2 times
- "getPropertyDescriptor" is defined 3 times
- "removeDropRequest" is defined 2 times
- "deleteUpdateRequest" is defined 2 times
- operation "uploadTemplate" has both formData and body parameters. Only one such In: type may be used for a given operation
```
Resolving manually by renaming the operationId's by prefixing the tags camelCased.

e.g.
```
"tags" : [ "controller-services" ],
"summary" : "Gets the state for a controller service",
"description" : "",
"operationId" : "getState", <-- orginal
```
should become:
```
"tags" : [ "controller-services" ],
"summary" : "Gets the state for a controller service",
"description" : "",
"operationId" : "controllerServicesGetState", <-- to become
```

Add a miniumum to RevisionDTO.Version, this will create an *int64, so 0 can be set (default unmarshalling behaviour golang)
```
"RevisionDTO" : {
    "type" : "object",
    "properties" : {
    "version" : {
        "format" : "int64",
        "minimum": 0
```

Bug in the code of NiFi, when posting a 201 status code is returned, not a 200, so all the posts needs to have a 201.

TODO: UploadTemplate for now I removed the body, which isn't tested at all, should followup on this